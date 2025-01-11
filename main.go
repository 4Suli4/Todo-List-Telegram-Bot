package main

import (
	"context"
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"todo_list_telegram/db"
	"todo_list_telegram/model"
	"todo_list_telegram/postgresql"
)

func GetTelegramToken() (string, error) {
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		return "", errors.New("TELEGRAM_TOKEN environment variable not set")
	}
	return token, nil
}

func main() {

	errLoadEnv := LoadEnv()
	if errLoadEnv != nil {
		log.Fatal("Error loading .env file")
	}

	pgClient, errConnectPg := postgresql.NewPostgresClient()
	if errConnectPg != nil {
		log.Fatal(errConnectPg)
	}

	db := db.NewTodoDb(pgClient)

	errConnectPg = pgClient.Ping(context.Background())
	if errConnectPg != nil {
		log.Fatal(errConnectPg)
	}

	telegramToken, errGetToken := GetTelegramToken()
	if errGetToken != nil {
		log.Fatal(errGetToken)
	}

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.Text == "/create" {
			err := db.AddTodo(context.Background(), model.Todo{
				UUID:  "",
				Title: "from telegram bot",
				Done:  1,
			})
			if err != nil {
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, err.Error()))
				return
			}
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
