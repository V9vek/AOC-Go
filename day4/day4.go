package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := int64(0)
	mp2 := make(map[int]int)
	mp := make(map[int]int)

	cardNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		cards := strings.Split(line, ":")
		wh := strings.Split(cards[1], "|")
		win := strings.Fields(wh[0])
		have := strings.Fields(wh[1])

		cnt := 0
		for _, h := range have {
			for _, w := range win {
				if h == w {
					cnt++
				}
			}
		}

		mp[cardNum] = cnt

		sum := 1
		for key, val := range mp2 {
			if (key + mp[key]) >= cardNum {
				sum += val
			}
		}
		mp2[cardNum] = sum
		ans += int64(mp2[cardNum])

		cardNum++
	}

	fmt.Printf("ans = %d\n", ans)
}
