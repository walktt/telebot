package main

import (
	"fmt"
	"log"
	"net/http"

	"gopkg.in/telegram-bot-api.v4"
)

const (
	WebHookURL = "https://8ff3f72c.ngrok.io"
	BotToken   = "659311466:AAGWVDghpVlLdjZ3hN_BpQSptT4zIoMQ7xU"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(BotToken)
	if err != nil {
		panic(err)
	}
	// bot.Debug = true
	fmt.Printf("Authorized on account %s\n", bot.Self.UserName)
	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebHookURL))
	if err != nil {
		panic(err)
	}
	updates := bot.ListenForWebhook("/")
	go http.ListenAndServe(":8080", nil)
	fmt.Println("start listen :8080")

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

//t.me/eem_prod_test_bot. You can now add a description, about section and profile picture for your bot, see /help for a list of commands. By the way, when you've finished creating your cool bot, ping our Bot Support if you want a better username for it. Just make sure the bot is fully operational before you do this.

//Use this token to access the HTTP API:
//659311466:AAGWVDghpVlLdjZ3hN_BpQSptT4zIoMQ7xU
