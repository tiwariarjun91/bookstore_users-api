package app

import(
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default()
)

func StartApplication(){
	mapUrls()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"Message" : "Pong",})
	})

	router.Run()
}