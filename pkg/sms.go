package pkg

import (
	"fmt"
	"log"
	"os"

	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

// SendMessage sends a text message with the specified text content.
// It uses the Twilio API to send the message.
func SendMessage(text string) error {
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

	log.Printf("Sending SMS: %s", text)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending SMS: %v", err)
		return err
	}

	log.Println("SMS sent successfully!")
	return nil
}

// Alert sends an alert message notifying that user received a Craigslist post.
// It uses the Twilio API to send the message.
func Alert() error {
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
	params.SetBody("Drew received a Craigslist post")

	log.Println("Sending alert SMS: received a Craigslist post")

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Printf("Error sending alert SMS: %v", err)
		return err
	}

	log.Println("Alert SMS sent successfully!")
	return nil
}
