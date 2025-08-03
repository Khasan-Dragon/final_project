package main

import (
	dbpkg "final_project/pgk/db"
	"final_project/pgk/server"
	"log"
)

func main() {
	// Инициализация базы данных
	dbFile := "./scheduler.db"
	if err := dbpkg.Init(dbFile); err != nil {
		log.Fatal("Ошибка инициализации базы данных:", err)
	}

	// Запуск веб-сервера
	server.StartServer()
}
