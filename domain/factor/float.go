package factor

import "fmt"

type Float interface {
	Correct(d float64) float64
}

func NewFloat(m float64) (Float, error) {
	if m < 0 {
		return nil, fmt.Errorf("%f not supported.", m)
	}
	return &floatImpl{m}, nil
}

type floatImpl struct {
	m float64
}

func (f *floatImpl) Correct(d float64) float64 {
	return d * f.m
}
