package requests

type UpdatePackRequest struct {
	PackSizes []int `json:"packSizes"`
}

type CalculatePackRequest struct {
	Total int `form:"total" binding:"required"`
}
