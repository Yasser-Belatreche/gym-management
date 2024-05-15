package application_specific

type PaginatedQuery struct {
	Page    *int
	PerPage *int
}

type PaginatedQueryResponse[T interface{}] struct {
	List       []T
	Pagination *Pagination
}

type Pagination struct {
	Page       int64
	PerPage    int64
	Total      int64
	TotalPages int64
	HasMore    bool
	FirstPage  int64
	LastPage   int64
	NextPage   int64
	PrevPage   int64
}

type PaginationOptions struct {
	Page    int
	PerPage int
	Skip    int
}

const (
	DefaultPage    = 1
	DefaultPerPage = 10
)

func NewPaginationOptions(query *PaginatedQuery) *PaginationOptions {
	page := DefaultPage
	perPage := DefaultPerPage

	if query.Page != nil && *query.Page >= 1 {
		page = *query.Page
	}

	if query.PerPage != nil && *query.PerPage >= 1 {
		perPage = *query.PerPage
	}

	skip := (page - 1) * perPage

	return &PaginationOptions{
		Page:    page,
		PerPage: perPage,
		Skip:    skip,
	}
}

func NewPaginatedResponse[T interface{}, I interface{}](options *PaginationOptions, total int64, list []I, mapper func(item I) T) PaginatedQueryResponse[T] {
	pagination := newPagination(int64(options.Page), int64(options.PerPage), total)

	var mappedList []T = make([]T, 0)
	for _, item := range list {
		mappedList = append(mappedList, mapper(item))
	}

	return PaginatedQueryResponse[T]{
		List:       mappedList,
		Pagination: pagination,
	}
}

func newPagination(page, perPage, total int64) *Pagination {
	totalPages := total / perPage
	if total%perPage > 0 {
		totalPages++
	}

	hasMore := page < totalPages
	firstPage := int64(1)
	lastPage := totalPages

	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}

	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 1
	}

	return &Pagination{
		Page:       page,
		PerPage:    perPage,
		Total:      total,
		TotalPages: totalPages,
		HasMore:    hasMore,
		FirstPage:  firstPage,
		LastPage:   lastPage,
		NextPage:   nextPage,
		PrevPage:   prevPage,
	}
}
