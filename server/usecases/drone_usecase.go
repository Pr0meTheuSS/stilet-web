package usecase

import (
	"drone_server/models"
	repository "drone_server/repositories"
)

type DroneUsecase interface {
	GetAllDrones() []*models.Drone
	AddNewDrone(drone *models.Drone)
	UpdateDroneInfo(id string, drone *models.Drone)
}

type droneUsecase struct {
	repo repository.DroneRepository
}

func NewDroneUsecase(repo repository.DroneRepository) DroneUsecase {
	return &droneUsecase{repo: repo}
}

func (uc *droneUsecase) GetAllDrones() []*models.Drone {
	return uc.repo.GetDrones()
}

func (uc *droneUsecase) AddNewDrone(drone *models.Drone) {
	uc.repo.AddDrone(drone)
}

func (uc *droneUsecase) UpdateDroneInfo(id string, drone *models.Drone) {
	uc.repo.UpdateDrone(id, drone)
}
