package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocketConnection struct {
	*websocket.Conn
}

type WsPayload struct {
	Action      string              `json:"action"`
	Message     string              `json:"message"`
	UserName    string              `json:"username"`
	MessageType string              `json:"message_type"`
	UserID      int                 `json:"user_id"`
	Conn        WebSocketConnection `json:"-"`
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WsJsonResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	UserID  int    `json:"user_id"`
}

// clients keeps track of every client who is connected
var clients = make(map[WebSocketConnection]string)

var wsChan = make(chan WsPayload)

func (app *application) WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	app.infoLog.Println(fmt.Sprintf("Client connected from %s", r.RemoteAddr))

	var response WsJsonResponse
	response.Message = "Connected to server"

	err = ws.WriteJSON(response)
	if err != nil {
		app.errorLog.Println(err)
		return
	}

	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""

	go app.ListenForWS(&conn)
}

func (app *application) ListenForWS(conn *WebSocketConnection) {

	/* We want to make sure, if sth goes wrong, this application doesn't die and we recover gracefully. */
	defer func() {
		if r := recover(); r != nil {
			app.errorLog.Println("ERROR:", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func (app *application) ListenToWsChannel() {
	var response WsJsonResponse

	for {
		// e stands for event
		e := <-wsChan

		switch e.Action {
		case "deleteUser":
			// specify what the end user should do
			response.Action = "logout"

			response.Message = "Your account has been deleted"
			response.UserID = e.UserID

			// broadcast this to all connected clients
			app.broadcastToAll(response)

		// this is just to make things clear!
		default:
		}
	}
}

func (app *application) broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		// broadcast to every connected client
		err := client.WriteJSON(response)

		/* If there's an error here, in the vast majority of cases it means this client has disconnected. */
		if err != nil {
			app.errorLog.Printf("Websocket err on %s: %s", response.Action, err)
			_ = client.Close()

			delete(clients, client)
		}
	}
}
