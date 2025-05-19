package dto

type UpdatePackRequest struct {
	PackSizes []uint `json:"packSizes"`
}

type CalculatePackRequest struct {
	Total uint `form:"total" binding:"required"`
}
