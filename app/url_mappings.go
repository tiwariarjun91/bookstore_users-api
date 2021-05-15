package app

import(
	"github.com/tiwariarjun91/bookstore_users-api/controllers/users"
	"github.com/tiwariarjun91/bookstore_users-api/controllers/ping"

	//"github.com/gin-gonic/gin"
)

func mapUrls(){

	router.GET("/ping", ping.Ping) // we are passing reference to the function being called Ping() would be incorrect in this context as it wil execute the function there

	//router.GET("/users/search", controllers.SearchUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}