package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseErrorResponse struct {
	Error interface{} `json:"error"`
}

func ErrorHandler(c *gin.Context, err any) {
	if e, ok := err.(error); ok {
		httpResponse := GenerateErrorResponse(e)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	c.AbortWithStatusJSON(
		http.StatusInternalServerError,
		GenerateErrorResponse(fmt.Errorf("%v", err)),
	)
}

func GenerateErrorResponse(err interface{}) *BaseErrorResponse {
	return &BaseErrorResponse{
		Error: err,
	}
}
