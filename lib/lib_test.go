package lib_test

import (
	"testing"

	"none.com/equ_solve/lib"
)

func Test_noKH_twoKH(t *testing.T) {
	inp1 := []string{"+", "2", "+", "3"}
	inp2 := []string{"+", "2", "-", "1"}
	correct := []string{"5"}

	actual := lib.NoKH_twoKH(inp1, inp2)
	if len(actual) != len(correct) {
		t.Fail()
	}
	for i, _ := range correct {
		if correct[i] != actual[i] {
			t.Fail()
		}
	}
}
