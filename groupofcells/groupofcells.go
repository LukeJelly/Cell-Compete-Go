package groupofcells

import "fmt"

const (
	//Active is a constant that represents what an active cell looks like.
	Active uint8 = 1
	//Inactive is a constant that represents what an inactive cell looks like.
	Inactive uint8 = 0
)

//GroupOfCells is a linkedlist type implementation to the GroupOfCells problem.
type GroupOfCells struct{
	farLeftCell *cell
	size int
}

//NewGroupOfCells creates a new group of cells from the given array.
func NewGroupOfCells(arr []uint8) GroupOfCells{
	index := 0
	GPC := GroupOfCells{
		farLeftCell: makeOneCell(arr[index],nil,nil),
		size: len(arr),
	}

	index++
	createGroupOfCellsFromArray(arr, &GPC, index)
	return GPC
}

func makeOneCell(value uint8, left *cell, right *cell) *cell{
	c := cell{
		value: value,
		left: left,
		right: right,
	}
	return &c
}

func createGroupOfCellsFromArray(arr []uint8, groupOfCells *GroupOfCells, index int){
	var lastUsedCell *cell
	lastUsedCell = groupOfCells.farLeftCell
	var currentCell *cell
	for ; index < len(arr); index++ {
		value := arr[index]
		currentCell = makeOneCell(value, lastUsedCell, nil)
		lastUsedCell.right = currentCell
		lastUsedCell = currentCell
	}
}

func (gpc GroupOfCells) String() string{
	var output string
	pointer := gpc.farLeftCell
	for pointer.right != nil {
		output += fmt.Sprintf("%v,",pointer.value)
		pointer = pointer.right
	}
	output += fmt.Sprintf("%v",pointer.value)

	return output
}

//Compete advances the current group of cells by the given amount
func (gpc GroupOfCells) Compete(days int){
	for i := 0; i < days; i++ {
		iterateGroupOfCellsOneDay(&gpc)
	}
}

func iterateGroupOfCellsOneDay(gpc *GroupOfCells){
	currentCellPointer := gpc.farLeftCell
	var leftPreviousValue uint8

	for currentCellPointer != nil {
		currentCellValue := currentCellPointer.value
		mutateACell(gpc, leftPreviousValue, currentCellPointer)
		leftPreviousValue = currentCellValue
		currentCellPointer = currentCellPointer.right
	}
}

func mutateACell(gpc *GroupOfCells, leftPreviousValue uint8, currentCellPointer *cell){
	rightValue := findRightCellsValue(currentCellPointer)
	
	if(leftPreviousValue == rightValue){
		currentCellPointer.value = Inactive
	}else{
		currentCellPointer.value = Active
	}
}

func findRightCellsValue(currentCellPointer *cell) uint8 {
	var rightValue uint8
	if(currentCellPointer.right == nil){
		rightValue = Inactive
	}else{
		rightValue = currentCellPointer.right.value
	}

	return rightValue
}

//GetGroupAsArray returned a groupofcells as an array.
func (gpc GroupOfCells) GetGroupAsArray() []uint8{
	pointer := gpc.farLeftCell
	returnedArray := make([]uint8, gpc.size)
	for i:=0; i < gpc.size; i++ {
		returnedArray[i] = pointer.value
		pointer = pointer.right
	}
	return returnedArray
}

type cell struct{
	value uint8
	left *cell
	right *cell
}

