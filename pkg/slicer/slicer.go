package slicer

import (
	"math"

	"github.com/phoebetronic/model/pkg/change"
)

type Slicer struct {
	His int
	Lis []float64
}

func (s *Slicer) Add(f float64) {
	{
		s.Lis = append(s.Lis, f)
	}

	if len(s.Lis) > s.His {
		copy(s.Lis[0:], s.Lis[1:])
		s.Lis[len(s.Lis)-1] = 0
		s.Lis = s.Lis[:len(s.Lis)-1]
	}
}

func (s *Slicer) Avg() float64 {
	var sum float64

	for _, x := range s.Lis {
		sum += x
	}

	return sum / float64(len(s.Lis))
}

func (s *Slicer) Cng() float64 {
	inc, dec := change.Abs(s.Lis...)
	return inc - dec
}

func (s *Slicer) Equ() bool {
	if !s.Red() {
		return false
	}

	var fir float64
	{
		fir = s.Lis[0]
	}

	for _, x := range s.Lis {
		if x != fir {
			return false
		}
	}

	return true
}

func (s *Slicer) Max() float64 {
	var max float64
	{
		max = -math.MaxFloat64
	}

	for _, x := range s.Lis {
		if x > max {
			max = x
		}
	}

	return max
}

func (s *Slicer) Min() float64 {
	var min float64
	{
		min = +math.MaxFloat64
	}

	for _, x := range s.Lis {
		if x < min {
			min = x
		}
	}

	return min
}

func (s *Slicer) Red() bool {
	return len(s.Lis) == s.His
}
