package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/go-cache"
)

type Response struct {
	Action string
	X      int
	Y      int
	Answer int
	Cached bool
}

var response Response

var ca = cache.New(time.Minute, 2*time.Minute)

func main() {
	router := gin.Default()

	router.GET("/:action", func(c *gin.Context) {
		if data, cached := ca.Get(c.Request.URL.String()); cached {
			response := data.(*Response)
			response.Cached = true
		} else {
			x, errX := strconv.Atoi(c.Query("x"))
			y, errY := strconv.Atoi(c.Query("y"))
			if errX != nil || errY != nil {
				log.Fatal(errX, errY)
			} else {
				response.X = x
				response.Y = y
			}
			switch c.Param("action") {
			case "add":
				response.Answer = (x + y)
			case "subtract":
				response.Answer = (x - y)
			case "multiply":
				response.Answer = (x * y)
			case "divide":
				response.Answer = (x / y)
			}
			response.Action = c.Param("action")
			response.Cached = false
			ca.Set(c.Request.URL.String(), &response, time.Minute)
		}
		c.JSON(200, response)
	})
	router.Run(":8080")
}
