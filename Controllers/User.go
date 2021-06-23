package Controllers

import (
	"go-rest-api/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []Models.User
	err := Models.GetAllUser(&users)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user Models.User

	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
