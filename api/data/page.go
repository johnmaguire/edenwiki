package data

import "time"

type Page struct {
	PageContent

	CreatedAt time.Time

	History PageHistory
}

type PageContent struct {
	Title string
	Body  string

	UpdatedAt time.Time
}

type PageHistory []PageContent
