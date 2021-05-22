package users

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/tiwariarjun91/bookstore_users-api/domain/users"
	"io/ioutil"
	//"strings"
	"fmt"
)

func CreateUser(c *gin.Context){
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println(user, "user")
	fmt.Println(err, "err")
	fmt.Println(string(bytes), "bytes")
	c.String(http.StatusNotImplemented, "Not Implemented bro")
}

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Not Implemented")
	fmt.Print("Test test test")
}

/*func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Not Implemented")
}*/