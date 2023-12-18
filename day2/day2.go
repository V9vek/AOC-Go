package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Split(r rune) bool {
	return r == ',' || r == ';'
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	ans := int(0)

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ":")
		colors := strings.FieldsFunc(game[1], Split)
		rgbMap := make(map[byte]int)
		invalid := false

		for _, item := range colors {
			v := strings.Split(item, " ")
			// fmt.Println(v[1], v[2])
			val, _ := strconv.Atoi(v[1])
			if (v[2][0] == 'r' && val > 12) || (v[2][0] == 'g' && val > 13) || (v[2][0] == 'b' && val > 14) {
				invalid = true
				break
			}
			// rgbMap[v[2][0]] += val
		}

		fmt.Println(rgbMap)

		if invalid {
			fmt.Println("--------------------------------")
			continue
		}

		gameid := strings.Split(game[0], " ")
		fmt.Println(gameid[1])
		id, _ := strconv.Atoi(gameid[1])
		fmt.Printf("gameId = %d\n", id)
		ans += int(id)
		fmt.Println("--------------------------------")
	}

	fmt.Printf("ans = %d", ans)
}
