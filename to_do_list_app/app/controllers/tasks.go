package controllers

import (
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"to_do_list_app/app/helpers"
	"to_do_list_app/app/models"
	"to_do_list_app/app/types"
)

var task types.TaskPayload

// Create - handles the POST request to create a new task.
func Create(w http.ResponseWriter, r *http.Request) {
	err := helpers.ParseRequest(r, &task)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	noteRecord, err := models.CreateTask(&task)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(w, http.StatusInternalServerError, "There is a problem on the server")
		return
	}

	helpers.SendResponse(w, http.StatusCreated, noteRecord)
}

// Index - handles the GET request to retrieve a list of all tasks.
func Index(w http.ResponseWriter, r *http.Request) {
	notes, err := models.GetAllTasks()
	if err != nil {
		log.Printf("failed to get notes from DB %v", err)
		helpers.SendResponse(w, http.StatusInternalServerError, "There is a problem on the server")
		return
	}

	helpers.SendResponse(w, http.StatusOK, notes)
}

// Show -handles the GET request to retrieve a task by its ID.
func Show(w http.ResponseWriter, r *http.Request) {
	str, ok := mux.Vars(r)["taskID"]
	if !ok {
		log.Println("missing task id")
		helpers.SendResponse(w, http.StatusBadRequest, errors.New("task id is required"))
		return
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		log.Println("failed to convert task id to int")
		helpers.SendResponse(w, http.StatusInternalServerError, "There is a problem on the server")
		return
	}

	task, err := models.GetTaskByID(value)
	if err != nil {
		log.Printf("failed to get task from DB %v", err)
		helpers.SendResponse(w, http.StatusNotFound, "task not found")
		return
	}

	helpers.SendResponse(w, http.StatusOK, task)
}

// Update - handles the PUT request to update an existing task by its ID.
func Update(w http.ResponseWriter, r *http.Request) {
	str, ok := mux.Vars(r)["taskID"]
	if !ok {
		log.Println("missing task id")
		helpers.SendResponse(w, http.StatusBadRequest, "task not found")
		return
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		log.Println("failed to convert task id to int")
		helpers.SendResponse(w, http.StatusInternalServerError, "There is a problem on the server")
		return
	}
	task.ID = value
	err = helpers.ParseRequest(r, &task)
	if err != nil {
		log.Printf("failed to parse json: %v", err)
		helpers.SendResponse(w, http.StatusBadRequest, "incorrect data format")
		return
	}

	taskRecord, err := models.UpdateTask(&task)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(w, http.StatusNotFound, "task not found")
		return
	}

	helpers.SendResponse(w, http.StatusOK, taskRecord)
}

// Delete - handles the DELETE request to remove a task by its ID.
func Delete(w http.ResponseWriter, r *http.Request) {
	str, ok := mux.Vars(r)["taskID"]
	if !ok {
		log.Println("missing task id")
		helpers.SendResponse(w, http.StatusBadRequest, "task not found")
		return
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		log.Println("failed to convert task id to int")
		helpers.SendResponse(w, http.StatusInternalServerError, "there is a problem on the server")
		return
	}

	err = models.DeleteTusk(value)
	if err != nil {
		log.Printf("failed to delete task from DB %v", err)
		helpers.SendResponse(w, http.StatusNotFound, "task not found")
		return
	}

	helpers.SendResponse(w, http.StatusNoContent, "task deleted")
}
