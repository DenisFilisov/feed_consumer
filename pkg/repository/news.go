package repository

import (
	"context"
	"feed_consumer/pkg/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strconv"
	"time"
)

type NewsMongo struct {
	db *mongo.Collection
}

type newsItem struct {
	ID          string   `bson:"_id"`
	TeamID      string   `bson:"teamId"`
	OptaMatchID string   `bson:"optaMatchId,omitempty"`
	Title       string   `bson:"title"`
	Type        []string `bson:"type"`
	Teaser      string   `bson:"teaser,omitempty"`
	Content     string   `bson:"content"`
	URL         string   `bson:"url"`
	ImageURL    string   `bson:"imageUrl"`
	GalleryURLs []string `bson:"galleryUrls,omitempty"`
	VideoURL    string   `bson:"videoUrl,omitempty"`
	Published   string   `bson:"published"`
}

func (l *NewsMongo) FindNewsById(id int) (feed model.GetNewsById, err error) {
	filter := bson.M{"_id": strconv.Itoa(id)}
	var n newsItem
	err = l.db.FindOne(context.Background(), filter).Decode(&n)
	if err != nil {
		logrus.Info(err)
		return model.GetNewsById{Status: "fail"}, err
	}

	return model.GetNewsById{
		Status:   "success",
		Metadata: model.Metadata{time.Now()},
		Data:     buildNewsModel(n),
	}, nil
}

func (l *NewsMongo) GetAllNews() (feeds model.GetAllNews, err error) {
	filter := bson.D{}
	cursor, err := l.db.Find(context.Background(), filter)
	if err != nil {
		logrus.Info(err)
	}

	var data []model.Data

	for cursor.Next(context.Background()) {
		var n newsItem
		if err := cursor.Decode(&n); err != nil {
			logrus.Info("Can't decode news")
			return model.GetAllNews{Status: "fail"}, err
		}
		data = append(data, buildNewsModel(n))
	}
	return model.GetAllNews{
		Status: "success",
		MetadataForAllNews: model.MetadataForAllNews{
			CreatedAt:  time.Now(),
			TotalItems: len(data),
			Sort:       "-published",
		},
		Data: data,
	}, err
}

func buildNewsModel(newsFromMongo newsItem) model.Data {
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, newsFromMongo.Published)
	if err != nil {
		logrus.Info("Error parsing time: %v", err)
	}
	return model.Data{
		ID:          newsFromMongo.ID,
		TeamID:      newsFromMongo.TeamID,
		OptaMatchID: newsFromMongo.OptaMatchID,
		Title:       newsFromMongo.Title,
		Type:        newsFromMongo.Type,
		Teaser:      newsFromMongo.Teaser,
		Content:     newsFromMongo.Content,
		URL:         newsFromMongo.URL,
		ImageURL:    newsFromMongo.ImageURL,
		GalleryURLs: newsFromMongo.GalleryURLs,
		VideoURL:    newsFromMongo.GalleryURLs,
		Published:   parsedTime,
	}
}

func (l *NewsMongo) SaveAllNews(feeds model.NewListInformation) error {
	for _, feed := range feeds.NewsletterNewsItems.News {
		newsItem := newsItem{
			ID:          feed.NewsArticleID,
			TeamID:      "",
			OptaMatchID: feed.OptaMatchId,
			Title:       feed.Title,
			Type:        []string{},
			Teaser:      feed.TeaserText,
			Content:     "",
			URL:         feed.ArticleURL,
			ImageURL:    feed.ThumbnailImageURL,
			GalleryURLs: []string{},
			VideoURL:    "",
			Published:   feed.PublishDate,
		}
		log, err := l.db.InsertOne(context.Background(), newsItem)
		if err != nil {
			logrus.Info(err)
			return err
		}
		logrus.Info("News saved with ID: ", log.InsertedID)
	}
	return nil
}

func NewNewsMongo(db *mongo.Collection) *NewsMongo {
	return &NewsMongo{db: db}
}
