package models

import (
	"log"

	"github.com/Elison12/workoutAPI/db"
)

func Update(w *db.Workout, id string) (err error) {
	// fmt.Println(w)
	// Config.DB.Save(w)
	// return nil

	conn, err := db.OpenConnection()
	if err != nil {
		log.Println(err)

		return
	}
	if err = conn.Save(w).Error; err != nil {
		log.Println(err)

		return err
	}
	return nil
}
