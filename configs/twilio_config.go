package configs

import (
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

func CreateTwilioAuth() *twilio.RestClient {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return client
}
