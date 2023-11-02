package main

import (
	"flag"
	"fmt"
	"log"
	httpDelivery "file-upload-api/delivery/http"
	port "file-upload-api/delivery/http"
	"file-upload-api/domain"
	"file-upload-api/pkg/logger"
	"file-upload-api/repository/mongodb"
	"os"
)

func init() {

}

func main() {
	l, err := logger.InitLogger()

	if err != nil {
		panic(err)
	}

	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "dev"
	}

	l.Info(fmt.Sprintf("Loading %s env", env))

	domain.GetSecrets(l)

	_ = mongodb.New(l)

	httpConfig := httpDelivery.Config{
	}

	app := port.RunHttpServer(httpConfig)

	port := os.Getenv("PORT")

	if port == "" {
		port = "6001"
	}

	addr := flag.String("addr", fmt.Sprintf(":%s", port), "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
