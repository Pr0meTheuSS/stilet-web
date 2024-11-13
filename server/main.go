package main

import (
	"drone_server/handlers"
	repository "drone_server/repositories"
	usecase "drone_server/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	repo := repository.NewInMemoryDroneRepository()
	usecase := usecase.NewDroneUsecase(repo)
	droneHandler := handlers.NewDroneHandler(usecase)
	wsHandler := handlers.NewWebSocketHandler(usecase)

	r.GET("/drones", droneHandler.GetDrones)
	r.POST("/drones", droneHandler.AddDrone)
	r.GET("/ws", wsHandler.HandleWebSocket)

	r.Run(":8080")
}
