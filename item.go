package main

import "time"

type Item struct {
	Id string
	Title string
	Description string
	DueTo time.Time
	Completed bool
}
