# math/stats

The _math/stats_ package contains statistical functions. The following are currently implemented:

## Descriptive statistics
- `Sum([]float64) float64` Calculate the sum of a set of values
- `Product([]float64) float64` Calculate the product of a set of values
- `Min(values ... int) (int, float64)` Calculate the minimum of a set of values
- `Max(values ... int) (int, float64)` Calculate the maximum of a set of values
- `Average([]float64) (float64, error)` Synonym for ArithmeticMean.
- `ArithmeticMean([]float64) (float64, error)` Calculate the arithmetic mean of a set of values
- `HarmonicMean([]float64) (float64, error)` Calculate the harmonic mean of a set of values
- `QuadraticMean([]float64) (float64, error)` Calculate the quadratic mean of a set of values
- `GeometricMean([]float64) (float64, error)` Calculate the geometric mean of a set of values
- `Mean(data []float64, meanNumber int) (float64, error)` Calculate a specified mean for a set of values
- `Accumulate(data []float64, initial float64, fn func(acc, val float64)) float64` Generic accumulator that applies a function to each value

## Inferential statistics
- `PairedT(data [2][]float64) (float64, error)` Calculates the t statistic for paired (related) data
- `UnpairedT(data [2][]float64) (float64, error)` Calculates the t statistic for unpaired (unrelated) data
