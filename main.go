package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Verkury/ScheduleBot/client"
	"github.com/Verkury/ScheduleBot/server"
)

func main() {
	mode := flag.String("mode", "", "Запуск в режиме 'client' или 'server'")
	flag.Parse()

	userInput := strings.ToLower(*mode)

	if userInput == "" {
		fmt.Println("=== ScheduleBOT Install Wizard ===")
		fmt.Print("Выберите режим установки (c - client / s - server): ")
		fmt.Scanln(&userInput)
		userInput = strings.ToLower(userInput)
	}

	switch userInput {
	case "s", "server":
		fmt.Println("Запуск сервера...")
		server.StartServer()
	case "c", "client":
		fmt.Println("Запуск клиента...")
		client.StartClient()
	default:
		fmt.Fprintln(os.Stderr, "Ошибка: Некорректный режим. Используйте 'client' или 'server'.")
		os.Exit(1)
	}
}