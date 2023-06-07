package routes

import (
	"go-crud-notes/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterNoteRoutes = func(router *mux.Router) {
	router.HandleFunc("/note/", controllers.CreateNote).Methods("POST")
	router.HandleFunc("/note/", controllers.GetNotes).Methods("GET")
	router.HandleFunc("/note/{noteId}", controllers.GetNoteById).Methods("GET")
	//router.HandleFunc("/note/{noteId}", controllers.UpdateNote).Methods("PUT")
	//router.HandleFunc("/note/{noteId}", controllers.DeleteNote).Methods("DELETE")
}
