package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
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

	taskRecord = new(types.Task)
	for rows.Next() {
		taskRecord, err = scanIntoTask(rows)
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
	query := "INSERT INTO tasks (title, description, due_date) VALUES ($1, $2, $3) RETURNING id, title, description, due_date, created_at, updated_at"
	log.Print("pq: ", query, task.Title, task.Description, task.DueDate)
	rows, err := db.Db.Query(query, task.Title, task.Description, task.DueDate)
	if err != nil {
		return nil, err
	}
	log.Println(task.DueDate)
	taskRecord := new(types.Task)
	for rows.Next() {
		taskRecord, err = scanIntoTask(rows)
		if err != nil {
			return nil, err
		}
	}

	return taskRecord, nil
}

func GetAllTasks() ([]types.Task, error) {
	var tasks []types.Task

	query := "SELECT * FROM tasks"

	log.Print("pq: ", query)
	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := new(types.Task)
		task, err = scanIntoTask(rows)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, *task)
	}

	return tasks, nil
}

func UpdateTask(task *types.TaskPayload) (*types.Task, error) {

	query := `
		UPDATE tasks
		SET title = $2,
			description = $3,
			due_date = $4, 
			updated_at = $5
		WHERE id = $1
		RETURNING id, title, description, due_date, created_at, updated_at;
    `

	log.Printf("pq: %s, %d, %s, %s, %s", query, task.ID, task.Title, task.Description, task.DueDate, time.Now())

	rows, err := db.Db.Query(query, task.ID, task.Title, task.Description, task.DueDate, time.Now())
	if err != nil {
		log.Println("Failed to update task")
		return nil, err
	}

	taskRecord := new(types.Task)
	for rows.Next() {
		taskRecord, err = scanIntoTask(rows)
		if err != nil {
			return nil, err
		}
	}

	return taskRecord, nil
}

func DeleteTusk(ID int) error {
	query := "DELETE FROM tasks WHERE \"id\" = $1"
	log.Print("pq: ", query, ID)
	_, err := db.Db.Exec(query, ID)

	if err != nil {
		return err
	}

	return nil
}

func scanIntoTask(rows *sql.Rows) (*types.Task, error) {
	task := new(types.Task)

	err := rows.Scan(
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

	return task, nil
}
