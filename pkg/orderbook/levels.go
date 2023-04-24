package orderbook

import "encoding/json"

type Levels struct {
	Price json.Number
	Size  json.Number
}
