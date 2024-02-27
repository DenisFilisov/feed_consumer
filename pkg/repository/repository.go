package repository

import (
	"feed_consumer/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type News interface {
	FindNewsById(id int) (feed model.GetNewsById, err error)
	GetAllNews() (feeds model.GetAllNews, err error)
	SaveAllNews(feeds model.NewListInformation) error
}

type Repository struct {
	News
}

func NewRepository(db *mongo.Collection) *Repository {
	return &Repository{
		News: NewNewsMongo(db),
	}
}
