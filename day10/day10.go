package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type Tuple struct {
	x, y, val int
}

func enqueue(queue []Tuple, element Tuple) []Tuple {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}
func dequeue(queue []Tuple) (Tuple, []Tuple) {
	element := queue[0]       // The first element is the one to be dequeued.
	return element, queue[1:] // Slice off the element once it is dequeued.
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matrix := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		temp := []string{}
		for _, r := range line {
			temp = append(temp, string(r))
		}
		matrix = append(matrix, temp)
	}

	// output(matrix)

	sx, sy := getStart(matrix)
	nrow := len(matrix)
	ncol := len(matrix[0])

	valueMatrix := make([][]int, nrow)
	for i := range valueMatrix {
		valueMatrix[i] = make([]int, ncol)
	}

	maxScore := 0
	queue := []Tuple{}
	queue = enqueue(queue, Tuple{sx, sy, 0})

	// L R T B
	dx := []int{0, 0, -1, 1}
	dy := []int{-1, 1, 0, 0}

	maybe_s := []string{"|", "-", "F", "J", "L", "7"}

	for len(queue) != 0 {
		cell, deq := dequeue(queue)
		queue = deq

		x := cell.x
		y := cell.y
		val := cell.val
		valueMatrix[x][y] = val
		maxScore = max(maxScore, val)

		if matrix[x][y] == "S" {
			// LRTB
			nx := x + dx[0]
			ny := y + dy[0]
			str := "FL-"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
				for _, s := range str {
					if slices.Contains(maybe_s, string(s)) {
						maybe_s = remove(string(s), maybe_s)
					}
				}
			}

			nx = x + dx[1]
			ny = y + dy[1]
			str = "-J7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
				for _, s := range str {
					if slices.Contains(maybe_s, string(s)) {
						maybe_s = remove(string(s), maybe_s)
					}
				}
			}

			nx = x + dx[2]
			ny = y + dy[2]
			str = "|F7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
				for _, s := range str {
					if slices.Contains(maybe_s, string(s)) {
						maybe_s = remove(string(s), maybe_s)
					}
				}
			}

			nx = x + dx[3]
			ny = y + dy[3]
			str = "|LJ"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
				for _, s := range str {
					if slices.Contains(maybe_s, string(s)) {
						maybe_s = remove(string(s), maybe_s)
					}
				}
			}
		}
		if matrix[x][y] == "|" {
			// T B
			nx := x + dx[2]
			ny := y + dy[2]
			str := "|F7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}

			nx = x + dx[3]
			ny = y + dy[3]
			str = "|LJ"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}
		}
		if matrix[x][y] == "-" {
			// L R
			nx := x + dx[0]
			ny := y + dy[0]
			str := "FL-"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}

			nx = x + dx[1]
			ny = y + dy[1]
			str = "-J7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}
		}
		if matrix[x][y] == "L" {
			// T R
			nx := x + dx[2]
			ny := y + dy[2]
			str := "|F7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}

			nx = x + dx[1]
			ny = y + dy[1]
			str = "-J7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}
		}
		if matrix[x][y] == "J" {
			// L T
			nx := x + dx[0]
			ny := y + dy[0]
			str := "FL-"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}

			nx = x + dx[2]
			ny = y + dy[2]
			str = "|F7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}
		}
		if matrix[x][y] == "7" {
			// L B
			nx := x + dx[0]
			ny := y + dy[0]
			str := "FL-"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}

			nx = x + dx[3]
			ny = y + dy[3]
			str = "|LJ"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}
		}
		if matrix[x][y] == "F" {
			// R B
			nx := x + dx[1]
			ny := y + dy[1]
			str := "-J7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}

			nx = x + dx[3]
			ny = y + dy[3]
			str = "|LJ"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}
		}
	}

	fmt.Println("maybe_s = ", maybe_s)
	matrix[sx][sy] = maybe_s[0]

	// output2(valueMatrix)
	// fmt.Println(maxScore)

	// part2
	vis := make([][]int, nrow)
	for i := range vis {
		vis[i] = make([]int, ncol)
	}
	for i := 0; i < nrow; i++ {
		if valueMatrix[i][0] == 0 && vis[i][0] == 0 {
			dfs(i, 0, valueMatrix, &vis, nrow, ncol)
		}
	}
	for i := 0; i < nrow; i++ {
		if valueMatrix[i][ncol-1] == 0 && vis[i][ncol-1] == 0 {
			dfs(i, ncol-1, valueMatrix, &vis, nrow, ncol)
		}
	}
	for j := 0; j < ncol; j++ {
		if valueMatrix[0][j] == 0 && vis[0][j] == 0 {
			dfs(0, j, valueMatrix, &vis, nrow, ncol)
		}
	}
	for j := 0; j < ncol; j++ {
		if valueMatrix[nrow-1][j] == 0 && vis[nrow-1][j] == 0 {
			dfs(nrow-1, j, valueMatrix, &vis, nrow, ncol)
		}
	}

	// output2(vis)

	count := 0
	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			if vis[i][j] == 0 && valueMatrix[i][j] == 0 && checkInside(i, j, ncol, valueMatrix) {
				count++
			}
		}
	}

	fmt.Println(count)
}

func checkInside(i int, j int, ncol int, valueMatrix [][]int) bool {
	// fmt.Printf("i = %d, j = %d", i, j)
	cnt := 0
	for ; j < ncol; j++ {
		if valueMatrix[i][j] != 0 {
			cnt++
		}
	}
	// fmt.Printf(" cnt = %d\n", cnt)
	return cnt%2 != 0
}

func dfs(i int, j int, valueMatrix [][]int, vis *[][]int, nrow int, ncol int) {
	if (*vis)[i][j] == 1 {
		return
	}
	(*vis)[i][j] = 1

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, 1, -1}

	for k := 0; k < 4; k++ {
		nx := i + dx[k]
		ny := j + dy[k]
		if isValid(nx, ny, nrow, ncol) && (*vis)[nx][ny] == 0 && valueMatrix[nx][ny] == 0 {
			dfs(nx, ny, valueMatrix, vis, nrow, ncol)
		}
	}
}

func isValid(nx int, ny int, nrow int, ncol int) bool {
	return nx >= 0 && nx < nrow && ny >= 0 && ny < ncol
}

func getStart(matrix [][]string) (int, int) {
	for i, row := range matrix {
		for j, col := range row {
			if col == "S" {
				return i, j
			}
		}
	}
	return -1, -1
}

func output(matrix [][]string) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func output2(matrix [][]int) {
	for _, row := range matrix {
		for _, col := range row {
			fmt.Printf("%v ", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func remove(name string, nations []string) []string {
	i := 0
	for idx, item := range nations {
		if item != name {
			nations[i] = nations[idx]
			i++
		}
	}
	return nations[:i]
}
