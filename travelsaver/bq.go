package travelsaver

import (
	"time"
)

type BQ struct {
	ProjectID string
	DatasetID string
	Tables    Tables
}

type Tables struct {
	PaymentPlan              string
	TravelPlan               string
	StartPaymentPlanInterval string
	PaymentPlanIntervalEnded string
	ContributeToTravelPlan   string
	Transfer                 string
}

type paymentPlanBQ struct {
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

type travelPlanBQ struct {
	ID                int       `bigquery:"ID"`
	Owner             string    `bigquery:"Owner"`
	OperatorPlanID    int       `bigquery:"OperatorPlanID"`
	OperatorUserID    int       `bigquery:"OperatorUserID"`
	ContributedAmount float64   `bigquery:"ContributedAmount"`
	CreatedAt         time.Time `bigquery:"CreatedAt"`
	ClaimedAt         time.Time `bigquery:"ClaimedAt"`
	Claimed           bool      `bigquery:"Claimed"`
}

type startPaymentPlanIntervalBQ struct {
	ID         int       `bigquery:"ID"`
	CallableOn time.Time `bigquery:"CallableOn"`
	Amount     int       `bigquery:"Amount"`
	IntervalNo int       `bigquery:"IntervalNo"`
}

type paymentPlanIntervalEndedBQ struct {
	ID         int       `bigquery:"ID"`
	IntervalNo int       `bigquery:"IntervalNo"`
	TS         time.Time `bigquery:"TS"`
}

type contributeToTravelPlanBQ struct {
	ID          int       `bigquery:"ID"`
	Contributor string    `bigquery:"Contributor"`
	Amount      float64   `bigquery:"Amount"`
	TS          time.Time `bigquery:"TS"`
}

type transferBQ struct {
	From   string    `bigquery:"From"`
	To     string    `bigquery:"To"`
	Amount float64   `bigquery:"Amount"`
	TS     time.Time `bigquery:"TS"`
	TX     string    `bigquery:"TX"`
}
