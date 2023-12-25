package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInput(path string) [][]string {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	patterns := [][]string{}
	pattern := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			patterns = append(patterns, pattern)
			pattern = nil
			continue
		}
		pattern = append(pattern, line)
	}
	patterns = append(patterns, pattern)
	return patterns
}

func main() {
	patterns := getInput("./input.txt")
	score := 0
	r, c := 0, 0

	for _, pattern := range patterns {
		mr := map[int]string{}
		for i := 0; i < len(pattern); i++ {
			mr[i] = pattern[i]
		}

		mc := map[int]string{}
		for j := 0; j < len(pattern[0]); j++ {
			col := ""
			for i := 0; i < len(pattern); i++ {
				col += string(pattern[i][j])
			}
			mc[j] = col
		}

		// row mirror
		for i := 0; i < len(pattern)-1; i++ {
			up := i
			down := i + 1
			found := false
			for up >= 0 && down < len(pattern) {
				if mr[up] != mr[down] {
					if count(mr[up], mr[down]) {
						found = true
						up--
						down++
					}
					break
				}
				up--
				down++
			}
			if found && (up < 0 || down >= len(pattern)) {
				r += (i + 1)
				break
			}
		}

		// col mirror
		for j := 0; j < len(pattern[0])-1; j++ {
			left := j
			right := j + 1
			found := false
			for left >= 0 && right < len(pattern[0]) {
				if mc[left] != mc[right] {
					if count(mc[left], mc[right]) {
						found = true
						left--
						right++
					}
					break
				}
				left--
				right++
			}
			if found && (left < 0 || right >= len(pattern[0])) {
				c += (j + 1)
				break
			}
		}
		// fmt.Println(r, c)
	}

	score = r*100 + c
	fmt.Printf("%#v \n", score)

}

func count(s1 string, s2 string) bool {
	cnt := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			cnt++
		}
	}
	return cnt == 1
}
