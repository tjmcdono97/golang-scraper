package pkg

import (
	twilio "github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"os"

	"go.uber.org/zap"
)

func SendMessage(text string, logger *zap.Logger) error {
	return sendSMS(text, "Sending SMS", logger)
}

func Alert(logger *zap.Logger) error {
	return sendSMS("Received a new post", "Sending alert SMS: received a post", logger)
}

func sendSMS(text, logMessage string, logger *zap.Logger) error {
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

	logger.Info(logMessage, zap.String("Message", text))

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		logger.Error("Error sending SMS", zap.Error(err))
		return err
	}

	logger.Info("SMS sent successfully!")
	return nil
}
