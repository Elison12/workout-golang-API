package main

import (
	Routers "github.com/Elison12/workoutAPI/Routers"
	"github.com/Elison12/workoutAPI/configs"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	// r := Routers.SetupRouter()
	r := Routers.SetupRouter()
	// running
	r.Run()
}
