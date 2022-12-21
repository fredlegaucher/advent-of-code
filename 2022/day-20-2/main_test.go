package main

import (
	"testing"
)

func TestSolveExample(t *testing.T){
	a,b,c := solveProblem("input_test.txt")	
    
    if a != 811589153 || b != 2434767459 || c!= -1623178306 {
        t.Fatalf("Expected 811589153,2434767459,-1623178306 not %v,%v,%v",a,b,c)
    }
}

func TestSolveProblem(t *testing.T){
	a,b,c := solveProblem("input.txt")	
    
    if a != -2633606801485 || b != -1482773382531 || c!= 4978287864502 {
        t.Fatalf("Expected -2633606801485,-1482773382531,4978287864502 not %v,%v,%v",a,b,c)
    }
}
