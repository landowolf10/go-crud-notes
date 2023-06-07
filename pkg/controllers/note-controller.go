package controllers

import (
	"encoding/json"
	"fmt"
	"go-crud-notes/pkg/models"
	"go-crud-notes/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var newNote models.Note

/*
var result models.CreatedNote
var singleResult models.Note
var note models.Note
*/
var noteId string

func GetNotes(w http.ResponseWriter, r *http.Request) {
	noteList := models.GetNotes()
	//res, _ := json.Marshal(noteList)
	var response models.Response

	response.Status = 200
	response.Message = "Success!"
	response.Data = noteList

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	//w.Write(res)
}

func GetNoteById(w http.ResponseWriter, r *http.Request) {
	noteId = mux.Vars(r)["noteid"]

	ID, err := strconv.ParseInt(noteId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing")
	}

	noteDetails, _ := models.GetNoteById(ID)

	res, _ := json.Marshal(noteDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNote(w http.ResponseWriter, r *http.Request) {
	noteBody := &models.CreatedNote{}
	utils.ParseBody(r, noteBody)
	note := noteBody.CreateNote()
	res, _ := json.Marshal(note)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
