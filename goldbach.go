package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

// getPrimes generates prime numbers concurrently up to the "maxNumber".
func getPrimes(maxNumber int) []int {

	primes := make(chan int) // Channel to send prime numbers.
	var primeSlice []int     // Slice(list) to store the prime numbers.
	var wg sync.WaitGroup    // Wait group to synchronize goroutines.

	// Start a goroutine to generate prime numbers.
	wg.Add(1)
	go func() {
		defer wg.Done()     // Mark this goroutine as done.
		defer close(primes) // Close the channel when done.

		// Generate prime numbers and send them to the channel.
		primes <- 2 // First prime number.
		for num := 3; num <= maxNumber; num += 2 {
			isPrime := true
			for i := 2; i*i <= num; i++ { // Check divisibility up to the square root of (num).
				if num%i == 0 { // If it's divisible, it's not prime.
					isPrime = false
					break
				}
			}
			if isPrime {
				primes <- num // Send to the channel if it's a prime.
			}
		}
	}()

	// Take from the channel & build the list of primes.
	for prime := range primes {
		primeSlice = append(primeSlice, prime) // Add to the list.
	}

	wg.Wait()         // Wait until all goroutines complete before proceeding.
	return primeSlice // Return the list of primes.
}

// The Goldbach function finds Goldbach pairs for the current value
// using our list of prime numbers.
func goldbach(value int, primes []int) []int {
	var result []int
	if value >= 4 && value%2 == 0 { // Only even numbers.
		for _, prime := range primes {
			if prime > value/2 { // No valid pairs beyond `value / 2`
				break
			}

			// For each prime number, if the value - prime == a value in primes[]
			// append the Goldbach pair to our result[] list.
			difference := value - prime
			if contains(primes, difference) {
				result = append(result, prime)
			}
		}
	}
	return result
}

// contains is a method used for checking if the current number
// is in a given slice.
func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

// Main function to calculate Goldbach pairs. Used to read from the "data.txt" file.
func main() {
	// Read data from the "data.txt" file.
	data, err := readfile("data.txt")

	// If there's an error reading "data.txt" or it doesn't exist, use default values.
	if err != nil || len(data) == 0 {
		data = []int{3, 4, 14, 26, 100}
	}

	maxNumber := 0
	for _, num := range data {
		if num > maxNumber {
			maxNumber = num
		}
	}

	primes := getPrimes(maxNumber)

	for _, value := range data {
		goldbachPairs := goldbach(value, primes)
		fmt.Printf("We found %d Goldbach pair(s) for %d:\n", len(goldbachPairs), value)
		for _, pair := range goldbachPairs {
			difference := value - pair
			fmt.Printf("%d = %d + %d\n", value, pair, difference)
		}
		fmt.Println()
	}
}

// Function to read data from a file.
func readfile(filename string) ([]int, error) {
	var data []int
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
