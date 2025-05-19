package handlers

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHealthHandler() *Handler {
	return &Handler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Check if the service is up and running
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} string "i_am_working"
// @ID get-services-health
// @Router /health [get]
func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "I'm working :)",
	})
}
