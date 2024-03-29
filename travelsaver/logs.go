package travelsaver

import (
	"context"
	"math/big"
	"time"

	"cloud.google.com/go/bigquery"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type LogCreatedPaymentPlan struct {
	ID          big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
}

func (log LogCreatedPaymentPlan) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.PaymentPlan).Inserter()

	items := []*paymentPlanBQ{

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

type LogCreatedTravelPlan struct {
	ID         *big.Int
	Owner      common.Address
	TravelPlan TravelSaverTravelPlan
}

func (log LogCreatedTravelPlan) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.TravelPlan).Inserter()

	items := []*travelPlanBQ{

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

type LogStartPaymentPlanInterval struct {
	ID         *big.Int
	CallableOn *big.Int
	Amount     *big.Int
	IntervalNo *big.Int
}

func (log LogStartPaymentPlanInterval) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.StartPaymentPlanInterval).Inserter()

	items := []*startPaymentPlanIntervalBQ{

		{ID: int(log.ID.Int64()),
			CallableOn: time.Time(time.Unix(log.CallableOn.Int64(), 0)),
			Amount:     int(log.Amount.Int64()),
			IntervalNo: int(log.IntervalNo.Int64()),
		},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

type LogContributeToTravelPlan struct {
	ID          *big.Int
	Contributor common.Address
	Amount      *big.Int
}

func (log LogContributeToTravelPlan) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.ContributeToTravelPlan).Inserter()

	items := []*contributeToTravelPlanBQ{

		{ID: int(log.ID.Int64()),
			Contributor: string(log.Contributor.Hex()),
			Amount:      float64(log.Amount.Int64()),
			TS:          time.Now()},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

type LogClaimTravelPlan struct {
	ID     *big.Int
	Owner  common.Address
	Amount *big.Int
}

func (log LogClaimTravelPlan) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context, t TravelSaverTravelPlan) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.TravelPlan).Inserter()

	items := []*travelPlanBQ{

		{ID: int(t.ID.Int64()),
			Owner:             string(t.Owner.Hex()),
			OperatorPlanID:    int(t.OperatorPlanID.Int64()),
			OperatorUserID:    int(t.OperatorUserID.Int64()),
			ContributedAmount: float64(t.ContributedAmount.Int64()),
			CreatedAt:         time.Time(time.Unix(t.CreatedAt.Int64(), 0)),
			Claimed:           true,
			ClaimedAt:         time.Now(),
		},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

func (log LogClaimTravelPlan) instertClaim(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.ClaimedTravelPlan).Inserter()

	items := []*claimedTravelPlanBQ{

		{ID: int(log.ID.Int64()),
			Owner:  string(log.Owner.Hex()),
			Amount: float64(log.Amount.Int64()),
			TS:     time.Now(),
		},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

type LogTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
}

func (log LogTransfer) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context, tx string) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.Transfer).Inserter()

	items := []*transferBQ{

		{From: string(log.From.Hex()),
			To:     string(log.To.Hex()),
			Amount: float64(log.Amount.Int64()),
			TS:     time.Now(),
			TX:     tx},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

type LogCancelPaymentPlan struct {
	ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
}

func (log LogCancelPaymentPlan) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.PaymentPlan).Inserter()

	items := []*paymentPlanBQ{

		{ID: int(log.PaymentPlan.ID.Int64()),
			Owner:              string(log.PaymentPlan.Sender.Hex()),
			TravelPlanID:       int(log.PaymentPlan.TravelPlanID.Int64()),
			TotalAmount:        float64(log.PaymentPlan.TotalAmount.Int64()),
			AmountSent:         float64(log.PaymentPlan.AmountSent.Int64()),
			AmountPerInterval:  float64(log.PaymentPlan.AmountPerInterval.Int64()),
			TotalIntervals:     int(log.PaymentPlan.TotalIntervals.Int64()),
			IntervalsProcessed: int(log.PaymentPlan.IntervalsProcessed.Int64()),
			Alive:              false,
			CancelledAt:        time.Now()},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

type LogPaymentPlanIntervalEnded struct {
	ID         *big.Int
	IntervalNo *big.Int
}

func (log LogPaymentPlanIntervalEnded) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.PaymentPlanIntervalEnded).Inserter()

	items := []*paymentPlanIntervalEndedBQ{

		{ID: int(log.ID.Int64()),
			IntervalNo: int(log.IntervalNo.Int64()),
			TS:         time.Now(),
		},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

type LogEndPaymentPlan struct {
	ID          *big.Int
	Owner       common.Address
	PaymentPlan TravelSaverPaymentPlan
}

func (log LogEndPaymentPlan) instert(bq *BQ, clientBQ *bigquery.Client, ctx context.Context) error {

	inserter := clientBQ.Dataset(bq.DatasetID).Table(bq.Tables.PaymentPlan).Inserter()

	items := []*paymentPlanBQ{

		{ID: int(log.PaymentPlan.ID.Int64()),
			Owner:              string(log.PaymentPlan.Sender.Hex()),
			TravelPlanID:       int(log.PaymentPlan.TravelPlanID.Int64()),
			TotalAmount:        float64(log.PaymentPlan.TotalAmount.Int64()),
			AmountSent:         float64(log.PaymentPlan.AmountSent.Int64()),
			AmountPerInterval:  float64(log.PaymentPlan.AmountPerInterval.Int64()),
			TotalIntervals:     int(log.PaymentPlan.TotalIntervals.Int64()),
			IntervalsProcessed: int(log.PaymentPlan.IntervalsProcessed.Int64()),
			Alive:              false,
			CancelledAt:        time.Now()},
	}
	if err := inserter.Put(ctx, items); err != nil {
		return err
	}
	return nil

}

type logSigHash struct {
	logCreatedPaymentPlanSigHash       common.Hash
	logCreatedTravelPlanSigHash        common.Hash
	logStartPaymentPlanIntervalSigHash common.Hash
	logContributeToTravelPlanSigHash   common.Hash
	logClaimTravelPlanSigHash          common.Hash
	logTransferSigHash                 common.Hash
	logCancelPaymentPlanSigHash        common.Hash
	logPaymentPlanIntervalEndedSigHash common.Hash
	logEndPaymentPlanSigHash           common.Hash
}

func newLogSigHash() *logSigHash {

	logCreatedPaymentPlanSig := []byte("CreatedPaymentPlan(uint256,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))")
	logCreatedTravelPlanSig := []byte("CreatedTravelPlan(uint256,address,(address,uint256,uint256,uint256,uint256,uint256,uint256,bool))")
	logStartPaymentPlanIntervalSig := []byte("StartPaymentPlanInterval(uint256,uint256,uint256,uint256)")
	LogContributeToTravelPlanSig := []byte("ContributeToTravelPlan(uint256,address,uint256)")
	logClaimTravelPlanSig := []byte("ClaimTravelPlan(uint256,address,uint256)")
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logCancelPaymentPlanSig := []byte("CancelPaymentPlan(uint256,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))")
	logPaymentPlanIntervalEndedSig := []byte("PaymentPlanIntervalEnded(uint256,uint256)")
	logEndPaymentPlanSig := []byte("EndPaymentPlan(uint256,address,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,bool))")

	return &logSigHash{
		logCreatedPaymentPlanSigHash:       crypto.Keccak256Hash(logCreatedPaymentPlanSig),
		logCreatedTravelPlanSigHash:        crypto.Keccak256Hash(logCreatedTravelPlanSig),
		logStartPaymentPlanIntervalSigHash: crypto.Keccak256Hash(logStartPaymentPlanIntervalSig),
		logContributeToTravelPlanSigHash:   crypto.Keccak256Hash(LogContributeToTravelPlanSig),
		logClaimTravelPlanSigHash:          crypto.Keccak256Hash(logClaimTravelPlanSig),
		logTransferSigHash:                 crypto.Keccak256Hash(logTransferSig),
		logCancelPaymentPlanSigHash:        crypto.Keccak256Hash(logCancelPaymentPlanSig),
		logPaymentPlanIntervalEndedSigHash: crypto.Keccak256Hash(logPaymentPlanIntervalEndedSig),
		logEndPaymentPlanSigHash:           crypto.Keccak256Hash(logEndPaymentPlanSig),
	}
}
