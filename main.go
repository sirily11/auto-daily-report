package main

import (
	"auto-daily-report/src/router"
	"io"
	"log"

	"github.com/google/logger"
)

func main() {
	logger.Init("Logger", true, false, io.Discard)

	route := router.Router()

	err := route.Run()
	if err != nil {
		log.Fatal(err)
	}
}
