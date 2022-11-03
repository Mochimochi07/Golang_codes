package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var Persons = []Person{
	{ID: "1", Name: "Ajdampil", Age: 52, Address: "Taal"},
	{ID: "2", Name: "ADampil", Age: 32, Address: "Lemery"},
	{ID: "3", Name: "Aj", Age: 23, Address: "Calaca"},
}

func getPersons(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Persons)
}

func main() {
	router := gin.Default()
	router.GET("/persons", getPersons)
	router.Run("localhost:8080")
}
