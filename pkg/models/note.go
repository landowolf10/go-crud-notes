package models

import (
	"go-crud-notes/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Note struct {
	Noteid  int    `form:"noteid" json:"noteid"`
	Userid  int    `form:"userid" json:"userid"`
	Owner   string `form:"owner" json:"owner"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

type UpdateNote struct {
	Owner   string `form:"owner" json:"owner"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

type CreatedNote struct {
	UserId  int    `form:"userid" json:"userid"`
	Owner   string `form:"owner" json:"owner"`
	Title   string `form:"title" json:"title"`
	Content string `form:"content" json:"content"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Note
}

type SingleResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Note
}

type CreatedResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    CreatedNote
}

type DeletedResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func init() {
	config.Connection()
	db = config.GetDB()
}

func (note *CreatedNote) CreateNote() *CreatedNote {
	db.NewRecord(note)
	db.Create(&note)

	return note
}

func GetNotes() []Note {
	var Notes []Note

	db.Find(&Notes)

	return Notes
}

func GetNoteById(noteId int64) (*Note, *gorm.DB) {
	var getNote Note
	db := db.Where("noteid=?", noteId).Find(&getNote)

	return &getNote, db
}

func DeleteNote(noteId int64) Note {
	var note Note
	db.Where("noteid=?", noteId).Delete(&note)

	return note
}
