package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gstpsk/skessage/models"
)

var chatRooms []models.ChatRoom

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/room/new", createNewRoom)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewRoom(w http.ResponseWriter, r *http.Request) {
	room := models.ChatRoom{
		Id:   "testjen",
		Name: "coole kamer",
	}

	chatRooms = append(chatRooms, room)

	http.Redirect(w, r, "/room/"+room.Id, http.StatusFound)
}
