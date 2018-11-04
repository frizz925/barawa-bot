package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	messageHandler "github.com/Frizz925/barawa-bot/handler/message"
	raven "github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	godotenv.Load()

	channelSecret := getenv("CHANNEL_SECRET")
	channelToken := getenv("CHANNEL_TOKEN")
	bot, err := linebot.New(channelSecret, channelToken)

	sentryDsn := getenv("SENTRY_DSN")
	raven.SetDSN(sentryDsn)

	host := getenv("HOST", "")
	port := getenv("PORT", "8080")
	address := fmt.Sprintf("%s:%s", host, port)

	http.HandleFunc("/", createWebhookHandler(bot))
	err = http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}

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

func createWebhookHandler(bot *linebot.Client) func(http.ResponseWriter, *http.Request) {
	return raven.RecoveryHandler(func(w http.ResponseWriter, r *http.Request) {
		events, err := bot.ParseRequest(r)
		if err != nil {
			log.Println(err)
			return
		}

		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeMessage:
				handleMessageEvent(bot, event)
				break
			}
		}
	})
}

func handleMessageEvent(bot *linebot.Client, event *linebot.Event) {
	err := logMessageEvent(event)
	if err != nil {
		log.Println(err)
	}
	message := event.Message
	textMessage, ok := message.(*linebot.TextMessage)
	if !ok {
		return
	}
	response := messageHandler.ProcessMessage(textMessage.Text)
	if len(response) > 0 {
		reply := linebot.NewTextMessage(response)
		logReply(event.Source, "reply", textMessage.ID, reply.Text)
		_, err := bot.ReplyMessage(event.ReplyToken, reply).Do()
		if err != nil {
			log.Panicln(err)
		}
	}
}

func logMessageEvent(event *linebot.Event) error {
	source := event.Source
	message := event.Message
	switch message.(type) {
	case *linebot.TextMessage:
		textMessage := message.(*linebot.TextMessage)
		logMessage(source, "text", textMessage.ID, textMessage.Text)
		break
	case *linebot.ImageMessage:
		imageMessage := message.(*linebot.ImageMessage)
		logMessage(source, "image", imageMessage.ID, imageMessage.OriginalContentURL)
		break
	case *linebot.StickerMessage:
		stickerMessage := message.(*linebot.StickerMessage)
		stickerStr := fmt.Sprintf("%s/%s", stickerMessage.PackageID, stickerMessage.StickerID)
		logMessage(source, "sticker", stickerMessage.ID, stickerStr)
		break
	case *linebot.VideoMessage:
		videoMessage := message.(*linebot.VideoMessage)
		logMessage(source, "video", videoMessage.ID, videoMessage.OriginalContentURL)
		break
	case *linebot.LocationMessage:
		locationMessage := message.(*linebot.LocationMessage)
		locationStr := fmt.Sprintf("%s: %f %f", locationMessage.Title, locationMessage.Latitude, locationMessage.Longitude)
		logMessage(source, "location", locationMessage.ID, locationStr)
		break
	}
	return nil
}

func logMessage(source *linebot.EventSource, messageType string, id string, text string) {
	log.Println(formatMessage(source, "<-", messageType, id, text))
}

func logReply(source *linebot.EventSource, messageType string, id string, text string) {
	log.Println(formatMessage(source, "->", messageType, id, text))
}

func formatMessage(source *linebot.EventSource, prefix string, messageType string, id string, text string) string {
	return fmt.Sprintf("[%s][%s][%s][%s][%s] %s", prefix, source.Type, getSourceID(source), id, messageType, text)
}

func getSourceID(source *linebot.EventSource) string {
	switch source.Type {
	case linebot.EventSourceTypeGroup:
		return source.GroupID
	case linebot.EventSourceTypeRoom:
		return source.RoomID
	case linebot.EventSourceTypeUser:
		return source.UserID
	default:
		return "xxx"
	}
}

func getMessageID(message linebot.Message) string {
	return "xxx"
}
