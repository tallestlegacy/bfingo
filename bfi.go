package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/tallestlegacy/bfingo/interpreter"
)

func main() {
	args := os.Args
	bfFile, inputError := os.Open(args[1])

	if inputError != nil {
		fmt.Println("No such file : ", os.Args[1])
		return
	}

	defer bfFile.Close()

	var bfSource string
	inputReader := bufio.NewReader(bfFile)

	for {
		inputString, readerError := inputReader.ReadString('\n')
		bfSource += inputString
		if readerError == io.EOF {
			break
		}
	}

	interpreter.Interprete(bfSource)
}
