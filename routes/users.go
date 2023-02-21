package routes

import (
	"fmt"
	"net/http"

	ur "github.com/GorginZ/community-library/user"
	"github.com/gin-gonic/gin"
)

var (
	userService *ur.UserService = ur.NewUserService(ur.WithFakeUserRepository())
)

func HandleUsers(c *gin.Context) {
	Users, err := getUsers()
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(http.StatusOK, Users)
}

func getUsers() ([]ur.User, error) {
	users, err := userService.GetAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func AddUser(c *gin.Context) {
	//create user object from the request body
	var user ur.User
	err := c.BindJSON(&user) // BindJSON binds the request body into a struct.
	fmt.Println("user :", user)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	//add user to the database
	err = userService.UserRepository.AddUser(user)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
}
