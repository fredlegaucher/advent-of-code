package main

import (
	"testing"
	"fmt"
)

func TestScoreCommonLetter(t *testing.T){
    testMap := map[string]int{"p":16,"L":38,"P":42,"v":22,"t":20,"s":19}

    for letter,expectedScore := range testMap {
    score := scoreCommonLetter(letter)
        if score != expectedScore {
        t.Fatalf(fmt.Sprintf("Expected common item type %d to be %d ", score , expectedScore))
    }
}  
}

func TestSolveProblem(t *testing.T){
    if score:= solveProblem("input_test.txt"); score != 70 {
        t.Fatalf(fmt.Sprintf("Expected a score of 70 rather than %d ", score))
    }
}







