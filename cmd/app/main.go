package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var host string = os.Getenv("POSTGRES_HOST")
var port string = os.Getenv("POSTGRES_PORT")
var user string = os.Getenv("POSTGRES_USER")
var password string = os.Getenv("POSTGRES_PASSWORD")
var dbname string = os.Getenv("POSTGRES_DB")
var token string = os.Getenv("BOT_TOKEN")

func postgresPin(pin tgbotapi.PinChatMessageConfig, type_pull string) (int64, int) {

	// Коннект
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Пинг
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Достать из бд последний закрепленный
	selectStmt, err := db.Prepare("SELECT * FROM pin_pull WHERE $1 = chat_id AND $2 = type_pull")
	if err != nil {
		panic(err)
	}
	defer selectStmt.Close()

	rows, err := selectStmt.Query(pin.ChatID, type_pull)
	if err != nil {
		panic(err)
	}

	var unpinChatID int64
	var unpinMessageID int

	for rows.Next() {
		var id int
		var createdAt time.Time
		var typePull string

		if err := rows.Scan(&id, &unpinChatID, &unpinMessageID, &createdAt, &typePull); err != nil {
			panic(err)
		}
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	// Подготовка SQL-запросов
	insertStmt, err := db.Prepare("INSERT INTO pin_pull(chat_id, message_id, type_pull, created_at) VALUES($1, $2, $3, $4)")
	if err != nil {
		panic(err)
	}
	defer insertStmt.Close()

	deleteStmt, err := db.Prepare("DELETE FROM pin_pull WHERE chat_id = $1 AND message_id = $2 AND type_pull = $3")
	if err != nil {
		panic(err)
	}
	defer deleteStmt.Close()

	// Удаление данных из таблицы
	_, err = deleteStmt.Exec(unpinChatID, unpinMessageID, type_pull)
	if err != nil {
		panic(err)
	}

	// Текущая дата
	currentDate := time.Now().Format("2006-01-02")

	//  Добавление данных в таблицу
	_, err = insertStmt.Exec(pin.ChatID, pin.MessageID, type_pull, currentDate)
	if err != nil {
		panic(err)
	}

	return unpinChatID, unpinMessageID
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
				msg.Text = "Вы что с калькулятора, сударь!1?"
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
