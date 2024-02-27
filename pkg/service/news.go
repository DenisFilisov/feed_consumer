package service

import (
	"encoding/xml"
	"feed_consumer/pkg/model"
	"feed_consumer/pkg/repository"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type NewsService struct {
	repo repository.News
}

func (n *NewsService) FindNewsById(id int) (model.GetNewsById, error) {
	return n.repo.FindNewsById(id)
}

func (n *NewsService) GetAllNews() (model.GetAllNews, error) {
	return n.repo.GetAllNews()
}

func (n *NewsService) SaveAllNews() error {
	s := "https://www.htafc.com/api/incrowd/getnewlistinformation?count=50"
	resp, err := http.Get(s)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	var feeds model.NewListInformation

	if err != nil {
		return err
	}

	if err := xml.Unmarshal(body, &feeds); err != nil {
		logrus.Infof("Error unmarshaling XML: %v", err)
		return err
	}
	logrus.Info("Response Status Code:", resp.StatusCode)

	return n.repo.SaveAllNews(feeds)
}

func NewNewsService(repo repository.News) *NewsService {
	return &NewsService{repo: repo}
}
