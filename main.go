package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func remove(arr []string, index, count int) {
	if arr[index] == "(up)" || arr[index] == "(low)" || arr[index] == "(cap)" || arr[index] == "(bin)" || arr[index] == "(hex)" {
		arr[index] = ""
	} else if strings.Contains(arr[index], "(up,") || strings.Contains(arr[index], "(low,") || strings.Contains(arr[index], "(cap,") {
		arr[index] = ""
		arr[index+1] = ""
	}
}

// Remove empty spaces in slice.

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

		// PunctuationFixer(s, i)

		AOrAnChecker(s, i)

		switch s[i] {

		case "(hex)":
			result := ToDecimal(s[i-1], 16)
			s[i-1] = result
			remove(s, i, 1)

		case "(bin)":
			result := ToDecimal(s[i-1], 2)
			s[i-1] = result
			remove(s, i, 1)

		case "(up)", "(up,":
			ToStringMethod(s, i, strings.ToUpper)

		case "(low)", "(low,":
			ToStringMethod(s, i, strings.ToLower)

		case "(cap)", "(cap,":
			ToStringMethod(s, i, strings.Title)
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

func ToStringMethod(s []string, i int, fn func(string) string) {

	num, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])

	if err != nil {
		s[i-1] = fn(s[i-1])
		remove(s, i, 1)
	} else {
		// fmt.Println("Num converted successfully:", num)
		for index := i - 1; index >= i-num; index-- {
			s[index] = fn(s[index])
			fmt.Println("num", num)
		}
		remove(s, i, num)
	}
}

func PunctuationFixer(arr []string, i int) {

	puncArr := []rune(arr[i])

	for _, c := range puncArr {
		if unicode.IsPunct(c) {
			if c != '\'' {
				// fmt.Println("rune", string(c))
				// arr[i-1] = arr[i-1] + arr[i]
				// arr[i] = ""
				getEachPunc := arr[i]
				fmt.Println(string(getEachPunc[0]))
				arr[i-1] = arr[i-1] + string(getEachPunc[0])
				removeExtraPunc := getEachPunc[1:]
				arr[i] = removeExtraPunc
			}
		}
	}

}

func AOrAnChecker(arr []string, i int) {

	switch strings.ToLower(arr[i]) {
	case "a":
		nextWord := arr[i+1]
		firstLetter := string(nextWord[0])
		if strings.ContainsAny(firstLetter, "aeiouh") {
			arr[i] = arr[i] + "n"
		}
	}
}

func DataHandler() []string {

	switch {
	case len(os.Args) > 3:
		fmt.Println("Too many arguments.")
		return nil
	case len(os.Args) == 2:
		fmt.Println("You need one more argument for result text file.")
		return nil
	case len(os.Args) != 1:
		data, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("File Not Found!")
		}
		dataArr := strings.Fields(string(data))
		fmt.Println("original: ", dataArr)
		return dataArr
	default:
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
