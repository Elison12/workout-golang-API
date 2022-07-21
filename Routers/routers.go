package routers

import (
	"github.com/Elison12/workoutAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("workout", controllers.ListWorkouts)
	r.POST("workout", controllers.AddNewWorkout)
	r.GET("workout/:id", controllers.GetOneWorkout)
	r.PUT("workout/:id", controllers.UpdateWorkout)
	r.DELETE("workout/:id", controllers.DeleteWorkout)

	return r
}
