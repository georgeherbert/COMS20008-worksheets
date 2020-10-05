package main

//import "fmt"

func getNeighbours(world [][]byte, column int, row int) []byte {
	rowAbove := 0
	rowBelow := 0
	columnLeft := 0
	columnRight := 0

	if row == 0 {
		rowAbove = len(world[0]) - 1
	} else if row == len(world[0]) - 1 {
		rowBelow = 0
	} else {
		rowAbove = row - 1
		rowBelow = row + 1
	}

	if column == 0 {
		columnLeft = len(world[0]) - 1
	} else if column == len(world[0]) - 1 {
		columnRight = 0
	} else {
		columnLeft = column - 1
		columnRight = column + 1
	}

	neighbours := []byte{}
	neighbours = append(neighbours, world[columnLeft][rowAbove])
	neighbours = append(neighbours, world[columnLeft][row])
	neighbours = append(neighbours, world[columnLeft][rowBelow])
	neighbours = append(neighbours, world[column][rowAbove])
	neighbours = append(neighbours, world[column][rowBelow])
	neighbours = append(neighbours, world[columnRight][rowAbove])
	neighbours = append(neighbours, world[columnRight][row])
	neighbours = append(neighbours, world[columnRight][rowBelow])

	//fmt.Println(neighbours)
	return neighbours
}

func calculateLiveNeighbours(neighbours []byte) int {
	liveNeighbours := 0
	for _, j := range neighbours {
		if j == 255 {
			liveNeighbours += 1
		}
	}
	//fmt.Println(liveNeighbours)
	return liveNeighbours
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	worldCopy := [][]byte{}
	for i, layer := range world {
		worldCopy = append(worldCopy, []byte{})
		for j, cell := range layer {

			neighbours := getNeighbours(world, i, j)
			liveNeighbours := calculateLiveNeighbours(neighbours)

			if cell == 255 {
				if liveNeighbours < 2 || liveNeighbours > 3 {
					worldCopy[i] = append(worldCopy[i], 0)
				} else {
					worldCopy[i] = append(worldCopy[i], 255)
				}
			} else {
				if liveNeighbours == 3 {
					worldCopy[i] = append(worldCopy[i], 255)
				} else {
					worldCopy[i] = append(worldCopy[i], 0)
				}
			}
		}
	}
	return worldCopy
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}
	for i, layer := range world {
		for j, item := range layer {
			if item == 255 {
				aliveCells = append(aliveCells, cell{j, i})
			}
		}
	}
	return aliveCells
}
