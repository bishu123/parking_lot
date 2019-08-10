package main

import (
	"os"

	"parking_lot/processor"
)

const  (
	ExitingAsErrorInArgs = 1
)

func main(){
	args := os.Args
	switch len(args) {
	case 1:
		// Interactive Mode
		processor.ProcessFromConsole()
	case 2:
		// Read Input From File
		inputFile := os.Args[1]
		processor.ProcessFromFile(inputFile)
	default:
		os.Exit(ExitingAsErrorInArgs)
	}
	return
}
