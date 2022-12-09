package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
	
    if results != 21 {
        t.Fatalf(fmt.Sprintf("Expected a score of 21 rather than %d ", results))
    }
}







