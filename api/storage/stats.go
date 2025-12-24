package storage

import (
	"math"
)

type ZScore struct {
	values    []float64
	threshold float64
}

var ZS *ZScore

func InitZScore(threshold float64) *ZScore {
	ZS = &ZScore{
		values:    []float64{},
		threshold: threshold,
	}

	return ZS
}

func (z *ZScore) Add(value float64) {
	z.values = append(z.values, value)
}

func (z *ZScore) Mean() float64 {
	sum := 0.0
	for _, v := range z.values {
		sum += v
	}

	return sum / float64(len(z.values))
}

func (z *ZScore) Std() float64 {
	mean := z.Mean()
	sum := 0.0

	for _, v := range z.values {
		diff := v - mean

		sum += diff * diff
	}

	return math.Sqrt(sum / float64(len(z.values)))
}

func (z *ZScore) ZScore(value float64) float64 {
	std := z.Std()
	if std == 0 {
		return 0
	}
	return (value - z.Mean()) / std
}

func (z *ZScore) IsAnomaly(value float64) (float64, bool) {
	zscore := z.ZScore(value)
	return zscore, math.Abs(zscore) > z.threshold
}
