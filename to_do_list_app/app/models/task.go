package models

import (
	"fmt"
	"log"
	"to_do_list_app/app/types"
	"to_do_list_app/db"
)

func GetTaskByID(ID int) (*types.Task, error) {
	taskRecord := new(types.Task)

	query := "SELECT * FROM tasks WHERE \"id\" = $1"

	log.Print("pq: ", query, ID)
	rows, err := db.Db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&taskRecord.ID,
			&taskRecord.Title,
			&taskRecord.Description,
			&taskRecord.DueDate,
			&taskRecord.CreatedAt,
			&taskRecord.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	if taskRecord.ID == 0 {
		return nil, fmt.Errorf("task not found")
	}

	return taskRecord, nil
}

func CreateTask(task *types.TaskPayload) (*types.Task, error) {
	taskRecord := new(types.Task)

	query := "INSERT INTO tasks (title, description, due_date) VALUES ($1, $2, $3) RETURNING id, title, description, due_date, created_at, updated_at"
	log.Print("pq: ", query, task.Title, task.Description, task.DueDate)
	rows, err := db.Db.Query(query, task.Title, task.Description, task.DueDate)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&taskRecord.ID,
			&taskRecord.Title,
			&taskRecord.Description,
			&taskRecord.DueDate,
			&taskRecord.CreatedAt,
			&taskRecord.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return taskRecord, nil
}

func GetAllTasks() ([]types.Task, error) {
	var notes []types.Task

	query := "SELECT * FROM tasks"

	log.Print("pq: ", query)
	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task types.Task
		err = rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.DueDate,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		notes = append(notes, task)
	}

	return notes, nil
}
