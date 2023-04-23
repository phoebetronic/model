package main

import (
	"github.com/phoebetronic/model/cmd"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/tracer"
)

func main() {
	err := mainE()
	if err != nil {
		tracer.Panic(err)
	}
}

func mainE() error {
	var err error

	var c *cobra.Command
	{
		c, err = cmd.New()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		err = c.Execute()
		if err != nil {
			return tracer.Mask(err)
		}
	}

	return nil
}
