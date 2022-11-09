package uniswapapi

import (
	"encoding/json"
	"strconv"
	"sync"
	"time"
)

func RequestPairs(target chan string, limit int, skip int) {
	query := QueryPairs(limit, skip)
	request(query, target)
}

func SwapsByDays(target chan string, limit int, address string) {
	var results Swaps
	var temp Swaps
	i := 0
	now := time.Now()
	for {
		ch := make(chan string)
		query := QuerySwaps(address, 1000, i*1000)
		go request(query, ch)

		msg := <-ch
		var swaps Swaps
		json.Unmarshal([]byte(msg), &swaps)
		if len(swaps.Data.Swaps) == 0 {
			tg, _ := json.Marshal(results)
			target <- string(tg)
			return
		}
		endIndex := len(swaps.Data.Swaps) - 1
		if endIndex < 0 {
			endIndex = 0
		}

		lastInt, _ := strconv.ParseInt(swaps.Data.Swaps[endIndex].Timestamp, 10, 64)
		lastTime := time.Unix(lastInt, 0)
		period := now.Sub(lastTime)

		if period.Hours() >= 24*float64(limit) {
			temp = swaps
			break
		} else {
			results.Data.Swaps = append(results.Data.Swaps, swaps.Data.Swaps...)
		}
		i += 1
	}

	for i = 0; i < len(temp.Data.Swaps); i++ {
		lastInt, _ := strconv.ParseInt(temp.Data.Swaps[i].Timestamp, 10, 64)
		lastTime := time.Unix(lastInt, 0)
		period := now.Sub(lastTime)
		if period.Hours() > 24*float64(limit) {
			break
		} else {
			results.Data.Swaps = append(results.Data.Swaps, temp.Data.Swaps[i])
		}
	}

	tg, _ := json.Marshal(results)
	target <- string(tg)
}

func SwapsByCounts(target chan string, limit int, address string) {
	var results Swaps
	length := limit/1000 + 1
	var wg sync.WaitGroup
	wg.Add(length)

	for i := 0; i < length; i++ {
		ch := make(chan string)
		counts := 1000
		skip := i * 1000
		if limit < 1000 {
			counts = limit
		}
		if (i+1)*1000 > limit {
			counts = limit - i*counts
		}

		query := QuerySwaps(address, counts, skip)
		go request(query, ch)

		msg := <-ch
		var swaps Swaps
		json.Unmarshal([]byte(msg), &swaps)
		results.Data.Swaps = append(results.Data.Swaps, swaps.Data.Swaps...)

		wg.Done()
	}

	wg.Wait()

	tg, _ := json.Marshal(results)
	target <- string(tg)
}
