package handlers

import (
    "net/http"
    "drone_server/models"
    "drone_server/usecases"

    "github.com/gin-gonic/gin"
)

type DroneHandler struct {
    usecase usecase.DroneUsecase
}

func NewDroneHandler(usecase usecase.DroneUsecase) *DroneHandler {
    return &DroneHandler{usecase: usecase}
}

func (h *DroneHandler) GetDrones(c *gin.Context) {
    drones := h.usecase.GetAllDrones()
    c.JSON(http.StatusOK, drones)
}

func (h *DroneHandler) AddDrone(c *gin.Context) {
    var drone models.Drone
    if err := c.ShouldBindJSON(&drone); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    h.usecase.AddNewDrone(&drone)
    c.JSON(http.StatusCreated, drone)
}
