package structs

type PaginationLink struct {
	URL    *string `json:"url"`
	Label  string  `json:"label"`
	Active bool    `json:"active"`
}

type PaginatedResponse[T any] struct {
	CurrentPage  int              `json:"current_page"`
	Data         []T              `json:"data"`
	FirstPageURL string           `json:"first_page_url"`
	From         int              `json:"from"`
	LastPage     int              `json:"last_page"`
	LastPageURL  string           `json:"last_page_url"`
	Links        []PaginationLink `json:"links"`
	NextPageURL  *string          `json:"next_page_url"`
	Path         string           `json:"path"`
	PerPage      int              `json:"per_page"`
	PrevPageURL  *string          `json:"prev_page_url"`
	To           int              `json:"to"`
	Total        int64            `json:"total"`
}
