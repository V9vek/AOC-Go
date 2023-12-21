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

		k := 0
		for {
			for i := len(history) - 1; i > k; i-- {
				history[i] = history[i] - history[i-1]
			}
			k++
			if allEqual(history, k) {
				break
			}
		}

		prediction := int64(0)
		for i := k; i >= 0; i-- {
			prediction = history[i] - prediction
		}
		score += prediction
	}

	fmt.Println(score)
}

func allEqual(arr []int64, length int) bool {
	for i := len(arr) - 1; i > length; i-- {
		if arr[i] != arr[i-1] {
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
