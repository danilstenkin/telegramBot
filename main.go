package main

import (
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получаем токен из переменной окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("Telegram bot token not set in environment variables")
	}

	// Создаем нового бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	// Включаем режим логирования
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Создаем обновление для получения новых сообщений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получаем сообщения
	updates := bot.GetUpdatesChan(u)

	// Обрабатываем все сообщения
	for update := range updates {
		// Проверяем, что это сообщение не пустое
		if update.Message == nil {
			continue
		}

		// Логируем входящие сообщения (можно отключить или изменить для отладки)
		log.Printf("Received message: %s", update.Message.Text)

		// Если сообщение содержит "Hello" (не чувствительно к регистру)
		if strings.ToLower(update.Message.Text) == "hello" {
			// Создаем новый ответ
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет!")
			// Отправляем ответ
			if _, err := bot.Send(msg); err != nil {
				log.Println("Error sending message:", err)
			}
		}
	}
}
