package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		num := int64(10)
		// logic
		runes := []rune(line)
		for i := 0; i < len(runes); i++ {
			if unicode.IsNumber(runes[i]) {
				num *= int64(runes[i] - 48)
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsNumber(runes[i]) {
				num += int64(runes[i] - 48)
				break
			}
		}
		fmt.Printf("num = %d\n", num)
		sum += int64(num)
	}

	fmt.Printf("sum = %d", sum)
}
