package entity

import "sync"

type Book struct {
	Id string `json:"id"`

	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   int    `json:"year"`

	Mu sync.Mutex `json:"-"`
}

type GetBookFilter struct {
	Author string

	Page   int
	Limit  int
	Offset int
}
