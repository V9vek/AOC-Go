package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	cfg  string
	nums []int
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	score := 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		cfg := split[0]
		cfg = part2c(cfg)
		fmt.Println(cfg)

		nums := getInts(strings.Split(split[1], ","))
		nums = part2n(nums)
		fmt.Println(nums)

		m := map[*Pair]int{}
		ans := helper(cfg, nums, m)
		fmt.Println(line)
		fmt.Println(ans)
		score += ans
	}

	fmt.Printf("%#v\n", score)
}

func part2c(cfg string) string {
	o := cfg
	for i := 0; i < 4; i++ {
		cfg += "?"
		cfg += o
	}
	return cfg
}

func part2n(nums []int) []int {
	o := nums
	for i := 0; i < 4; i++ {
		for _, val := range o {
			nums = append(nums, val)
		}
	}
	return nums
}

func helper(cfg string, nums []int, m map[*Pair]int) int {
	// fmt.Println(cfg)
	// fmt.Println()
	if val, ok := m[&Pair{cfg, nums}]; ok {
		return val
	}

	if len(cfg) == 0 {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if cfg[0] == '.' {
		return helper(cfg[1:], nums, m)
	}

	if cfg[0] == '#' {
		ans := hs(cfg, nums, m)
		m[&Pair{cfg, nums}] = ans
		return ans
		// if nums[0] <= len(cfg) && !strings.Contains(cfg[:nums[0]], ".") && (nums[0] == len(cfg) || cfg[nums[0]] != '#') {
		// 	result += helper(cfg[nums[0]+1:], nums[1:])
		// }
	}

	if cfg[0] == '?' {
		// ? -> .
		// ? -> #
		ans := helper(cfg[1:], nums, m) + hs(cfg, nums, m)
		m[&Pair{cfg, nums}] = ans
		return ans
	}

	return 0
}

func hs(cfg string, nums []int, m map[*Pair]int) int {
	if val, ok := m[&Pair{cfg, nums}]; ok {
		return val
	}

	if len(nums) == 0 {
		return 0
	}
	if len(cfg) < nums[0] {
		return 0
	}
	for i := 0; i < nums[0]; i++ {
		if cfg[i] == '.' {
			return 0
		}
	}
	if len(cfg) == nums[0] {
		if len(nums) == 1 {
			return 1
		}
		return 0
	}
	if cfg[nums[0]] == '#' {
		return 0
	}
	return helper(cfg[nums[0]+1:], nums[1:], m)
}

func getInts(a []string) []int {
	ints := []int{}
	for _, val := range a {
		v, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		ints = append(ints, v)
	}
	return ints
}
