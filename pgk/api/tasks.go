package api

import (
	"final_project/pgk/db"
	"net/http"
)

// TasksResp представляет ответ с списком задач
type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// tasksHandler обрабатывает GET запросы для получения списка задач
func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.Tasks(50) // максимальное количество записей
	if err != nil {
		writeErrorJson(w, "Ошибка получения задач из базы данных", http.StatusInternalServerError)
		return
	}

	writeJson(w, TasksResp{
		Tasks: tasks,
	})
}
