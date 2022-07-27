package connect

import (
	"fruites_crud/connect/handlers"

	"github.com/gin-gonic/gin"
)

//

func SetupRoutes(r *gin.Engine) {
	// r.GET("/user", GetAllUser)
	r.POST("/fruites", handlers.CreateFruites)

}
