package orderbook

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Orderbook_AskLevel(t *testing.T) {
	testCases := []struct {
		ob *Orderbook
		ix int
		ap json.Number
	}{
		// Case 0
		{
			ob: &Orderbook{
				ask: []Levels{
					{Price: "4"},
					{Price: "3"},
					{Price: "2"},
					{Price: "1"},
				},
			},
			ix: +1,
			ap: "1",
		},
		// Case 1
		{
			ob: &Orderbook{
				ask: []Levels{
					{Price: "4"},
					{Price: "3"},
					{Price: "2"},
					{Price: "1"},
				},
			},
			ix: +2,
			ap: "2",
		},
		// Case 2
		{
			ob: &Orderbook{
				ask: []Levels{
					{Price: "4"},
					{Price: "3"},
					{Price: "2"},
					{Price: "1"},
				},
			},
			ix: +3,
			ap: "3",
		},
		// Case 3
		{
			ob: &Orderbook{
				ask: []Levels{
					{Price: "4"},
					{Price: "3"},
					{Price: "2"},
					{Price: "1"},
				},
			},
			ix: +4,
			ap: "4",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var ap json.Number
			{
				ap = tc.ob.AskLevel(tc.ix).Price
			}

			if !reflect.DeepEqual(tc.ap, ap) {
				t.Fatalf("ap\n\n%s\n", cmp.Diff(tc.ap, ap))
			}
		})
	}
}

func Test_Orderbook_BidLevel(t *testing.T) {
	testCases := []struct {
		ob *Orderbook
		ix int
		bp json.Number
	}{
		// Case 0
		{
			ob: &Orderbook{
				bid: []Levels{
					{Price: "1"},
					{Price: "2"},
					{Price: "3"},
					{Price: "4"},
				},
			},
			ix: -1,
			bp: "1",
		},
		// Case 1
		{
			ob: &Orderbook{
				bid: []Levels{
					{Price: "1"},
					{Price: "2"},
					{Price: "3"},
					{Price: "4"},
				},
			},
			ix: -2,
			bp: "2",
		},
		// Case 2
		{
			ob: &Orderbook{
				bid: []Levels{
					{Price: "1"},
					{Price: "2"},
					{Price: "3"},
					{Price: "4"},
				},
			},
			ix: -3,
			bp: "3",
		},
		// Case 3
		{
			ob: &Orderbook{
				bid: []Levels{
					{Price: "1"},
					{Price: "2"},
					{Price: "3"},
					{Price: "4"},
				},
			},
			ix: -4,
			bp: "4",
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var bp json.Number
			{
				bp = tc.ob.BidLevel(tc.ix).Price
			}

			if !reflect.DeepEqual(tc.bp, bp) {
				t.Fatalf("bp\n\n%s\n", cmp.Diff(tc.bp, bp))
			}
		})
	}
}
