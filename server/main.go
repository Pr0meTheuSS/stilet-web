package main

import (
	"drone_server/events"
	"drone_server/handlers"
	repository "drone_server/repositories"
	usecase "drone_server/usecases"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Укажите фронтенд-домен
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Разрешаем методы
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Разрешаем заголовки
		AllowCredentials: true,                                                // Разрешаем передачу cookies
	}))

	eventEmitter := events.NewEventEmitter()
	repo := repository.NewInMemoryDroneRepository()
	usecase := usecase.NewDroneUsecase(repo, eventEmitter)
	droneHandler := handlers.NewDroneHandler(usecase)
	wsHandler := handlers.NewWebSocketHandler(usecase, eventEmitter)

	r.GET("/drones", droneHandler.GetDrones)
	r.POST("/drones", droneHandler.AddDrone)
	r.GET("/ws", wsHandler.HandleWebSocket)

	r.Run(":8080")
}
