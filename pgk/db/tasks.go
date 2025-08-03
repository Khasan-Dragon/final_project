package db

// Tasks возвращает список задач, отсортированных по дате в порядке возрастания
func Tasks(limit int) ([]*Task, error) {
	// SQL запрос для получения задач с сортировкой по дате
	query := `SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date ASC LIMIT ?`

	rows, err := db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*Task

	// Читаем результаты
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	// Проверяем ошибки итерации
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// Если слайс равен nil, возвращаем пустой слайс
	if tasks == nil {
		tasks = []*Task{}
	}

	return tasks, nil
}
