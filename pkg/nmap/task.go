package nmap

import (
	"github.com/jeotic/nmapper/pkg"
	"log"
)

type Task struct {
	Id        int64
	RunId     int64
	Time      float64
	Task      string
	Type      string
	ExtraInfo JsonNullString
}

func SelectTasks(env pkg.ENV, query string, args ...interface{}) ([]Task, error) {
	rows, err := env.DB.Query(query, args...)
	var tasks []Task

	if err != nil {
		return tasks, err
	}

	for rows.Next() {
		var task Task

		if err := rows.Scan(&task.Id, &task.RunId, &task.Task, &task.Time, &task.Type, &task.ExtraInfo); err != nil {
			log.Println(err.Error())
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTasks(env pkg.ENV, run_id int64) ([]Task, error) {
	return SelectTasks(env, "SELECT * FROM task WHERE run_id = ?", run_id)
}

func GetTask(env pkg.ENV, run_id int64, id int64) (*Task, error) {
	runs, err := SelectTasks(env, "SELECT * FROM task WHERE run_id = ? AND id = ?", run_id, id)

	if err != nil {
		return nil, err
	}

	return &runs[0], nil
}
