package arb

import (
	"github.com/spf13/cobra"
)

const (
	use = "arb"
	sho = "Execute a specific arbitrage strategy."
	lon = "Execute a specific arbitrage strategy."
)

type Config struct{}

func New(config Config) (*cobra.Command, error) {
	var c *cobra.Command
	{
		c = &cobra.Command{
			Use:   use,
			Short: sho,
			Long:  lon,
			Run:   (&run{}).run,
		}
	}

	return c, nil
}
