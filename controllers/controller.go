package controllers

import (
	"log"

	apihelpers "github.com/Elison12/workoutAPI/ApiHelpers"
	"github.com/Elison12/workoutAPI/db"
	"github.com/Elison12/workoutAPI/models"
	"github.com/gin-gonic/gin"
)

func AddNewWorkout(c *gin.Context) {
	var workout db.Workout
	c.BindJSON(&workout)
	err := models.Post(&workout)
	if err != nil {
		apihelpers.RespondJSON(c, 404, workout)
		log.Println(err)
	} else {
		apihelpers.RespondJSON(c, 200, workout)
		log.Println(err)
	}
}

func GetOneWorkout(c *gin.Context) {
	id := c.Params.ByName("id")
	var workout db.Workout
	err := models.GetOne(&workout, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, workout)
		log.Println(err)

	} else {
		apihelpers.RespondJSON(c, 200, workout)
		log.Println(err)

	}
}

func ListWorkouts(c *gin.Context) {
	var workouts []db.Workout

	err := models.GetAll(&workouts)
	if err != nil {
		apihelpers.RespondJSON(c, 404, workouts)
		log.Println(err)

	} else {
		apihelpers.RespondJSON(c, 200, workouts)
		log.Println(err)

	}
}

func DeleteWorkout(c *gin.Context) {
	var workout db.Workout
	id := c.Params.ByName("id")
	err := models.Delete(&workout, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, workout)
		log.Println(err)

	} else {
		apihelpers.RespondJSON(c, 200, workout)
		log.Println(err)

	}
}

func UpdateWorkout(c *gin.Context) {
	var workout db.Workout
	id := c.Params.ByName("id")
	err := models.GetOne(&workout, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, workout)
		log.Println(err)
	}
	c.BindJSON(&workout)
	err = models.Update(&workout, id)
	if err != nil {
		apihelpers.RespondJSON(c, 404, workout)
		log.Println(err)

	} else {
		apihelpers.RespondJSON(c, 200, workout)
		log.Println(err)
	}
}
