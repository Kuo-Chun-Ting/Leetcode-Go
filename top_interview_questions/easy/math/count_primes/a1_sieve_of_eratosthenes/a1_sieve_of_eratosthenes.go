package a1_sieve_of_eratosthenes

import "math"

func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	primesTable := initPrimesTable(n)
	squartRoot := int(math.Ceil(math.Sqrt(float64(n))))

	for i := 2; i <= squartRoot; i++ {
		if !primesTable[i] {
			continue
		}

		for num := i * i; num < n; num += i {
			primesTable[num] = false
		}
	}

	primesCount := countPrimesInTable(primesTable)
	return primesCount
}

func initPrimesTable(n int) []bool {
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}
	return primes
}

func countPrimesInTable(primesTable []bool) int {
	primesCount := 0
	for i := 1; i < len(primesTable); i++ {
		if primesTable[i] {
			primesCount++
		}
	}
	return primesCount
}
