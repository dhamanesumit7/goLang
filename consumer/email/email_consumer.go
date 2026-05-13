package main

import (
	"context"
	"encoding/json"
	"fmt"

	"golang-api/models"

	"github.com/segmentio/kafka-go"
)

func main() {

	reader := kafka.NewReader(
		kafka.ReaderConfig{
			Brokers: []string{"localhost:9092"},

			Topic: "user-email",

			GroupID: "email-group",
		},
	)

	fmt.Println(
		"Email Consumer Started...",
	)

	for {

		msg, err := reader.ReadMessage(
			context.Background(),
		)

		if err != nil {

			fmt.Println(
				"Consumer error:",
				err,
			)

			continue
		}

		var event models.EmailEvent

		err = json.Unmarshal(
			msg.Value,
			&event,
		)

		if err != nil {

			fmt.Println(
				"JSON parse error:",
				err,
			)

			continue
		}

		fmt.Println(
			"Sending email to:",
			event.Email,
		)

		fmt.Println(
			"Message:",
			event.Message,
		)
	}
}
