package entity

import "sync"

type Book struct {
	Id string `json:"id"`

	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`

	Mu sync.Mutex `json:"-"`
}
