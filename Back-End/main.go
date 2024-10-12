package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type textMessage struct {
	number  string
	message string
}

func CreateSMS(number string, message string) textMessage {
	return textMessage{
		number:  number,
		message: message,
	}
}

type TextResponse struct {
	Success        bool `json:"success"`
	Textid         string
	QuotaRemaining int `json:"quotaRemaining"`
}

type ErrorTextResponse struct {
	Success        bool   `json:"success"`
	Error          string `json:"error"`
	QuotaRemaining int    `json:"quotaRemaining"`
}

type Sender interface {
	send(textMessage) TextResponse
}

type messenger struct {
	url    string
	apiKey string
}

func NewMessenger(baseUrkey string, key string) Sender {
	return &messenger{
		url:    baseUrkey,
		apiKey: key,
	}
}

type emailInfo struct {
	recipient string
	body      string
}

func newmail(reciept string, body string) emailInfo {
	return emailInfo{
		recipient: reciept,
		body:      body,
	}
}

type Mailer struct {
	sender    string
	password  string
	customers []string
}

func NewMailer() Mailer {
	return Mailer{
		sender:    os.Getenv("SENDEREMAIL"),
		password:  "tiwk yiyk pbch enbb",
		customers: make([]string, 1),
	}
}

func (m *Mailer) Sendmail(email emailInfo) {
	sendEmail(m.sender, m.password, email.recipient, SubjectAlarm3, email.body)
}

type ModelResponse struct {
	Description string `json:"description"`
	Link        string `json:"link"`
	Issue       string `json:"issue"`
}
type RequestData struct {
	Email  string `json:"email"`
	Number string `json:"number"`
}

type ClientDataResponse struct {
	NumberCLients []string
	MailClients   []string
	TotalCount    int `json:"total"`
}

var ClientData_Number map[string][]string
var ClientData_Mail map[string][]string

const (
	SubjectAlarm1 = "URGENT: Immediate Medical Attention Required for [Family Member]"
	SubjectAlarm2 = "ALERT: Medical Emergency Detected - Immediate Help Needed"
	SubjectAlarm3 = "EMERGENCY: [Family Member] is Experiencing a Critical Health Event"

	MedicalAlertChoking = "URGENT: [Family Member] appears to be choking. They are unable to breathe or speak. Please perform the Heimlich maneuver if trained, or call emergency services immediately."

	MedicalAlertStroke = "CRITICAL: [Family Member] is exhibiting signs of a stroke, such as slurred speech, confusion, or weakness on one side of their body. Every second counts. Please call 911 and provide immediate assistance."

	MedicalAlertAllergicReaction = "EMERGENCY: [Family Member] is experiencing a severe allergic reaction, including difficulty breathing and swelling. Administer an EpiPen if available, and seek emergency medical help immediately."
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Verify if a specific environment variable is loaded (for example, TEXTBELTKEY)
	if os.Getenv("TEXTBELTKEY") == "" {
		log.Fatal("TEXTBELTKEY environment variable not set.")
	}

	fmt.Println("Environment variables successfully loaded")
}

func clientInfo() (map[string][]string, map[string][]string) {
	if ClientData_Number != nil && ClientData_Mail != nil {
		return ClientData_Number, ClientData_Mail
	}
	ClientData_Number = make(map[string][]string, 1)
	ClientData_Mail = make(map[string][]string, 1)
	return ClientData_Number, ClientData_Mail
}
func (m *messenger) send(textMessage textMessage) TextResponse {
	formData := url.Values{
		"phone":   {textMessage.number},
		"message": {textMessage.message},
		"key":     {m.apiKey},
	}

	fmt.Println("1: Sending SMS to:", textMessage.number)
	fmt.Println("Request to textBelt -> ", formData)
	// Create a new POST request
	resp, err := http.PostForm(m.url, formData)
	fmt.Println("2: Request sent")

	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return TextResponse{}
	}

	// Check if resp is nil to avoid defer panic
	if resp == nil {
		fmt.Println("Error: Received nil response")
		return TextResponse{}
	}
	fmt.Println("3: Received response with status:", resp.Status)

	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return TextResponse{}
	}

	// Step 2: Copy the body into a temporary clone
	tempClone := make([]byte, len(bodyBytes))
	copy(tempClone, bodyBytes)

	// Step 3: Unmarshal the byte array into an empty interface
	var result interface{}
	err = json.Unmarshal(tempClone, &result)
	if err != nil {
		fmt.Printf("Error unmarshalling response body into interface: %v\n", err)
		return TextResponse{}
	}
	var ErrorRespo TextResponse

	// Step 4: Type assertion to read the properties (example with "success" property)
	if resMap, ok := result.(map[string]interface{}); ok {
		if success, ok := resMap["success"].(bool); ok {
			fmt.Printf("Success: %v\n", success)
			ErrorRespo.Success = success
		}
		if errorMsg, ok := resMap["error"].(string); ok {
			fmt.Printf("Error: %s\n", errorMsg)
		}
		if quotaRemaining, ok := resMap["quotaRemaining"].(float64); ok { // JSON numbers are float64 by default
			fmt.Printf("Quota Remaining: %d\n", int(quotaRemaining))
			ErrorRespo.QuotaRemaining = int(quotaRemaining)
		}
		return ErrorRespo
	} else {
		fmt.Println("Error: Response is not a valid JSON object.")
	}

	var response TextResponse
	fmt.Println("4: Decoding response")
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Printf("Error decoding JSON response: %v\n", err)
		return TextResponse{}
	}
	fmt.Println("5: Response decoded successfully:", response)
	return response
}

