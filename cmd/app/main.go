package main

import (
	// "database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"io/ioutil"

	_ "github.com/lib/pq"

	// "github.com/robfig/cron/v3"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	host     string
	port     string
	user     string
	password string
	dbname   string
	botToken string
	pinFilePath = "/app/shared/pin_messages.txt"
)

func init() {
	host = os.Getenv("POSTGRES_HOST")
	port = os.Getenv("POSTGRES_PORT")
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DB")
	botToken = os.Getenv("BOT_TOKEN")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" || botToken == "" {
		log.Fatal("Some environment variables are not set")
	}
}

func savePinnedMessage(chatID int64, messageID int, typePull string) error {
	data := fmt.Sprintf("%d %d %s\n", chatID, messageID, typePull)
	return ioutil.WriteFile(healthcheckFilePath, []byte(data), 0644)
}

func getPinnedMessage(typePull string) (int64, int, error) {
	content, err := ioutil.ReadFile(healthcheckFilePath)
	if err != nil {
		return 0, 0, err
	}

	var chatID int64
	var messageID int
	var savedTypePull string

	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fmt.Sscanf(line, "%d %d %s", &chatID, &messageID, &savedTypePull)
		if savedTypePull == typePull {
			return chatID, messageID, nil
		}
	}

	return 0, 0, fmt.Errorf("no pinned message found for type: %s", typePull)
}

func main() {

	const WHERE_LOCATION string = "where_location"
	const TOWN_LOCATION string = "town_location"
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

    // // Создание экземпляра cron
	// c := cron.New()

	// // Добавление задачи, которая будет выполняться каждую минуту
	// _, err := c.AddFunc("0 1 * * *", func() {
	// 	// log.Println("Задача выполнена:", time.Now())
	// })

	// if err != nil {
	// 	log.Printf("Ошибка при добавлении функции в cron: %sn", err)
	// 	return
	// }
	
	// // Запуск планировщика задач
	// c.Start()

	bot, err := tgbotapi.NewBotAPI(botToken)

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
				msg.Text = "Вы что с калькулятора, сударь!?"
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

					unpinChatID, unpinMessageID := postgresPin(pin, WHERE_LOCATION)

					unpin := tgbotapi.UnpinChatMessageConfig{
						ChatID:    unpinChatID,
						MessageID: unpinMessageID,
					}

					bot.Send(poll)
					bot.Send(unpin)
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

					unpinChatID, unpinMessageID := postgresPin(pin, TOWN_LOCATION)

					unpin := tgbotapi.UnpinChatMessageConfig{
						ChatID:    unpinChatID,
						MessageID: unpinMessageID,
					}

					bot.Send(poll)
					bot.Send(unpin)
					bot.Send(pin)
				}
			}
		}
	}

}
