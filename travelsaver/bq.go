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
	PaymentPlan string
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
