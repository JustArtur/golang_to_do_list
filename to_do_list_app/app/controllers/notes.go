package controllers

import (
	"log"
	"net/http"
	"to_do_list_app/app/helpers"
	"to_do_list_app/app/models"
	"to_do_list_app/app/types"
)

var note types.TaskPayload

func Create(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := helpers.ParseRequest(r, &note)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	noteRecord, err := models.CreateTask(&note)
	if err != nil {
		log.Println(err)
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusCreated, noteRecord)
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	notes, err := models.GetAllTasks()
	if err != nil {
		log.Printf("failed to get notes from DB %v", err)
		helpers.SendErrorResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, notes)
}

func Show(w http.ResponseWriter, r *http.Request) {

}

func Update(w http.ResponseWriter, r *http.Request) {

}

func Destroy(w http.ResponseWriter, r *http.Request) {

}
