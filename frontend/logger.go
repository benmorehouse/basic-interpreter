package main

import(
	logrus "github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Logger{ // returns a logger for good debugging
	f, err := os.OpenFile("basicInterpreterLogs.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		panic(err)
	}
	log := logrus.New()
	log.Out = f
	return log
}

