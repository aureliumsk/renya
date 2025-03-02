package main

import (
	"errors"
	"log"
	"os"
)

type lg struct {
	*log.Logger
	file *os.File
}

var logger lg

func createLogger() {
	f, err := os.Create("log.log")
	if err != nil {
		panic(err)
	}
	logger.file = f
	logger.Logger = log.New(f, "", log.LstdFlags)
}

func closeLogger() {
	if logger.file == nil {
		panic(errors.New("logger.file = <nil>"))
	}
	logger.file.Close()
}
