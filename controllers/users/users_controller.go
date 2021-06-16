package users

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/tiwariarjun91/bookstore_users-api/domain/users"
	"github.com/tiwariarjun91/bookstore_users-api/services"
	"io/ioutil"
	//"strings"
	"fmt"
	"encoding/json"
)

func CreateUser(c *gin.Context){
	var user users.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	err1 := json.Unmarshal(bytes, &user)
	if err1 != nil{
		fmt.Println(err1)
		return 
	}

	result, err := services.CreateUser(user) 
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(user, "user")
	fmt.Println(err, "err")
	fmt.Println(string(bytes), "bytes")
	c.JSON(http.StatusCreated, result) // can use c.XML to send the response in XML format
}

func GetUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Not Implemented")
	fmt.Print("Test test test")
}

/*func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Not Implemented")
}*/