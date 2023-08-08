package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	if len(s) != 1 {
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

			AOrAnChecker(s, i)
		}

		s = QuoteFixer(s)
		s = PunctuationFixer(s)
		finalArr := FilterArr(s)
		fmt.Println("final:", finalArr)
		s = finalArr
		os.WriteFile(os.Args[2], []byte(strings.Join(s, " ")), 0644)
		fmt.Println("<INFO> All done! Please check results.txt for final results.")
	} else {
		fmt.Println("<ERROR> Need to pass a word and an argument to start!")
	}
}

func ToDecimal(s string, base int) string {

	num, err := strconv.ParseInt(s, base, 64)

	if err != nil {
		fmt.Printf("<ERROR> Can't convert \"%s\" to a number!\n", s)
		os.Exit(0)
	}

	str := strconv.Itoa(int(num))
	return str
}

func ToStringMethod(s []string, i int, fn func(string) string) {

	if strings.Contains(s[i], "(up,") || strings.Contains(s[i], "(low,") || strings.Contains(s[i], "(cap,") {
		num, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
		if (num - 1) > i-1 {
			fmt.Println("<ERROR> Not enough words to convert according to the number passed.")
			os.Exit(0)
		}

		if err != nil {
			fmt.Printf("<ERROR> \"%s\" is a invalid number that is passed to the argument.\n", s[i+1][:len(s[i+1])-1])
			os.Exit(0)
		}

		for index := i - 1; index >= i-num; index-- {
			s[index] = strings.ToLower(s[index])
			s[index] = fn(s[index])
		}
		remove(s, i, num)

	} else {
		s[i-1] = strings.ToLower(s[i-1])
		s[i-1] = fn(s[i-1])
		remove(s, i, 1)
	}
}

func PunctuationFixer(arr []string) []string {

	// For cases where theres space before punctuation.
	firstCheck := regexp.MustCompile(`\s([\s.,!?:;]+)[.,!?:;]*`)
	fixedString := firstCheck.ReplaceAllString(strings.Join(arr, " "), "$1")

	// For cases where theres no space after punctuation.
	secondCheck := regexp.MustCompile(`([.,!?:;]+)([^.,!?;'])`)
	finalString := secondCheck.ReplaceAllString(fixedString, "$1 $2")

	finalArr := strings.Fields(finalString)
	return finalArr
}

func QuoteFixer(arr []string) []string {
	regex := regexp.MustCompile(`[']\s*(.*?)\s*[']\s*`)
	fixedString := regex.ReplaceAllString(strings.Join(arr, " "), "'$1' ")
	finalArr := strings.Fields(fixedString)
	return finalArr
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
		fmt.Println("<ERROR> Too many arguments.")
		os.Exit(0)
		return nil
	case len(os.Args) == 2:
		fmt.Println("<ERROR> You need one more argument for result text file.")
		os.Exit(0)
		return nil
	case len(os.Args) == 3:
		if !strings.HasSuffix(os.Args[1], ".txt") {
			fmt.Println("<ERROR> Sample file name must end in .txt!")
			os.Exit(0)
		}

		data, err := os.ReadFile(os.Args[1])

		if err != nil {
			fmt.Println("<ERROR> File Not Found!")
			os.Exit(0)
		}

		if !strings.HasSuffix(os.Args[2], ".txt") {
			fmt.Println("<ERROR> Result file name must end in .txt!")
			os.Exit(0)
		}

		dataArr := strings.Fields(string(data))

		if len(dataArr) == 0 && err == nil {
			fmt.Println("<ERROR> No text found in the sample file! Exiting the program.")
			os.Exit(0)
		}

		fmt.Println("original: ", dataArr)
		return dataArr
	default:
		fmt.Println("<ERROR> Missing file name and result text name.")
		os.Exit(0)
		return nil
	}
}

func main() {
	arr := DataHandler()
	CheckArgsAndRun(arr)
}
