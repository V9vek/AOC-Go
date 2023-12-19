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

func check2(matrix [][]rune, i int, j int, valMap map[Pair]int) int64 {
	dx := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	maxy := len(matrix)
	maxx := len(matrix[0])

	mp := make(map[int]int)
	for k := 0; k < 8; k++ {
		nx := i + dx[k]
		ny := j + dy[k]
		if isValidCell(nx, ny, maxx, maxy) && matrix[nx][ny] != '.' && unicode.IsNumber(matrix[nx][ny]) {
			// fmt.Printf("i = %d, j = %d ", nx, ny)
			mp[valMap[Pair{i: nx, j: ny}]]++
		}
	}

	fmt.Println("upar map = ", mp)

	mul := int64(1)
	if len(mp) == 2 {
		for k := range mp {
			mul *= int64(k)
		}
		return mul
	}
	return 0
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
	mp := make(map[Pair]int)
	for i := 0; i < len(matrix); i++ {
		num := ""
		found := false
		pairs := make([]Pair, 0)
		for j, char := range matrix[i] {
			if unicode.IsNumber(char) {
				num += string(char)
				pairs = append(pairs, Pair{i: i, j: j})
				if !found && check(matrix, i, j) {
					found = true
				}
			} else {
				if found && num != "" {
					fmt.Printf("%v ", num)
					val, _ := strconv.Atoi(num)
					// ans += int64(val)

					for _, p := range pairs {
						mp[p] = val
					}

					pairs = nil
				}
				num = ""
				found = false
				pairs = nil
			}
		}
		if found && num != "" {
			fmt.Printf("%v ", num)
			val, _ := strconv.Atoi(num)
			// ans += int64(val)

			for _, p := range pairs {
				mp[p] = val
			}
		}
		fmt.Println()
	}
	fmt.Println(mp)

	for i := 0; i < len(matrix); i++ {
		for j, char := range matrix[i] {
			if char == '*' {
				ans += check2(matrix, i, j, mp)
			}
		}
	}

	fmt.Printf("ans = %d", ans)
}
