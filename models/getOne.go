package models

import (
	"log"

	"github.com/Elison12/workoutAPI/db")


func GetOne(w *db.Workout, id string) (err error) {
	//instancia comunicação
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)

		return
	}

	if err = conn.Where("id = ?", id).First(w).Error; err != nil {
		log.Println(err)

		return err
	}

	return nil
}
