package db // функции для взаимодействия с базой данных SQLite.

import (
	"database/sql"
	"log"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX idx_scheduler_date ON scheduler(date);
`

func Init(dbFile string) error {
	// Проверяем существование файла базы данных
	install := false
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		install = true
		log.Printf("Файл базы данных не найден: %s. Будет создан новый файл.", dbFile)
	}

	// Открываем базу данных
	var err error
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	// Проверяем соединение
	if err = db.Ping(); err != nil {
		return err
	}

	// Если файл не существовал, создаем таблицы
	if install {
		log.Printf("Создание таблиц в базе данных...")
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
		log.Printf("Таблицы успешно созданы")
	}

	log.Printf("База данных успешно инициализирована: %s", dbFile)
	return nil
}
