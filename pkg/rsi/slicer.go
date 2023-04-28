package rsi

type Slicer struct {
	his int
	lis []float64
}

func (s *Slicer) Add(f float64) {
	{
		s.lis = append(s.lis, f)
	}

	if len(s.lis) > s.his {
		copy(s.lis[0:], s.lis[1:])
		s.lis[len(s.lis)-1] = 0
		s.lis = s.lis[:len(s.lis)-1]
	}
}

func (s *Slicer) Red() bool {
	return len(s.lis) == s.his
}
