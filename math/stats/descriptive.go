package stats

import (
	"errors"
	"math"
)

// convenience constants for meanNumber in Mean()
const (
	ArithmeticMeanNumber = 1
	HarmonicMeanNumber   = -1
	QuadraticMeanNumber  = 2
	GeometricMeanNumber  = 0
)

func Accumulate(data []float64, init float64, fn func(float64, float64) float64) float64 {
	for _, value := range data {
		init = fn(init, value)
	}

	return init
}

func Sum(data []float64) (sum float64) {
	for _, value := range data {
		sum += value
	}

	return
}

func Product(data []float64) float64 {
	prod := 1.0

	for _, value := range data {
		prod *= value
	}

	return prod
}

func Max(values ...float64) (float64, error) {
	if 0 == len(values) {
		return 0, errors.New("no values to compare")
	}

	a := values[0]

	for _, b := range values[1:] {
		if b > a {
			a = b
		}
	}

	return a, nil
}

func MaxInt(values ...int) (int, error) {
	if 0 == len(values) {
		return 0, errors.New("no values to compare")
	}

	a := values[0]

	for _, b := range values[1:] {
		if b > a {
			a = b
		}
	}

	return a, nil
}

func Min(values ...float64) (float64, error) {
	if 0 == len(values) {
		return 0, errors.New("no values to compare")
	}

	a := values[0]

	for _, b := range values[1:] {
		if b < a {
			a = b
		}
	}

	return a, nil
}

func MinInt(values ...int) (int, error) {
	if 0 == len(values) {
		return 0, errors.New("no values to compare")
	}

	a := values[0]

	for _, b := range values[1:] {
		if b < a {
			a = b
		}
	}

	return a, nil
}

func Mean(data []float64, meanNumber int) (float64, error) {
	n := len(data)

	if 0 == n {
		return 0.0, errors.New("empty array of data")
	}

	if 0 == meanNumber {
		return math.Pow(Product(data), 1.0/float64(n)), nil
	}

	return math.Pow(Accumulate(data, 0.0, func(acc, val float64) float64 {
		return acc + math.Pow(val, float64(meanNumber))
	})/float64(n), 1.0/float64(meanNumber)), nil
}

func Average(data []float64) (float64, error) {
	return Mean(data, ArithmeticMeanNumber)
}

func ArithmeticMean(data []float64) (float64, error) {
	return Mean(data, ArithmeticMeanNumber)
}

func HarmonicMean(data []float64) (float64, error) {
	return Mean(data, HarmonicMeanNumber)
}

func QuadraticMean(data []float64) (float64, error) {
	return Mean(data, QuadraticMeanNumber)
}

func GeometricMean(data []float64) (float64, error) {
	return Mean(data, GeometricMeanNumber)
}
