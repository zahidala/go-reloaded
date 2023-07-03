package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CheckArgsAndRun(s []string) {
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			ToDecimal(s[i-1], 16)
		} else if s[i] == "(bin)" {
			ToDecimal(s[i-1], 2)
		}
	}
}

func ToDecimal(s string, base int) {

	num, err := strconv.ParseInt(s, base, 64)

	if err != nil {
		fmt.Println("Error: Can't convert text to a num!")
	}

	fmt.Println(num)
}

func DataHandler() []string {
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("File Not Found!")
	}
	dataArr := strings.Fields(string(data))
	// fmt.Println(dataArr)
	return dataArr
}

func main() {
	arr := DataHandler()
	CheckArgsAndRun(arr)
}
