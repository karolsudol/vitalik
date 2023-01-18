package travelsaver

import (
	"context"
	"fmt"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

func (r *ReadWriter) Subscribe() error {
	ctx := context.Background()

	clientBQ, err := bigquery.NewClient(ctx, r.BQ.ProjectID)
	if err != nil {
		return fmt.Errorf("bigquery new client err: %v", err)
	}

	contractAddress := common.HexToAddress(r.ContractAddress)

	clientWSS, err := ethclient.Dial(r.WSS)
	if err != nil {
		return fmt.Errorf("eth new wss client dial err: %v", err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)
	// logs := make(chan types.Log, 1000)

	sub := event.Resubscribe(2*time.Second, func(ctx context.Context) (event.Subscription, error) {
		return clientWSS.SubscribeFilterLogs(ctx, query, logs)
	})

	// sub, err := clientWSS.SubscribeFilterLogs(ctx, query, logs)
	// if err != nil {
	// 	return fmt.Errorf("eth  wss client SubscribeFilterLogs err: %v", err)
	// }

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
				r.LogInfo.Println("CreatedPaymentPlan")
				var createdPaymentPlanEvent LogCreatedPaymentPlan
				err := contractAbi.UnpackIntoInterface(&createdPaymentPlanEvent, "CreatedPaymentPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("CreatedPaymentPlan log abi unpack err: %v", err)
				}

				r.LogInfo.Println(createdPaymentPlanEvent)
				err = createdPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("CreatedPaymentPlan BQ instert err: %v", err)
				}

			case l.logCreatedTravelPlanSigHash.Hex():
				r.LogInfo.Println("CreatedTravelPlan")
				var createdTravelPlanEvent LogCreatedTravelPlan
				err := contractAbi.UnpackIntoInterface(&createdTravelPlanEvent, "CreatedTravelPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("CreatedTravelPlan log abi unpack err: %v", err)
				}
				r.LogInfo.Println(createdTravelPlanEvent)
				err = createdTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("CreatedTravelPlan BQ instert err: %v", err)
				}

			case l.logStartPaymentPlanIntervalSigHash.Hex():
				r.LogInfo.Println("Log Name: StartPaymentPlanInterval")
				var startPaymentPlanIntervalEvent LogStartPaymentPlanInterval
				err := contractAbi.UnpackIntoInterface(&startPaymentPlanIntervalEvent, "StartPaymentPlanInterval", vLog.Data)
				if err != nil {
					return fmt.Errorf("StartPaymentPlanInterval log abi unpack err: %v", err)
				}
				startPaymentPlanIntervalEvent.ID = vLog.Topics[1].Big()
				startPaymentPlanIntervalEvent.CallableOn = vLog.Topics[2].Big()
				startPaymentPlanIntervalEvent.Amount = vLog.Topics[3].Big()

				r.LogInfo.Println(startPaymentPlanIntervalEvent)
				err = startPaymentPlanIntervalEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("StartPaymentPlanInterval BQ instert err: %v", err)
				}

				ID := int(startPaymentPlanIntervalEvent.ID.Int64())
				TS := int(startPaymentPlanIntervalEvent.CallableOn.Int64())
				err = r.publishIntervals(ID, TS)
				if err != nil {
					return fmt.Errorf("publishIntervals PUB/SUB instert err: %v", err)
				}

			case l.logContributeToTravelPlanSigHash.Hex():
				r.LogInfo.Println("ContributeToTravelPlan")
				var contributeToTravelPlanEvent LogContributeToTravelPlan
				err := contractAbi.UnpackIntoInterface(&contributeToTravelPlanEvent, "ContributeToTravelPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("ContributeToTravelPlan log abi unpack err: %v", err)
				}

				contributeToTravelPlanEvent.ID = vLog.Topics[1].Big()
				contributeToTravelPlanEvent.Contributor = common.HexToAddress(vLog.Topics[2].Hex())

				r.LogInfo.Println(contributeToTravelPlanEvent)

				err = contributeToTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("ContributeToTravelPlan BQ instert err: %v", err)
				}

			case l.logClaimTravelPlanSigHash.Hex():
				r.LogInfo.Println("ClaimTravelPlan")
				var claimTravelPlanEvent LogClaimTravelPlan
				err := contractAbi.UnpackIntoInterface(&claimTravelPlanEvent, "ClaimTravelPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("ClaimTravelPlan log abi unpack err: %v", err)
				}
				claimTravelPlanEvent.ID = vLog.Topics[1].Big()

				r.LogInfo.Println(claimTravelPlanEvent)

				err = claimTravelPlanEvent.instertClaim(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("ClaimedTravelPlan BQ instert err: %v", err)
				}

				// var createdTravelPlanEvent LogCreatedTravelPlan

				// createdTravelPlanEvent.TravelPlan, err = r.readTravelPlan(claimTravelPlanEvent.ID)
				// if err != nil {
				// 	return fmt.Errorf("reader readTravelPlan err: %v", err)
				// }
				// fmt.Println("GetTravelPlanDetails:")
				// prettyPrint(createdTravelPlanEvent.TravelPlan)

				// err = createdTravelPlanEvent.instert(&r.BQ, clientBQ, ctx)
				// if err != nil {
				// 	return fmt.Errorf("ClaimTravelPlan BQ instert err: %v", err)
				// }

			case l.logTransferSigHash.Hex():
				r.LogInfo.Println("Transfer")
				var transferEvent LogTransfer
				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					return fmt.Errorf("Transfer log abi unpack err: %v", err)
				}
				transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
				transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
				tx := vLog.TxHash.Hex()

				r.LogInfo.Println(transferEvent)

				err = transferEvent.instert(&r.BQ, clientBQ, ctx, tx)
				if err != nil {
					return fmt.Errorf("Transfer BQ instert err: %v", err)
				}

			case l.logCancelPaymentPlanSigHash.Hex():
				r.LogInfo.Println("CancelPaymentPlan")
				var cancelPaymentPlanEvent LogCancelPaymentPlan
				err := contractAbi.UnpackIntoInterface(&cancelPaymentPlanEvent, "CancelPaymentPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("CancelPaymentPlan log abi unpack err: %v", err)
				}
				r.LogInfo.Println(cancelPaymentPlanEvent)

				err = cancelPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("CancelPaymentPlan BQ instert err: %v", err)
				}

			case l.logPaymentPlanIntervalEndedSigHash.Hex():
				r.LogInfo.Println("PaymentPlanIntervalEnded")
				var paymentPlanIntervalEndedEvent LogPaymentPlanIntervalEnded
				err := contractAbi.UnpackIntoInterface(&paymentPlanIntervalEndedEvent, "PaymentPlanIntervalEnded", vLog.Data)
				if err != nil {
					return fmt.Errorf("PaymentPlanIntervalEnded log abi unpack err: %v", err)
				}

				paymentPlanIntervalEndedEvent.ID = vLog.Topics[1].Big()
				paymentPlanIntervalEndedEvent.IntervalNo = vLog.Topics[2].Big()

				r.LogInfo.Println(paymentPlanIntervalEndedEvent)
				err = paymentPlanIntervalEndedEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("PaymentPlanIntervalEnded BQ instert err: %v", err)
				}

			case l.logEndPaymentPlanSigHash.Hex():
				r.LogInfo.Println("EndPaymentPlan")
				var endPaymentPlanEvent LogEndPaymentPlan
				err := contractAbi.UnpackIntoInterface(&endPaymentPlanEvent, "EndPaymentPlan", vLog.Data)
				if err != nil {
					return fmt.Errorf("EndPaymentPlan log abi unpack err: %v", err)
				}
				r.LogInfo.Println(endPaymentPlanEvent)
				err = endPaymentPlanEvent.instert(&r.BQ, clientBQ, ctx)
				if err != nil {
					return fmt.Errorf("EndPaymentPlan BQ instert err: %v", err)
				}
			}

		}
	}

}
