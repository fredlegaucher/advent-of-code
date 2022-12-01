package main

import (
	"testing"
	"fmt"
)

func TestGetCompartments(t *testing.T) {
	left,right := getCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")

	if left != "vJrwpWtwJgWr" {
		t.Fatalf(fmt.Sprintf("Expected left compartiment %s to be %s",left,  "vJrwpWtwJgWr"))
    }

    if right != "hcsFMMfFFhFp" {
		t.Fatalf(fmt.Sprintf("Expected right compartment %s to be %s ", right , "hcsFMMfFFhFp"))
    }
}

func TestFindCommonItemtype(t *testing.T){
    commonItemType := findCommonItemtype("vJrwpWtwJgWr","hcsFMMfFFhFp")
    if commonItemType != "p" {
        t.Fatalf(fmt.Sprintf("Expected common item type %s to be %s ", commonItemType , "p"))
    }
}

func TestScoreCommonLetter(t *testing.T){
    testMap := map[string]int{"p":16,"L":38,"P":42,"v":22,"t":20,"s":19}

    for letter,expectedScore := range testMap {
    score := scoreCommonLetter(letter)
        if score != expectedScore {
        t.Fatalf(fmt.Sprintf("Expected common item type %d to be %d ", score , expectedScore))
    }
}  
}

func TestScoreRuckSack(t *testing.T){
    testMap := map[string]int{
        "vJrwpWtwJgWrhcsFMMfFFhFp":16,
        "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL":38,
        "PmmdzqPrVvPwwTWBwg":42,
        "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn":22,
        "ttgJtRGJQctTZtZT":20,
        "CrZsJsPPZsGzwwsLwLmpwMDw":19,
    }

    for ruckSack, expectedScore := range testMap{
        score := scoreRucksack(ruckSack)
        if score != expectedScore {
        t.Fatalf(fmt.Sprintf("Expected score of %d to be %d for rucksack %s", score , expectedScore, ruckSack))
    }
    }
}

func TestSolveProblem(t *testing.T){
    testMap := map[int]string{
        1: "vJrwpWtwJgWrhcsFMMfFFhFp",
        2: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
        3: "PmmdzqPrVvPwwTWBwg",
        4: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
        5: "ttgJtRGJQctTZtZT",
        6: "CrZsJsPPZsGzwwsLwLmpwMDw",
    }

    
    if score:= solveProblem(testMap); score != 157 {
        t.Fatalf(fmt.Sprintf("Expected a score of 157 rather than %d ", score))
    }
}


