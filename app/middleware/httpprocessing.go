package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	model "github.com/aniljaiswalcs/everphone/app/model"
)

type AssignHandler struct {
	Assigner model.EmployeeGiftAssigner
}

func NewAssignHandler(assigner model.EmployeeGiftAssigner) *AssignHandler {
	return &AssignHandler{
		Assigner: assigner,
	}
}

func (h *AssignHandler) HandleAssignGift(c *gin.Context) {
	gift, err := h.Assigner.AssignGift(c.Param("employeeName"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gift)
}
