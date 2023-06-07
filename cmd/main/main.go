package main

import (
	"go-crud-notes/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterNoteRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
