package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdiZarepoor/pack_service_assignment/configs"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/core/ports/packs_port"
	"github.com/mahdiZarepoor/pack_service_assignment/internal/driver/http/requests"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/logging"
	"github.com/mahdiZarepoor/pack_service_assignment/pkg/response"
)

type PackHandler struct {
	logging logging.Logger
	config  configs.Config
	packSrv packs_port.IPackService
}

func NewPackHandler(
	logging logging.Logger,
	config configs.Config,
	packSrv packs_port.IPackService,
) *PackHandler {
	return &PackHandler{
		logging: logging,
		config:  config,
		packSrv: packSrv,
	}
}

// List godoc
// @Summary List all packs
// @Description list all packs in system
// @Tags Packs
// @Accept json
// @Produce json
// @Param x-Client-Version header string true "Client version information" default(iOS-1.4.0)
// @Param x-Client header string true "Client Device information" default(iOS)
// @Param x-App-Version header string true "App version information" default(1.4.0)
// @Param language header string true "Language" default(en)
// @Success 200 {object} response.Response{data=[]uint}
// @Failure 400 {object} response.ErrorResponse "Failed response"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @ID Get-api-v1-packs
// @Router /api/v1/packs [get]
func (h *PackHandler) List(ctx *gin.Context) {
	packs, err := h.packSrv.List(ctx)
	if err != nil {
		h.logging.Error(logging.General, logging.InternalInfo, err.GetMessage(), nil)
		response.NewResponse(ctx, StatusCodeMapping).Error(err).Echo()
		return
	}

	response.NewResponse(ctx).
		Payload(packs).
		Echo(http.StatusOK)
}

// Update godoc
// @Summary Update packs
// @Description update packs in system
// @Tags Packs
// @Accept json
// @Produce json
// @Param x-Client-Version header string true "Client version information" default(iOS-1.4.0)
// @Param x-Client header string true "Client Device information" default(iOS)
// @Param x-App-Version header string true "App version information" default(1.4.0)
// @Param language header string true "Language" default(en)
// @Param Request body requests.UpdatePackRequest true "Request"
// @Success 204  "Successful response"
// @Failure 400 {object} response.ErrorResponse "Failed response"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @ID Put-api-v1-packs
// @Router /api/v1/packs [put]
func (h *PackHandler) Update(ctx *gin.Context) {
	var requestBody requests.UpdatePackRequest
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		response.NewResponse(ctx).Validation(err).Echo(http.StatusUnprocessableEntity)
		return
	}

	err := h.packSrv.Update(ctx, requestBody.PackSizes)
	if err != nil {
		h.logging.Error(logging.General, logging.InternalInfo, err.GetMessage(), nil)
		response.NewResponse(ctx, StatusCodeMapping).Error(err).Echo()
		return
	}

	response.NewResponse(ctx).
		Echo(http.StatusNoContent)
}

// Calculate godoc
// @Summary Calculate packs
// @Description calculate packs in system
// @Tags Packs
// @Accept json
// @Produce json
// @Param x-Client-Version header string true "Client version information" default(iOS-1.4.0)
// @Param x-Client header string true "Client Device information" default(iOS)
// @Param x-App-Version header string true "App version information" default(1.4.0)
// @Param language header string true "Language" default(en)
// @Param total query string true "Total"
// @Success 200 {object} response.Response{data=map[int]int}
// @Failure 400 {object} response.ErrorResponse "Failed response"
// @Failure 500 {object} response.ErrorResponse "Internal server error"
// @ID GET-api-v1-packs-calculate
// @Router /api/v1/packs/calculate [get]
func (h *PackHandler) Calculate(ctx *gin.Context) {
	var requestForm requests.CalculatePackRequest
	if err := ctx.ShouldBindQuery(&requestForm); err != nil {
		response.NewResponse(ctx).Validation(err).Echo(http.StatusUnprocessableEntity)
		return
	}

	p, err := h.packSrv.Calculate(ctx, requestForm.Total)
	if err != nil {
		h.logging.Error(logging.General, logging.InternalInfo, err.GetMessage(), nil)
		response.NewResponse(ctx, StatusCodeMapping).Error(err).Echo()
		return
	}

	response.NewResponse(ctx).
		Payload(p).
		Echo(http.StatusOK)
}
