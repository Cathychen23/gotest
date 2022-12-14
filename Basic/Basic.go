package main

import (
	"fmt"
	"math"
)

func Add(x int, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Basic13(x, y int) uint {
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	return z
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// Basic16
func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func Flowcontrol3(sum int) int {
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	return sum
}

// Flowcontrol5
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// Flowcontrol7
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
