package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type cmdInOut struct { }

func newCmdInOut() *cmdInOut {
	return &cmdInOut{}
}

func (cmd *cmdInOut) Help() {
	if len(os.Args) > 1 && os.Args[1] == "-help"{
		fmt.Printf("Aplikacja czyta ze standardowego wejscia. \n  %s - zakoncz apke\n", quitCmd)
		os.Exit(0)
	}
}

func (cmd *cmdInOut) execute() {
	fin := os.Stdin
	reader := bufio.NewReader(fin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			logger.Errorf("Error when read line %v", err)
		}
		logger.Infof("line: %v", string(line))
		if cmd.quit(line, err) {
			fin.Close()
			break
		}
	}
}

func (cmd *cmdInOut) quit(line []byte, err error) bool {
	if string(line) == quitCmd || err == io.EOF {
		return true
	}
	return false
}
