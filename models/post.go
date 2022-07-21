package models

import (
	"log"

	"github.com/Elison12/workoutAPI/db")


func Post(w *db.Workout) (err error) {
	//instancia comunicação
	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)

		return
	}
	if err = conn.Create(w).Error; err != nil {
		log.Println(err)

		return err
	}
	return nil
}
