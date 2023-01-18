package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/logging"
	"github.com/karolsudol/vitalik/travelsaver"
)

func main() {

	const projectID = "gcp-project-ID"
	const datasetID = "gcp-bq-ds-id"

	var envirnoment string
	var token string
	var network string
	var wss string
	var https string
	var contractAddress string

	configPtr := flag.String("config", "X-X-X", "network-token-envirnoment")

	flag.Parse()

	ctx := context.Background()

	// Creates a client.
	clientLogger, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer clientLogger.Close()

	// Sets the name of the log to write to.
	logName := *configPtr

	loggerInfo := clientLogger.Logger(logName).StandardLogger(logging.Info)
	loggerErr := clientLogger.Logger(logName).StandardLogger(logging.Error)

	switch *configPtr {

	// *** CELO *****

	case "CELO-CUSD-TEST":

		envirnoment = "TEST"
		token = "CUSD"
		network = "CELO"
		wss = "wss://alfajores-forno.celo-testnet.org/ws"
		https = "https://alfajores-forno.celo-testnet.org"
		contractAddress = "0x54713127daf2bFD5129C980Ea800E3fCD616B547"

	case "CELO-CUSD-PROD":

		envirnoment = "PROD"
		token = "CUSD"
		network = "CELO"
		wss = "wss://forno.celo.org/ws"
		https = "https://forno.celo.org"
		contractAddress = "0x207856B02b264b7C60fdE304658d683184254330"

		// *** POLYGON *****

	case "POLYGON-USDC-TEST":

		envirnoment = "TEST"
		token = "USDC"
		network = "POLYGON"
		wss = "wss://rpc-mumbai.matic.today"
		https = "https://rpc-mumbai.matic.today"
		contractAddress = "0x54713127daf2bFD5129C980Ea800E3fCD616B547"

	case "POLYGON-USDC-PROD":

		envirnoment = "PROD"
		token = "USDC"
		network = "POLYGON"
		wss = "wss://rpc-mainnet.matic.network"
		https = "https://polygon-rpc.com"
		contractAddress = "0x6Bd249181BAdf2a389296D68f80A8B1c74fDDAC1"

		// *** BSC *****

	case "BSC-USDT-TEST":

		envirnoment = "TEST"
		token = "USDT"
		network = "BSC"
		wss = "wss://testnet-dex.binance.org/api/ws"
		https = "https://data-seed-pre-0-s1.binance.org:443"
		contractAddress = "0x54713127daf2bFD5129C980Ea800E3fCD616B547"

	case "BSC-USDT-PROD":

		envirnoment = "PROD"
		token = "USDT"
		network = "BSC"
		wss = "wss://dex.binance.org/api/ws"
		https = "https://dataseed1.binance.org:443"
		contractAddress = "0x6Bd249181BAdf2a389296D68f80A8B1c74fDDAC1"

	default:
		log.Fatalf("wrong config - exiting")
	}

	loggerInfo.Printf("started wss conn: %s for contract: %s \n ", wss, contractAddress)

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
			ClaimedTravelPlan: fmt.Sprintf("%s_%s_%s_%s", "ClaimedTravelPlan",
				network, token, envirnoment),
		},
	}

	rw := travelsaver.ReadWriter{
		WSS:             wss,
		HTTPS:           https,
		ContractAddress: contractAddress,
		BQ:              bq,
		LogInfo:         loggerInfo,
	}

	err = rw.New()
	if err != nil {
		loggerErr.Println(err)
	}

	err = rw.Subscribe()
	if err != nil {
		loggerErr.Println(err)
	}
}
