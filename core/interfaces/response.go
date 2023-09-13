package interfaces

type ResponseDefault struct {
	Data    interface{} `json:"data"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

type PaginateResponse struct {
	Data    *PaginateResult `json:"data"`
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Errors  interface{}     `json:"errors"`
}
