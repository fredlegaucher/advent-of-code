package pfs

import (
	"fmt"
	"testing"
)

func TestPrimeFactors(t *testing.T) {

	checkInteger(23,map[uint64]uint64{23:1},t)
	checkInteger(12,map[uint64]uint64{2:2,3:1},t)
	checkInteger(360,map[uint64]uint64{2:3,3:2,5:1},t)
	checkInteger(97,map[uint64]uint64{97:1},t)
}

func checkInteger(number uint64, expectedDecomposition map[uint64] uint64,t *testing.T){
	if fmt.Sprintf("%v",expectedDecomposition) != fmt.Sprintf("%v",PrimeFactors(number)){
		t.Error(number)
	}
}

func TestMultiplication(t *testing.T) {
	lhs := map[uint64]uint64{2:2,3:1}
	rhs:= map[uint64]uint64{2:3,3:2,5:1}
	expectedDecomposition := map[uint64]uint64{2:5,3:3,5:1}
	actualDecomposition := PfsMultiplication(lhs,rhs)
	if fmt.Sprintf("%v",expectedDecomposition) != fmt.Sprintf("%v",actualDecomposition){
		t.Error(expectedDecomposition)
	}
}

func TestAddInttoPFS(t *testing.T) {
	lhs := map[uint64]uint64{2:2,3:1}
	operand:= uint64(6)
	expectedDecomposition := map[uint64]uint64{2:1,3:2}
	actualDecomposition := AddInttoPFS(lhs,operand)
	if fmt.Sprintf("%v",expectedDecomposition) != fmt.Sprintf("%v",actualDecomposition){
		t.Error(expectedDecomposition)
	}
}

func TestDivisionTrue(t *testing.T) {
	lhs := map[uint64]uint64{2:2,3:1}
	rhs:= map[uint64]uint64{2:1,3:1}
	
	dividable := CanPFSBeDivided(lhs,rhs)
	if !dividable {
		t.Error(lhs)
	}
}

func TestDivisionFalse(t *testing.T) {
	lhs := map[uint64]uint64{2:2,3:1}
	rhs:= map[uint64]uint64{2:1,3:2}
	
	dividable := CanPFSBeDivided(lhs,rhs)
	if dividable{
		t.Error(lhs)
	}
}

