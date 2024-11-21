package handlers

import (
	"drone_server/models"
	usecase "drone_server/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DroneHandler struct {
	usecase usecase.DroneUsecase
}

func NewDroneHandler(usecase usecase.DroneUsecase) *DroneHandler {
	return &DroneHandler{usecase: usecase}
}

func (h *DroneHandler) GetDrones(c *gin.Context) {
	drones, err := h.usecase.GetDrones()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, drones)
}

func (h *DroneHandler) AddDrone(c *gin.Context) {
	var drone models.Drone
	if err := c.ShouldBindJSON(&drone); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.AddDrone(drone); err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusCreated, drone)
}
