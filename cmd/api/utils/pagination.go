package utils

// PagedResponse is the standard envelope for paginated list endpoints.
type PagedResponse[T any] struct {
	Data       []T   `json:"data"`
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

// NewPagedResponse builds a PagedResponse, computing total_pages from total and limit.
func NewPagedResponse[T any](data []T, page, limit int, total int64) PagedResponse[T] {
	totalPages := 0
	if limit > 0 {
		totalPages = int((total + int64(limit) - 1) / int64(limit))
	}
	return PagedResponse[T]{
		Data:       data,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}
}
