package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cnt := 0
	seeds := make([]int64, 0)
	ss := make([][]int64, 0)
	sf := make([][]int64, 0)
	fw := make([][]int64, 0)
	wl := make([][]int64, 0)
	lt := make([][]int64, 0)
	th := make([][]int64, 0)
	hl := make([][]int64, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			cnt++
			continue
		}

		if cnt == 0 && unicode.IsNumber(rune(line[0])) {
			s := strings.Split(line, " ")
			for _, val := range s {
				i, _ := strconv.Atoi(val)
				seeds = append(seeds, int64(i))
			}
		}

		if cnt > 0 && unicode.IsNumber(rune(line[0])) {
			s := strings.Split(line, " ")
			temp := make([]int64, 0)
			for _, val := range s {
				i, _ := strconv.Atoi(val)
				temp = append(temp, int64(i))
			}
			if cnt == 1 {
				ss = append(ss, temp)
			} else if cnt == 2 {
				sf = append(sf, temp)
			} else if cnt == 3 {
				fw = append(fw, temp)
			} else if cnt == 4 {
				wl = append(wl, temp)
			} else if cnt == 5 {
				lt = append(lt, temp)
			} else if cnt == 6 {
				th = append(th, temp)
			} else if cnt == 7 {
				hl = append(hl, temp)
			}
		}
	}

	ans := int64(math.MaxInt64)
	for _, seed := range seeds {
		loc := seed
		for _, row := range ss {
			if loc <= (row[1]+row[2]) && loc >= row[1] {
				loc += (row[0] - row[1])
				break
			}
		}
		// fmt.Println(loc)
		for _, row := range sf {
			if loc <= (row[1]+row[2]) && loc >= row[1] {
				loc += (row[0] - row[1])
				break
			}
		}
		// fmt.Println(loc)
		for _, row := range fw {
			if loc <= (row[1]+row[2]) && loc >= row[1] {
				loc += (row[0] - row[1])
				break
			}
		}
		// fmt.Println(loc)
		for _, row := range wl {
			if loc <= (row[1]+row[2]) && loc >= row[1] {
				loc += (row[0] - row[1])
				break
			}
		}
		// fmt.Println(loc)
		for _, row := range lt {
			if loc <= (row[1]+row[2]) && loc >= row[1] {
				loc += (row[0] - row[1])
				break
			}
		}
		// fmt.Println(loc)
		for _, row := range th {
			if loc <= (row[1]+row[2]) && loc >= row[1] {
				loc += (row[0] - row[1])
				break
			}
		}
		// fmt.Println(loc)
		for _, row := range hl {
			if loc <= (row[1]+row[2]) && loc >= row[1] {
				loc += (row[0] - row[1])
				break
			}
		}

		fmt.Printf("seed = %d loc = %d\n", seed, loc)
		ans = min(ans, loc)
	}

	fmt.Printf("ans = %d", ans)

	// fmt.Println(seeds)
	// fmt.Println(ss)
	// fmt.Println(sf)
	// fmt.Println(fw)
	// fmt.Println(wl)
	// fmt.Println(lt)
	// fmt.Println(th)
	// fmt.Println(hl)
}
