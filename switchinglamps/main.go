package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func switchLamps(n int32, trip int32) int32 {
	var count int32
	lamps := make([]bool, n)
	for i := 1; i <= int(trip); i++ {
		for j := i - 1; j < int(n); j += i {
			lamps[j] = !lamps[j]
		}
	}

	for _, lamp := range lamps {
		if lamp {
			count++
		}
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine("Enter number of lamps: ", reader), 10, 64)
	checkError(err)

	n := int32(nTemp)

	tripsTemp, err := strconv.ParseInt(readLine("Enter number of trips: ", reader), 10, 64)
	checkError(err)

	trips := int32(tripsTemp)

	result := switchLamps(n, trips)

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
