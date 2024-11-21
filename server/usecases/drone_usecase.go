package usecases

import (
	"drone_server/events"
	"drone_server/models"
	repository "drone_server/repositories"
	"math"
	"sync"
	"time"
)

type DroneUsecase interface {
	GetDrones() ([]models.Drone, error)
	AddDrone(drone models.Drone) error
	UpdateDrone(drone models.Drone) error
	DeleteDrone(id int) error
}

type droneUsecaseImpl struct {
	repo         repository.DroneRepository
	eventEmitter *events.EventEmitter
	mu           sync.Mutex
	stopChan     chan struct{}
}

func NewDroneUsecase(repo repository.DroneRepository, emitter *events.EventEmitter) DroneUsecase {
	uc := &droneUsecaseImpl{
		repo:         repo,
		eventEmitter: emitter,
		stopChan:     make(chan struct{}),
	}

	go uc.startDroneMovement()
	return uc
}

func (uc *droneUsecaseImpl) AddDrone(drone models.Drone) error {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	uc.repo.Add(drone)
	return nil
}

func (u *droneUsecaseImpl) UpdateDrone(drone models.Drone) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if err := u.repo.Update(drone); err != nil {
		return err
	}
	// Генерация события при обновлении дрона
	u.eventEmitter.Emit("drone_updated", drone)
	return nil
}

func (u *droneUsecaseImpl) GetDrones() ([]models.Drone, error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.repo.GetAll()
}

func (uc *droneUsecaseImpl) DeleteDrone(id int) error {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	uc.repo.Delete(id)
	return nil
}

// Симуляция движения дронов
func (uc *droneUsecaseImpl) startDroneMovement() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			uc.moveDrones()
		case <-uc.stopChan:
			return
		}
	}
}

// Остановка симуляции движения
func (uc *droneUsecaseImpl) Stop() {
	close(uc.stopChan)
}

func (uc *droneUsecaseImpl) moveDrones() {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	const centerLat = 54.843243
	const centerLon = 83.088801
	const degreePerMeter = 1.0 / 111000.0 // Приблизительно 1 градус ≈ 111 км

	drones, _ := uc.repo.GetAll()
	for i, drone := range drones {
		// Угол для текущего дрона
		angle := float64(time.Now().UnixNano()/1e9) + float64(i)*2*math.Pi/float64(len(drones))
		radius := 300.0 + float64(i)*50.0 // Радиус в метрах

		// Смещение по широте и долготе
		deltaLat := radius * math.Cos(angle) * degreePerMeter
		deltaLon := radius * math.Sin(angle) * degreePerMeter

		// Обновляем координаты
		drone.Latitude = centerLat + deltaLat
		drone.Longitude = centerLon + deltaLon

		// Сохраняем в репозиторий и отправляем обновление
		uc.repo.Update(drone)
		uc.eventEmitter.Emit("drone_updated", drone)
	}
}
