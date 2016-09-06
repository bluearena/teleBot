package main

import (
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("237745108:AAHgBTGsO-tHdleGMCMuhY8FGVYn-CAL2RQ")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("Got message [%s] %s", update.Message.From.UserName, update.Message.Text)
		_, repl := ProcessMessage(update.Message)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, repl)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
