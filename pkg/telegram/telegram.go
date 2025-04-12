package telegram

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func StartBot() {
	var err error
	Bot, err = tgbotapi.NewBotAPI("7770574747:AAHbM83zkXo4l407UdAcilXTlkxIaWEHvwc")
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", Bot.Self.UserName)

	// Создание конфигурации для получения обновлений
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 // Таймаут для получения обновлений

	// Канал обновлений
	updates := Bot.GetUpdatesChan(u)

	// Обработка обновлений
	for update := range updates {
		if update.Message != nil {
			handleMessage(update.Message)
		}
	}
}

func handleMessage(message *tgbotapi.Message) {
	if message.Text == "/status" {
		resp, err := http.Get("http://localhost:8080/status")
		if err != nil {
			log.Println("Ошибка получения данных с API:", err)
			sendMessage(message.Chat.ID, "Не удалось получить данные о метриках.")
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Ошибка чтения ответа API:", err)
			sendMessage(message.Chat.ID, "Ошибка при обработке данных.")
			return
		}

		sendMessage(message.Chat.ID, fmt.Sprintf("⚡️ Состояние системы:\n\n%s", string(body)))
	}
}

func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка отправки сообщения в Telegram:", err)
	}
}
