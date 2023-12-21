package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := int64(0)

	for scanner.Scan() {
		line := scanner.Text()

		history := getInts(line)

		// 1 3 6 10 15 21
		// 2 3 4 5 6 21
		// 1 1 1 1 6 21
		// 0 0 0 1 6 21

		k := 0
		for {
			for i := 1; i < len(history)-k; i++ {
				history[i-1] = history[i] - history[i-1]
			}
			k++
			if allEqual(history, len(history)-k) {
				break
			}
		}

		prediction := int64(0)
		for i := len(history) - 1; i >= len(history)-k-1; i-- {
			prediction += history[i]
		}
		score += prediction
	}

	fmt.Println(score)
}

func allEqual(arr []int64, length int) bool {
	for i := 1; i < length; i++ {
		if arr[i-1] != arr[i] {
			return false
		}
	}
	return true
}

func getInts(str string) []int64 {
	nums := strings.Split(str, " ")
	numInts := []int64{}
	for _, num := range nums {
		val, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}
		numInts = append(numInts, int64(val))
	}
	return numInts
}
