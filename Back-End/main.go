package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	Success        bool
	Textid         string
	QuotaRemaining int
}
type Sender interface {
	send(textMessage) TextResponse
}

type messenger struct {
	url    string
	apiKey string
}

func (m *messenger) send(textMessage textMessage) TextResponse {
	formData := url.Values{
		"phone":   {textMessage.number},
		"message": {textMessage.message},
		"key":     {m.apiKey},
	}
	fmt.Println("1: Sending SMS to:", textMessage.number)

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

func NewMessenger(baseUrkey string, key string) Sender {
	fmt.Printf("BaseURL %v  Apikey %v\n", baseUrkey, key)
	return &messenger{
		url:    baseUrkey,
		apiKey: key,
	}
}

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
func BatchMessages(input []textMessage, Provider Sender) []error {
	var mistakes []error
	for _, value := range input {
		response := Provider.send(value)
		println("7")
		if !response.Success {
			mistakes = append(mistakes, fmt.Errorf("erros has occured proccessing %s", response.Textid))
		}
		fmt.Println(response)
	}
	return mistakes
}
func main() {
	print("started\n")
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
	*/

	fmt.Println("Finished")
}
