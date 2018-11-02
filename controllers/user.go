package controllers

import (
	"log"
	"strconv"
	"user-api/models"

	"github.com/gin-gonic/gin"
	gmodels "github.com/golang-modules/users/models"
	gmysql "github.com/golang-modules/users/mysql"
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

	log.Println("In GetAll")

	ul := gmysql.New()
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

	ul := gmysql.New()
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

//Create to retrieve user
func (uc UserController) Create(c *gin.Context) {

	log.Println("In Create")

	u := gmodels.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.Error(err)
		return
	}

	log.Println(u)

	ul := gmysql.New()
	r, err := ul.Create(u)

	res := models.Response{}
	if err != nil {
		res.StatusCode = 1
		res.Message = err.Error()
	} else {
		res.StatusCode = 1000
		res.Message = "Success"
		res.Data = r
	}

	c.JSON(200, res)
}

//Update to retrieve user
func (uc UserController) Update(c *gin.Context) {
	u := gmodels.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.Error(err)
		return
	}

	log.Println(u)

	ul := gmysql.New()
	r, err := ul.Update(u)

	res := models.Response{}
	if err != nil {
		res.StatusCode = 1
		res.Message = err.Error()
	} else {
		res.StatusCode = 1000
		res.Message = "Success"
		res.Data = r
	}

	c.JSON(200, res)
}

//UpdateStatus to retrieve user
func (uc UserController) UpdateStatus(c *gin.Context) {
	u := gmodels.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.Error(err)
		return
	}

	log.Println(u)

	ul := gmysql.New()
	r, err := ul.UpdateStatus(u.ID, u.Active)

	res := models.Response{}
	if err != nil {
		res.StatusCode = 1
		res.Message = err.Error()
	} else {
		res.StatusCode = 1000
		res.Message = "Success"
		res.Data = r
	}

	c.JSON(200, res)
}

//UpdatePassword to retrieve user
func (uc UserController) UpdatePassword(c *gin.Context) {
	u := gmodels.User{}
	err := c.BindJSON(&u)
	if err != nil {
		c.Error(err)
		return
	}

	log.Println(u)

	ul := gmysql.New()
	r, err := ul.UpdatePassword(u.ID, u.Password)

	res := models.Response{}
	if err != nil {
		res.StatusCode = 1
		res.Message = err.Error()
	} else {
		res.StatusCode = 1000
		res.Message = "Success"
		res.Data = r
	}

	c.JSON(200, res)
}

//Delete to retrieve user
func (uc UserController) Delete(c *gin.Context) {
	id := c.Param("id")

	i, _ := strconv.ParseInt(id, 10, 64)

	ul := gmysql.New()
	u, err := ul.Delete(i)

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
