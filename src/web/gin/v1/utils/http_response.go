package utils

type HttpErrorResponse struct {
	Method        string                 `json:"method"`
	Path          string                 `json:"path"`
	Status        int                    `json:"status"`
	CorrelationId string                 `json:"correlationId"`
	Error         HttpErrorResponseError `json:"error"`
}

type HttpErrorResponseError struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Payload map[string]string `json:"payload"`
}

type HttpPaginatedRequest struct {
	Page    *int `form:"page"`
	PerPage *int `form:"perPage"`
}

type HttpPaginatedResponse[T interface{}] struct {
	List       []T                             `json:"list"`
	Pagination HttpPaginatedResponsePagination `json:"pagination"`
}

type HttpPaginatedResponsePagination struct {
	Page       int  `json:"page"`
	PerPage    int  `json:"perPage"`
	Total      int  `json:"total"`
	TotalPages int  `json:"totalPages"`
	HasMore    bool `json:"hasMore"`
	FirstPage  int  `json:"firstPage"`
	LastPage   int  `json:"lastPage"`
	NextPage   int  `json:"nextPage"`
	PrevPage   int  `json:"prevPage"`
}
