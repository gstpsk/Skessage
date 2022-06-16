package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/gstpsk/skessage/controllers"
)

var PORT string = "8000"

func main() {
	handleRequests()
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/room/new", controllers.CreateNewRoom)
	r.HandleFunc("/messages/send", controllers.SendMessage)

	log.Printf("HTTP server running on http://localhost:%s", PORT)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + PORT,
		WriteTimeout: 20 * time.Second,
		ReadTimeout:  20 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homePage")
	http.ServeFile(w, r, "templates/index.html")
}
