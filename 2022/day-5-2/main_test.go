package main

import (
	"fmt"
	"testing"
)

func TestSolveProblem(t *testing.T){
    firstStack := []string{"Z","N"}
	secondStack := []string{"M","C","D"}
	thirdStack := []string{"P"}

	indexToStackMap := make(map[int][]string)
	indexToStackMap[1] = firstStack
	indexToStackMap[2] = secondStack
	indexToStackMap[3] = thirdStack


    if score:= solveProblem("input_test.txt",indexToStackMap); score != "MCD" {
        t.Fatalf(fmt.Sprintf("Expected a score of MCD rather than %s ", score))
    }
}







