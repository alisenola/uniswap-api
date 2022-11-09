package uniswapapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func request(query map[string]string, target chan string) {
	jsonQuery, _ := json.Marshal(query)
	request, _ := http.NewRequest("POST", "https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2", bytes.NewBuffer(jsonQuery))
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		target <- ""
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		target <- string(data)
	}
	defer response.Body.Close()
}

func QueryBundles() map[string]string {
	return map[string]string{
		"query": `
			query bundles {
				bundles(where: { id: "1" }) {
					ethPrice
				}
			}
		`,
	}
}

func QuertyToken(address string) map[string]string {
	query := fmt.Sprintf(`
		query tokens {
			tokens(where: { id: "%s" }) {
				id
				name
				symbol
				derivedETH
				totalLiquidity
			}
		}
	`, address)
	return map[string]string{"query": query}
}

func QuerySwaps(address string, limit int, skip int) map[string]string {
	query := fmt.Sprintf(`
		query swaps {
			swaps(first: %d, skip: %d, orderBy: timestamp, orderDirection: desc, where:
				{ pair: "%s" }
			) {
				pair {
					token0 {
						symbol
						name
						derivedETH
					}
					token1 {
						symbol
						name
						derivedETH
					}
				}
				amount0In
				amount0Out
				amount1In
				amount1Out
				amountUSD
				to
				timestamp
				id
			}
		}
	`, limit, skip, address)
	return map[string]string{"query": query}
}

func QueryPairs(limit int, skip int) map[string]string {
	query := fmt.Sprintf(`
		query pairs {
			pairs(first: %d, skip: %d, orderBy: reserveUSD, orderDirection: desc) {
				id,
				token0 {
					symbol
				},
				token1 {
					symbol
				},
				token0Price,
				token1Price,
			}
		}
		`, limit, skip)
	return map[string]string{"query": query}
}
