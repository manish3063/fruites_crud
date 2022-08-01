package handlers

import (
	"fruites_crud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteFruites(c *gin.Context) {

	reqBody := pkg.Fruites{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	result := pkg.DeleteFruitsService(reqBody.ID)
	if result == false {
		ress := gin.H{
			"message": "error",
		}
		c.JSON(http.StatusBadRequest, ress)
		c.Abort()
		return
	}
	ress := gin.H{
		"message": "succefully deleted",
		"result":  result,
	}
	c.JSON(http.StatusOK, ress)

}
