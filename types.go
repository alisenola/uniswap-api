package uniswapapi

type Bundles struct {
	Data struct {
		Bundles []struct {
			EthPrice string `json:"ethPrice"`
		} `json:"bundles"`
	} `json:"data"`
}

type Tokens struct {
	Data struct {
		Tokens []struct {
			Id             string `json:"id"`
			Name           string `json:"name"`
			Symbol         string `json:"symbol"`
			DerivedETH     string `json:"derivedETH"`
			TotalLiquidity string `json:"totalLiquidity"`
		} `json:"tokens"`
	} `json:"data"`
}

type Swap struct {
	Amount0In  string `json:"amount0In"`
	Amount0Out string `json:"amount0Out"`
	Amount1In  string `json:"amount1In"`
	Amount1Out string `json:"amount1Out"`
	AmountUSD  string `json:"amountUSD"`
	Id         string `json:"id"`
	Pair       struct {
		Token0 struct {
			Symbol     string `json:"symbol"`
			Name       string `json:"name"`
			DerivedETH string `json:"derivedETH"`
		} `json:"token0"`
		Token1 struct {
			Symbol     string `json:"symbol"`
			Name       string `json:"name"`
			DerivedETH string `json:"derivedETH"`
		} `json:"token1"`
	} `json:"pair"`
	Timestamp string `json:"timestamp"`
	To        string `json:"to"`
}

type Swaps struct {
	Data struct {
		Swaps []Swap
	}
}

type Pair struct {
	Id     string `json:"id"`
	Token0 struct {
		Symbol string `json:"symbol"`
	} `json:"token0"`
	Token1 struct {
		Symbol string `json:"symbol"`
	} `json:"token1"`
	Token0Price string `json:"token0Price"`
	Token1Price string `json:"token1Price"`
}

type Pairs struct {
	Data struct {
		Pairs []Pair
	}
}
