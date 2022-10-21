package main

import (
	"log"
	"strings"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	locations := []string{
		"Советск",
		"Киров",
		"Йошка",
		"Мск",
		"На Київ",
		"Ушёл в тайгу",
		"Где-то в степях Казахстана",
		"Это секрет, товарищ майор",
		"Другое...",
	}

	answers := []string{
		"?",
		"хто",
		"кто",
	}

	wordsWhere := []string{
		"где",
	}

	wordsLocations := []string{
		"сов",
	}


	token := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true
	// log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			if msg.Text == "+" || msg.Text == "-" {
				msg.Text = "ты чё мля с калькулятора, сцука?"
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
						Type: "regular",
						Options: locations,
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
						Question: msg.Text,
						Type:     "quiz",
						Options: []string{"да", "нет"},
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
