package handlers

import (
	"drone_server/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	usecase  usecase.DroneUsecase
	upgrader websocket.Upgrader
}

func NewWebSocketHandler(usecase usecase.DroneUsecase) *WebSocketHandler {
	return &WebSocketHandler{
		usecase: usecase,
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
	}
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to establish WebSocket connection"})
		return
	}
	defer conn.Close()

	for {
		drones := h.usecase.GetAllDrones()
		if err := conn.WriteJSON(drones); err != nil {
			break
		}
	}
}
