package model

type ShortLink struct {
	ID       *string
	LongURL  string
	ShortURL string
}

type Repository interface {
	CreateLink(link ShortLink) error
	FindByID(id string) (*ShortLink, error)
}