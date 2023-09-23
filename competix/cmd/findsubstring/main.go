package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	me "github.com/toshim45/competix"
)

func main() {
	debugMode := os.Getenv("DEBUG") == "true"
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	//first line
	raw := me.ReadLineString(reader)

	//second line
	find := me.ReadLineString(reader)

	startTime := time.Now()
	result := findSubstring(raw, find)

	if debugMode {
		fmt.Println("execution time: ", time.Since(startTime))
	}

	fmt.Fprintf(writer, "%s\n", result)
	writer.Flush()
}

func findSubstring(raw, find string) string {
	lenInputRaw := len(raw)
	lenInputFind := len(find)

	curr := lenInputRaw
	for i := 0; i <= curr-lenInputFind; i++ {
		validStartCount := 0
		for j := 0; j < lenInputFind; j++ {
			if raw[i+j] == find[j] {
				validStartCount++
			}
		}
		if validStartCount == lenInputFind {
			return "found"
		}

		validEndCount := 0
		for j := 0; j < lenInputFind; j++ {
			// log.Printf("compare: %c with %c \r\n", raw[curr-lenInputFind+j], find[j])
			if raw[curr-lenInputFind+j] == find[j] {
				validEndCount++
			}
		}

		if validEndCount == lenInputFind {
			return "found"
		}
		curr--

		// log.Print("===")
	}

	return "not found"
}
