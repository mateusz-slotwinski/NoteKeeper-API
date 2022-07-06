package controllers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	Config "NoteKeeperAPI/src/config"
	Helpers "NoteKeeperAPI/src/helpers"
	Services "NoteKeeperAPI/src/services"
	Requests "NoteKeeperAPI/src/types/requests"
)

type AuthController struct {
	Service Services.AuthService
}

func (v AuthController) Register(c *gin.Context) {
	var req Requests.Register

	err := c.ShouldBindJSON(&req)
	Helpers.PrintError(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if !v.Service.ValidateRegister(req) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "existing_data",
		})

		return
	}

	req.Password = v.Service.HashPassword(req.Password)
	v.Service.Register(req)
	Token := v.Service.CreateToken(req)

	c.SetCookie("jwt", Token, Config.JWT_ExpiresIn, "/", Config.Host, false, false)
	c.JSON(http.StatusCreated, Token)
}

func (v AuthController) Login(c *gin.Context) {
	var req Requests.Login

	err := c.ShouldBindJSON(&req)
	Helpers.PrintError(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	User := v.Service.Login(req)

	if User == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_data",
		})

		return
	}

	Token := v.Service.CreateToken(User)
	c.SetCookie("jwt", Token, Config.JWT_ExpiresIn, "/", Config.Host, false, false)
	c.JSON(http.StatusOK, Token)

}

func (v AuthController) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "localhost", false, false)
	c.JSON(http.StatusOK, "Hello, how are you?")
}
