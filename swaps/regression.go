package swaps

import (
	regression "github.com/gaillard/go-online-linear-regression/v1"
	uniswap "github.com/hirokimoto/uniswap-api"
	"github.com/hirokimoto/uniswap-api/swap"
)

func Regression(swaps uniswap.Swaps) (float64, float64, float64) {
	r := regression.New(7)

	for i := 0; i < len(swaps.Data.Swaps); i++ {
		item := swaps.Data.Swaps[i]
		price, _ := swap.Price(item)
		r.Add(float64(i), price)
	}

	slope, intercept, stdError := r.CalculateWithStdError()
	return slope, intercept, stdError
}
