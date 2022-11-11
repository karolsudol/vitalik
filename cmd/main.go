package main

import (
	"fmt"

	"github.com/flywallet-io/evm-events-consumer/travelsaver"
)

func main() {

	// travelsaver.PrintEvents()

	const projectID = "flywallet-web"
	const datasetID = "LogsEVM"
	const envirnoment = "TEST"
	const token = "CUSD"
	const network = "CELO"

	bq := travelsaver.BQ{
		ProjectID: projectID,
		DatasetID: datasetID,
		Tables: travelsaver.Tables{
			PaymentPlan: fmt.Sprintf("%s_%s_%s_%s", "PaymentPlan",
				network, token, envirnoment),
		},
	}

	rw := travelsaver.ReadWriter{
		WSS:             "wss://alfajores-forno.celo-testnet.org/ws",
		HTTPS:           "https://alfajores-forno.celo-testnet.org",
		ContractAddress: "0xa883d9C6F7FC4baB52AcD2E42E51c4c528d7F7D3",
		BQ:              bq,
	}

	rw.Subscribe()

}
