package main

import (
	"context"
	_ "feed_consumer/docs"
	"feed_consumer/pkg/handler"
	"feed_consumer/pkg/repository"
	"feed_consumer/pkg/server"
	"feed_consumer/pkg/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title News Feeds
// @version 1.0
// @description API for consume and get news feeds

// @host localhost:8080
// @BasePath /

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrusChan := make(chan *logrus.Entry, 100)
	logrus.AddHook(handler.NewChannelHook(logrusChan))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error while initializing configuration %s", err.Error())
	}

	db, err := repository.NewMongoDB(repository.Config{
		Host:       viper.GetString("db.mongo_db_host"),
		Port:       viper.GetString("db.mongo_db_port"),
		Username:   viper.GetString("db.mongo_db_username"),
		Password:   viper.GetString("db.mongo_db_password"),
		DataBase:   viper.GetString("db.mongo_db_database"),
		Collection: viper.GetString("db.mongo_db_collection"),
	})

	if err != nil {
		logrus.Fatalf("Can't initialize db: %s", err.Error())
		logrus.Printf("Can't initialize db: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	//(Can be processed through Queue)
	//go func() {
	//	for entry := range logrusChan {
	// Print logs
	//logrus.Printf("Data: %v, Time: %v, Level: %v, Message: %v\n", entry.Data, entry.Time, entry.Level, entry.Message)
	//}
	//}()

	// Schedule fetchData to run every 5 minutes
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	go func() {
		handlers.ScheduleNewsFeeds(ticker)
	}()

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
			logrus.Fatal("Can't run server:", err)
		}
	}()
	logrus.Print("App started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error while shutdown process: %s", err)
	}

	close(logrusChan)
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
