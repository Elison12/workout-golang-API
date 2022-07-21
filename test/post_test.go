package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Elison12/workoutAPI/controllers"
	"github.com/Elison12/workoutAPI/db"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestPostWorkout(t *testing.T) {

	db.OpenConnection()
	router := gin.Default()

	values := map[string]interface{}{"Exercicio": "Banana", "series": 45, "Carga": 50}
	json_data, err := json.Marshal(values)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", "/workout", bytes.NewBuffer(json_data))
	router.POST("/workout", controllers.AddNewWorkout)

	var res map[string]interface{}

	w := httptest.NewRecorder()

	json.NewDecoder(w.Body).Decode(&res)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// assert.Equal(t, GetGroceryTest("1"), w.Body.String())
	// fmt.Println(GetGroceryTest("1"))
	fmt.Println(w.Body.String())

}
