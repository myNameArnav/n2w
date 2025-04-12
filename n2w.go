package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var units = []string{
	"zero", "one", "two", "three", "four", "five", "six", "seven", "eight",
	"nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen",
	"sixteen", "seventeen", "eighteen", "nineteen",
}

var tens = []string{
	"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy",
	"eighty", "ninety",
}

var thousands = []string{
	"", "thousand", "million", "billion", "trillion",
}

func convertChunk(number int) string {
	if number == 0 {
		return ""
	}
	parts := []string{}
	if number >= 100 {
		parts = append(parts, units[number/100], "hundred")
		number %= 100
	}
	if number > 0 {
		if number < 20 {
			parts = append(parts, units[number])
		} else {
			tensPart := tens[number/10]
			unitPart := number % 10
			if unitPart > 0 {
				parts = append(parts, fmt.Sprintf("%s-%s", tensPart, units[unitPart]))
			} else {
				parts = append(parts, tensPart)
			}
		}
	}
	return strings.Join(parts, " ")
}

func numberToWords(number int64) (string, error) {
	if number == 0 {
		return units[0], nil
	}
	if number < 0 {
		absWords, err := numberToWords(-number)
		if err != nil {
			return "", err
		}
		return "negative " + absWords, nil
	}
	parts := []string{}
	chunkIndex := 0
	for number > 0 {
		if number%1000 != 0 {
			chunk := int(number % 1000)
			chunkWords := convertChunk(chunk)
			if chunkIndex >= len(thousands) {
				return "", fmt.Errorf("number too large, exceeds defined thousands scale")
			}
			scaleWord := thousands[chunkIndex]
			if scaleWord != "" {
				parts = append(parts, fmt.Sprintf("%s %s", chunkWords, scaleWord))
			} else {
				parts = append(parts, chunkWords)
			}
		}
		number /= 1000
		chunkIndex++
	}
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}
	return strings.TrimSpace(strings.Join(parts, " ")), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		numStr := strings.TrimSpace(line)

		if numStr == "" {
			continue
		}

		number, err := strconv.ParseInt(numStr, 10, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Invalid integer input '%s' from stdin: %v\n", numStr, err)
			os.Exit(1)
		}

		words, err := numberToWords(number)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		if len(words) > 0 {
			words = strings.ToUpper(string(words[0])) + words[1:]
		}

		fmt.Println(words)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading standard input: %v\n", err)
		os.Exit(1)
	}
}
