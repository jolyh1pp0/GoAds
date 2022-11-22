package controller

import (
	"GoAds/usecase/interfactor"
	"github.com/gorilla/websocket"
	"log"
)

type chatController struct {
	chatInterfactor interfactor.ChatInterfactor
}

type ChatController interface {
	Chat(c Context) error
}

func NewChatController(ci interfactor.ChatInterfactor) ChatController {
	return &chatController{ci}
}

func (cc *chatController) Chat(c Context) error {
	userID := getUserID(c)
	username, err := cc.chatInterfactor.GetUser(userID)
	if err != nil {
		log.Print(err)
		return err
	}

	upgrader := websocket.Upgrader{}
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Print(err)
		return err
	}
	defer ws.Close()

	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello " + username + ", start chatting now!"))
	if err != nil {
		log.Print(err)
		return err
	}

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Print(err)
			return err
		}

		err = ws.WriteMessage(websocket.TextMessage, append([]byte(username+": "), msg...))
		if err != nil {
			log.Print(err)
			return err
		}
	}
}
