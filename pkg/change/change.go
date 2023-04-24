package change

import "github.com/phoebetronic/model/pkg/round"

func Abs(lis ...float64) (float64, float64) {
	var inc float64
	var dec float64

	for i := 1; i < len(lis); i++ {
		a := lis[i-1]
		b := lis[i]

		if a > b {
			dec += a - b
		}

		if b > a {
			inc += b - a
		}
	}

	return round.RoundP(inc, 5), round.RoundP(dec, 5)
}

func Avg(lis ...float64) (float64, float64) {
	var inc float64
	var dec float64
	{
		inc, dec = Abs(lis...)
	}

	{
		inc = inc / float64(len(lis))
		dec = dec / float64(len(lis))
	}

	return round.RoundP(inc, 5), round.RoundP(dec, 5)
}

func Per(lis ...float64) (float64, float64) {
	var inc float64
	var dec float64

	for i := 1; i < len(lis); i++ {
		a := lis[i-1]
		b := lis[i]

		if a > b {
			del := a - b
			dec += del * 100 / a
		}

		if b > a {
			del := b - a
			inc += del * 100 / a
		}
	}

	{
		inc = inc / float64(len(lis))
		dec = dec / float64(len(lis))
	}

	return round.RoundP(inc, 5), round.RoundP(dec, 5)
}
