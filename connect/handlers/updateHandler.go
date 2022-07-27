package handlers

import (
	"fruites_crud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateFruitHandler(c *gin.Context) {

	reqBody := pkg.Fruites{}

	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err.Error(),
		}
		c.JSON(http.StatusBadGateway, res)
	}
	result := pkg.FruitUpdate(reqBody)

	if result == false {
		res := gin.H{
			"error": "Something went wrong",
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"error": "Updated Successfully",
	}
	c.JSON(http.StatusOK, res)
	return

}
