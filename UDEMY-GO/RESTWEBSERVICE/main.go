package main

import(
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	r.GET("/ping", handlePing)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func handlePing(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
