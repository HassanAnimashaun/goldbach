package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// getPrimes is a method that iterates through all numbers up to the MAX number in the
// data set to retrieve all the prime numbers.
func getPrimes(maxNumber int) []int {

	// Initialize a list containing our prime numbers.
	var primes []int

	// Iterate from 2 to the MAX number in our data set, incrementing value by 1.
	for value := 2; value <= maxNumber; value++ {
		isPrime := true

		// For each prime number in the list primes [], if the current iterator (value) is evenly
		// divisible by any number in primes[], the number is not prime.
		for _, prime := range primes {
			if value%prime == 0 {
				isPrime = false
				break
			}
		}

		// If the number is prime, append to our list primes[].
		if isPrime {
			primes = append(primes, value)
		}
	}
	return primes
}

// The Goldbach function finds Goldbach pairs for the current value
// using our list of prime numbers in primes[].
func goldbach(value int, primes []int) []int {
	var result []int
	if value >= 4 && value%2 == 0 {
		for _, prime := range primes {
			if prime > value/2 {
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
// is present in a given array.
func contains(arr []int, num int) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

// Main function where the program starts. Used to read from the "data.txt" file.
func main() {
	data, _ := readfile("data.txt")

	// Loop to calculate the max number in the data set.
	maxNumber := 0
	for _, num := range data {
		if num > maxNumber {
			maxNumber = num
		}
	}

	// Use the getPrimes function to retrieve all prime numbers up to our max &
	// store it in primes.
	primes := getPrimes(maxNumber)

	for _, value := range data {
		goldbachPairs := goldbach(value, primes)
		fmt.Printf("We found %d Goldbach pair(s) for %d:\n", len(goldbachPairs), value)

		// Print the output of every Goldbach pair in the format of the "results.txt" file.
		for _, pair := range goldbachPairs {
			difference := value - pair
			fmt.Printf("%d = %d + %d\n", value, pair, difference)
		}
		fmt.Println()
	}
}

// Function to scan line by line & convert each item to an integer,
// appending them to a list called data.
func readfile(filename string) ([]int, error) {
	var data []int
	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		data = append(data, num)
	}

	return data, nil
}
