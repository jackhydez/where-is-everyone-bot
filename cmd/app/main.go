package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var host string = os.Getenv("POSTGRES_HOST")
var port string = os.Getenv("POSTGRES_PORT")
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = os.Getenv("POSTGRES_DB")

func main() {

	const ONE_MIN = 60

	locations := []string{
		"Советск",
		"Киров",
		"Йошка",
		"Мск",
		"Другое",
	}

	answers := []string{
		"?",
		"кто",
	}

	wordsWhere := []string{
		"где",
	}

	wordsLocations := []string{
		"сов",
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	token := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true
	// log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = ONE_MIN

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			if msg.Text == "+" || msg.Text == "-" {
				msg.Text = "вы что с калькулятора сударь!?"
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}

			text := strings.ToLower(msg.Text)
			isContain := false
			countAnswerWorld := 0

			for item := range answers {
				isContain = false
				isContain = strings.Contains(text, answers[item])
				if isContain {
					countAnswerWorld++
				}
			}

			if countAnswerWorld >= 2 {
				if strings.Contains(text, wordsWhere[0]) {
					poll := tgbotapi.SendPollConfig{
						BaseChat: tgbotapi.BaseChat{
							ChatID: update.Message.Chat.ID,
						},
						Question: msg.Text,
						Type:     "regular",
						Options:  locations,
					}

					pin := tgbotapi.PinChatMessageConfig{
						ChatID:              update.Message.Chat.ID,
						MessageID:           update.Message.MessageID + 1,
						DisableNotification: false,
					}

					bot.Send(poll)
					bot.Send(pin)
				}

				if strings.Contains(text, wordsLocations[0]) {
					poll := tgbotapi.SendPollConfig{
						BaseChat: tgbotapi.BaseChat{
							ChatID: update.Message.Chat.ID,
						},
						Question:        msg.Text,
						Type:            "regular",
						Options:         []string{"да", "нет", "пукнуть"},
						CorrectOptionID: 0,
					}

					pin := tgbotapi.PinChatMessageConfig{
						ChatID:              update.Message.Chat.ID,
						MessageID:           update.Message.MessageID + 1,
						DisableNotification: false,
					}

					bot.Send(poll)
					bot.Send(pin)
				}
			}
		}
	}

}
