package main

import (
	"flag"

	"log"

	tgClient "link-keeper-bot/clients/telegram"
	"link-keeper-bot/consumer/event-consumer"
	"link-keeper-bot/events/telegram"
	"link-keeper-bot/storage/files"
)

/* ToDo
Подключить бд
Лучше константы убрать и сделать как mustToken()
*/

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
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
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}
	return *token
}
