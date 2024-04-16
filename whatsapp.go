package main

import (
	"fmt"

	twilio "github.com/twilio/twilio-go"
	api "github.com/twilio/twilio-go/rest/api/v2010"
)


func main() {
	client := twilio.NewRestClient()

	params := &api.CreateMessageParams{}
	params.SetTo("whatsapp:+918217325610")
	params.SetFrom("whatsapp:+14155238886")
	params.SetBody("Hello from Golang!")

	//_, err := client.ApiV2010.CreateMessage(params)
	_, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Message sent successfully!")
	}
}

