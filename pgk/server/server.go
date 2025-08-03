package server //функция для запуска сервера

import (
	"final_project/pgk/api"
	"log"
	"net/http"
)

func StartServer() {
	// Инициализация API обработчиков
	api.Init()

	// Директория с веб-файлами
	webDir := "./web"

	// Настройка обработчика
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	// Запуск сервера на порту 7540
	port := ":7540"
	log.Printf("Сервер запущен на http://localhost%s", port)
	log.Printf("Веб-файлы обслуживаются из директории: %s", webDir)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}
