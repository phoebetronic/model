package arb

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cicdteam/go-deribit"
	"github.com/cicdteam/go-deribit/models"
	"github.com/phoebetronic/model/pkg/orderbook"
	"github.com/phoebetronic/model/pkg/slicer"
	"github.com/spf13/cobra"
)

const (
	thr = 0.3
)

type run struct{}

func (r *run) run(cmd *cobra.Command, args []string) {
	fmt.Println()
	fmt.Println("|===========================================|")
	fmt.Println("|      wss://www.deribit.com/ws/api/v2      |")
	fmt.Println("|===========================================|")
	fmt.Println()

	var err error

	var erc chan error
	var clo chan bool
	{
		erc = make(chan error)
		clo = make(chan bool)
	}

	var cli *deribit.Exchange
	{
		cli, err = deribit.NewExchange(false, erc, clo)
		if err != nil {
			panic(err)
		}

		err = cli.Connect()
		if err != nil {
			panic(err)
		}

		defer cli.Close()
	}

	go func() {
		err := <-erc
		clo <- true
		panic(err)
	}()

	var prp *orderbook.Orderbook
	{
		prp = orderbook.New()
	}

	{
		var msg chan *models.BookNotification
		{
			msg, err = cli.SubscribeBookGroup("ETH-PERPETUAL", "none", "10", "100ms")
			if err != nil {
				panic(err)
			}
		}

		go func() {
			for x := range msg {
				err = prp.Middleware(x)
				if err != nil {
					panic(err)
				}
			}
		}()
	}

	var spt *orderbook.Orderbook
	{
		spt = orderbook.New()
	}

	{
		var msg chan *models.BookNotification
		{
			msg, err = cli.SubscribeBookGroup("ETH_USDC-PERPETUAL", "none", "10", "100ms")
			if err != nil {
				panic(err)
			}
		}

		go func() {
			for x := range msg {
				err = spt.Middleware(x)
				if err != nil {
					panic(err)
				}
			}
		}()
	}

	var tic *time.Ticker
	{
		tic = time.NewTicker(200 * time.Millisecond)
	}

	for {
		if prp.Empty() || spt.Empty() {
			time.Sleep(100 * time.Millisecond)
		} else {
			break
		}
	}

	var pdr *slicer.Slicer
	var sdr *slicer.Slicer
	{
		pdr = &slicer.Slicer{His: 10}
		sdr = &slicer.Slicer{His: 10}
	}

	var equ *slicer.Slicer
	{
		equ = &slicer.Slicer{His: 10}
	}

	var bal float64

	for range tic.C {
		{
			pdr.Add(prp.MidPri())
			sdr.Add(spt.MidPri())
			equ.Add(spt.MidPri())
		}

		if !pdr.Red() || equ.Equ() {
			continue
		}

		pcg := pdr.Cng()
		scg := sdr.Cng()

		if pcg < +thr && scg < +thr {
			var sho float64
			{
				deblog(prp, spt)
				sho = musflo(spt.BidLevel(-1).Price)
			}

			var cou int
			for {
				if cou >= 5 {
					break
				}

				{
					cou++
				}

				{
					time.Sleep(5000 * time.Millisecond)
				}

				if sho == musflo(spt.BidLevel(-1).Price) {
					continue
				} else {
					break
				}
			}

			var lon float64
			{
				deblog(prp, spt)
				lon = musflo(spt.AskLevel(+1).Price)
			}

			{
				bal += sho - lon
			}

			{
				fmt.Printf("sho del %.2f (bid fir %.2f - ask las %.2f)\n", bal, sho, lon)
				fmt.Println()
				fmt.Println()
				fmt.Println()
			}

		} else if pcg > -thr && scg > -thr {
			var lon float64
			{
				deblog(prp, spt)
				lon = musflo(spt.AskLevel(+1).Price)
			}

			var cou int
			for {
				if cou >= 5 {
					break
				}

				{
					cou++
				}

				{
					time.Sleep(5000 * time.Millisecond)
				}

				if lon == musflo(spt.AskLevel(+1).Price) {
					continue
				} else {
					break
				}
			}

			var sho float64
			{
				deblog(prp, spt)
				sho = musflo(spt.BidLevel(-1).Price)
			}

			{
				bal += sho - lon
			}

			{
				fmt.Printf("lon del %.2f (bid las %.2f - ask fir %.2f)\n", bal, sho, lon)
				fmt.Println()
				fmt.Println()
				fmt.Println()
			}
		}

		{
			equ = &slicer.Slicer{His: 10}
			pdr = &slicer.Slicer{His: 10}
			sdr = &slicer.Slicer{His: 10}
		}
	}
}

func musflo(num json.Number) float64 {
	flo, err := num.Float64()
	if err != nil {
		panic(err)
	}

	return flo
}

func deblog(prp *orderbook.Orderbook, spt *orderbook.Orderbook) {
	var now time.Time
	{
		now = time.Now().Round(10 * time.Millisecond)
	}

	fmt.Printf("%32s\n", now)

	fmt.Println()

	fmt.Printf("%7s    %7s\n", prp.AskLevel(+5).Price, spt.AskLevel(+5).Price)
	fmt.Printf("%7s    %7s\n", prp.AskLevel(+4).Price, spt.AskLevel(+4).Price)
	fmt.Printf("%7s    %7s\n", prp.AskLevel(+3).Price, spt.AskLevel(+3).Price)
	fmt.Printf("%7s    %7s\n", prp.AskLevel(+2).Price, spt.AskLevel(+2).Price)
	fmt.Printf("%7s    %7s\n", prp.AskLevel(+1).Price, spt.AskLevel(+1).Price)
	fmt.Printf("------------------\n")
	fmt.Printf("%7s    %7s\n", prp.BidLevel(-1).Price, spt.BidLevel(-1).Price)
	fmt.Printf("%7s    %7s\n", prp.BidLevel(-2).Price, spt.BidLevel(-2).Price)
	fmt.Printf("%7s    %7s\n", prp.BidLevel(-3).Price, spt.BidLevel(-3).Price)
	fmt.Printf("%7s    %7s\n", prp.BidLevel(-4).Price, spt.BidLevel(-4).Price)
	fmt.Printf("%7s    %7s\n", prp.BidLevel(-5).Price, spt.BidLevel(-5).Price)

	fmt.Println()
}
