package commons

import (
	"example/kenva-be/configs"
	"log"

	"github.com/twilio/twilio-go"
)

func CreateTwilioAuth() *twilio.RestClient {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load config")
		return nil
	}
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: configs.TwilioAccountSid,
		Password: configs.TwilioAuthenToken,
	})
	return client
}
