package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func haveAnyDigit(str string, stringDigitMap map[string]int) int {
	if strings.Contains(str, "one") {
		return stringDigitMap["one"]
	}
	if strings.Contains(str, "two") {
		return stringDigitMap["two"]
	}
	if strings.Contains(str, "three") {
		return stringDigitMap["three"]
	}
	if strings.Contains(str, "four") {
		return stringDigitMap["four"]
	}
	if strings.Contains(str, "five") {
		return stringDigitMap["five"]
	}
	if strings.Contains(str, "six") {
		return stringDigitMap["six"]
	}
	if strings.Contains(str, "seven") {
		return stringDigitMap["seven"]
	}
	if strings.Contains(str, "eight") {
		return stringDigitMap["eight"]
	}
	if strings.Contains(str, "nine") {
		return stringDigitMap["nine"]
	}

	return -1
}

func reverseString(str string) (reverse string) {
	for _, v := range str {
		reverse = string(v) + reverse
	}
	return
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := int64(0)

	stringDigitMap := make(map[string]int)
	stringDigitMap["one"] = 1
	stringDigitMap["two"] = 2
	stringDigitMap["three"] = 3
	stringDigitMap["four"] = 4
	stringDigitMap["five"] = 5
	stringDigitMap["six"] = 6
	stringDigitMap["seven"] = 7
	stringDigitMap["eight"] = 8
	stringDigitMap["nine"] = 9

	for scanner.Scan() {
		line := scanner.Text()
		num := int64(10)
		str := ""
		// logic
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			str += string(runes[i])
			digit := haveAnyDigit(str, stringDigitMap)
			if digit != -1 {
				num *= int64(digit)
				break
			}
			if unicode.IsNumber(runes[i]) {
				num *= int64(runes[i] - 48)
				break
			}
		}

		str2 := ""
		for i := len(runes) - 1; i >= 0; i-- {
			str2 += string(runes[i])
			rev := reverseString(str2)
			digit := haveAnyDigit(rev, stringDigitMap)
			if digit != -1 {
				num += int64(digit)
				break
			}
			if unicode.IsNumber(runes[i]) {
				num += int64(runes[i] - 48)
				break
			}
		}
		// fmt.Printf("num = %d\n", num)
		sum += int64(num)
	}

	fmt.Printf("sum = %d", sum)
}
