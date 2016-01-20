package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/hulucat/utils"
	"net/http"
)

const (
	READ_TIMEOUT = 5
)

var WebSocketConn *websocket.Conn

func StartServer() (err error) {
	fmt.Println("Websocket server start now...")
	http.HandleFunc("/face", Handle)
	if err = http.ListenAndServe(":9090", nil); err != nil {
		utils.Errorf("Error serve http: %s", err.Error())
	}
	return
}

func Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(req *http.Request) bool {
			if req.Host == "127.0.0.1:9090" {
				return true
			} else {
				return false
			}
		},
	}
	var err error
	WebSocketConn, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.Errorf("Error upgrade http: %s", err.Error())
		return
	}
	for {
		mt, bytes, err := WebSocketConn.ReadMessage()
		if err != nil {
			utils.Errorf("Error read message: %v", err)
			WebSocketConn.Close()
			WebSocketConn = nil
			return
		}
		utils.Debugf("Client msg received: %d, %s", mt, string(bytes))
	}
}
