package main

const (
	// polecenie konczace aplikacje
	quitCmd = "quit"
)

func main() {
	cmdInOut := newCmdInOut()
	cmdInOut.Help()
	logger.Infof("--- start ---")
	cmdInOut.execute()
	logger.Infof("--- end ---")
}

