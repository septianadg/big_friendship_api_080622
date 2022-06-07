package routes

import (
	"Golang_latihan/big_friendship_api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUsers)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUsers)
	r.DELETE("users/:id", controllers.DeleteUsers)

	r.GET("/requests", controllers.FindRequests)
	r.POST("/requests", controllers.CreateRequest)
	r.GET("/requests/:id", controllers.FindRequest)
	r.PATCH("/requests/:id", controllers.UpdateRequest)
	r.DELETE("requests/:id", controllers.DeleteRequests)
	r.GET("/request_to_me", controllers.FindRequestToMe)

	r.GET("/statuss", controllers.FindStatuss)
	r.POST("/statuss", controllers.CreateStatus)
	r.GET("/statuss/:id", controllers.FindStatus)
	r.PATCH("/statuss/:id", controllers.UpdateStatus)
	r.DELETE("statuss/:id", controllers.DeleteStatus)
	return r
}
