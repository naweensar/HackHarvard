import smtplib
from email.mime.text import MIMEText
from email.mime.multipart import MIMEMultipart

def send_email(sender_email, sender_password, recipient_email, subject, body):
    try:
        # Create a multipart message
        msg = MIMEMultipart()
        msg['From'] = sender_email
        msg['To'] = recipient_email
        msg['Subject'] = subject

        # Attach the email body
        msg.attach(MIMEText(body, 'plain'))

        # Connect to Gmail's SMTP server
        server = smtplib.SMTP('smtp.gmail.com', 587)
        server.starttls()  # Upgrade the connection to a secure encrypted SSL/TLS
        server.login(sender_email, sender_password)  # Log in to the email account

        # Send the email
        text = msg.as_string()
        server.sendmail(sender_email, recipient_email, text)

        # Close the connection
        server.quit()

        print(f"Email sent to {recipient_email}")
    except Exception as e:
        print(f"Error sending email: {e}")

if __name__ == "__main__":
    # Fill in your details below
    print("starting")
    sender_email = "richiebbaah@gmail.com"  # Your Gmail address
    sender_password = "tiwk yiyk pbch enbb"  # Your Gmail password (or App Password if 2FA is enabled)
    
    recipient_email = "rbb98@scarletmail.rutgers.edu"  # Recipient's email address

    # Email subject and body
    subject = "Test Email from Python"
    body = "This is a test email sent using a Python script."

    # Send the email
    send_email(sender_email, sender_password, recipient_email, subject, body)
