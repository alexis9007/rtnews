package models

import "time"

type News struct {
	Title       string
	Content     string
	PublishedAt time.Time
	Author      string
}
