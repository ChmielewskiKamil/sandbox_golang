package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 50
	height = 50
)

type Grid [][]bool

// Create a new empty grid (all cells dead)
func newGrid() Grid {
	grid := make(Grid, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}

// Randomly initialize the grid with live cells
func (g Grid) randomize() {
	for y := range g {
		for x := range g[y] {
			g[y][x] = rand.Float64() < 0.3 // 30% chance of a live cell
		}
	}
}

// Display the grid in the terminal
func (g Grid) print() {
	for _, row := range g {
		for _, cell := range row {
			if cell {
				fmt.Print("O ") // Alive
			} else {
				fmt.Print(". ") // Dead
			}
		}
		fmt.Println()
	}
}

// Get the number of alive neighbors for a given cell
func (g Grid) aliveNeighbors(x, y int) int {
	neighbors := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			nx, ny := x+i, y+j
			if nx >= 0 && nx < width && ny >= 0 && ny < height && g[ny][nx] {
				neighbors++
			}
		}
	}
	return neighbors
}

// Update the grid based on the rules of the game
func (g Grid) update() Grid {
	newGrid := newGrid()
	for y := range g {
		for x := range g[y] {
			aliveNeighbors := g.aliveNeighbors(x, y)
			if g[y][x] {
				// Cell is alive: survives with 2 or 3 neighbors
				newGrid[y][x] = aliveNeighbors == 2 || aliveNeighbors == 3
			} else {
				// Cell is dead: becomes alive with exactly 3 neighbors
				newGrid[y][x] = aliveNeighbors == 3
			}
		}
	}
	return newGrid
}

// Clear the terminal (ANSI escape code)
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	grid := newGrid()
	grid.randomize()

	for {
		clearScreen()
		grid.print()
		grid = grid.update()
		time.Sleep(300 * time.Millisecond) // Control the speed of the simulation
	}
}
