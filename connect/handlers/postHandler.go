package handlers

import (
	"fruites_crud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFruites(c *gin.Context) {

	reqBody := pkg.Fruites{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	result := pkg.FruitInsertService(reqBody)

	ress := gin.H{

		"data": result,
	}

	c.JSON(http.StatusBadRequest, ress)
	return

}
