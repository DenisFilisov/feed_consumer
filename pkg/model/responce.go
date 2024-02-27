package model

import (
	"time"
)

type GetNewsById struct {
	Status   string   `json:"status"`
	Data     Data     `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Data struct {
	ID          string      `json:"id"`
	TeamID      string      `json:"teamId"`
	OptaMatchID interface{} `json:"optaMatchId"`
	Title       string      `json:"title"`
	Type        []string    `json:"type"`
	Teaser      interface{} `json:"teaser"`
	Content     string      `json:"content"`
	URL         string      `json:"url"`
	ImageURL    string      `json:"imageUrl"`
	GalleryURLs interface{} `json:"galleryUrls"`
	VideoURL    interface{} `json:"videoUrl"`
	Published   time.Time   `json:"published"`
}

type Metadata struct {
	CreatedAt time.Time `json:"createdAt"`
}

type GetAllNews struct {
	Status             string             `json:"status"`
	Data               []Data             `json:"data"`
	MetadataForAllNews MetadataForAllNews `json:"metadata"`
}

type MetadataForAllNews struct {
	CreatedAt  time.Time `json:"createdAt"`
	TotalItems int       `json:"totalItems"`
	Sort       string    `json:"sort"`
}
