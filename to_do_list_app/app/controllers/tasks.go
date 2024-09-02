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

func Create(w http.ResponseWriter, r *http.Request) {
	err := helpers.ParseRequest(r, &task)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	noteRecord, err := models.CreateTask(&task)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusCreated, noteRecord)
}

func Index(w http.ResponseWriter, r *http.Request) {
	notes, err := models.GetAllTasks()
	if err != nil {
		log.Printf("failed to get notes from DB %v", err)
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, notes)
}

func Show(w http.ResponseWriter, r *http.Request) {
	str, ok := mux.Vars(r)["taskID"]
	if !ok {
		log.Println("missing task id")
		helpers.SendErrorResponse(w, http.StatusBadRequest, errors.New("task id is required"))
		return
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		log.Println("failed to convert task id to int")
		helpers.SendErrorResponse(w, http.StatusInternalServerError, errors.New("invalid task id"))
		return
	}

	task, err := models.GetTaskByID(value)
	if err != nil {
		log.Printf("failed to get task from DB %v", err)
		helpers.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, task)
}

func Update(w http.ResponseWriter, r *http.Request) {
	str, ok := mux.Vars(r)["taskID"]
	if !ok {
		log.Println("missing task id")
		helpers.SendErrorResponse(w, http.StatusBadRequest, errors.New("task id is required"))
		return
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		log.Println("failed to convert task id to int")
		helpers.SendErrorResponse(w, http.StatusInternalServerError, errors.New("invalid task id"))
		return
	}
	task.ID = value
	err = helpers.ParseRequest(r, &task)
	if err != nil {
		log.Printf("failed to parse json: %v", err)
		helpers.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	taskRecord, err := models.UpdateTask(&task)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, taskRecord)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	str, ok := mux.Vars(r)["taskID"]
	if !ok {
		log.Println("missing task id")
		helpers.SendErrorResponse(w, http.StatusBadRequest, errors.New("task id is required"))
		return
	}

	value, err := strconv.Atoi(str)
	if err != nil {
		log.Println("failed to convert task id to int")
		helpers.SendErrorResponse(w, http.StatusInternalServerError, errors.New("invalid task id"))
		return
	}

	err = models.DeleteTusk(value)
	if err != nil {
		log.Printf("failed to delete task from DB %v", err)
		helpers.SendErrorResponse(w, http.StatusNotFound, err)
		return
	}

	helpers.SendResponse(w, http.StatusNoContent, nil)
}
