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
            neighbour_coord= (row+coord[0],col+coord[1])
            if neighbour_coord[0] in [-1, len(self.grid)]:
                continue
            elif neighbour_coord[1] in [-1, len(self.grid[0])]:
                continue
            else:
                neighbour = self.grid[neighbour_coord[0]][neighbour_coord[1]]
            if neighbour.alive:
                count+=1
        
        return count


    def update_cell_state(self,row,col):
        cell = self.grid[row][col]
        living_neighbour_count = self.count_living_neighbours(row,col)
        print(living_neighbour_count)

        if cell.alive:
            if living_neighbour_count <2:
                cell.alive = False # cell dies if 0 or 1 neighbours
            elif living_neighbour_count >3:
                cell.alive = False # cell dies if 4 or more neighbours
            else:
                pass
                 # do nothing if 2 or 3 neighbours
        else:
            if living_neighbour_count == 3:
                cell.alive = True # cell comes to life if 3 neighbours
        

    def generate_next_grid(self):
        for row in range(len(self.grid)):
            for col in range(len(self.grid[0])):
                self.update_cell_state(row,col)
        
        self.display_grid()
        
    


class Cell:
    def __init__(self, col, row):
        self.col = col
        self.row = row
        self.alive = False
        


    def __str__(self):
        return '#' if self.alive else '.'



if __name__ == "__main__":  
    grid = Grid(5,5)
     
    grid.grid[2][2].alive = True
    grid.grid[3][2].alive = True
    grid.grid[1][2].alive = True


    grid.display_grid()

    print('\n\n')

    grid.generate_next_grid()


