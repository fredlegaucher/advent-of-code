package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
    if score:= solveProblem("input_test.txt"); score != 4 {
        t.Fatalf(fmt.Sprintf("Expected a score of 4 rather than %d ", score))
    }
}







