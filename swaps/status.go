package swaps

import (
	"errors"
	"math"

	uniswap "github.com/hirokimoto/uniswap-api"
	"github.com/hirokimoto/uniswap-api/swap"
)

func State(swaps uniswap.Swaps) (string, string, error) {
	old, _ := swap.Old(swaps.Data.Swaps[0])
	price, _ := swap.Price(swaps.Data.Swaps[0])
	average := AveragePrice(swaps.Data.Swaps)

	// Filter out some tokens which is in the active trading in recent3 days.
	if old < 3*24 && price > 0.0001 {
		slope, _, _ := Regression(swaps)
		var isGoingUp = slope > 0
		var isGoingDown = slope < 0
		var isStable = math.Abs((average-price)/price) < 0.1
		var isUnStable = math.Abs((average-price)/price) > 0.1

		target := ""
		updown := ""
		if isUnStable {
			target = "unstable"
		}
		if isStable {
			target = "stable"
		}
		if isGoingUp {
			updown = "up"
		}
		if isGoingDown {
			updown = "down"
		}

		return target, updown, nil
	}

	return "", "", errors.New("dead token")
}
