package routes

import (
	"contacts-pay/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/sync", controllers.SyncContacts)
	r.GET("/lookup/:phone", controllers.LookupContact)

	return r
}
