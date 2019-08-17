package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// using websocket to connect to NodeMCU

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Handler a websocket handler
func Handler(w http.ResponseWriter, r *http.Request) {

	wsUpgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	reader(conn)
}

func reader(conn *websocket.Conn) {
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}
