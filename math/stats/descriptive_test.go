package stats

import (
	"fmt"
	"testing"
	"github.com/darrenedale/go/util"
)

const (
	ComparisonPrecision = 9 // compare to this dp
)

func TestMean(test *testing.T) {
	cases := []struct {
		data           []float64
		meanNumber     int
		expectedResult float64
		expectedErr    bool
	}{
		{make([]float64, 0), ArithmeticMeanNumber, 0.0, true},
		{[]float64{1.0, 2.0, 3.0}, ArithmeticMeanNumber, 2.0, false},
		{[]float64{1.0, 2.0, 3.0}, HarmonicMeanNumber, 1.63636363636364, false},
		{[]float64{1.0, 2.0, 3.0}, QuadraticMeanNumber, 2.16024689946929, false},
		{[]float64{1.0, 2.0, 3.0}, GeometricMeanNumber, 1.81712059283214, false},
	}

	// comparing floats with == is unreliable: test equality to specific dp
	comparisonFormat := fmt.Sprintf("%%0.%df", ComparisonPrecision)

	for i, testCase := range cases {
		actualResult, err := Mean(testCase.data, testCase.meanNumber)
		fmt.Printf("Test case %d: Mean(%v, %d) => "+comparisonFormat+", %s\n", i, testCase.data, testCase.meanNumber, actualResult, util.Choose(nil != err, "error", "no error"))

		if err == nil && testCase.expectedErr {
			test.Errorf("Test case %d: mean(%v, %d): expected error state %v, actual error state %v", i, testCase.data, testCase.meanNumber, testCase.expectedErr, (err == nil))
		} else if err == nil && fmt.Sprintf(comparisonFormat, actualResult) != fmt.Sprintf(comparisonFormat, testCase.expectedResult) {
			test.Errorf("Test case %d: mean(%v, %d): expected "+comparisonFormat+", actual "+comparisonFormat, i, testCase.data, testCase.meanNumber, testCase.expectedResult, actualResult)
		}
	}
}

func TestMedian(test *testing.T) {
	cases := []struct {
		data           []float64
		expectedResult float64
		expectedErr    bool
	}{
		{[]float64{}, 0.0, true},
		{[]float64{1.0, 2.0, 3.0}, 2.0, false},
		{[]float64{1.0, 2.0}, 1.5, false},
		{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}, 3.0, false},
		{[]float64{1.0, 3.0, 5.0, 4.0, 2.0}, 3.0, false},
	}

	// comparing floats with == is unreliable, so we test equality when formatted to specific dp
	comparisonFormat := fmt.Sprintf("%%0.%df", ComparisonPrecision)

	for i, testCase := range cases {
		actualResult, err := Median(testCase.data)
		fmt.Printf("Test case %d: Median(%v) => "+comparisonFormat+", %s\n", i, testCase.data, actualResult, util.Choose(nil != err, "error", "no error"))

		if err == nil && testCase.expectedErr {
			test.Errorf("Test case %d: Median(%v): expected error state %v, actual error state %v", i, testCase.data, testCase.expectedErr, (err == nil))
		} else if err == nil && fmt.Sprintf(comparisonFormat, actualResult) != fmt.Sprintf(comparisonFormat, testCase.expectedResult) {
			test.Errorf("Test case %d: Median(%v): expected "+comparisonFormat+", actual "+comparisonFormat, i, testCase.data, testCase.expectedResult, actualResult)
		}
	}
}
