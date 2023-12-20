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

	tim := int64(0)
	dst := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")
		if splits[0] == "Time" {
			tim = getInt(strings.ReplaceAll(splits[1], " ", ""))
		}
		if splits[0] == "Distance" {
			dst = getInt(strings.ReplaceAll(splits[1], " ", ""))
		}
	}
	// fmt.Println(tim)
	// fmt.Println(dst)

	score := int64(0)
	for i := int64(1); i <= tim/2; i++ {
		newDst := i * (tim - i)
		if newDst > dst {
			score++
		}
	}

	score *= 2
	if score%2 == 0 {
		score--
	}
	fmt.Printf("%v", score)
}

func getInt(str string) int64 {
	val, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return int64(val)
}

func getInts(strs []string) []int {
	ints := []int{}
	for _, str := range strs {
		val, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, val)
	}
	return ints
}
