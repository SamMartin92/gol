package main

import (
	"fmt"
	"time"
)

var neighbourCoords [8][2]int = [8][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

	

type Grid struct {
	rows, cols int
	grid [][]Cell
	neighbourCoords [8][2]int
}


func (g Grid) InitiateCells() {
	for row:= range g.rows {
		for col:= range g.cols {
			cell := Cell{row, col, 0, false}
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

// this can be much better, make this a switch case, abstract the coords
func (g Grid) CountLivingNeighbours(row, col int) int {
	count:=0
	for coord := range g.neighbourCoords {
		var neighbourCoord [2]int

		if row + g.neighbourCoords[coord][0] == -1 {
			neighbourCoord[0] = len(g.grid)-1 
		} else if row + g.neighbourCoords[coord][0] != len(g.grid) {
			neighbourCoord[0] = row + g.neighbourCoords[coord][0] 
		} else {
			neighbourCoord[0] = 0
		}

		if col + g.neighbourCoords[coord][1] == -1 {
			neighbourCoord[1] = len(g.grid)-1
		} else if col + g.neighbourCoords[coord][1] != len(g.grid){
			neighbourCoord[1] = col + g.neighbourCoords[coord][1]
		}else {
			neighbourCoord[1] = 0
		}

		neighbour:= g.grid[neighbourCoord[0]][neighbourCoord[1]]
		if neighbour.alive {
			count+=1
		}
	}
	return count
}


func (g Grid) UpdateCellState(row,col int) bool {
	cell := g.grid[row][col]

	cell.neighbours = g.CountLivingNeighbours(row,col)

	if cell.alive {
		if cell.neighbours == 2 || cell.neighbours == 3 {
			return true
		} else {
			return false
		}
	} else {
		if cell.neighbours == 3 {
			return true
		}		
	}

	return cell.alive
}


func (g Grid) GenerateNextGrid() Grid {
	new_grid := Grid{g.rows, g.cols, make([][]Cell,g.rows), neighbourCoords}
	new_grid.InitiateCells()
	
	for row := range g.rows {
		for col := range g.cols {
			new_grid.grid[row][col].alive = g.UpdateCellState(row,col)
		}
	}

	return new_grid
}


type Cell struct {
	row, col, neighbours int
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
	grid.grid[9][9].alive = true
	grid.grid[0][0].alive = true
	grid.grid[1][1].alive = true
	grid.grid[1][2].alive = true
	grid.grid[2][2].alive = true
	
	grid.DisplayGrid()


	for x:= true; x; {
		time.Sleep(1 * time.Second)
		grid = grid.GenerateNextGrid()
		grid.DisplayGrid()
	}
	


}