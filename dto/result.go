package dto

type Result struct {
	Status      int         `json:"status"`
	Message     string      `json:"message"`
	TotalData   int64       `json:"totalData,omitempty"`
	TotalPages  int         `json:"totalPages,omitempty"`
	CurrentPage int         `json:"currentPage,omitempty"`
	Data        interface{} `json:"data,omitempty"`
}
