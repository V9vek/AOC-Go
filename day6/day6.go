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

	tim := []int{}
	dst := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")
		if splits[0] == "Time" {
			tim = getInts(strings.Fields(splits[1]))
		}
		if splits[0] == "Distance" {
			dst = getInts(strings.Fields(splits[1]))
		}
	}
	// fmt.Println(tim)
	// fmt.Println(dst)

	score := 1
	for i := 0; i < len(tim); i++ {
		cnt := 0
		for j := 1; j <= tim[i]; j++ {
			newDst := j * (tim[i] - j)
			if newDst > dst[i] {
				cnt++
			}
		}
		score *= cnt
	}

	fmt.Printf("%v", score)
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
