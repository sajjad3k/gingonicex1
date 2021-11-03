package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sajjad3k/models"
)

//var records []models.User

var records = []models.User{
	{Id: 1, Name: "ser", Phone: "65768", City: "chittoor", Email: "singht@kio.com"},
	{Id: 2, Name: "nft", Phone: "65909", City: "puttor", Email: "nft@kio.com"},
	{Id: 3, Name: "cvb", Phone: "65123", City: "mumbai", Email: "lpu@kio.com"},
}

//Get all users
func Getusers(c *gin.Context) {

	if records != nil {
		c.JSON(http.StatusOK, records)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}

}

//Get user for the given ID
func Getuserbyid(c *gin.Context) {

	reqid, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusNotAcceptable)
	}

	for i, val := range records {
		if val.Id == reqid {
			c.JSON(http.StatusOK, records[i])
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	}

}

//Create a user
func Createuser(c *gin.Context) {
	var p models.User
	if c.Request.PostForm == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	c.BindJSON(&p)
	records = append(records, p)
	c.JSON(http.StatusOK, p)
}

//update the user details for a given id
func Updateuser(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		name := c.Params.ByName("name")
		var s bool = false
		var o models.User
		for i, val := range records {
			if val.Id == id {
				records[i].Name = name
				s = true
				o = records[i]
			}
		}
		if s != false {
			c.JSON(http.StatusOK, o)
		}
	}

}

//Delete the user at with given ID
func Deleteuser(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		var v models.User
		var b bool
		for i, val := range records {
			if val.Id == id {
				v = val
				records[i] = records[i+1]
				b = true
			}
		}
		if b != false {
			c.JSON(http.StatusOK, v)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	}
}
