package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sockMerchant(n int32, ar []int32) int32 {
	var pairCount int32 = 0
	colorsSeen := map[int32]bool{}
	var i int32
	for i = 0; i < n; i++ {
		_, ok := colorsSeen[ar[i]]
		if ok {
			pairCount++
			delete(colorsSeen, ar[i])
		} else {
			colorsSeen[ar[i]] = true
		}
	}
	return pairCount
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine("Enter number of socks: ", reader), 10, 64)
	checkErr(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine("Enter socks (separated by space): ", reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkErr(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
