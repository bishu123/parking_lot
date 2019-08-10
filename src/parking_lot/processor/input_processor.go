package processor

import (
	"bufio"
	"io"
	"os"
)

func ProcessFromFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ReadInput(f)
}

func ProcessFromConsole() {
	ReadInput(os.Stdin)
}

func ReadInput(reader io.Reader) {
	input := bufio.NewScanner(reader)
	for input.Scan() {
		inputStr := input.Text()
		if inputStr == "exit" {
			return
		} else {
			processInput(inputStr)
		}
	}
}
