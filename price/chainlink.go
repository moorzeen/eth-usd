package price

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

type ChainlinkPriceFeed struct {
	client   *ethclient.Client
	feedAddr common.Address
	decimals int32
	abi      abi.ABI
}

func NewChainlinkPriceFeed(client *ethclient.Client, feedAddress string) *ChainlinkPriceFeed {
	parsedABI, err := abi.JSON(strings.NewReader(ChainlinkPriceFeedABI))
	if err != nil {
		panic(err)
	}

	return &ChainlinkPriceFeed{
		client:   client,
		feedAddr: common.HexToAddress(feedAddress),
		decimals: 8, // standard number of decimals
		abi:      parsedABI,
	}
}

func (p *ChainlinkPriceFeed) GetLatestPrice(ctx context.Context) (decimal.Decimal, error) {
	callOpts := &bind.CallOpts{Context: ctx}

	latestRoundData, err := p.getLatestRoundData(callOpts)
	if err != nil {
		return decimal.Zero, err
	}

	// convert price to decimal
	price := decimal.NewFromBigInt(latestRoundData.Answer, -p.decimals)
	return price, nil
}

// get the data of the latest round
func (p *ChainlinkPriceFeed) getLatestRoundData(opts *bind.CallOpts) (*struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {

	// prepare data for the call
	data, err := p.abi.Pack("latestRoundData")
	if err != nil {
		return nil, err
	}

	msg := ethereum.CallMsg{
		To:   &p.feedAddr,
		Data: data,
	}

	result, err := p.client.CallContract(opts.Context, msg, nil)
	if err != nil {
		return nil, err
	}

	var roundData struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	}

	err = p.abi.UnpackIntoInterface(&roundData, "latestRoundData", result)
	if err != nil {
		return nil, err
	}

	return &roundData, nil
}
