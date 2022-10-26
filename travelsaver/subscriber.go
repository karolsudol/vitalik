package travelsaver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func prettyPrint(d ...interface{}) {
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
	ID *big.Int
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

type TravelPlanBQ struct {
	ID                int       `bigquery:"ID"`
	Owner             string    `bigquery:"Owner"`
	OperatorPlanID    int       `bigquery:"OperatorPlanID"`
	OperatorUserID    int       `bigquery:"OperatorUserID"`
	ContributedAmount float64   `bigquery:"ContributedAmount"`
	CreatedAt         time.Time `bigquery:"CreatedAt"`
	ClaimedAt         time.Time `bigquery:"ClaimedAt"`
	Claimed           bool      `bigquery:"Claimed"`
}

func insertRowsBQtravelPlan(ctx context.Context, projectID, datasetID, tableID string, clientBQ *bigquery.Client, log LogCreatedTravelPlan) error {

	inserter := clientBQ.Dataset(datasetID).Table(tableID).Inserter()
	items := []*TravelPlanBQ{

		{ID: int(log.TravelPlan.ID.Int64()),
			Owner:             string(log.TravelPlan.Owner.Hex()),
			OperatorPlanID:    int(log.TravelPlan.OperatorPlanID.Int64()),
			OperatorUserID:    int(log.TravelPlan.OperatorUserID.Int64()),
			ContributedAmount: float64(log.TravelPlan.ContributedAmount.Int64()),
			CreatedAt:         time.Time(time.Unix(log.TravelPlan.CreatedAt.Int64(), 0)),
			Claimed:           log.TravelPlan.Claimed,
		},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil
}

type PaymentPlanBQ struct {
	ID                 int       `bigquery:"ID"`
	Owner              string    `bigquery:"Owner"`
	TravelPlanID       int       `bigquery:"TravelPlanID"`
	TotalAmount        float64   `bigquery:"TotalAmount"`
	AmountSent         float64   `bigquery:"AmountSent"`
	AmountPerInterval  float64   `bigquery:"AmountPerInterval"`
	TotalIntervals     int       `bigquery:"TotalIntervals"`
	IntervalsProcessed int       `bigquery:"IntervalsProcessed"`
	NextTransferOn     time.Time `bigquery:"NextTransferOn"`
	Alive              bool      `bigquery:"Alive"`
	CreatedAt          time.Time `bigquery:"CreatedAt"`
	CancelledAt        time.Time `bigquery:"CancelledAt"`
}

func insertRowsBQpaymentPlan(ctx context.Context, projectID, datasetID, tableID string, clientBQ *bigquery.Client, log LogCreatedPaymentPlan) error {

	inserter := clientBQ.Dataset(datasetID).Table(tableID).Inserter()
	items := []*PaymentPlanBQ{

		{ID: int(log.PaymentPlan.ID.Int64()),
			Owner:              string(log.PaymentPlan.Sender.Hex()),
			TravelPlanID:       int(log.PaymentPlan.TravelPlanID.Int64()),
			TotalAmount:        float64(log.PaymentPlan.TotalAmount.Int64()),
			AmountSent:         float64(log.PaymentPlan.AmountSent.Int64()),
			AmountPerInterval:  float64(log.PaymentPlan.AmountPerInterval.Int64()),
			TotalIntervals:     int(log.PaymentPlan.TotalIntervals.Int64()),
			IntervalsProcessed: int(log.PaymentPlan.IntervalsProcessed.Int64()),
			NextTransferOn:     time.Time(time.Unix(log.PaymentPlan.NextTransferOn.Int64(), 0)),
			Alive:              log.PaymentPlan.Alive,
			CreatedAt:          time.Now()},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil
}

func PrintEvents() {

	ctx := context.Background()
	projectID := "flywallet-web"
	datasetID := "LogsEVM"

	tablePaymentPlanCeloCUSD := "PaymentPlanCeloCUSD"
	tableTravelPlanCeloCUSD := "flywallet-web.LogsEVM.TravelPlanCeloCUSD"

	clientBQ, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("bigquery.NewClient: %v", err)
	}
	defer clientBQ.Close()

	// datasetBQ := clientBQ.Dataset(datasetID)
	// tableBQpaymentPlanCeloCUSD := datasetBQ.Table(tablePaymentPlanCeloCUSD)
	// tableBQtravelPlanCeloCUSD := datasetBQ.Table(tableTravelPlanCeloCUSD)

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
	logContributeToTravelPlanSigHash := crypto.Keccak256Hash(LogContributeToTravelPlanSig)
	logClaimTravelPlanSigHash := crypto.Keccak256Hash(logClaimTravelPlanSig)
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logCancelPaymentPlanSigHash := crypto.Keccak256Hash(logCancelPaymentPlanSig)
	logPaymentPlanIntervalEndedSigHash := crypto.Keccak256Hash(logPaymentPlanIntervalEndedSig)
	logEndPaymentPlanSigHash := crypto.Keccak256Hash(logEndPaymentPlanSig)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			// fmt.Println(vLog.Data)
			switch vLog.Topics[0].Hex() {
			case logCreatedPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: CreatedPaymentPlan\n")
				var createdPaymentPlanEvent LogCreatedPaymentPlan
				err := contractAbi.UnpackIntoInterface(&createdPaymentPlanEvent, "CreatedPaymentPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				// createdPaymentPlanEvent.ID = vLog.Topics[1].String()

				// createdPaymentPlanEvent.Owner = common.HexToAddress(vLog.Topics[2].Hex())

				prettyPrint(createdPaymentPlanEvent)

				err = insertRowsBQpaymentPlan(ctx, projectID, datasetID, tablePaymentPlanCeloCUSD, clientBQ, createdPaymentPlanEvent)
				if err != nil {
					log.Printf("Failed to insert BQ PaymentPlan: %v", err)
				}

			case logCreatedTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: CreatedTravelPlan\n")
				var createdTravelPlanEvent LogCreatedTravelPlan
				err := contractAbi.UnpackIntoInterface(&createdTravelPlanEvent, "CreatedTravelPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(createdTravelPlanEvent)

				err = insertRowsBQtravelPlan(ctx, projectID, datasetID, tableTravelPlanCeloCUSD, clientBQ, createdTravelPlanEvent)
				if err != nil {
					log.Printf("Failed to insert BQ TravelPlan: %v", err)
				}

			case logStartPaymentPlanIntervalSigHash.Hex():
				fmt.Printf("Log Name: StartPaymentPlanInterval\n")
				var startPaymentPlanIntervalEvent LogStartPaymentPlanInterval
				err := contractAbi.UnpackIntoInterface(&startPaymentPlanIntervalEvent, "StartPaymentPlanInterval", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(startPaymentPlanIntervalEvent)

			case logContributeToTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: ContributeToTravelPlan\n")
				var contributeToTravelPlanEvent LogContributeToTravelPlan
				err := contractAbi.UnpackIntoInterface(&contributeToTravelPlanEvent, "ContributeToTravelPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(contributeToTravelPlanEvent)

			case logClaimTravelPlanSigHash.Hex():
				fmt.Printf("Log Name: ClaimTravelPlan\n")
				var claimTravelPlanEvent LogClaimTravelPlan
				err := contractAbi.UnpackIntoInterface(&claimTravelPlanEvent, "ClaimTravelPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(claimTravelPlanEvent)

			case logTransferSigHash.Hex():
				fmt.Printf("Log Name: Transfer\n")
				var transferEvent LogTransfer
				err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(transferEvent)

			case logCancelPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: CancelPaymentPlan\n")
				var cancelPaymentPlanEvent LogCancelPaymentPlan
				err := contractAbi.UnpackIntoInterface(&cancelPaymentPlanEvent, "CancelPaymentPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(cancelPaymentPlanEvent)

			case logPaymentPlanIntervalEndedSigHash.Hex():
				fmt.Printf("Log Name: PaymentPlanIntervalEnded\n")
				var paymentPlanIntervalEndedEvent LogPaymentPlanIntervalEnded
				err := contractAbi.UnpackIntoInterface(&paymentPlanIntervalEndedEvent, "PaymentPlanIntervalEnded", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(paymentPlanIntervalEndedEvent)

			case logEndPaymentPlanSigHash.Hex():
				fmt.Printf("Log Name: EndPaymentPlan\n")
				var endPaymentPlanEvent LogEndPaymentPlan
				err := contractAbi.UnpackIntoInterface(&endPaymentPlanEvent, "EndPaymentPlan", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				prettyPrint(endPaymentPlanEvent)
			}

		}
	}

}

func insertBQpaymentPlanRow(ctx context.Context, projectID, datasetID, tableID string, clientBQ *bigquery.Client, log LogCreatedPaymentPlan) error {
	ins := clientBQ.Dataset(datasetID).Table(tableID).Inserter()

	schemaBQpaymentPlan := bigquery.Schema{
		{Name: "ID", Required: false, Type: bigquery.IntegerFieldType},
		{Name: "Owner", Required: false, Type: bigquery.StringFieldType},
		{Name: "TravelPlanID", Required: false, Type: bigquery.IntegerFieldType},
		{Name: "TotalAmount", Required: false, Type: bigquery.FloatFieldType},
		{Name: "AmountSent", Required: false, Type: bigquery.FloatFieldType},
		{Name: "AmountPerInterval", Required: false, Type: bigquery.FloatFieldType},
		{Name: "TotalIntervals", Required: false, Type: bigquery.IntegerFieldType},
		{Name: "IntervalsProcessed", Required: false, Type: bigquery.IntegerFieldType},
		{Name: "NextTransferOn", Required: false, Type: bigquery.TimestampFieldType},
		{Name: "Alive", Required: false, Type: bigquery.BooleanFieldType},
		{Name: "CreatedAt", Required: false, Type: bigquery.TimestampFieldType},
		{Name: "CancelledAt", Required: false, Type: bigquery.TimestampFieldType},
	}

	schema, err := bigquery.InferSchema(schemaBQpaymentPlan)
	if err != nil {
		return fmt.Errorf("Failed to infer schema: %v", err)
	}

	ss := bigquery.StructSaver{
		Schema: schema,
		Struct: &PaymentPlanBQ{ID: int(log.PaymentPlan.ID.Int64()),
			Owner:              string(log.PaymentPlan.Sender.Hex()),
			TravelPlanID:       int(log.PaymentPlan.TravelPlanID.Int64()),
			TotalAmount:        float64(log.PaymentPlan.TotalAmount.Int64()),
			AmountSent:         float64(log.PaymentPlan.AmountSent.Int64()),
			AmountPerInterval:  float64(log.PaymentPlan.AmountPerInterval.Int64()),
			TotalIntervals:     int(log.PaymentPlan.TotalIntervals.Int64()),
			IntervalsProcessed: int(log.PaymentPlan.IntervalsProcessed.Int64()),
			NextTransferOn:     time.Time(time.Unix(log.PaymentPlan.NextTransferOn.Int64(), 0)),
			Alive:              log.PaymentPlan.Alive,
			CreatedAt:          time.Now(),
			// CancelledAt: null,
		},
	}
	if err := ins.Put(ctx, ss); err != nil {
		return fmt.Errorf("Failed to insert row: %v", err)
	}
	return nil
}

// schemaBQtravelPlan := bigquery.Schema{
// 	{Name: "ID", Required: false, Type: bigquery.IntegerFieldType},
// 	{Name: "Owner", Required: false, Type: bigquery.StringFieldType},
// 	{Name: "OperatorPlanID", Required: false, Type: bigquery.IntegerFieldType},
// 	{Name: "OperatorUserID", Required: false, Type: bigquery.IntegerFieldType},
// 	{Name: "ContributedAmount", Required: false, Type: bigquery.FloatFieldType},
// 	{Name: "CreatedAt", Required: false, Type: bigquery.TimestampFieldType},
// 	{Name: "ClaimedAt", Required: false, Type: bigquery.TimestampFieldType},
// 	{Name: "Claimed", Required: false, Type: bigquery.BooleanFieldType},
// }
