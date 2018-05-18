package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	messageHandler "github.com/Frizz925/barawa-bot/handler/message"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func getenv(args ...string) string {
	key := args[0]
	value := os.Getenv(key)
	if len(value) == 0 {
		if len(args) >= 2 {
			fallback := args[1]
			return fallback
		}
		panic(fmt.Sprintf("Environment variable \"%s\" is not set", key))
	}
	return value
}

func handleMessageEvent(bot *linebot.Client, event *linebot.Event) {
	message := event.Message
	textMessage, ok := message.(*linebot.TextMessage)
	if !ok {
		return
	}
	response := messageHandler.ProcessMessage(textMessage.Text)
	if len(response) > 0 {
		reply := linebot.NewTextMessage(response)
		_, err := bot.ReplyMessage(event.ReplyToken, reply).Do()
		if err != nil {
			log.Println(err)
		}
	}
}

func createWebhookHandler(bot *linebot.Client) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		events, err := bot.ParseRequest(r)
		if err != nil {
			log.Println(err)
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				handleMessageEvent(bot, &event)
				break
			}
		}
	}
}

func main() {
	godotenv.Load()

	channelSecret := getenv("CHANNEL_SECRET")
	channelToken := getenv("CHANNEL_TOKEN")
	bot, err := linebot.New(channelSecret, channelToken)

	host := getenv("HOST", "")
	port := getenv("PORT", "8080")
	address := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/", createWebhookHandler(bot))
	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
