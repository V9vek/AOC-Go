package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Pair struct {
	i, j int
}

func isValidCell(nx int, ny int, maxx int, maxy int) bool {
	return nx >= 0 && ny >= 0 && nx < maxx && ny < maxy
}

func check(matrix [][]rune, i int, j int) bool {
	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	maxy := len(matrix)
	maxx := len(matrix[0])

	for k := 0; k < 8; k++ {
		nx := i + dx[k]
		ny := j + dy[k]
		if isValidCell(nx, ny, maxx, maxy) && matrix[nx][ny] != 46 && !unicode.IsNumber(matrix[nx][ny]) {
			// fmt.Printf("i = %d, j = %d ", nx, ny)
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
	}

	ans := int64(0)
	for i := 0; i < len(matrix); i++ {
		num := ""
		found := false
		for j, char := range matrix[i] {
			if unicode.IsNumber(char) {
				num += string(char)
				if !found && check(matrix, i, j) {
					found = true
				}
			} else {
				if found && num != "" {
					fmt.Printf("%v ", num)
					val, _ := strconv.Atoi(num)
					ans += int64(val)
				}
				num = ""
				found = false
			}
		}
		if found && num != "" {
			fmt.Printf("%v ", num)
			val, _ := strconv.Atoi(num)
			ans += int64(val)
		}
		fmt.Println()
	}

	fmt.Printf("ans = %d", ans)
}
