package main

import (
	"github.com/jsurnow/airports"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	airports.Load()
	//airports.LoadDB()
	router := gin.Default()

	router.GET("/airport/:iata", func(c *gin.Context) {
		iata := c.Param("iata")
		airport := airports.Get(iata)
		if airport.IATA != "" {
			c.IndentedJSON(http.StatusOK, airport)
		} else {
			c.Status(http.StatusNotFound)
		}
	})

	router.Run()
}
