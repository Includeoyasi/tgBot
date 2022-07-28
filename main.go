package main

import (
	"flag"
	"log"

	"github.com/Includeoyasi/tgbot/clients/telegram"
)

func main() {

	host, token := mustHost(), mustToken()

	tgClient := telegram.New(host, token)

	// fetcher := fetcher.New()

	// processor := processor.New()

	// customer.Start(fetcher, processor)
}

func mustHost() string {
	host := flag.String(
		"host",
		"",
		"host for tg-api requests",
	)

	flag.Parse()

	if *host == "" {
		log.Fatal("tg host not found!")
	}

	return *host
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for auth tg bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("tg auth token not found!")
	}

	return *token
}
