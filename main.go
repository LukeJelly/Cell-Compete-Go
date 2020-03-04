package main

import (
	"github.com/LukeJelly/Cell-Compete-Go/groupofcells"
	"fmt"
)


func main(){
	var original = []uint8{1,0,0,0,0,1,0,0}
	var expected = []uint8{0,1,0,0,1,0,1,0}
	GPC := groupofcells.NewGroupOfCells(original)
	fmt.Println(GPC)
	GPC.Compete(1)
	fmt.Println(GPC)
	fmt.Println(expected)

}
