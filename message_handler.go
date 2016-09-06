package main

import (
	"citiesBase"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strings"
)

// Classifies message type(command, message, media, etc.) and transfers it to apropriate handler
func ProcessMessage(message tgbotapi.Message) (bool, string) {
	var result bool
	var answer string
	if message.IsCommand() {
		result, answer = processCommand(message)
	}
	return result, answer
}

func processCommand(message tgbotapi.Message) (bool, string) {
	return true, message.Command() + " processed"
}

func processText(message tgbotapi.Message) {
	tok := strings.Split(message.Text, " ")
	if len(tok) != 1 {
		return "Please specify city name"
	}
	req := tok[0]
	log.Printf("Received req: <%s>", req)

	citiesBase.Find
}
