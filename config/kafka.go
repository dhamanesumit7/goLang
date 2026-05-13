package config

import (
	"golang-api/logger"

	"github.com/segmentio/kafka-go"
)

var LoginWriter *kafka.Writer

var EmailWriter *kafka.Writer

func ConnectKafka() {

	LoginWriter = &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),

		Topic: "user-login",

		Balancer: &kafka.LeastBytes{},
	}

	EmailWriter = &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),

		Topic: "user-email",

		Balancer: &kafka.LeastBytes{},
	}

	logger.InfoLogger.Println(
		"Kafka connected successfully",
	)
}
