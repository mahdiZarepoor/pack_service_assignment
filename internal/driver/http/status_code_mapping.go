package http

import (
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/consts"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	// General
	consts.ServerError:         http.StatusInternalServerError,
	consts.RecordNotFound:      http.StatusNotFound,
	consts.CacheNotInitialized: http.StatusInternalServerError,
}
