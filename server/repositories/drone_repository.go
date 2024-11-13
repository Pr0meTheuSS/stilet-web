package repository

import "drone_server/models"

type DroneRepository interface {
	GetDrones() []*models.Drone
	AddDrone(drone *models.Drone)
	UpdateDrone(id string, drone *models.Drone)
}

type InMemoryDroneRepository struct {
	drones map[string]*models.Drone
}

func NewInMemoryDroneRepository() *InMemoryDroneRepository {
	return &InMemoryDroneRepository{
		drones: make(map[string]*models.Drone),
	}
}

func (repo *InMemoryDroneRepository) GetDrones() []*models.Drone {
	var result []*models.Drone
	for _, drone := range repo.drones {
		result = append(result, drone)
	}
	return result
}

func (repo *InMemoryDroneRepository) AddDrone(drone *models.Drone) {
	repo.drones[drone.ID] = drone
}

func (repo *InMemoryDroneRepository) UpdateDrone(id string, drone *models.Drone) {
	if _, exists := repo.drones[id]; exists {
		repo.drones[id] = drone
	}
}
