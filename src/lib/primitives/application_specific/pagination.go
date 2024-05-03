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
	Page       int
	PerPage    int
	Total      int
	TotalPages int
	HasMore    bool
	FirstPage  int
	LastPage   int
	NextPage   int
	PrevPage   int
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

func NewPagination(page, perPage, total int) *Pagination {
	totalPages := total / perPage
	if total%perPage > 0 {
		totalPages++
	}

	hasMore := page < totalPages
	firstPage := 1
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
