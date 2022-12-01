package main

import (
	"testing"
	"fmt"
)

func TestSolveProblem(t *testing.T){
    if score:= solveProblem("input_test.txt"); score != 2 {
        t.Fatalf(fmt.Sprintf("Expected a score of 2 rather than %d ", score))
    }
}







