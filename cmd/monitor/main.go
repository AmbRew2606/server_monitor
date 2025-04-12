package main

import (
	"github.com/AmbRew2606/server_monitor/pkg/api"
	"github.com/AmbRew2606/server_monitor/pkg/telegram"
)

func main() {
	// Запускаем сервер API
	go api.StartServer()

	// Запускаем Telegram-бота
	go telegram.StartBot()

	// Чтобы программа не завершалась сразу, оставляем основную горутину "живой"
	select {}
}
