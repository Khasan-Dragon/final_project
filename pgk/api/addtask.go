package api

import (
	"encoding/json"
	"final_project/pgk/db"
	"fmt"
	"net/http"
	"time"
)

// checkDate проверяет корректность даты и обрабатывает логику с правилами повторения
func checkDate(task *db.Task) error {
	now := time.Now()

	// Если дата не указана, используем сегодняшнее число
	if task.Date == "" {
		task.Date = now.Format(DateFormat)
		return nil
	}

	// Проверяем корректность формата даты
	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}

	// Если есть правило повторения, проверяем его корректность
	if task.Repeat != "" {
		_, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return err
		}
	}

	// Если дата меньше сегодняшней
	if !afterNow(t, now) {
		if task.Repeat == "" {
			// Если правила повторения нет, берем сегодняшнее число
			task.Date = now.Format(DateFormat)
		} else {
			// Иначе вычисляем следующую дату
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}

	return nil
}

// addTaskHandler обрабатывает POST запросы для добавления задач
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Десериализуем JSON в структуру Task
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeErrorJson(w, "Ошибка десериализации JSON", http.StatusBadRequest)
		return
	}

	// Проверяем обязательное поле title
	if task.Title == "" {
		writeErrorJson(w, "Не указан заголовок задачи", http.StatusBadRequest)
		return
	}

	// Проверяем и обрабатываем дату
	if err := checkDate(&task); err != nil {
		writeErrorJson(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Добавляем задачу в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		writeErrorJson(w, "Ошибка добавления задачи в базу данных", http.StatusInternalServerError)
		return
	}

	// Возвращаем идентификатор созданной задачи
	writeJson(w, map[string]string{"id": fmt.Sprintf("%d", id)})
}
