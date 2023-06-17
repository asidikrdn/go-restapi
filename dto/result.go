package dto

type SuccessResult struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	TotalData   int         `json:"totalData,omitempty"`
	TotalPages  int         `json:"totalPages,omitempty"`
	CurrentPage int         `json:"currentPage,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}

type ErrorResult struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
