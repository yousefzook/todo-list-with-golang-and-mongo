package main

import (
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(true)
}

func main() {
	logrus.Info("-----------started!---------")
	restController := Controller{&RestController{}}
	restController.start(":8000")
}
