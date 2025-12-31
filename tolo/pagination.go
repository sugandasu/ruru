package tolo

type PaginationRequest struct {
	PerPage int `query:"per_page" validate:"omitempty,min=1,max=100"`
	Page    int `query:"page" validate:"omitempty,min=1"`
}

type Pagination struct {
	Total int `json:"total"`
}

type PaginationResponse struct {
	Items      any        `json:"items"`
	Pagination Pagination `json:"pagination"`
}

func NewPaginationResponse(items any, total int) PaginationResponse {
	return PaginationResponse{
		Items: items,
		Pagination: Pagination{
			Total: total,
		},
	}
}
