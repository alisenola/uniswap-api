package swaps

import (
	"strconv"
	"time"

	uniswap "github.com/hirokimoto/uniswap-api"
)

// CheckUp lets you know if up sticks is larger than down sticks.
func CheckUp(swaps uniswap.Swaps) bool {
	avg := 0.0
	checkUp := 0
	checkDown := 0
	counts := len(swaps.Data.Swaps) - 1
	entryTime, _ := strconv.ParseInt(swaps.Data.Swaps[counts].Timestamp, 10, 64)

	var empty []uniswap.Swap
	var temp []uniswap.Swap

	for i := counts; i > 0; i-- {
		createdat, _ := strconv.ParseInt(swaps.Data.Swaps[i].Timestamp, 10, 64)
		t := time.Unix(createdat, 0).UTC()
		createdAt := t.Round(time.Hour).UTC().Unix()
		if entryTime != createdAt {
			cavg := AveragePrice(temp)
			if cavg > avg {
				checkUp += 1
			} else {
				checkDown += 1
			}
			avg = cavg
			entryTime = createdAt
			temp = empty
		}
		temp = append(temp, swaps.Data.Swaps[i])
	}

	return checkUp > 2*checkDown
}

// CheckDown let you know if down sticks is larger than up sticks.
func CheckDown(swaps uniswap.Swaps) bool {
	avg := 10000000.0
	checkUp := 0
	checkDown := 0
	counts := len(swaps.Data.Swaps) - 1
	entryTime, _ := strconv.ParseInt(swaps.Data.Swaps[counts].Timestamp, 10, 64)

	var empty []uniswap.Swap
	var temp []uniswap.Swap

	for i := counts; i > 0; i-- {
		createdat, _ := strconv.ParseInt(swaps.Data.Swaps[i].Timestamp, 10, 64)
		t := time.Unix(createdat, 0).UTC()
		createdAt := t.Round(time.Hour).UTC().Unix()
		if entryTime != createdAt {
			cavg := AveragePrice(temp)
			if cavg > avg {
				checkUp += 1
			} else {
				checkDown += 1
			}
			avg = cavg
			entryTime = createdAt
			temp = empty
		}
		temp = append(temp, swaps.Data.Swaps[i])
	}

	return checkUp < 2*checkDown
}
