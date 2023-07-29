package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func remove(arr []string, index, count int) {
	switch arr[index] {
	case "(up)", "(low)", "(cap)", "(bin)", "(hex)":
		arr[index] = ""
	}

	if strings.Contains(arr[index], "(up,") || strings.Contains(arr[index], "(low,") || strings.Contains(arr[index], "(cap,") {
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

		PunctuationFixer(s, i)
		AOrAnChecker(s, i)
	}

	finalArr := FilterArr(s)
	fmt.Println("final:", finalArr)
	s = finalArr
}

func ToDecimal(s string, base int) string {

	num, err := strconv.ParseInt(s, base, 64)

	if err != nil {
		fmt.Println("<ERROR> Can't convert text to a number!")
		return "<ERROR>"
	}

	str := strconv.Itoa(int(num))
	return str
}

func ToStringMethod(s []string, i int, fn func(string) string) {

	if strings.Contains(s[i], "(up,") || strings.Contains(s[i], "(low,") || strings.Contains(s[i], "(cap,") {
		num, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
		for index := i - 1; index >= i-num; index-- {
			s[index] = fn(s[index])
		}
		remove(s, i, num)

		if err != nil {
			fmt.Println("<ERROR> Invalid number passed to argument...Converting just the word before the argument instead!")
			s[i-1] = fn(s[i-1])
			remove(s, i, 1)
		}

	} else {
		s[i-1] = fn(s[i-1])
		remove(s, i, 1)
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
				// fmt.Println("punc", string(puncRecheckArr))
				if unicode.IsPunct(puncArr[0]) {
					// fmt.Println(string(getEachPunc[0]))
					arr[i-1] = arr[i-1] + string(getEachPunc[0])
					removeExtraPunc := getEachPunc[1:]
					// fmt.Println(removeExtraPunc)
					arr[i] = removeExtraPunc
				}
			} else {
				fmt.Println(string(c))
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
	arr := DataHandler()
	CheckArgsAndRun(arr)
	os.WriteFile("result.txt", []byte(strings.Join(arr, " ")), 0644)
	// fmt.Println("<INFO> All done! Please check results.txt for final results.")
}
