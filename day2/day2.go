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
	ans := int64(0)

	for scanner.Scan() {
		line := scanner.Text()
		game := strings.Split(line, ":")
		colors := strings.FieldsFunc(game[1], Split)
		r, g, b := -1, -1, -1

		for _, item := range colors {
			v := strings.Split(item, " ")
			// fmt.Println(v[1], v[2])
			val, _ := strconv.Atoi(v[1])
			if v[2][0] == 'r' {
				r = max(r, val)
			}
			if v[2][0] == 'g' {
				g = max(g, val)
			}
			if v[2][0] == 'b' {
				b = max(b, val)
			}
		}

		ans += (int64(r) * int64(g) * int64(b))

		// gameid := strings.Split(game[0], " ")
		// fmt.Println(gameid[1])
		// id, _ := strconv.Atoi(gameid[1])
		// fmt.Printf("gameId = %d\n", id)
		// ans += int(id)
		// fmt.Println("--------------------------------")
	}

	fmt.Printf("ans = %d", ans)
}
