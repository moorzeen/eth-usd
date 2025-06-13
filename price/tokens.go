package price

type TokenConfig struct {
	Symbol    string
	Address   string
	PriceFeed string
	Decimals  int32
}

var SupportedTokens = []TokenConfig{
	{
		Symbol:    "ETH",
		Address:   "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
		PriceFeed: "0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419",
		Decimals:  18,
	},
	{
		Symbol:    "WETH",
		Address:   "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		PriceFeed: "0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419",
		Decimals:  18,
	},
	{
		Symbol:    "USDC",
		Address:   "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
		PriceFeed: "0x8fFfFfd4AfB6115b954Bd326cbe7B4BA576818f6",
		Decimals:  6,
	},
	{
		Symbol:    "USDT",
		Address:   "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		PriceFeed: "0x3E7d1eAB13ad0104d2750B8863b489D65364e32D",
		Decimals:  6,
	},
	{
		Symbol:    "DAI",
		Address:   "0x6B175474E89094C44Da98b954EedeAC495271d0F",
		PriceFeed: "0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9",
		Decimals:  18,
	},
	{
		Symbol:    "LINK",
		Address:   "0x514910771AF9Ca656af840dff83E8264EcF986CA",
		PriceFeed: "0x2c1d072e956AFFC0D435Cb7AC38EF18d24d9127c",
		Decimals:  18,
	},
	{
		Symbol:    "WBTC",
		Address:   "0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599",
		PriceFeed: "0xF4030086522a5bEEa4988F8cA5B36dbC97BeE88c",
		Decimals:  8,
	},
	{
		Symbol:    "AAVE",
		Address:   "0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9",
		PriceFeed: "0x547a514d5e3769680Ce22B2361c10Ea13619e8a9",
		Decimals:  18,
	},
	{
		Symbol:    "UNI",
		Address:   "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984",
		PriceFeed: "0x553303d460EE0afB37EdFf9bE42922D8FF63220e",
		Decimals:  18,
	},
	{
		Symbol:    "SNX",
		Address:   "0xC011a73ee8576Fb46F5E1c5751cA3B9Fe0af2a6F",
		PriceFeed: "0xDC3EA94CD0AC27d9A86C180091e7f78C683d3699",
		Decimals:  18,
	},
}
