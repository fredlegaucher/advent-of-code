package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	a,b,c := solveProblem("input_test.txt")	
    
    if a != 6 || b != 8 || c!= 0 {
        t.Fatalf("Expected 6,8,0 not %v,%v,%v",a,b,c)
    }
}
