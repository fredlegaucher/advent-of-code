package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
	
    if results != 95437 {
        t.Fatalf(fmt.Sprintf("Expected a score of 95437 rather than %d ", results))
    }
}







