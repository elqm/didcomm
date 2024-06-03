package controllers

import (
	"didcomm_web/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocket struct {
	wsWeb  *websocket.Conn
	wsBack *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var config = utils.InitConfig()
var wsUrl = "ws://" + config.BackServerURL + config.WebSocketEndPoint

func NewWebSocket() *WebSocket {
	ws := &WebSocket{
		wsWeb:  &websocket.Conn{},
		wsBack: &websocket.Conn{},
	}
	return ws
}

// 페이지 웹소켓 연결
func (ws *WebSocket) HandleWebSocket2(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("upgrade err: %v\n", err)
	}
	defer conn.Close()

	ws.wsWeb = conn

	// 서버 웹소켓 연결
	ws.ConnectWebSocketServer()

	// 메시지 수신
	// 고루틴을 사용하니 메시지를 받아오지 못함? 잘 모르겠음. 일단 for문만 사용하는 걸로...
	// go func() {
	for {
		_, msg, err := ws.wsWeb.ReadMessage()
		if err != nil {
			log.Printf("readMessage err: %v\n", err)
			break
		} else {
			fmt.Printf("클라이언트로부터 수신한 메시지: %s\n", msg)

			ws.SendMessageToBack(string(msg))
		}
	}
	// }()
}

func (ws *WebSocket) SendMessageToWeb(msg string) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("json err: %v\n", err)
	}

	err = ws.wsWeb.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		log.Printf("writeMessage err: %v\n", err)
	}
	fmt.Println("클라이언트에게 메시지 전송 성공.")
	return nil
}

func (ws *WebSocket) ConnectWebSocketServer() {
	client, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		log.Printf("dial err: %v\n", err)
	}
	defer client.Close()

	ws.wsBack = client

	// 메시지 수신
	// go func() {
	for {
		_, msg, err := ws.wsBack.ReadMessage()
		if err != nil {
			log.Printf("readMessage err: %v\n", err)
			break
		} else {
			fmt.Printf("백엔드 서버로부터 수신한 메시지: %s\n", msg)

			ws.SendMessageToWeb(string(msg))
		}
	}
	// }()
}

// 백엔드 서버로 메시지 전송
func (ws *WebSocket) SendMessageToBack(msg string) error {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		log.Printf("json err: %v\n", err)
	}

	err = ws.wsBack.WriteMessage(websocket.TextMessage, jsonMsg)
	if err != nil {
		log.Printf("writeMessage err: %v\n", err)
	}
	fmt.Println("백엔드 서버로 메시지 전송 성공.")
	return nil
}
