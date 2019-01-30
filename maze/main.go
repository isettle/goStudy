package main

import (
	"fmt"
	"os"
)

type point struct {
	i, j int
}

var dires = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) move(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	// 行越界
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	// 列越界
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	var row, col int

	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	queue := []point{start}

	for len(queue) > 0 {
		// 队头
		cur := queue[0]
		queue = queue[1:]

		if cur == end {
			break
		}

		for _, dire := range dires {
			next := cur.move(dire)

			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val2, ok2 := next.at(steps)
			if !ok2 || val2 != 0 {
				continue
			}

			if next == start {
				continue
			}

			curStep, _ := cur.at(steps)
			steps[next.i][next.j] = curStep + 1
			queue = append(queue, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze.in")

	for _, row := range maze {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}

	fmt.Println("--")
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
