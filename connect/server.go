package connect

import (
	"github.com/gin-gonic/gin"
)

func Connect() {
	r := gin.Default()
	SetupRoutes(r)
	r.Run()
	//	connect.SetupRoutes()
}
