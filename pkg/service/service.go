package service

import (
	"feed_consumer/pkg/model"
	"feed_consumer/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type News interface {
	FindNewsById(id int) (model.GetNewsById, error)
	GetAllNews() (model.GetAllNews, error)
	SaveAllNews() error
}

type Service struct {
	News
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		News: NewNewsService(repos.News),
	}
}
