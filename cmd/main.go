package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Получаем токен бота из переменнной окружения
	token := os.Getenv("TOKEN")
	if token == "" {
		panic("TOKEN environment variable is empty")
	}
	fmt.Println("Token:", token)

	// Создаем нового бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}
	// Включает режим отладки. Когда Debug установлен в true, библиотека будет выводить в лог все отправляемые и получаемые от Telegram API запросы и ответы
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Создаем обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	// Цикл обработки обновлений
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Логируем сообщение
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Ответим на сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.FirstName)
		bot.Send(msg)
	}
}
