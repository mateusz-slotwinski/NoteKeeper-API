package controllers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	Helpers "NoteKeeperAPI/src/helpers"
	Services "NoteKeeperAPI/src/services"
	Requests "NoteKeeperAPI/src/types/requests"
)

type NotesController struct {
	Service Services.NotesService
}

func (v NotesController) GetNotes(c *gin.Context) {
	UserID := c.Param("id")

	Notes := v.Service.GetNotes(UserID)
	c.JSON(http.StatusOK, Notes)
}

func (v NotesController) GetSingleNote(c *gin.Context) {
	NoteID := c.Param("note")

	Note := v.Service.GetSingleNote(NoteID)

	if Note == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_data",
		})

		return
	}

	c.JSON(http.StatusOK, Note)
}

func (v NotesController) CreateNote(c *gin.Context) {
	var req Requests.CreateNote

	err := c.ShouldBindJSON(&req)
	Helpers.PrintError(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	Note := v.Service.CreateNote(req)

	if Note == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_data",
		})

		return
	}

	c.JSON(http.StatusCreated, req)
}

func (v NotesController) UpdateNote(c *gin.Context) {
	var req Requests.CreateNote

	err := c.ShouldBindJSON(&req)
	Helpers.PrintError(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	NoteID := c.Param("note")

	Note := v.Service.UpdateNote(req, NoteID)

	if Note == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_data",
		})

		return
	}

	c.JSON(http.StatusOK, Note)
}

func (v NotesController) DeleteNote(c *gin.Context) {
	NoteID := c.Param("note")

	ok := v.Service.DeleteNote(NoteID)

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid_data",
		})

		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
