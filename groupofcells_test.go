package testing

import (
	"github.com/LukeJelly/Cell-Compete-Go/groupofcells"
	"testing"
)

func TestGetGroupAsArray(t *testing.T) {
	var original = []uint8{1,0,0,0,0,1,0,0}
	testGPC := groupofcells.NewGroupOfCells(original)
	actual := testGPC.GetGroupAsArray()
	var expected = []uint8{1,0,0,0,0,1,0,0}
	compareTwoGroups(expected, actual, t)
}

func TestCaseOneDay(t *testing.T){
	var days int = 1;
	var original = []uint8{1,0,0,0,0,1,0,0}
	var expected = []uint8{0,1,0,0,1,0,1,0}
	runCompete(days, original, expected, t)
}

func TestCaseTwoDays(t *testing.T){
	var days = 2
	var original = []uint8{1,1,1,0,1,1,1,1}
	var expected = []uint8{0,0,0,0,0,1,1,0}
	runCompete(days, original, expected, t)
}

func runCompete(days int, original []uint8, expected []uint8, t *testing.T){
	testGPC := groupofcells.NewGroupOfCells(original)
	testGPC.Compete(days)
	actual := testGPC.GetGroupAsArray()
	compareTwoGroups(expected, actual, t)
}

func compareTwoGroups(expected []uint8, actual []uint8, t *testing.T){
	lenExpected := len(expected)
	lenActual := len(actual)
	if (lenExpected != lenActual) {
		t.Errorf("Expected a length of %d but actual length was %d", lenExpected, lenActual)
	}
	for i := 0; i < lenExpected; i++ {
		expectedValue := expected[i]
		actualValue := actual[i]
		if (expectedValue != actualValue){
			t.Errorf("At index %d, expected %d but was %d", i, expectedValue, actualValue)
		}
	}
}