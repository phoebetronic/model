package orderbook

import "github.com/montanaflynn/stats"

func Correlation(a *Orderbook, b *Orderbook) float64 {
	var err error

	var cor float64
	{
		cor, err = stats.Correlation(a.Prices(), b.Prices())
		if err != nil {
			panic(err)
		}
	}

	return cor
}
