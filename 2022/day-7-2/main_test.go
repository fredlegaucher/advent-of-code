package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
	results := solveProblem("input_test.txt")	
	
    if results != 24933642 {
        t.Fatalf(fmt.Sprintf("Expected a score of 24933642 rather than %d ", results))
    }
}







