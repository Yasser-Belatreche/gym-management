package utils

import "gym-management-gyms/src/lib/primitives/application_specific"

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
	Page    *int `form:"page" json:"page"`
	PerPage *int `form:"perPage" json:"perPage"`
}

type HttpPaginatedResponse[T interface{}] struct {
	List       []T                             `json:"list"`
	Pagination HttpPaginatedResponsePagination `json:"pagination"`
}

type HttpPaginatedResponsePagination struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"perPage"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"totalPages"`
	HasMore    bool  `json:"hasMore"`
	FirstPage  int64 `json:"firstPage"`
	LastPage   int64 `json:"lastPage"`
	NextPage   int64 `json:"nextPage"`
	PrevPage   int64 `json:"prevPage"`
}

func GetHttpPaginatedResponse[T interface{}, Q interface{}](res *application_specific.PaginatedQueryResponse[T], mapper func(item T) Q) *HttpPaginatedResponse[Q] {
	var list = make([]Q, 0)

	for _, item := range res.List {
		list = append(list, mapper(item))
	}

	return &HttpPaginatedResponse[Q]{
		List: list,
		Pagination: HttpPaginatedResponsePagination{
			Page:       res.Pagination.Page,
			PerPage:    res.Pagination.PerPage,
			Total:      res.Pagination.Total,
			TotalPages: res.Pagination.TotalPages,
			HasMore:    res.Pagination.HasMore,
			FirstPage:  res.Pagination.FirstPage,
			LastPage:   res.Pagination.LastPage,
			NextPage:   res.Pagination.NextPage,
			PrevPage:   res.Pagination.PrevPage,
		},
	}
}
