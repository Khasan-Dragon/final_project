package api

import (
	"encoding/json"
	"final_project/pgk/db"
	"fmt"
	"net/http"
	"time"
)

// writeJson сериализует data в JSON и отправляет ответ
func writeJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(data)
}

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
		writeJson(w, map[string]string{"error": "Ошибка десериализации JSON"})
		return
	}

	// Проверяем обязательное поле title
	if task.Title == "" {
		writeJson(w, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	// Проверяем и обрабатываем дату
	if err := checkDate(&task); err != nil {
		writeJson(w, map[string]string{"error": err.Error()})
		return
	}

	// Добавляем задачу в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, map[string]string{"error": "Ошибка добавления задачи в базу данных"})
		return
	}

	// Возвращаем идентификатор созданной задачи
	writeJson(w, map[string]string{"id": fmt.Sprintf("%d", id)})
}
