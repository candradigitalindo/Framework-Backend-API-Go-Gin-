package helpers

import (
	"candra/backend-api/structs"
	"strconv"
)

func BuildPaginationResponse(
	baseURL string,
	page, limit int,
	total int64,
	data []structs.UserResponse,
) structs.PaginatedUserResponse {
	lastPage := int((total + int64(limit) - 1) / int64(limit))
	from := (page-1)*limit + 1
	to := (page-1)*limit + len(data)
	if total == 0 {
		from = 0
		to = 0
	}

	// Build links array
	var links []structs.PaginationLink
	for i := 1; i <= lastPage; i++ {
		url := baseURL + "?page=" + strconv.Itoa(i) + "&limit=" + strconv.Itoa(limit)
		links = append(links, structs.PaginationLink{
			URL:    &url,
			Label:  strconv.Itoa(i),
			Active: i == page,
		})
	}
	// Previous link
	var prevURL *string
	if page > 1 {
		url := baseURL + "?page=" + strconv.Itoa(page-1) + "&limit=" + strconv.Itoa(limit)
		prevURL = &url
	}
	// Next link
	var nextURL *string
	if page < lastPage {
		url := baseURL + "?page=" + strconv.Itoa(page+1) + "&limit=" + strconv.Itoa(limit)
		nextURL = &url
	}
	links = append([]structs.PaginationLink{{
		URL:    prevURL,
		Label:  "&laquo; Previous",
		Active: false,
	}}, links...)
	links = append(links, structs.PaginationLink{
		URL:    nextURL,
		Label:  "Next &raquo;",
		Active: false,
	})

	return structs.PaginatedUserResponse{
		CurrentPage:  page,
		Data:         data,
		FirstPageURL: baseURL + "?page=1&limit=" + strconv.Itoa(limit),
		From:         from,
		LastPage:     lastPage,
		LastPageURL:  baseURL + "?page=" + strconv.Itoa(lastPage) + "&limit=" + strconv.Itoa(limit),
		Links:        links,
		NextPageURL:  nextURL,
		Path:         baseURL,
		PerPage:      limit,
		PrevPageURL:  prevURL,
		To:           to,
		Total:        total,
	}
}
