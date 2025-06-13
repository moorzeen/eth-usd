package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"eth-usd/balance"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

const (
	ethereumNode = "https://eth.llamarpc.com"
)

func main() {
	address := flag.String("address", "", "Ethereum address to check")
	flag.Parse()

	if *address == "" {
		log.Fatal("provide an Ethereum address using -address flag")
	}

	client, err := ethclient.Dial(ethereumNode)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	balances, err := balance.GetTokenBalances(context.Background(), client, common.HexToAddress(*address))
	if err != nil {
		log.Fatalf("failed to get token balances: %v", err)
	}

	fmt.Printf("Address: %s\n\n", *address)
	fmt.Println("Token balances:")

	var totalUsdValue float64
	for _, b := range balances {
		price, err := b.PriceFeed.GetLatestPrice(context.Background())
		if err != nil {
			fmt.Printf("%s: %s (Failed to get price)\n", b.Symbol, b.Balance.String())
			continue
		}

		if b.Symbol == "USDC" || b.Symbol == "USDT" || b.Symbol == "DAI" {
			price = decimal.NewFromInt(1)
		}

		usdValue := b.Balance.Mul(price).InexactFloat64()
		totalUsdValue += usdValue
		fmt.Printf("%s: %s (≈ $%.2f)\n", b.Symbol, b.Balance.String(), usdValue)
	}

	fmt.Printf("\nTotal USD Value: $%.2f\n", totalUsdValue)
}
