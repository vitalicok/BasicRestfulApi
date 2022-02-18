package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type car struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Price float64 `json:"price"`
}

var cars = []car{
	{ID: "1", Name: "Ford", Model: "Mustang", Price: 10000},
	{ID: "2", Name: "Tesla", Model: "X", Price: 45000},
	{ID: "3", Name: "Mazda", Model: "MX-5", Price: 4500},
}

func main() {
	router := gin.Default()
	router.GET("/cars", getCars)
	router.GET("/cars/:id", getCarsByID)
	router.POST("/cars", postCars)

	router.Run("localhost:8081")
}

func getCars(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cars)
}

func postCars(c *gin.Context) {
	var newCar car

	if err := c.BindJSON(&newCar); err != nil {
		return
	}

	cars = append(cars, newCar)
	c.IndentedJSON(http.StatusCreated, newCar)
}

func getCarsByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range cars {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "car not found"})
}
