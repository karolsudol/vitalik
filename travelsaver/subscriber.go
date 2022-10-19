package travelsaver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func prettyPrint(d ...interface{})  {
    b, err := json.MarshalIndent(d, "", "  ")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(b))
}



type LogCreatedPaymentPlan struct {
    ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
}

type LogCreatedTravelPlan struct {
	ID         *big.Int
	Owner      common.Address
	TravelPlan TravelSaverTravelPlan
}

type LogStartPaymentPlanInterval struct {
    ID         *big.Int
	CallableOn *big.Int
	Amount     *big.Int
	IntervalNo *big.Int
}

type LogContributeToTravelPlan struct {
    ID          *big.Int
	Contributor common.Address
	Amount      *big.Int

}

type LogClaimTravelPlan struct {
    ID  *big.Int
    
}


type LogTransfer struct {
    From   common.Address
	To     common.Address
	Amount *big.Int
    
}

type LogCancelPaymentPlan struct {
    ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
    
}

type LogPaymentPlanIntervalEnded struct {
    ID         *big.Int
	IntervalNo *big.Int
    
}

type LogEndPaymentPlan struct {
    ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan

}

func PrintEvents()  {

	client, err := ethclient.Dial("wss://alfajores-forno.celo-testnet.org/ws")
	if err != nil {
	log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0xa883d9C6F7FC4baB52AcD2E42E51c4c528d7F7D3")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }


    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(TravelSaverABI)))
	if err != nil {
		log.Fatal(err)
	}

    logCreatedPaymentPlanSig := []byte("CreatedPaymentPlan(uint256,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))")
    logCreatedTravelPlanSig := []byte("CreatedTravelPlan(uint256,address,(address,uint256,uint256,uint256,uint256,uint256,uint256,bool))")
    logStartPaymentPlanIntervalSig := []byte("StartPaymentPlanInterval(uint256,uint256,uint256,uint256)")
    LogContributeToTravelPlanSig := []byte("ContributeToTravelPlan(uint256,address,uint256)")
    logClaimTravelPlanSig := []byte("ClaimTravelPlan(uint256)")
    logTransferSig := []byte("Transfer(address,address,uint256)")
    logCancelPaymentPlanSig := []byte("CancelPaymentPlan(uint256,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))")
    logPaymentPlanIntervalEndedSig := []byte("PaymentPlanIntervalEnded(uint256,uint256)")
    logEndPaymentPlanSig := []byte("EndPaymentPlan(uint256,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))")

    logCreatedPaymentPlanSigHash := crypto.Keccak256Hash(logCreatedPaymentPlanSig)
    logCreatedTravelPlanSigHash := crypto.Keccak256Hash(logCreatedTravelPlanSig)
    logStartPaymentPlanIntervalSigHash := crypto.Keccak256Hash(logStartPaymentPlanIntervalSig)
    LogContributeToTravelPlanSigHash := crypto.Keccak256Hash(LogContributeToTravelPlanSig)
    logClaimTravelPlanSigHash := crypto.Keccak256Hash(logClaimTravelPlanSig)
    logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
    logCancelPaymentPlanSigHash := crypto.Keccak256Hash(logCancelPaymentPlanSig)
    logPaymentPlanIntervalEndedSigHash := crypto.Keccak256Hash(logPaymentPlanIntervalEndedSig)
    LogEndPaymentPlanSigHash := crypto.Keccak256Hash(logEndPaymentPlanSig)



    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
            // fmt.Println(vLog.Data)
            switch vLog.Topics[0].Hex(){
            case logCreatedPaymentPlanSigHash.Hex():
                fmt.Printf("Log Name: CreatedPaymentPlan\n")
                var createdPaymentPlanEvent LogCreatedPaymentPlan
                 err := contractAbi.UnpackIntoInterface(&createdPaymentPlanEvent,"CreatedPaymentPlan", vLog.Data)
                if err != nil {
                    log.Fatal(err)
                }
                prettyPrint(createdPaymentPlanEvent)
 
            
            case logCreatedTravelPlanSigHash.Hex():
                fmt.Printf("Log Name: CreatedTravelPlan\n")
            
            case logStartPaymentPlanIntervalSigHash.Hex():
                fmt.Printf("Log Name: StartPaymentPlanInterval\n")
            
            case LogContributeToTravelPlanSigHash.Hex():
                fmt.Printf("Log Name: ContributeToTravelPlan\n")
            
            case logClaimTravelPlanSigHash.Hex():
                fmt.Printf("Log Name: ClaimTravelPlan\n")
            
            case logTransferSigHash.Hex():
                fmt.Printf("Log Name: Transfer\n")
            
            case logCancelPaymentPlanSigHash.Hex():
                fmt.Printf("Log Name: CancelPaymentPlan\n")

            case logPaymentPlanIntervalEndedSigHash.Hex():
                fmt.Printf("Log Name: PaymentPlanIntervalEnded\n")

            case LogEndPaymentPlanSigHash.Hex():
                fmt.Printf("Log Name: EndPaymentPlan\n")
            }
            
            

        }
    }

}


