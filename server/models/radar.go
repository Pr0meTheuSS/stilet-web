package models

type Radar struct {
    ID        string  `json:"id"`
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Coverage  float64 `json:"coverage"` // сектор покрытия в градусах
}
