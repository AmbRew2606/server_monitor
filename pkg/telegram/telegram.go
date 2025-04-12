package telegram

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

type Metrics struct {
	CPUUsage float64 `json:"cpu_usage"`
	RAMUsage float64 `json:"ram_usage"`
}

func StartBot() {
	var err error
	Bot, err = tgbotapi.NewBotAPI("")
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
	if message.Text == "/start" {
		// Отправка клавиатуры
		sendKeyboard(message.Chat.ID)
	}

	// /status
	if message.Text == "Получить статус системы" {
		// Данные с API
		resp, err := http.Get("http://89.104.70.198:8080/status")
		if err != nil {
			log.Println("Ошибка получения данных с API:", err)
			sendMessage(message.Chat.ID, "Не удалось получить данные о метриках.")
			return
		}
		defer resp.Body.Close()

		// Чтение и демаршализация JSON
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Ошибка чтения ответа API:", err)
			sendMessage(message.Chat.ID, "Ошибка при обработке данных.")
			return
		}

		var metrics Metrics
		err = json.Unmarshal(body, &metrics)
		if err != nil {
			log.Println("Ошибка декодирования JSON:", err)
			sendMessage(message.Chat.ID, "Ошибка при обработке данных.")
			return
		}

		// Формирование ответа для пользователя
		messageText := fmt.Sprintf("⚡️ Состояние системы:\n\nCPU: %.2f%%\nRAM: %.2f%%", metrics.CPUUsage, metrics.RAMUsage)
		sendMessage(message.Chat.ID, messageText)
	}
}

// Функция для отправки сообщений в Telegram
func sendMessage(chatID int64, text string) {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка отправки сообщения в Telegram:", err)
	}
}

// Функция для отправки клавиатуры
func sendKeyboard(chatID int64) {
	// Создание кнопки на клавиатуре
	button := tgbotapi.NewKeyboardButton("Получить статус системы")

	// Создание клавиатуры
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(button),
	)

	// Отправка клавиатуры
	msg := tgbotapi.NewMessage(chatID, "Привет! Я сервер Егора! Выбери опцию, которая есть ниже в клвиатуре!")
	msg.ReplyMarkup = keyboard
	_, err := Bot.Send(msg)
	if err != nil {
		log.Println("Ошибка отправки клавиатуры в Telegram:", err)
	}
}
