package orderbook

import (
	"encoding/json"
	"sort"
	"strconv"
	"sync"
	"time"
)

// https://docs.deribit.com/#book-instrument_name-group-depth-interval
type Orderbook struct {
	ask []Levels
	bid []Levels
	mut sync.Mutex
	tim time.Time
}

func New() *Orderbook {
	return &Orderbook{}
}

func (o *Orderbook) AskLevel(ind int) Levels {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	return o.ask[len(o.ask)-ind]
}

func (o *Orderbook) BidLevel(ind int) Levels {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	return o.bid[(-ind)-1]
}

func (o *Orderbook) Empty() bool {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	return len(o.ask) == 0 || len(o.bid) == 0
}

func (o *Orderbook) MarshalJSON() ([]byte, error) {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	return json.Marshal(&struct {
		Ask []Levels  `json:"ask"`
		Bid []Levels  `json:"bid"`
		Tim time.Time `json:"tim"`
	}{
		Ask: o.ask,
		Bid: o.bid,
		Tim: o.tim,
	})
}

func (o *Orderbook) Middleware(upd Response) error {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	o.ask = make([]Levels, len(upd.Asks))
	o.bid = make([]Levels, len(upd.Bids))
	o.tim = time.Unix(int64(upd.Timestamp), 0)

	for i, x := range upd.Asks {
		o.ask[i] = Levels{Price: musnum(x[0]), Size: musnum(x[1])}
	}
	sort.Slice(o.ask, func(i, j int) bool {
		return o.ask[i].Price > o.ask[j].Price
	})

	for i, x := range upd.Bids {
		o.bid[i] = Levels{Price: musnum(x[0]), Size: musnum(x[1])}
	}
	sort.Slice(o.bid, func(i, j int) bool {
		return o.bid[i].Price > o.bid[j].Price
	})

	return nil
}

func (o *Orderbook) MidPri() float64 {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	ask, err := o.ask[len(o.ask)-1].Price.Float64()
	if err != nil {
		panic(err)
	}
	bid, err := o.bid[0].Price.Float64()
	if err != nil {
		panic(err)
	}

	return (bid + ask) / 2
}

func (o *Orderbook) Prices() []float64 {
	{
		o.mut.Lock()
		defer o.mut.Unlock()
	}

	var pri []float64

	for _, x := range o.ask {
		flo, err := x.Price.Float64()
		if err != nil {
			panic(err)
		}

		pri = append(pri, flo)
	}

	for _, x := range o.bid {
		flo, err := x.Price.Float64()
		if err != nil {
			panic(err)
		}

		pri = append(pri, flo)
	}

	return pri
}

func musnum(flo float64) json.Number {
	return json.Number(strconv.FormatFloat(flo, 'f', -1, 64))
}
