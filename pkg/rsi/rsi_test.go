package rsi

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_RSI_Active(t *testing.T) {
	testCases := []struct {
		upd float64
		act bool
	}{
		// Case 0
		{
			upd: 32,
			act: false,
		},
		// Case 1
		{
			upd: 72,
			act: false,
		},
		// Case 2
		{
			upd: 12,
			act: true,
		},
		// Case 3
		{
			upd: 82,
			act: true,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var r *RSI
			{
				r = New(7, 20)
			}

			var a bool
			{
				a = r.Active(tc.upd)
			}

			if a != tc.act {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.act, a))
			}
		})
	}
}

func Test_RSI_Update(t *testing.T) {
	testCases := []struct {
		pri []float64
		val float64
	}{
		// Case 0
		{
			pri: []float64{
				32,
			},
			val: 0,
		},
		// Case 1
		{
			pri: []float64{
				32,
				34,
			},
			val: 0,
		},
		// Case 2
		{
			pri: []float64{
				32,
				34,
				37,
				33,
				31,
				36,
				39,
				38,
			},
			val: 63.13547001113969,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var r *RSI
			{
				r = New(7, 20)
			}

			var v float64
			for _, x := range tc.pri {
				v = r.Update(x)
			}

			if v != tc.val {
				t.Fatalf("\n\n%s\n", cmp.Diff(tc.val, v))
			}
		})
	}
}
