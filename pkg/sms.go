package pkg

import (
	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
)

// SendMessage sends a text message with the specified text content.
// It uses the Twilio API to send the message.
func SendMessage(text string) error {
	return sendSMS(text, "Sending SMS: %s")
}

// Alert sends an alert message notifying that user received a post.
// It uses the Twilio API to send the message.
func Alert() error {
	return sendSMS("Received a new post", "Sending alert SMS: received a post")
}

func sendSMS(text, logMessage string) error {
	// removed duplicate code to a separate function
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	apiSid := os.Getenv("TWILIO_SID")
	apiSecret := os.Getenv("TWILIO_SECRET")
	to := os.Getenv("RECIPIENT_PHONE_NUMBER")
	from := os.Getenv("TWILIO_PHONE_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   apiSid,
		Password:   apiSecret,
		AccountSid: accountSid,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(text)

	log.Printf(logMessage, text)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		return err
	}

	log.Println("SMS sent successfully!")
	return nil
}
