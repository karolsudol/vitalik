package travelsaver

import (
	"context"
	"fmt"
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

func (r *ReadWriter) Subscribe() error {
	ctx := context.Background()

	reader := reader{
		HTTPS:           r.HTTPS,
		contractAddress: r.ContractAddress,
	}
	err := reader.new()
	if err != nil {
		return fmt.Errorf("smart contract new reader err: %v", err)
	}

	clientBQ, err := bigquery.NewClient(ctx, r.BQ.ProjectID)
	if err != nil {
		return fmt.Errorf("bigquery new client err: %v", err)
	}

	contractAddress := common.HexToAddress(r.ContractAddress)

	clientWSS, err := ethclient.Dial(r.WSS)
	if err != nil {
		return fmt.Errorf("eth new client dial err: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := clientWSS.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("eth  wss client SubscribeFilterLogs err: %v", err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(TravelSaverABI)))
	if err != nil {
		return fmt.Errorf("contract's abi json err: %v", err)
	}

	l := newLogSigHash()

	for {
		select {
		case err := <-sub.Err():
			return fmt.Errorf("sub log err: %v", err)
		case vLog := <-logs:

			switch vLog.Topics[0].Hex() {
			case l.logCreatedPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: CreatedPaymentPlan\n")
				var createdPaymentPlanEvent LogCreatedPaymentPlan
				err := contractAbi.UnpackIntoInterface(&createdPaymentPlanEvent, "CreatedPaymentPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("CreatedPaymentPlan log abi unpack err: %v", err)
				}

				prettyPrint(createdPaymentPlanEvent)
				err = createdPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("CreatedPaymentPlan BQ instert err: %v", err)
				}

			case l.logCreatedTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: CreatedTravelPlan\n")
				var createdTravelPlanEvent LogCreatedTravelPlan
				err := contractAbi.UnpackIntoInterface(&createdTravelPlanEvent, "CreatedTravelPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("CreatedTravelPlan log abi unpack err: %v", err)
				}
				prettyPrint(createdTravelPlanEvent)
				err = createdTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("CreatedTravelPlan BQ instert err: %v", err)
				}

			case l.logStartPaymentPlanIntervalSigHash.Hex():
				fmt.Printf("Log Name: StartPaymentPlanInterval\n")
				var startPaymentPlanIntervalEvent LogStartPaymentPlanInterval
				err := contractAbi.UnpackIntoInterface(&startPaymentPlanIntervalEvent, "StartPaymentPlanInterval", vLog.Data)
				if err != nil {
					return fmt.Errorf("StartPaymentPlanInterval log abi unpack err: %v", err)
				}
				startPaymentPlanIntervalEvent.ID = vLog.Topics[1].Big()
				startPaymentPlanIntervalEvent.CallableOn = vLog.Topics[2].Big()
				startPaymentPlanIntervalEvent.Amount = vLog.Topics[3].Big()

				prettyPrint(startPaymentPlanIntervalEvent)
				err = startPaymentPlanIntervalEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("StartPaymentPlanInterval BQ instert err: %v", err)
				}

			case l.logContributeToTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: ContributeToTravelPlan\n")
				var contributeToTravelPlanEvent LogContributeToTravelPlan
				err := contractAbi.UnpackIntoInterface(&contributeToTravelPlanEvent, "ContributeToTravelPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("ContributeToTravelPlan log abi unpack err: %v", err)
				}

				contributeToTravelPlanEvent.ID = vLog.Topics[1].Big()
				contributeToTravelPlanEvent.Contributor = common.HexToAddress(vLog.Topics[2].Hex())

				prettyPrint(contributeToTravelPlanEvent)

				err = contributeToTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("ContributeToTravelPlan BQ instert err: %v", err)
				}

			case l.logClaimTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: ClaimTravelPlan\n")
				var claimTravelPlanEvent LogClaimTravelPlan
				err := contractAbi.UnpackIntoInterface(&claimTravelPlanEvent, "ClaimTravelPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("ClaimTravelPlan log abi unpack err: %v", err)
				}
				claimTravelPlanEvent.ID = vLog.Topics[1].Big()

				var createdTravelPlanEvent LogCreatedTravelPlan

				createdTravelPlanEvent.TravelPlan, err = reader.readTravelPlan(claimTravelPlanEvent.ID)
				if err != nil {
					return fmt.Errorf("reader readTravelPlan err: %v", err)
				}

				prettyPrint(createdTravelPlanEvent.TravelPlan)

				err = createdTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("ClaimTravelPlan BQ instert err: %v", err)
				}

			case l.logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")
				var transferEvent LogTransfer
				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					return fmt.Errorf("Transfer log abi unpack err: %v", err)
				}
				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
				tx := vLog.TxHash.Hex()

				prettyPrint(transferEvent)

				err = transferEvent.instert(&r.BQ, clientBQ, ctx, tx)
				if err != nil {
					return fmt.Errorf("Transfer BQ instert err: %v", err)
				}

			case l.logCancelPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: CancelPaymentPlan\n")
				var cancelPaymentPlanEvent LogCancelPaymentPlan
				err := contractAbi.UnpackIntoInterface(&cancelPaymentPlanEvent, "CancelPaymentPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("CancelPaymentPlan log abi unpack err: %v", err)
				}
				prettyPrint(cancelPaymentPlanEvent)

				err = cancelPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("CancelPaymentPlan BQ instert err: %v", err)
				}

			case l.logPaymentPlanIntervalEndedSigHash.Hex():
				fmt.Printf("Log Name: PaymentPlanIntervalEnded\n")
				var paymentPlanIntervalEndedEvent LogPaymentPlanIntervalEnded
				err := contractAbi.UnpackIntoInterface(&paymentPlanIntervalEndedEvent, "PaymentPlanIntervalEnded", vLog.Data)
				if err != nil {
					return fmt.Errorf("PaymentPlanIntervalEnded log abi unpack err: %v", err)
				}

				paymentPlanIntervalEndedEvent.ID = vLog.Topics[1].Big()
				paymentPlanIntervalEndedEvent.IntervalNo = vLog.Topics[2].Big()

				prettyPrint(paymentPlanIntervalEndedEvent)
				err = paymentPlanIntervalEndedEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("PaymentPlanIntervalEnded BQ instert err: %v", err)
				}

			case l.logEndPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: EndPaymentPlan\n")
				var endPaymentPlanEvent LogEndPaymentPlan
				err := contractAbi.UnpackIntoInterface(&endPaymentPlanEvent, "EndPaymentPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("EndPaymentPlan log abi unpack err: %v", err)
				}
				prettyPrint(endPaymentPlanEvent)
				err = endPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("EndPaymentPlan BQ instert err: %v", err)
				}
			}

		}
	}

}
