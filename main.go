package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func remove(arr []string, index, count int) {
	temp := append(arr[:index], arr[index+count:]...)
	fmt.Println("temp: ", temp)
	arr = temp
}

func FilterArr(arr []string) []string {
	var filteredArr []string
	for _, s := range arr {
		if s != "" {
			filteredArr = append(filteredArr, s)
		}
	}
	return filteredArr
}

func CheckArgsAndRun(s []string) {
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			result := ToDecimal(s[i-1], 16)
			s[i-1] = result
			remove(s, i, 1)
			s[len(s)-1] = ""
		} else if s[i] == "(bin)" {
			result := ToDecimal(s[i-1], 2)
			s[i-1] = result
			remove(s, i, 1)
			s[len(s)-1] = ""
		}
	}
}

func ToDecimal(s string, base int) string {

	num, err := strconv.ParseInt(s, base, 64)

	if err != nil {
		fmt.Println("Error: Can't convert text to a num!")
	}

	str := strconv.Itoa(int(num))
	return str
}

func DataHandler() []string {

	if len(os.Args) > 3 {
		fmt.Println("Too many arguments.")
		return nil
	} else if len(os.Args) == 2 {
		fmt.Println("You need one more argument for result text file.")
		return nil
	} else if len(os.Args) != 1 {

		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("File Not Found!")
		}
		dataArr := strings.Fields(string(data))
		fmt.Println("original: ", dataArr)
		return dataArr
	} else {
		fmt.Println("Error missing file name and result text name.")
		return nil
	}
}

func main() {
	initialArr := DataHandler()
	CheckArgsAndRun(initialArr)
	finalArr := FilterArr(initialArr)
	fmt.Println("final: ", finalArr)
	os.WriteFile("result.txt", []byte(strings.Join(finalArr, " ")), 0644)
}
