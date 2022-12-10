package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
	
    if results != 13140 {
        t.Fatalf(fmt.Sprintf("Expected a score of 13140 rather than %d ", results))
    }
}










