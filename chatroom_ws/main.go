package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type ChatRoom struct {
	RoomId   string
	Username string
	Ws       *websocket.Conn
}

// Global map to store rooms and their connections
var rooms = make(map[string][]*ChatRoom)
var roomsMutex sync.RWMutex

func wsChatRoom(c echo.Context) error {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	roomId := c.Param("roomId")
	username := c.Param("username")

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	chatRoom := &ChatRoom{
		RoomId:   roomId,
		Username: username,
		Ws:       ws,
	}

	// Add user to room
	roomsMutex.Lock()
	rooms[roomId] = append(rooms[roomId], chatRoom)
	roomsMutex.Unlock()

	// Handle messages
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message from client", err)
			// Remove user from room when they disconnect
			removeUserFromRoom(roomId, chatRoom)
			break
		}
		log.Println(messageType, string(message))

		// Broadcast message to other users in the same room
		broadcastMessage(roomId, chatRoom, messageType, message)
	}

	return nil
}

func broadcastMessage(roomId string, sender *ChatRoom, messageType int, message []byte) {
	roomsMutex.RLock()
	roomUsers := rooms[roomId]
	roomsMutex.RUnlock()

	response := []byte(fmt.Sprintf("%s: %s", sender.Username, string(message)))

	for _, user := range roomUsers {
		if user != sender { // Don't send to the sender
			err := user.Ws.WriteMessage(messageType, response)
			if err != nil {
				log.Println("Error broadcasting message to user", err)
			}
		}
	}
}

func removeUserFromRoom(roomId string, userToRemove *ChatRoom) {
	roomsMutex.Lock()
	defer roomsMutex.Unlock()

	roomUsers := rooms[roomId]
	for i, user := range roomUsers {
		if user == userToRemove {
			rooms[roomId] = append(roomUsers[:i], roomUsers[i+1:]...)
			break
		}
	}
}

func main() {
	e := echo.New()
	e.GET("/ws/chat/:roomId/user/:username", wsChatRoom)
	e.Logger.Fatal(e.Start(":8080"))
}
