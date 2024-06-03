package main

import (
	"fmt"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"testapi/bd"
	"testapi/pkg/loger"
	"testapi/routes"
	"time"
)

func main() {
	// Инициализация Sentry
	// s
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              "https://d015837fbcc25c593682e791bc2cbdac@o4507311082962944.ingest.de.sentry.io/4507333767725136",
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	// Инициализация Sentry хука
	hook, err := logrus_sentry.NewSentryHook("https://d015837fbcc25c593682e791bc2cbdac@o4507311082962944.ingest.de.sentry.io/4507333767725136", []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})
	if err != nil {
		log.Fatalf("sentry hook error: %s", err)
	}
	loger.Logrus.Hooks.Add(hook)

	bd.InitDB()
	defer bd.DB.Close()

	server := gin.Default()
	routes.RegisterRoutes(server)

	if err := server.Run(":8080"); err != nil {
		loger.Logrus.Fatalf("Failed to run server: %v", err)

	}

	defer sentry.Flush(2 * time.Second)
}
