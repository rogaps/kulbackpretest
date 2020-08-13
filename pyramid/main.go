package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func pyramidNumbers(s string) string {
	var builder strings.Builder
	re, _ := regexp.Compile("[^0-9]+")
	numStr := re.ReplaceAllString(s, "")
	numLen := len(numStr)

	for i, num := range numStr {
		builder.WriteRune(num)
		builder.WriteString(strings.Repeat("0", numLen-(i+1)))
		builder.WriteRune('\n')
	}

	return builder.String()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	s := readLine("Enter input: ", reader)

	result := pyramidNumbers(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(prompt string, reader *bufio.Reader) string {
	fmt.Print(prompt)
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
