package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"drone_server/events"
	"drone_server/usecases"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	usecase      usecases.DroneUsecase
	eventEmitter *events.EventEmitter
	broadcast    chan []byte
	clients      map[*websocket.Conn]bool
}

func NewWebSocketHandler(uc usecases.DroneUsecase, emitter *events.EventEmitter) *WebSocketHandler {
	handler := &WebSocketHandler{
		usecase:      uc,
		eventEmitter: emitter,
		broadcast:    make(chan []byte),
		clients:      make(map[*websocket.Conn]bool),
	}

	go handler.listenToDroneUpdates()

	return handler
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	h.clients[conn] = true

	// Удаление клиента при отключении
	defer func() {
		delete(h.clients, conn)
	}()

	// Отправка всех текущих дронов подключённому клиенту
	h.sendAllDronesToClient(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received message: %s\n", message)
	}
}

func (h *WebSocketHandler) listenToDroneUpdates() {
	updateChannel := make(chan interface{})
	h.eventEmitter.Subscribe("drone_updated", updateChannel)

	for update := range updateChannel {
		message, err := json.Marshal(map[string]interface{}{
			"action": "drone_updated",
			"data":   update,
		})
		if err != nil {
			log.Println("Failed to marshal update:", err)
			continue
		}

		h.broadcastToClients(message)
	}
}

func (h *WebSocketHandler) broadcastToClients(message []byte) {
	for client := range h.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("Failed to send message to client:", err)
			client.Close()
			delete(h.clients, client)
		}
	}
}

// Отправка всех текущих дронов подключённому клиенту
func (h *WebSocketHandler) sendAllDronesToClient(conn *websocket.Conn) {
	drones, err := h.usecase.GetDrones()
	if err != nil {
		log.Println("Failed to fetch drones:", err)
		conn.WriteJSON(map[string]interface{}{
			"error": "Failed to fetch drones",
		})
		return
	}

	message, err := json.Marshal(map[string]interface{}{
		"action": "get_all_drones",
		"data":   drones,
	})
	if err != nil {
		log.Println("Failed to marshal drones:", err)
		return
	}

	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("Failed to send drones to client:", err)
	}
}
