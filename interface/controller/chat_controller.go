package controller

import (
	"GoAds/usecase/interfactor"
	"bytes"
	"github.com/gorilla/websocket"
	"log"
)

type chatController struct {
	userInterfactor interfactor.UserInterfactor
}

type ChatController interface {
	ServeWs(c Context, hub *Hub) error
}

func NewChatController(ui interfactor.UserInterfactor) ChatController {
	return &chatController{ui}
}

const (
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan []byte
	username string
}

func (c *Client) readPump() {
	defer func() {
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		msg := c.username + ": " + string(message)
		c.hub.broadcast <- []byte(msg)
	}
}

func (c *Client) writePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(err)
				return
			}
			w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func (cc chatController) ServeWs(c Context, hub *Hub) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}

	userID := getUserID(c)
	username, err := cc.userInterfactor.GetUser(userID)
	if err != nil {
		log.Print(err)
	}

	client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), username: username}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()

	return nil
}
