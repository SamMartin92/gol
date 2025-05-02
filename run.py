import time

class Grid:
    def __init__(self, cols, rows):
        self.cols = cols
        self.rows = rows
        self.grid = [[Cell(col,row) for col in range(cols)] for row in range(rows)]
        self.neighbour_coords = [(-1,-1), (-1,0), (-1,1),
                                (0,-1), (0,1),
                                (1,-1), (1,0), (1,1)]
    

    def display_grid(self):
        for row in range(len(self.grid)):
            print()
            for col in range(len(self.grid[0])):
                cell = self.grid[row][col]
                print(cell, end="")
        print()
    

    def count_living_neighbours(self,row,col):
        count=0
        for coord in self.neighbour_coords:
            neighbour_coord= (row+coord[0] if row+coord[0] != len(self.grid) else 0,
                              col+coord[1] if col+coord[1] != len(self.grid) else 0)
            neighbour = self.grid[neighbour_coord[0]][neighbour_coord[1]]
            if neighbour.alive:
                count+=1
        return count


    def update_cell_state(self,row,col):
        cell = self.grid[row][col]

        cell.neighbours = self.count_living_neighbours(row,col)

        if cell.alive:
            if cell.neighbours <2:
                return False # cell dies if 0 or 1 neighbours
            elif cell.neighbours >3:
                return False # cell dies if 4 or more neighbours
            else:
                return True
                 #do nothing if 2 or 3 neighbours
        else:
            if cell.neighbours == 3:
                return  True # cell comes to life if 3 neighbours
        
        




    def generate_next_grid(self):
        new_grid = Grid(self.rows,self.cols)

        for row in range(len(self.grid)):
            for col in range(len(self.grid[0])):
                new_grid.grid[row][col].alive = grid.update_cell_state(row,col)

        
        return new_grid
    


class Cell:
    def __init__(self, col, row):
        self.col = col
        self.row = row
        self.alive = False
        self.neighbours = 0
        

    def __str__(self):
        return '#' if self.alive else '.'



if __name__ == "__main__":  
    grid = Grid(10,10)
     
     
    grid.grid[2][2].alive = True
    grid.grid[3][2].alive = True
    grid.grid[1][2].alive = True
    grid.grid[0][1].alive = True
    grid.grid[1][3].alive = True
    grid.grid[4][4].alive = True
    grid.grid[1][3].alive = True
    grid.grid[4][4].alive = True
    grid.grid[1][3].alive = True
    grid.grid[6][6].alive = True
    grid.grid[6][7].alive = True
    grid.grid[7][8].alive = True
    grid.grid[6][8].alive = True
    grid.grid[8][9].alive = True
    grid.grid[9][8].alive = True
    grid.grid[1][1].alive = True
    grid.grid[1][2].alive = True
    grid.grid[2][1].alive = True
    grid.grid[2][2].alive = True

    
    grid.display_grid()
    while True:
        time.sleep(1)
        grid = grid.generate_next_grid()
        grid.display_grid()



