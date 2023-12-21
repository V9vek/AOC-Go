package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Tuple struct {
	x, y, val int
}

func enqueue(queue []Tuple, element Tuple) []Tuple {
	queue = append(queue, element) // Simply append to enqueue.
	// fmt.Println("Enqueued:", element)
	return queue
}
func dequeue(queue []Tuple) (Tuple, []Tuple) {
	element := queue[0] // The first element is the one to be dequeued.
	// fmt.Println("Dequeued:", element)
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
			}

			nx = x + dx[1]
			ny = y + dy[1]
			str = "-J7"
			if isValid(nx, ny, nrow, ncol) && valueMatrix[nx][ny] == 0 && strings.Contains(str, matrix[nx][ny]) {
				queue = enqueue(queue, Tuple{nx, ny, val + 1})
			}

			nx = x + dx[2]
			ny = y + dy[2]
			str = "|F7"
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

	// output2(valueMatrix)
	fmt.Println(maxScore)
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
