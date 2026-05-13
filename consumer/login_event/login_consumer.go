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

			Topic: "user-login",

			GroupID: "login-group",
		},
	)

	fmt.Println(
		"Login Consumer Started...",
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

		var event models.LoginEvent

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
			"User Logged In:",
			event.Email,
		)
	}
}
