package app

import(
	"github.com/tiwariarjun91/bookstore_users-api/controllers"
	//"github.com/gin-gonic/gin"
)

func mapUrls(){

	router.GET("/ping", controllers.Ping)
	
}