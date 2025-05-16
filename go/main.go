package main

import (
	"fmt"
)


type Grid struct {
	cols, rows int
	grid [][]Cell
}


func (g Grid) InitiateCells() {
	for row:= range g.rows {
		for col:= range g.cols {
			cell := Cell{col, row, 0, false}
			g.grid[row] = append(g.grid[row], cell)
			// cell.col = col
			// cell.row = row
			// cell.neighbours = 0
			// cell.alive = false
		}
	}
}



func (g Grid) DisplayGrid(){
	for row := range g.grid {
		fmt.Println()
		for col := range g.grid[0] {
			cell:= g.grid[row][col]
			fmt.Print(cell.Rep())
		}
	}
	fmt.Println()
}


type Cell struct {
	col, row, neighbours int
	alive bool
}


func (c Cell) Rep() string {
	if c.alive {
		return "x"
	} else {
		return "."
	}
}


func main(){
	grid := Grid{10, 10, make([][]Cell, 10)}
	//fmt.Println(grid.grid)
	//fmt.Println(len(grid.grid))
	grid.InitiateCells()
	grid.DisplayGrid()

	// cell := Cell{10,10,5,false}
	// fmt.Println(cell.Rep())
}