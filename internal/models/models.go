package models

import "time"

// Для MongoDB
type Url struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Url       string    `json:"url"`
	ShortUrl  string    `json:"shorturl"`
	CreatedAt time.Time `json:"createdat"`
	LastVisit time.Time `json:"lastvisit"`
}

// Для ответов
type Response struct {
	ShortUrl string `json:"shorturl"`
}
