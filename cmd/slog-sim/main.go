package main

import (
	"log/slog"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

func main() {
	initLogs()
	for {
		randomMessage()
		time.Sleep(3 * time.Second)
	}
}

func randomMessage() {
	message := faker.Paragraph(options.WithRandomStringLength(40))
	name := faker.Name()
	domain := faker.DomainName()
	slog.Info(message, "name", name, "domain", domain)
}
