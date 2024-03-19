package handler

import (
	"feed_consumer/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouters() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		news := api.Group("/news")
		{
			news.GET("/", h.GetAllNews)
			news.GET("/:id", h.GetNewsById)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func (h *Handler) ScheduleNewsFeeds(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			// Fetch data
			if err := h.services.News.SaveAllNews(); err != nil {
				logrus.Println("Error fetching data: ", err)
			}
		}
	}
}
