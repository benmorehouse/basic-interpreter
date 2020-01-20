package main

import(
	logrus "github.com/sirupsen/logrus"
	"os"
)

func NewLogger(fileName string) (*logrus.Logger, error){ // returns a logger for good debugging
	f, err := os.OpenFile("basicInterpreterLogs.log", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil{
		return nil, err
	}
	log := logrus.Logger{
		Out: f,
	}

	return &log, nil
}

