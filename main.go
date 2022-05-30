package main

import (
	"flag"
	"log"
	tgClient "paper_adviser_bot/clients/telegram"
	"paper_adviser_bot/events/telegram"
	"paper_adviser_bot/storage/files"
	"paper_adviser_bot/consumer/event_consumer"
)

const(
	tgBotHost = "api.telegram.org"
	storagePath = "file_storage"
	batchSize = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()), 
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token", 
		"", 
		"token for tg bot access",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}