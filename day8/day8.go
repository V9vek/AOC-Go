package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Pair struct {
	left  string
	right string
}

func main() {
	file, err := os.Open("./input2.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	dirs := []int{}
	nodes := map[string]Pair{}
	starts := []string{}
	for scanner.Scan() {
		line := scanner.Text()

		if len(dirs) == 0 {
			for _, lr := range line {
				if lr == 'L' {
					dirs = append(dirs, 0)
				} else {
					dirs = append(dirs, 1)
				}
			}
			continue
		}

		if len(line) != 0 {
			mappings := strings.Split(line, " = ")
			mappings[1] = strings.Replace(mappings[1], "(", "", -1)
			mappings[1] = strings.Replace(mappings[1], ")", "", -1)
			lr := strings.Split(mappings[1], ", ")
			nodes[mappings[0]] = Pair{left: lr[0], right: lr[1]}

			if mappings[0][2] == 'A' {
				starts = append(starts, mappings[0])
			}
		}
	}

	cycles := []int64{}

	for i := 0; i < len(starts); i++ {
		steps := int64(0)
		for true {
			for _, dir := range dirs {
				steps++
				if dir == 0 {
					starts[i] = nodes[starts[i]].left
				} else {
					starts[i] = nodes[starts[i]].right
				}
				if starts[i][2] == 'Z' {
					break
				}
			}
			if starts[i][2] == 'Z' {
				break
			}
		}
		cycles = append(cycles, steps)
	}

	// no. of steps needed to reach first 'Z' for each starting point
	// for all the starting point to reach 'Z', lcm of all the steps would be the answer

	lcm := cycles[0]
	for _, n := range cycles {
		lcm = (lcm * n) / gcd(lcm, n)
	}

	// fmt.Println(cycles)
	fmt.Println(lcm)
}

func gcd(a int64, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func reached(starts []string) bool {
	for _, val := range starts {
		if val[2] != 'Z' {
			return false
		}
	}
	return true
}
