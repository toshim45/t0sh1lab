package competix

import (
	"bufio"
	"io"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadLineString(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
