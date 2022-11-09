package swaps

import (
	"strconv"
	"time"

	uniswap "github.com/hirokimoto/uniswap-api"
	"github.com/hirokimoto/uniswap-api/swap"
)

// Duration returns how long it would take in the swaps.
func Duration(swaps uniswap.Swaps) (time.Time, time.Time, float64) {
	var duration float64
	if swaps.Data.Swaps != nil && len(swaps.Data.Swaps) > 0 {
		timestamp1, _ := strconv.ParseInt(swaps.Data.Swaps[0].Timestamp, 10, 64)
		timestapm2, _ := strconv.ParseInt(swaps.Data.Swaps[len(swaps.Data.Swaps)-1].Timestamp, 10, 64)
		first := time.Unix(timestamp1, 0)
		last := time.Unix(timestapm2, 0)
		unixduration := first.Sub(last)
		duration = unixduration.Hours()
		return first, last, duration
	}
	now := time.Now()
	return now, now, 0.0
}

// MinMaxPrice returns min and max price of swaps.
func MinMaxPrice(swaps uniswap.Swaps) (float64, float64) {
	min := 0.0
	max := 0.0
	for _, item := range swaps.Data.Swaps {
		price, _ := swap.Price(item)
		if min == 0.0 || max == 0.0 {
			min = price
			max = price
		}
		if price < min {
			min = price
		}
		if price > max {
			max = price
		}
	}
	return min, max
}

// AveragePrice returns average price of swaps.
func AveragePrice(swaps []uniswap.Swap) float64 {
	sum := 0.0
	for _, item := range swaps {
		price, _ := swap.Price(item)
		sum += price
	}
	return sum / float64(len(swaps))
}
