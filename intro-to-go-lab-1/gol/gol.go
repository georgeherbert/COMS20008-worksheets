package main

//import "fmt"

func getNeighbours(world [][]byte, row int, column int) []byte {
	rowAbove := row - 1
	rowBelow := row + 1
	if row == 0 {
		rowAbove = len(world[0]) - 1
	} else if row == len(world[0]) - 1 {
		rowBelow = 0
	}
	columnLeft := column - 1
	columnRight := column + 1
	if column == 0 {
		columnLeft = len(world[0]) - 1
	} else if column == len(world[0]) - 1 {
		columnRight = 0
	}
	neighbours := []byte{world[rowAbove][columnLeft], world[rowAbove][column], world[rowAbove][columnRight],
		world[row][columnLeft], world[row][columnRight], world[rowBelow][columnLeft], world[rowBelow][column],
		world[rowBelow][columnRight]}
	return neighbours
}

func calculateLiveNeighbours(neighbours []byte) int {
	liveNeighbours := 0
	for _, neighbour := range neighbours {
		if neighbour == 255 {
			liveNeighbours += 1
		}
	}
	return liveNeighbours
}

func calculateValue(item byte, liveNeighbours int) byte {
	calculatedValue := byte(0)
	if item == 255 {
		if liveNeighbours == 2 || liveNeighbours == 3 {
			calculatedValue = byte(255)
		}
	} else {
		if liveNeighbours == 3 {
			calculatedValue = byte(255)
		}
	}
	return calculatedValue
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	nextWorld := [][]byte{}
	for i, row := range world {
		nextWorld = append(nextWorld, []byte{})
		for j, item := range row {
			neighbours := getNeighbours(world, i, j)
			liveNeighbours := calculateLiveNeighbours(neighbours)
			valueCalculated := calculateValue(item, liveNeighbours)
			nextWorld[i] = append(nextWorld[i], valueCalculated)
		}
	}
	return nextWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	aliveCells := []cell{}
	for i, row := range world {
		for j, item := range row {
			if item == 255 {
				aliveCells = append(aliveCells, cell{j, i})
			}
		}
	}
	return aliveCells
}
