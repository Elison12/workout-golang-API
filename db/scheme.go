package db

import "gorm.io/gorm"

type Workout struct {
	gorm.Model
	// ID        int64  `json: "id"`
	Exercicio string `json: "exercicio"`
	Series    int64  `json: "series"`
	Carga     int64  `json:"carga"`
}
