package main

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}

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
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	}
}
