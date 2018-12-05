package main

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

const (
	quitCmd = "quit"
)

func extLog() *log.Entry {
	log.SetReportCaller(true)
	log := log.WithFields(log.Fields{
		"prefix": "cmd",
	})
	return log
}

func args() {
	if len(os.Args) > 1 && os.Args[1] == "-help"{
		fmt.Printf("Aplikacja czyta ze standardowego wejscia. \n  %s - zakoncz apke\n", quitCmd)
		os.Exit(0)
	}
}

func main() {
	args()
	formatter := &log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
	}

	log.SetFormatter(formatter)
	//log := extLog()

	log.Infof("--- start ---")
	fin := os.Stdin
	reader := bufio.NewReader(fin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Errorf("Error when read line %v", err)
		}
		log.Infof("line: %v", string(line))
		if quit(line, err) {
			fin.Close()
			break
		}
	}
	log.Infof("--- end ---")
}

func quit(line []byte, err error) bool {
	if string(line) == quitCmd || err == io.EOF {
		return true
	}
	return false
}
