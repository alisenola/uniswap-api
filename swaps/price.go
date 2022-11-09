package swaps

import (
	"strconv"
	"time"

	uniswap "github.com/hirokimoto/uniswap-api"
	"github.com/hirokimoto/uniswap-api/swap"
)

// WholePriceChanges returns how much has been changed in the whole swaps.
func WholePriceChanges(swaps uniswap.Swaps) (price float64, change float64) {
	if swaps.Data.Swaps != nil && len(swaps.Data.Swaps) > 0 {
		first, _ := swap.Price(swaps.Data.Swaps[0])
		last, _ := swap.Price(swaps.Data.Swaps[len(swaps.Data.Swaps)-1])
		change = first - last
		return first, change
	}
	return 0.0, 0.0
}

// LastPriceChanges returns the amount of price changes of the last 2 swaps.
// price, change, duration
func LastPriceChanges(swaps uniswap.Swaps) (float64, float64, float64) {
	if swaps.Data.Swaps != nil && len(swaps.Data.Swaps) > 0 {
		first, _ := swap.Price(swaps.Data.Swaps[0])
		second, _ := swap.Price(swaps.Data.Swaps[len(swaps.Data.Swaps)-1])
		change := first - second

		timestamp1, _ := strconv.ParseInt(swaps.Data.Swaps[0].Timestamp, 10, 64)
		timestapm2, _ := strconv.ParseInt(swaps.Data.Swaps[len(swaps.Data.Swaps)-1].Timestamp, 10, 64)
		time1 := time.Unix(timestamp1, 0)
		time2 := time.Unix(timestapm2, 0)
		unixduration := time1.Sub(time2)
		duration := unixduration.Hours()
		return first, change, duration
	}
	return 0.0, 0.0, 0.0
}

// Average returns the average price of swaps
func Average(swaps []uniswap.Swap) float64 {
	sum := 0.0
	for _, item := range swaps {
		price, _ := swap.Price(item)
		sum += price
	}
	return sum / float64(len(swaps))
}
