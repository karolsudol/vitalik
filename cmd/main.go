package main

import (
	"fmt"
	"log"

	"github.com/karolsudol/vitalik/travelsaver"
)

func main() {

	const projectID = "flywallet-web"
	const datasetID = "LogsEVM"
	const envirnoment = "TEST"
	const token = "CUSD"
	const network = "CELO"
	const wss = "wss://alfajores-forno.celo-testnet.org/ws"
	const https = "https://alfajores-forno.celo-testnet.org"
	const contractAddress = "0xa883d9C6F7FC4baB52AcD2E42E51c4c528d7F7D3"

	bq := travelsaver.BQ{
		ProjectID: projectID,
		DatasetID: datasetID,
		Tables: travelsaver.Tables{
			PaymentPlan: fmt.Sprintf("%s_%s_%s_%s", "PaymentPlan",
				network, token, envirnoment),
			TravelPlan: fmt.Sprintf("%s_%s_%s_%s", "TravelPlan",
				network, token, envirnoment),
			StartPaymentPlanInterval: fmt.Sprintf("%s_%s_%s_%s", "StartPaymentPlanInterval",
				network, token, envirnoment),
			PaymentPlanIntervalEnded: fmt.Sprintf("%s_%s_%s_%s", "PaymentPlanIntervalEnded",
				network, token, envirnoment),
			ContributeToTravelPlan: fmt.Sprintf("%s_%s_%s_%s", "ContributeToTravelPlan",
				network, token, envirnoment),
			Transfer: fmt.Sprintf("%s_%s_%s_%s", "Transfer",
				network, token, envirnoment),
		},
	}

	rw := travelsaver.ReadWriter{
		WSS:             wss,
		HTTPS:           https,
		ContractAddress: contractAddress,
		BQ:              bq,
	}

	err := rw.New()
	if err != nil {
		log.Println(err)
	}

	err = rw.Subscribe()
	if err != nil {
		log.Println(err)
	}
}
