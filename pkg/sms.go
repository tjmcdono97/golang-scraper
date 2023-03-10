package pkg

import (
	"fmt"
	"log"
	"os"

	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)


func SendMessage(text string) error {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	apiSid := os.Getenv("TWILIO_SID_DREWS_MAIN")
	apiSecret := os.Getenv("TWILIO_SECRET_DREWS_MAIN")
	to := os.Getenv("RECIPIENT_PHONE_NUMBER")
	from := os.Getenv("TWILIO_PHONE_NUMBER")
	fmt.Println(apiSid)
	fmt.Println(apiSecret)
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   apiSid,
		Password:   apiSecret,
		AccountSid: accountSid,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody(text)
	log.Println(text)
	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("SMS sent successfully!")
	}
	return nil
}


// Sending a text message to a phone number.
func Alert() error {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	apiSid := os.Getenv("TWILIO_SID_DREWS_MAIN")
	apiSecret := os.Getenv("TWILIO_SECRET_DREWS_MAIN")
	to := os.Getenv("RECIPIENT_PHONE_NUMBER")
	from := os.Getenv("TWILIO_PHONE_NUMBER")
	fmt.Println(apiSid)
	fmt.Println(apiSecret)
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username:   apiSid,
		Password:   apiSecret,
		AccountSid: accountSid,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("drew recieved a craigslist post")
	_, err := client.Api.CreateMessage(params)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("SMS sent successfully!")
	}
	return nil
}
