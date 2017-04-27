package main

import (
	"fmt"
)

type Cell struct {
	Count int
	Parent *Cell
}

func NewCell() *Cell {
	return &Cell{0, nil}
}

func (c *Cell) isConnected() bool {
	return c.Parent != nil
}

func (c *Cell) setParent(p *Cell) {
	c.Parent = p
	c.Parent.Count += 1
}

//func (c *Cell) updateConnectedParent(gridMap map[int]*Cell, matrix [][]int, matrixCell [][]*Cell, posx, posy, colSize int) map[int]*Cell {
//	switch {
//	case posx > 0 && posy > 0 && matrixCell[posx-1][posy-1].Parent != nil:
//		c.setParent(matrixCell[posx-1][posy-1].Parent)
//	case posx > 0 && matrixCell[posx-1][posy].Parent != nil:
//		c.setParent(matrixCell[posx-1][posy].Parent)
//	case posy > 0 && matrixCell[posx][posy-1].Parent != nil:
//		c.setParent(matrixCell[posx][posy-1].Parent)
//	case posx > 0 && posy+1 < colSize && matrixCell[posx-1][posy+1].Parent != nil:
//		c.setParent(matrixCell[posx-1][posy+1].Parent)
//	default:
//		c.setParent(c)
//		gridMap[10*posx + posy] = c
//	}
//	return gridMap
//}

func (c *Cell) updateConnectedParent(gridMap map[int]*Cell, matrix [][]int, matrixCell [][]*Cell, x, y, row, col int) {
	if(c.Parent != nil) {
		return
	}
	selfParent := false
	var parentCell *Cell
	switch {
	case x > 0 && y > 0 && matrix[x-1][y-1] == 1:
		fmt.Println("1")
		parentCell = matrixCell[x-1][y-1]
		x--
		y--
	case x > 0 && matrix[x-1][y] == 1:
		fmt.Println("2")
		parentCell = matrixCell[x-1][y]
		x--
	case x > 0 && y+1 < col && matrix[x-1][y+1] == 1:
		fmt.Println("3")
		parentCell = matrixCell[x-1][y+1]
		x--
		y++
	case y > 0 && matrix[x][y-1] == 1:
		fmt.Println("4")
		parentCell = matrixCell[x][y-1]
		y--
	case y+1 < col && matrix[x][y+1] == 1:
		fmt.Println("5")
		parentCell = matrixCell[x][y+1]
		y++
	case x+1 < row && y > 0 && matrix[x+1][y-1] == 1:
		fmt.Println("6")
		parentCell = matrixCell[x+1][y-1]
		x++
		y--
	case x+1 < row && matrix[x+1][y] == 1:
		fmt.Println("7")
		parentCell = matrixCell[x+1][y]
		x++
	case x+1 < row && y+1 < col && matrix[x+1][y+1] == 1:
		fmt.Println("8")
		parentCell = matrixCell[x+1][y+1]
		x++
		y++
	default:
		parentCell = c
		gridMap[10*x + y] = c
		selfParent = true
	}
	if(selfParent == false && parentCell.Parent == nil) {
		parentCell.updateConnectedParent(gridMap, matrix, matrixCell, x, y, row, col)
	}
	c.setParent(parentCell.Parent)
}

func max(a1, a2 int) int {
	if(a1 > a2) {
		return a1
	} else {
		return a2
	}
}

func main() {
	fmt.Println("Implementing connected cells in a grid")
	row := 5
	col := 5
	//matrix := make([row][col]int)
	matrixa := [][]int {
		{0, 1, 1, 1, 1},
		{1, 0, 0, 0, 1},
		{1, 1, 0, 1, 1},
		{0, 1, 0, 1, 1},
		{0, 1, 1, 1, 0},
	}
	//0 1 1 1 1
	//1 0 0 0 1
	//1 1 0 1 0
	//0 1 0 1 1
	//0 1 1 1 0
	//var row, col int
	//_, _ = fmt.Scan(&row, &col)

	//matrix := make([][]int, row)
	//for i := 0; i < row; i++ {
	//	matrix[i] = make([]int, col)
	//	for j := 0; j < col; j++ {
	//		var aa int
	//		_, _ = fmt.Scan(&aa)
	//		matrix[i][j] = aa
	//	}
	//}

	matrixCell := make([][]*Cell, row)
	for i := 0; i < row; i++ {
		matrixCell[i] = make([]*Cell, col)
	}

	gridMap := make(map[int]*Cell)
	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			newCell := &Cell{0, nil}
			matrixCell[x][y] = newCell
		}
	}

	for x := 0; x < row; x++ {
		for y := 0; y < col; y++ {
			mcell := matrixCell[x][y]
			if(matrixa[x][y] == 1) {
				mcell.updateConnectedParent(gridMap, matrixa, matrixCell, x, y, row, col)
			}
		}
	}

	maxCells := 0
	for _, val := range gridMap {
		maxCells = max(maxCells, val.Count)
	}
	fmt.Println("Number of Grids: ", len(gridMap))
	fmt.Println(maxCells)
}
