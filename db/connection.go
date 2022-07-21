package db

import (
	"fmt"
	"log"

	"github.com/Elison12/workoutAPI/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenConnection() (*gorm.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := gorm.Open(postgres.Open(sc), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = conn.AutoMigrate(&Workout{}); err != nil {
		log.Println(err)
	}
	return conn, err
}
