package controllers

import (
	"log"
	"strconv"

	"github.com/golang-modules/users/mysql"

	"github.com/gin-gonic/gin"
)

//UserController defines user type
type UserController struct {
}

//NewUserController -
func NewUserController() *UserController {
	return &UserController{}
}

//GetAll to retrieve user
func (uc UserController) GetAll(c *gin.Context) {

	ul := mysql.New()
	u, err := ul.GetAll()

	log.Println(err)

	c.JSON(200, u)
}

//Get to retrieve user
func (uc UserController) Get(c *gin.Context) {
	id := c.Param("id")

	i, _ := strconv.ParseInt(id, 10, 64)

	ul := mysql.New()
	u, _ := ul.Get(i)

	c.JSON(200, u)
}
