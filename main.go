package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetReportCaller(true)
}

func main() {
	t := time.Time{}
	t = time.Now()
	fmt.Println(t)
	logrus.Info("-----------started!---------")
	c := RestController{}
	c.init()
	c.run(":8000")
}
