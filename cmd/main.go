package main

import (
	"Test_task1/cmd/server"
	"Test_task1/onlineStore/handler"
	"Test_task1/onlineStore/repository"
	"Test_task1/onlineStore/service"
	"Test_task1/pkg/database"
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	if err := initConfig(); err != nil {
		log.Fatalf("init config err: %v", err)
	}

	db, err := database.NewPostgresDB(database.DBCfg{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("init db err: %v", err)
	}
	defer db.Close()
	log.Printf("init db success")

	srv := server.Server{}
	srvConfig := server.SrvCfg{
		Port:                viper.GetString("server.port"),
		ReaderHeaderTimeOut: viper.GetDuration("server.reader_header_time_out"),
		IdleTimeout:         viper.GetDuration("server.idle_timeout"),
		WriterHeaderTimeOut: viper.GetDuration("server.writer_header_time_out"),
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	go func() {
		if err := srv.RunServer(srvConfig, handlers.MapRoutes()); err != nil {
			log.Fatalf("run server err: %v", err)
		}
	}()

	log.Printf("run server success on port %v", srvConfig.Port)

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt)
	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown err: %v", err)
	}

	log.Print("server shutdown success")
}

func initConfig() error {
	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
