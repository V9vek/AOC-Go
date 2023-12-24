package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pair struct {
	i, j int64
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	space := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		temp := []string{}
		for _, char := range line {
			temp = append(temp, string(char))
		}
		space = append(space, temp)
	}

	g := []Pair{}
	for i := 0; i < len(space); i++ {
		for j := 0; j < len(space[0]); j++ {
			if space[i][j] != "." {
				g = append(g, Pair{int64(i), int64(j)})
			}
		}
	}

	er := make([]int64, len(space))
	ec := make([]int64, len(space[0]))

	for i := 0; i < len(space); i++ {
		if isAllEmpty(space[i]) {
			er[i] = 999999
		}
	}

	for j := 0; j < len(space[0]); j++ {
		found := false
		for i := 0; i < len(space); i++ {
			if space[i][j] != "." {
				found = true
				break
			}
		}
		if !found {
			ec[j] = 999999
		}
	}

	er = preSum(er)
	ec = preSum(ec)

	for i := 0; i < len(g); i++ {
		x := g[i].i
		y := g[i].j
		g[i].i += er[x]
		g[i].j += ec[y]
	}

	score := int64(0)
	for i := 0; i < len(g); i++ {
		for j := i + 1; j < len(g); j++ {
			score += absDiffInt(g[i].i, g[j].i) + absDiffInt(g[i].j, g[j].j)
		}
	}

	fmt.Printf("%#v\n", score)
}

func preSum(a []int64) []int64 {
	for i := 1; i < len(a); i++ {
		a[i] += a[i-1]
	}
	return a
}

func insert(a []string, index int, value string) []string {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func isAllEmpty(s []string) bool {
	for _, char := range s {
		if char != "." {
			return false
		}
	}
	return true
}

func absDiffInt(x, y int64) int64 {
	if x < y {
		return y - x
	}
	return x - y
}
