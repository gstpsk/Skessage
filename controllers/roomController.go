package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gstpsk/skessage/models"
	"github.com/gstpsk/skessage/util"
)

var chatRooms []models.ChatRoom

func CreateNewRoom(w http.ResponseWriter, r *http.Request) {
	room := models.ChatRoom{
		Id:   "testjen",
		Name: "coole kamer",
	}

	chatRooms = append(chatRooms, room)

	log.Printf("Created new room: %s", room.Name)

	http.Redirect(w, r, "/room/"+room.Id, http.StatusFound)
}

// Returns a room with the given id, returns an empty room object if not found.
func GetRoomById(id string) models.ChatRoom {
	for _, room := range chatRooms {
		if room.Id == id {
			return room
		}
	}

	return models.ChatRoom{}
}

func SendMessage(w http.ResponseWriter, r *http.Request) {
	//username := r.Context().Value("username").(string)
	msg := &models.Message{}

	r.ParseForm()
	roomId := r.PostFormValue("room-id")
	room := GetRoomById(roomId)

	msg.Id = util.RandString(16)
	msg.Body = r.PostFormValue("body")
	msg.Time = time.Now()
	msg.Username = r.PostFormValue("username")

	room.Messages = append(room.Messages, *msg)
	log.Printf("Message from %s: %s\n", msg.Username, msg.Body)
	util.StatusCodeResponse(w, 200)
}

func GetMessagesForRoom(w http.ResponseWriter, r *http.Request) {
	//r.Parse
}
