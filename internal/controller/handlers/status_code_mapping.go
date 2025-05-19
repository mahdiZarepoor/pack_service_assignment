package handlers

import (
	"github.com/mahdiZarepoor/pack_service_assignment/internal/consts"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	// General
	consts.ServerError:    http.StatusInternalServerError,
	consts.RecordNotFound: http.StatusNotFound,
}
