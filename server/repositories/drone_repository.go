package repository

import (
	"drone_server/models"
	"fmt"
)

type DroneRepository interface {
	GetAll() ([]models.Drone, error)
	Add(drone models.Drone)
	Update(drone models.Drone) error
	Delete(id int)
}

type InMemoryDroneRepository struct {
	drones map[int]models.Drone
}

func NewInMemoryDroneRepository() *InMemoryDroneRepository {
	return &InMemoryDroneRepository{
		drones: make(map[int]models.Drone),
	}
}

func (repo *InMemoryDroneRepository) GetAll() ([]models.Drone, error) {
	fmt.Println("GetAll called")

	result := []models.Drone{
		{
			ID:        1,
			Name:      "Drone A",
			IsActive:  true,
			Status:    []string{"Connected"},
			Latitude:  55.55,
			Longitude: 83.55,
		},
		{
			ID:        2,
			Name:      "Drone B",
			IsActive:  false,
			Status:    []string{"Connected"},
			Latitude:  54.55,
			Longitude: 82.55,
		},
	}

	// for _, drone := range repo.drones {
	// 	result = append(result, drone)
	// }
	return result, nil
}

func (repo *InMemoryDroneRepository) Add(drone models.Drone) {
	repo.drones[drone.ID] = drone
}

func (repo *InMemoryDroneRepository) Update(drone models.Drone) error {
	if _, exists := repo.drones[drone.ID]; exists {
		repo.drones[drone.ID] = drone
	}
	return nil
}

func (repo *InMemoryDroneRepository) Delete(id int) {
	delete(repo.drones, id)
}
