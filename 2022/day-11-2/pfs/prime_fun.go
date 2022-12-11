package pfs

import (
	"fmt"
	"math"
)

// Get all prime factors of a given number n
func PrimeFactors(n uint64) (pfs map[uint64]uint64) {
	if n == 0 {
		return nil
	}

	pfs = make(map[uint64]uint64)
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs[2]++
		n = n / 2
	}

	// n must be odd at this pouint64. so we can skip one element
	// (note i = i + 2)
	for i := uint64(3); float64(i) <= math.Sqrt(float64(n)); i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs[i]++
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs[n]++
	}

	return
}

func PfsMultiplication(lhs map[uint64]uint64, rhs map[uint64]uint64) map[uint64]uint64{
	for factor,count := range rhs{
		lhs[factor] += count
	}
	return lhs
}

func AddInttoPFS(pfs map[uint64]uint64,operand uint64) map[uint64]uint64{
	currentValue := uint64(1)
	afterValue  := uint64(1)
	for factor,count := range pfs {
		afterValue = currentValue * uint64(math.Pow(float64(factor), float64(count)))
		if afterValue  < currentValue {
			panic(fmt.Sprintf("We have looped %v with pfs %v",currentValue,pfs))
		}
		currentValue = afterValue
		
	}

	if afterValue + operand < afterValue {
		panic(fmt.Sprintf("We have looped %v with operand %v",currentValue,operand))
	}


	return PrimeFactors(currentValue + operand)
}

func CanPFSBeDivided(lhs map[uint64]uint64, rhs map[uint64]uint64) bool {
	result := true
	for factor,count := range rhs {
		if lhs[factor] < count {
			result = false
			break
		}
	}

	return result
}