func BatchMessages(input []textMessage, Provider Sender) []error {
	var mistakes []error
	for _, value := range input {
		response := Provider.send(value)
		println("7")
		if !response.Success {
			mistakes = append(mistakes, fmt.Errorf("erros has occured %v", response))
		}
	}
	return mistakes
}

// Function to send an email
func sendEmail(senderEmail, senderPassword, recipientEmail, subject, body string) {
	// Set up authentication information.
	auth := smtp.PlainAuth("", senderEmail, senderPassword, "smtp.gmail.com")

	// Format the message
	msg := "From: " + senderEmail + "\n" +
		"To: " + recipientEmail + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// Send the email
	err := smtp.SendMail(
		"smtp.gmail.com:587",     // SMTP server and port
		auth,                     // Authentication
		senderEmail,              // Sender email
		[]string{recipientEmail}, // Recipient email
		[]byte(msg),              // Message body as byte array
	)
	if err != nil {
		fmt.Printf("Error sending email: %v\n", err)
		return
	}

	fmt.Printf("Email successfully sent to %s\n", recipientEmail)
}

func Createtxt(numbers []string, msg string) []textMessage {
	var txt []textMessage
	for _, number := range numbers {
		tempStruct := textMessage{
			number:  number,
			message: msg,
		}
		txt = append(txt, tempStruct)
	}
	return txt
}
func SetNotifications(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Url hit -> %v\n", req.URL)
	meth := req.Method
	queryParams := req.URL.Query()
	clientID := queryParams.Get("client")
	fmt.Println(clientID)

	w.Header().Set("Content-Type", "application/json")

	// Fetch the current maps
	number_Map, mail_Map := clientInfo()

	// Get the existing slices for the client ID
	numberSlice := number_Map[clientID]
	mail_Slice := mail_Map[clientID]

	switch meth {
	case "GET":
		// Create the response struct
		resp := ClientDataResponse{
			NumberCLients: numberSlice,
			MailClients:   mail_Slice,
			TotalCount:    len(numberSlice) + len(mail_Slice),
		}

		// Convert the response struct to JSON
		jsonResponse, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "Error converting to JSON", http.StatusInternalServerError)
			return
		}

		// Write the JSON response
		w.Write(jsonResponse)

	case "PUT":
		var data RequestData

		// Parse JSON body
		err := json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Append to the slices
		numberSlice = append(numberSlice, data.Number)
		mail_Slice = append(mail_Slice, data.Email)

		// Update the global maps with the new slices
		number_Map[clientID] = numberSlice
		mail_Map[clientID] = mail_Slice

		// Respond with success
		w.WriteHeader(200)
		fmt.Fprintf(w, "Received email: %s and number: %s", data.Email, data.Number)
		fmt.Println("Current Numberslice ", numberSlice, "Current Mail slice", mail_Slice)
	case "POST":
		// Send out messages to clients mail and sms
		var Request_data ModelResponse

		// Parse JSON body
		err := json.NewDecoder(req.Body).Decode(&Request_data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(numberSlice) == 0 || len(mail_Slice) == 0 {
			w.Write([]byte("Client has no people to send out notifications too"))
			w.WriteHeader(204)
			return
		}
		SMSSender := NewMessenger(os.Getenv("TEXTBELTURL"), os.Getenv("TEXTBELTKEY"))
		batch := Createtxt(numberSlice, Request_data.Description)
		Ers := BatchMessages(batch, SMSSender)
		fmt.Println("Erros ", Ers)
		PostMan := NewMailer()
		for _, Email := range mail_Slice {
			mail := newmail(Email, Request_data.Issue+"        Latest Image from live feed ->"+Request_data.Link)
			fmt.Printf("Message sent to %v", Email)
			go PostMan.Sendmail(mail)
		}
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
		w.Write([]byte("Invalid Request type"))
	}
}
func helloWord(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Url hit -> %v\n", req.URL)
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func ProcessVideoHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Url hit -> %v\n", req.URL)

}
func main() {
	fmt.Println("started")

	// Map client ID to relatives TWo maps clientID -> [array of phone numbers]    clientID -> [array of emails]
	// iterate through both and send Notifications to both groups.

	http.HandleFunc("/", helloWord)
	http.HandleFunc("/Alert", SetNotifications) //
	http.HandleFunc("/Video", ProcessVideoHandler)
	print("starting \n")
	http.ListenAndServe(":8080", nil)
	// Email credentials
	/*
			SMSSender := NewMessenger(os.Getenv("TEXTBELTURL"), os.Getenv("TEXTBELTKEY"))
			message1 := CreateSMS("9085252880", "One step closer")
			message2 := CreateSMS("6479132144", "Were getting a little closer")
			message3 := CreateSMS("6463596966", "A couple steps closer")
			message4 := CreateSMS("8574980409", "Another couple of steps")
			GroupMessage := []textMessage{message1, message2, message3, message4}
			fmt.Println("about to process")
			errs := BatchMessages(GroupMessage, SMSSender)
			fmt.Println(errs)


		TeamMails := []string{"rbb98@scarletmail.rutgers.edu", "naweensarwari@gmail.com", "mingyu@bu.edu", "aymane.omari@nyu.edu"}
		PostMan := NewMailer()
		for _, Email := range TeamMails {
			mail := newmail(Email, MedicalAlertStroke)
			fmt.Printf("Message sent to %v", Email)
			PostMan.Sendmail(mail)
		}
		//mail := newmail("rbb98@scarletmail.rutgers.edu", "mill due")
		//PostMan.Sendmail(mail)
	*/

	fmt.Println("done")
}
