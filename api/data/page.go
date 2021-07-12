package data

// Page is used to return a page response from the API.
type Page struct {
	Body string
}

// PageList is used to return a list of pages names from the API.
type PageList struct {
	Pages []string
}
