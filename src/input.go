package main

import (
	"bufio"
	"log"
	"os"
)

const maxLineLenght int = 60

func getText() ([]string, int) {
	if len(os.Args) > 1 && cmd_flags.file == "" {
		array := os.Args[1:]
		return normalizeText(array)
	}

	stream := getIOStream(cmd_flags.file)
	lines := make([]string, 0)
	for stream.Scan() {
		lines = append(lines, stream.Text())
	}

	return normalizeText(lines)
}

func normalizeText(lines []string) ([]string, int) {
	foundLineLenght := 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) > maxLineLenght {
			foundLineLenght = maxLineLenght

			//splitting the line
			line, newLine := lines[i][:maxLineLenght], lines[i][maxLineLenght:]

			//creating space for the next line
			lines = append(lines[:i+1], lines[i:]...)

			//insert the new line
			lines[i] = line
			lines[i+1] = newLine
		}
		if len(lines[i]) > foundLineLenght {
			foundLineLenght = len(lines[i])
		}
	}

	return lines, foundLineLenght
}

func getIOStream(fileName string) *bufio.Scanner {
	if fileName != "" {
		file, err := os.Open(fileName)

		if err != nil {
			log.Panicf("Could not open file to read, check your permissions:\n%e", err)
		}

		return bufio.NewScanner(file)
	}

	return bufio.NewScanner(os.Stdin)
}
