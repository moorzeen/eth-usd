package balance

import (
	"context"
	"math/big"
	"strings"

	"eth-usd/price"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

type TokenBalance struct {
	Symbol    string
	Address   common.Address
	Balance   decimal.Decimal
	Decimals  int32
	PriceFeed *price.ChainlinkPriceFeed
}

func GetTokenBalances(ctx context.Context, client *ethclient.Client, address common.Address) ([]TokenBalance, error) {
	var balances []TokenBalance

	// parse ABI for ERC20
	parsedABI, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}

	// check each supported token
	for _, token := range price.SupportedTokens {
		var balance decimal.Decimal

		if token.Symbol == "ETH" {
			// for ETH, get the balance directly
			ethBalance, err := client.BalanceAt(ctx, address, nil)
			if err != nil {
				return nil, err
			}
			balance = decimal.NewFromBigInt(ethBalance, -token.Decimals)
		} else {
			// for ERC20 tokens call balanceOf
			tokenAddress := common.HexToAddress(token.Address)
			data, err := parsedABI.Pack("balanceOf", address)
			if err != nil {
				return nil, err
			}

			msg := ethereum.CallMsg{
				To:   &tokenAddress,
				Data: data,
			}

			result, err := client.CallContract(ctx, msg, nil)
			if err != nil {
				continue // skip token if balance cannot be retrieved
			}

			var tokenBalance *big.Int
			err = parsedABI.UnpackIntoInterface(&tokenBalance, "balanceOf", result)
			if err != nil {
				continue
			}

			balance = decimal.NewFromBigInt(tokenBalance, -token.Decimals)
		}

		// add token to list only if balance > 0
		if !balance.IsZero() {
			priceFeed := price.NewChainlinkPriceFeed(client, token.PriceFeed)
			balances = append(balances, TokenBalance{
				Symbol:    token.Symbol,
				Address:   common.HexToAddress(token.Address),
				Balance:   balance,
				Decimals:  token.Decimals,
				PriceFeed: priceFeed,
			})
		}
	}

	return balances, nil
}
