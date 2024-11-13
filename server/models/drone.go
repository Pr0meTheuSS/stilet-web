package models

type Drone struct {
    ID        string   `json:"id"`
    Name      string   `json:"name"`
    IsActive  bool     `json:"is_active"`
    Status    []string `json:"status"`
    Latitude  float64  `json:"latitude"`
    Longitude float64  `json:"longitude"`
}
