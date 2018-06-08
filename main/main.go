package main

import "github.com/gin-gonic/gin"
import "github.com/jochasinga/requests"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/db", func(c *gin.Context) {
		res, err := requests.Get("http://db:8080/db")
		if err != nil {
				panic(err)
		}

		c.JSON(200, gin.H{
			"message": res.Body,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
