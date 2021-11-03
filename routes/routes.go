package routes

import (
	//"first/controllers"

	"github.com/gin-gonic/gin"
	"github.com/sajjad3k/controllers"
)

func Setroutes() *gin.Engine {

	r := gin.Default()

	api := r.Group("/v1")
	{
		api.GET("/user", controllers.Getusers)
		api.GET("/user/:id", controllers.Getuserbyid)
		api.POST("/user", controllers.Createuser)
		api.PUT("/user/:id", controllers.Updateuser)
		api.DELETE("/user/:id", controllers.Deleteuser)
	}
	return r
}
