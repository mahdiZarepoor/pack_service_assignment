package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ctx               *gin.Context
	response          map[string]interface{}
	statusCodeMapping map[string]int
	error             Error
}

type ErrorResponse struct {
	Error string `json:"error" example:"error message"`
}

func NewResponse(ctx *gin.Context, statusCodeMappings ...map[string]int) *Response {

	var statusCodeMapping = make(map[string]int)

	if len(statusCodeMappings) > 0 {
		statusCodeMapping = statusCodeMappings[0]
	}

	return &Response{
		ctx:               ctx,
		response:          make(map[string]interface{}),
		statusCodeMapping: statusCodeMapping,
	}
}

func (r *Response) Validation(err error) *Response {
	r.response["validationErrors"] = err.Error()
	return r
}

func (r *Response) Payload(data interface{}) *Response {
	r.response["data"] = data
	return r
}

func (r *Response) Meta(data interface{}) *Response {
	r.response["meta"] = data
	return r
}

func (r *Response) Error(err Error) *Response {
	r.error = err
	var errorResponse = ErrorResponse{
		Error: err.GetMessage(),
	}
	r.response["error"] = errorResponse.Error
	return r
}

func (r *Response) ErrorMsg(err string) *Response {
	var errorResponse = ErrorResponse{
		Error: err,
	}
	r.response["error"] = errorResponse.Error
	return r
}

func (r *Response) Message(msg string) *Response {
	r.response["message"] = msg
	return r
}

func (r *Response) Echo(overrideStatusCodes ...int) {

	var statusCode int

	if len(overrideStatusCodes) > 0 {
		statusCode = overrideStatusCodes[0]
	} else {
		if r.error != nil {
			if val, ok := r.statusCodeMapping[r.error.GetMessage()]; !ok {
				statusCode = http.StatusInternalServerError
			} else {
				statusCode = val
			}
		}
	}

	if statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices {
		r.ctx.JSON(statusCode, r.response)
		return
	}

	r.ctx.AbortWithStatusJSON(statusCode, r.response)
}
