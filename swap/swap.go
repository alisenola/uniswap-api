package swap

import (
	"fmt"
	"strconv"
	"time"

	uniswap "github.com/hirokimoto/uniswap-api"
)

// Name returns token name in a swap.
func Name(swap uniswap.Swap) (name string) {
	if swap.Pair.Token0.Symbol == "WETH" {
		name = swap.Pair.Token1.Name
	} else {
		name = swap.Pair.Token0.Name
	}
	return name
}

// Price returns token price of in a swap.
func Price(swap uniswap.Swap) (price float64, target string) {
	amountUSD, _ := strconv.ParseFloat(swap.AmountUSD, 32)
	amountToken, _ := strconv.ParseFloat(swap.Amount0Out, 32)

	if swap.Pair.Token0.Symbol == "WETH" {
		if swap.Amount0In != "0" && swap.Amount1Out != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount1Out, 32)
			target = "BUY"
		} else if swap.Amount0Out != "0" && swap.Amount1In != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount1In, 32)
			target = "SELL"
		}
	} else {
		if swap.Amount0Out != "0" && swap.Amount1In != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount0Out, 32)
			target = "BUY"
		} else if swap.Amount0In != "0" && swap.Amount1Out != "0" {
			amountToken, _ = strconv.ParseFloat(swap.Amount0In, 32)
			target = "SELL"
		}
	}

	price = amountUSD / amountToken
	return price, target
}

// Old returns how many hours have been passed from the latest swap.
func Old(swap uniswap.Swap) (float64, string) {
	timestamp, _ := strconv.ParseInt(swap.Timestamp, 10, 64)
	unixTimestamp := time.Unix(timestamp, 0)
	now := time.Now()
	old := now.Sub(unixTimestamp)
	olds := ""
	if old.Minutes() <= 1 {
		olds = fmt.Sprintf("%f", old.Seconds()) + "seconds"
	}
	if old.Minutes() > 1 && old.Hours() < 1 {
		olds = fmt.Sprintf("%f", old.Minutes()) + "minutes"
	}
	if old.Hours() >= 1 {
		olds = fmt.Sprintf("%f", old.Hours()) + "hours"
	}
	return old.Hours(), olds
}

func Trade(swap uniswap.Swap) (price float64, target string, amount string, amount1 string, amount2 string) {
	price, target = Price(swap)
	_amount, _ := strconv.ParseFloat(swap.AmountUSD, 32)
	amount = fmt.Sprintf("$%.2f", _amount)

	_amount1 := 0.1
	if swap.Amount0In == "0" {
		_amount1, _ = strconv.ParseFloat(swap.Amount0Out, 32)
	} else {
		_amount1, _ = strconv.ParseFloat(swap.Amount0In, 32)
	}
	amount1 = fmt.Sprintf("%.4f %s", _amount1, swap.Pair.Token0.Symbol)

	_amount2 := 0.1
	if swap.Amount1In == "0" {
		_amount2, _ = strconv.ParseFloat(swap.Amount1Out, 32)
	} else {
		_amount2, _ = strconv.ParseFloat(swap.Amount1In, 32)
	}
	amount2 = fmt.Sprintf("%.4f %s", _amount2, swap.Pair.Token1.Symbol)

	return price, target, amount, amount1, amount2
}
