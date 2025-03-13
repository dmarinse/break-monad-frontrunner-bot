package main

import (
	"fmt"
	"os"

	"github.com/FastLane-Labs/break-monad-frontrunner-bot/contract/frontrunner"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	frontrunnerContractAddress = common.HexToAddress("0x9EaBA701a49adE7525dFfE338f0C7E06Eca7Cf07")
)

func main() {
	ethClient, err := ethclient.Dial(os.Getenv("RPC_URL"))
	if err != nil {
		fmt.Println("failed to connect to Ethereum client", "error", err)
		os.Exit(1)
	}

	frontrunnerContract, err := frontrunner.NewFrontrunner(frontrunnerContractAddress, ethClient)
	if err != nil {
		fmt.Println("failed to create Frontrunner contract", "error", err)
		os.Exit(1)
	}

	scores, err := frontrunnerContract.GetScores(nil)
	if err != nil {
		fmt.Println("failed to get scores", "error", err)
		os.Exit(1)
	}

	fmt.Println("Address", "Wins/Losses") 

	for _, score := range scores {
		fmt.Printf("%s %s/%s\n", score.Address.Hex(), score.Wins, score.Losses)
	}
}
