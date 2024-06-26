package models

import "time"

type URL struct {
	ID           uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ShortenedURL string    `json:"shortened_url"`
	OriginalURL  string    `json:"original_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateURLInput struct {
	OriginalURL string `json:"url"`
}

type NavigateURLInput struct {
	ShortenedURL string `json:"shortened_url"`
}

type GetUniqueIdResponse struct {
	UniqueID string `json:"uid"`
}

func (URL) TableName() string {
	return "url_record"
}
