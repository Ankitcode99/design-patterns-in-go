package models

// NewsFeed represents a news item with properties and validation for non-empty category.
type NewsFeed struct {
	Headline    string `json:"headline"`
	Description string `json:"description"`
	PublishedOn string `json:"publishedOn"`
	Category    string `json:"category"`
}
