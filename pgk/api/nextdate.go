package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Константа для формата даты
const DateFormat = "20060102"

// NextDate вычисляет следующую дату для задачи в соответствии с правилом повторения
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Проверка на пустое правило повторения
	if repeat == "" {
		return "", fmt.Errorf("пустое правило повторения")
	}

	// Парсинг исходной даты
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("некорректная дата: %v", err)
	}

	// Разбор правила повторения
	parts := strings.Split(strings.TrimSpace(repeat), " ")

	// Обработка правила "y" (ежегодно)
	if len(parts) == 1 && parts[0] == "y" {
		return calculateNextYearlyDate(date, now), nil
	}

	// Обработка правила "d <число>" (дни)
	if len(parts) == 2 && parts[0] == "d" {
		return calculateNextDailyDate(date, now, parts[1])
	}

	// Проверка на неподдерживаемые форматы
	if len(parts) >= 1 {
		switch parts[0] {
		case "w", "m":
			return "", fmt.Errorf("неподдерживаемый формат")
		}
	}

	// Неверный формат правила
	return "", fmt.Errorf("неверный формат правила повторения")
}

// afterNow проверяет, больше ли первая дата второй (игнорируя время)
func afterNow(date, now time.Time) bool {
	// Нормализуем даты до начала дня для корректного сравнения
	dateNormalized := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowNormalized := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return dateNormalized.After(nowNormalized)
}

// calculateNextYearlyDate вычисляет следующую ежегодную дату
func calculateNextYearlyDate(date, now time.Time) string {
	// Увеличиваем дату на год до тех пор, пока она не станет больше now
	for {
		date = date.AddDate(1, 0, 0)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format(DateFormat)
}

// calculateNextDailyDate вычисляет следующую дату для правила дней
func calculateNextDailyDate(date, now time.Time, intervalStr string) (string, error) {
	// Парсинг интервала
	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		return "", fmt.Errorf("некорректный интервал дней: %v", err)
	}

	// Проверка максимального допустимого интервала
	if interval <= 0 || interval > 400 {
		return "", fmt.Errorf("интервал должен быть от 1 до 400 дней")
	}

	// Увеличиваем дату на указанное количество дней до тех пор, пока она не станет больше now
	for {
		date = date.AddDate(0, 0, interval)
		if afterNow(date, now) {
			break
		}
	}

	return date.Format(DateFormat), nil
}

// NextDayHandler обрабатывает запросы для вычисления следующей даты
func NextDayHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовок для JSON ответа
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Получаем параметры из запроса
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	// Если параметр now не указан, используем текущую дату
	var now time.Time
	if nowStr == "" {
		now = time.Now()
	} else {
		var err error
		now, err = time.Parse(DateFormat, nowStr)
		if err != nil {
			http.Error(w, fmt.Sprintf("{\"error\": \"некорректная дата now: %v\"}", err), http.StatusBadRequest)
			return
		}
	}

	// Проверяем обязательные параметры
	if dateStr == "" {
		http.Error(w, "{\"error\": \"параметр date обязателен\"}", http.StatusBadRequest)
		return
	}

	if repeatStr == "" {
		http.Error(w, "{\"error\": \"параметр repeat обязателен\"}", http.StatusBadRequest)
		return
	}

	// Вызываем функцию NextDate
	nextDate, err := NextDate(now, dateStr, repeatStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("{\"error\": \"%v\"}", err), http.StatusBadRequest)
		return
	}

	// Возвращаем результат
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(nextDate))
}
