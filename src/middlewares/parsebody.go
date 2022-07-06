package middlewares

import (
	// http "net/http"

	gin "github.com/gin-gonic/gin"
)

func ParseBody(c *gin.Context) {
	// body:=Body{}

	// err := c.ShouldBindJSON(&body)
	// Helpers.PrintError(err)

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": err.Error(),
	// 	})
	// }

	c.Next()
}
