package main

import "time"

type Document struct {
	Name 		string
	Content 	string
	CreatedAt  time.Time
}

func CreateDocument(name string, content string) (*Document) {
	document := Document{
		Name: name,
		Content: content,
		CreatedAt: time.Now(),
	}
	return &document
}
