package main

import (
	"fmt"
)

var neighbourCoords [8][2]int = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

	

type Grid struct {
	cols, rows int
	grid [][]Cell
	neighbourCoords [8][2]int
}


func (g Grid) InitiateCells() {
	for row:= range g.rows {
		for col:= range g.cols {
			cell := Cell{col, row, 0, false}
			g.grid[row] = append(g.grid[row], cell)
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


func (g Grid) CountLivingNeighbours(row, col int) int {
	count:=0

	for coord := range g.neighbourCoords {
		var neighbourCoord [2]int
		//can make this a switch case
		if row + g.neighbourCoords[coord][0] != len(g.grid) {
			neighbourCoord[0] = row + g.neighbourCoords[coord][0] 
		} else {
			neighbourCoord[0] = 0
		}

		if col + g.neighbourCoords[coord][1] != len(g.grid) {
			neighbourCoord[1] = col + g.neighbourCoords[coord][1]
		} else {
			neighbourCoord[1] = 0
		}

		neighbour:= g.grid[neighbourCoord[0]][neighbourCoord[1]]
		fmt.Println(neighbour.row, neighbour.col)
		if neighbour.alive {
			count+=1
		}
	}
	return count
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
	
	grid := Grid{10, 10, make([][]Cell,10), neighbourCoords}
	grid.InitiateCells()

	grid.grid[2][2].alive = true
	grid.grid[3][2].alive = true
	grid.grid[1][2].alive = true
	grid.grid[0][1].alive = true
	grid.grid[1][3].alive = true
	grid.grid[4][4].alive = true
	grid.grid[1][3].alive = true
	grid.grid[4][4].alive = true
	grid.grid[1][3].alive = true
	grid.grid[6][6].alive = true
	grid.grid[6][7].alive = true
	grid.grid[7][8].alive = true
	grid.grid[6][8].alive = true
	grid.grid[8][9].alive = true
	grid.grid[9][8].alive = true
	grid.grid[1][1].alive = true
	grid.grid[1][2].alive = true
	grid.grid[2][1].alive = true
	grid.grid[2][2].alive = true
	
	grid.DisplayGrid()

	fmt.Println(grid.CountLivingNeighbours(4,5))


}