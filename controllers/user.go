package controllers

import (
	"strconv"
	"user-api/models"

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

	res := models.Response{}
	if err != nil {
		res.StatusCode = 1
		res.Message = err.Error()
	} else {
		res.StatusCode = 1000
		res.Message = "Success"
		res.Data = u
	}

	c.JSON(200, res)
}

//Get to retrieve user
func (uc UserController) Get(c *gin.Context) {
	id := c.Param("id")

	i, _ := strconv.ParseInt(id, 10, 64)

	ul := mysql.New()
	u, err := ul.Get(i)

	res := models.Response{}
	if err != nil {
		res.StatusCode = 1
		res.Message = err.Error()
	} else {
		res.StatusCode = 1000
		res.Message = "Success"
		res.Data = u
	}

	c.JSON(200, res)
}
