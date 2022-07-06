package middlewares

import (
	fmt "fmt"
	http "net/http"
	strings "strings"

	gin "github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt"

	Config "NoteKeeperAPI/src/config"
)

func UseAuthorization(c *gin.Context) {
	AuthHeader := c.Request.Header["Authorization"]

	if AuthHeader == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		c.Abort()

		return
	}

	if AuthHeader[0] == "" || !strings.Contains(AuthHeader[0], " ") || strings.Split(AuthHeader[0], " ")[0] != "token" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid_token",
		})
		c.Abort()

		return
	}

	tokenString := strings.Split(AuthHeader[0], " ")[1]

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(Config.JWT_Secret), nil
	})

	if token == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid_token",
		})
		c.Abort()

		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		JWT_ID := claims["data"].(map[string]interface{})["_id"]
		PARAM_ID := c.Param("id")

		if JWT_ID != PARAM_ID {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "unauthorized",
			})
			c.Abort()

			return
		}

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		c.Abort()

		return
	}

	c.Next()
}
