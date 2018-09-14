package stats

import (
	"errors"
	"math"
)

func PairedT(data [2][]float64) (t float64, err error) {
	n := len(data[0])

	if 0 == n {
		err = errors.New("no observations")
		return
	}

	if n != len(data[1]) {
		err = errors.New("paired t-test requires paired observations")
		return
	}

	diffs := make([]float64, n)
	diffsSquared := make([]float64, n)

	for i := 0; i < n; i++ {
		diffs[i] = data[1][i] - data[0][i]
		diffsSquared[i] = (diffs[i] * diffs[i])
	}

	sumDiffs := Sum(diffs)
	sumDiffsSuared := Sum(diffsSquared)
	t = sumDiffs / (math.Sqrt(((float64(n) * sumDiffsSuared) - (sumDiffs * sumDiffs)) / float64(n-1)))
	return t, nil
}

func UnpairedT(data [2][]float64) (t float64, err error) {
	n1 := len(data[0])

	if 0 == n1 {
		err = errors.New("no observations in first condition")
		return
	}

	n2 := len(data[1])

	if 0 == n2 {
		err = errors.New("no observations in second condition")
		return
	}

	sum1 := Sum(data[0])
	sum2 := Sum(data[1])
	mean1 := sum1 / float64(n1)
	mean2 := sum2 / float64(n2)

	sum1SquaredDiffs := Accumulate(data[0], 0.0, func(sum float64, val float64) float64 {
		val -= mean1
		return val * val
	})

	sum2SquaredDiffs := Accumulate(data[1], 0.0, func(sum float64, val float64) float64 {
		val -= mean2
		return val * val
	})

	sum1SquaredDiffs /= float64(n1)
	sum2SquaredDiffs /= float64(n2)

	t = (mean1 - mean2) / math.Sqrt((sum1SquaredDiffs/float64(n1-1))+(sum2SquaredDiffs/float64(n2-1)))

	if 0 > t {
		t = -t
	}

	return
}
