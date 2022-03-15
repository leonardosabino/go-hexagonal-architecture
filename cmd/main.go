package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hexagonal/template/internal/src/config"
	"hexagonal/template/internal/src/transport_layer/http/router"

	"github.com/joomcode/errorx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	time.Local = time.UTC

	startServer()
	signalStopServer()
}

func startServer() {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	server.Use(middleware.Gzip())
	server.Use(middleware.RemoveTrailingSlash())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET},
	}))

	router.Register(server)

	go func() {
		log.Println("Starting http server at " + config.GetConfig().ServerHost)
		error := server.Start(config.GetConfig().ServerHost)
		errorx := errorx.Decorate(error, "Failed to start server")
		log.Fatal(errorx)
	}()
}

func signalStopServer() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		log.Println("Received SIGINT, stopping...")
	case syscall.SIGTERM:
		log.Println("Received SIGTERM, stopping...")
	}
}
