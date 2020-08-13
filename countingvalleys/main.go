package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func countingValleys(n int32, s string) int32 {
	paths := strings.Replace(s, " ", "", -1)
	var sum int = 0
	var count int32 = 0
	var i int32
	for i = 0; i < n; i++ {
		if paths[i] == 'U' {
			sum++
			if sum == 0 {
				count++
			}
		} else {
			sum--
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine("Number of steps: ", reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine("Paths: ", reader)

	result := countingValleys(n, s)

	fmt.Fprintf(writer, "%d\n", result)

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
