package travelsaver

import (
	"context"
	"fmt"
	"log"
	"strings"

	"cloud.google.com/go/bigquery"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ReadWriter struct {
	WSS             string
	HTTPS           string
	ContractAddress string
	BQ              BQ
}

func (r *ReadWriter) Subscribe() {
	ctx := context.Background()

	clientBQ, err := bigquery.NewClient(ctx, r.BQ.ProjectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}

	contractAddress := common.HexToAddress(r.ContractAddress)

	clientWSS, err := ethclient.Dial(r.WSS)
	if err != nil {
		log.Fatalf("ethClient WSS dial err: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := clientWSS.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalf("ethClient WSS sub err: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(TravelSaverABI)))
	if err != nil {
		log.Fatalf("abi json err: %v", err)
	}

	l := newLogSigHash()

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:

			switch vLog.Topics[0].Hex() {
			case l.logCreatedPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: CreatedPaymentPlan\n")
				var createdPaymentPlanEvent LogCreatedPaymentPlan
				err := contractAbi.UnpackIntoInterface(&createdPaymentPlanEvent, "CreatedPaymentPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				prettyPrint(createdPaymentPlanEvent)
				createdPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)

			case l.logCreatedTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: CreatedTravelPlan\n")
				var createdTravelPlanEvent LogCreatedTravelPlan
				err := contractAbi.UnpackIntoInterface(&createdTravelPlanEvent, "CreatedTravelPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(createdTravelPlanEvent)
				createdTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)

			case l.logStartPaymentPlanIntervalSigHash.Hex():
				fmt.Printf("Log Name: StartPaymentPlanInterval\n")
				var startPaymentPlanIntervalEvent LogStartPaymentPlanInterval
				err := contractAbi.UnpackIntoInterface(&startPaymentPlanIntervalEvent, "StartPaymentPlanInterval", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				startPaymentPlanIntervalEvent.ID = vLog.Topics[1].Big()
				startPaymentPlanIntervalEvent.CallableOn = vLog.Topics[2].Big()
				startPaymentPlanIntervalEvent.Amount = vLog.Topics[3].Big()

				prettyPrint(startPaymentPlanIntervalEvent)
				startPaymentPlanIntervalEvent.instert(&r.BQ, clientBQ, ctx)

			case l.logContributeToTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: ContributeToTravelPlan\n")
				var contributeToTravelPlanEvent LogContributeToTravelPlan
				err := contractAbi.UnpackIntoInterface(&contributeToTravelPlanEvent, "ContributeToTravelPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				contributeToTravelPlanEvent.ID = vLog.Topics[1].Big()
				contributeToTravelPlanEvent.Contributor = common.HexToAddress(vLog.Topics[2].Hex())
				prettyPrint(contributeToTravelPlanEvent)
				contributeToTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)

			case l.logClaimTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: ClaimTravelPlan\n")
				var claimTravelPlanEvent LogClaimTravelPlan
				err := contractAbi.UnpackIntoInterface(&claimTravelPlanEvent, "ClaimTravelPlan", vLog.Data)
				if err != nil {
					log.Printf("Failed to insert BQ tableTravelPlanContributionsCeloUSD: %v", err)
				}
				claimTravelPlanEvent.ID = vLog.Topics[1].Big()

				r.readTravelPlan(claimTravelPlanEvent.ID)

				prettyPrint(claimTravelPlanEvent)

			case l.logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")
				var transferEvent LogTransfer
				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
				tx := vLog.TxHash.Hex()

				prettyPrint(transferEvent)
				transferEvent.instert(&r.BQ, clientBQ, ctx, tx)

			case l.logCancelPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: CancelPaymentPlan\n")
				var cancelPaymentPlanEvent LogCancelPaymentPlan
				err := contractAbi.UnpackIntoInterface(&cancelPaymentPlanEvent, "CancelPaymentPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(cancelPaymentPlanEvent)
				cancelPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)

			case l.logPaymentPlanIntervalEndedSigHash.Hex():
				fmt.Printf("Log Name: PaymentPlanIntervalEnded\n")
				var paymentPlanIntervalEndedEvent LogPaymentPlanIntervalEnded
				err := contractAbi.UnpackIntoInterface(&paymentPlanIntervalEndedEvent, "PaymentPlanIntervalEnded", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				paymentPlanIntervalEndedEvent.ID = vLog.Topics[1].Big()
				paymentPlanIntervalEndedEvent.IntervalNo = vLog.Topics[2].Big()

				prettyPrint(paymentPlanIntervalEndedEvent)
				paymentPlanIntervalEndedEvent.instert(&r.BQ, clientBQ, ctx)

			case l.logEndPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: EndPaymentPlan\n")
				var endPaymentPlanEvent LogEndPaymentPlan
				err := contractAbi.UnpackIntoInterface(&endPaymentPlanEvent, "EndPaymentPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(endPaymentPlanEvent)
				endPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)

			}

		}
	}

}
