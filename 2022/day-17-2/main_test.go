package main

import (
	"testing"
)

// func TestSolveProblem(t *testing.T){
// 	results := solveProblem("input_test.txt", 1000000000000)
//     if results != 1514285714288 {
//         t.Fatalf("Expected 1514285714288, not %v",results)
//     }
// }

func TestSolveProblemPart1Example(t *testing.T){
	results := solveProblem("input_test.txt", 2022)	
    if results != 3068 {
        t.Fatalf("Expected 3068, not %v",results)
    }
}

func TestSolvePart1Actual(t *testing.T){
	results := solveProblem("input.txt", 2022)	
    if results != 3177 {
        t.Fatalf("Expected 3177, not %v",results)
    }
}












