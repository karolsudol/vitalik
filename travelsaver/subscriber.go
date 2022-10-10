package travelsaver

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func PrintEvents()  {

	client, err := ethclient.Dial("wss://alfajores-forno.celo-testnet.org/ws")
	if err != nil {
	log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xb93A7A27651F8CfF100a32D8762E1eA51Fc56d72")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }


    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            fmt.Println(vLog) // pointer to event log
        }
    }

}


