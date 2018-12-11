package main

import (
	log "github.com/sirupsen/logrus"
)

//var logger *log.Entry
var logger *log.Logger

func init() {
	formatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	logger = log.StandardLogger()
	//logger = extLog()
}

func extLog() *log.Entry {
	log.SetReportCaller(true)
	log := log.WithFields(log.Fields{
		"prefix": "cmd",
	})
	return log
}

