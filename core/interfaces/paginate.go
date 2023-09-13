package interfaces

type PaginateResult struct {
	Pagination *PaginateTotal `json:"pagination,omitempty"`
	Results    interface{}    `json:"results" swaggertype:"array,object"`
}

type PaginateTotal struct {
	Total int64 `json:"total,omitempty"`
}
