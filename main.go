package main

import (
	"log"
	"os"
	"strings"
	"telegram-interlocutor-bot/core"

	"telegram-interlocutor-bot/entity"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text := findPhrase(update.Message.Text, core.GetRules())

		if len(text) > 1 {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
		}
	}
}

func findPhrase(message string, phrases []entity.Phrase) string {
	for i := range phrases {
		if strings.Contains(strings.ToLower(message), phrases[i].Word) {
			return phrases[i].Answer
		}
	}

	return ""
}
