package rsi

import "github.com/phoebetronic/model/pkg/change"

type RSI struct {
	sli *Slicer
	thr float64
}

func New(his int, thr float64) *RSI {
	return &RSI{
		sli: &Slicer{his: his},
		thr: thr,
	}
}

func (r *RSI) Active(upd float64) bool {
	return upd <= r.thr || upd >= 100-r.thr
}

func (r *RSI) Update(pri float64) float64 {
	{
		r.sli.Add(pri)
	}

	if !r.sli.Red() {
		return 0
	}

	var inc float64
	var dec float64
	{
		inc, dec = change.Per(r.sli.lis...)
	}

	return 100 - (100 / (1 + ((inc / float64(r.sli.his)) / (dec / float64(r.sli.his)))))
}
