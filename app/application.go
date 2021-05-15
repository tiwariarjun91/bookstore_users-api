package app

import(
	"github.com/gin-gonic/gin"
)

var(
	router = gin.Default() //router is a pointer to the Engine that already has logger and recovery middleware
)

func StartApplication(){
	mapUrls()
	

	router.Run(":8080")
}