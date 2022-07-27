package handlers

import (
	"fruites_crud/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFruites(c *gin.Context) {

	reqBody := pkg.Fruites{}
	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	data := pkg.GetFruite()

	ress := gin.H{

		"detail_of_all_fruites": data,
	}

	c.JSON(http.StatusBadRequest, ress)
	return

}
