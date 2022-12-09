package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
	
    if results != 1 {
        t.Fatalf(fmt.Sprintf("Expected a score of 1 rather than %d ", results))
    }
}

func TestSolveProblem2(t *testing.T){
	results := solveProblem("input_test_2.txt")	
	
    if results != 36 {
        t.Fatalf(fmt.Sprintf("Expected a score of 36 rather than %d ", results))
    }
}







