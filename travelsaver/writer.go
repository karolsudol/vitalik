package travelsaver

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Write()  {

	client, err := ethclient.Dial("wss://alfajores-forno.celo-testnet.org/ws")
	if err != nil {
	log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0x147B8eb97fD247D06C4006D269c90C1908Fb5D54")
	instance, err := NewTravelSaver(address, client)
	if err != nil {
		log.Fatal(err)
	}

	operatorPlanID_ := big.NewInt(1)
	operatorUserID_ := big.NewInt(1)
	amountPerInterval := big.NewInt(1)
	totalIntervals := big.NewInt(1)
	intervalLength := big.NewInt(1)
	// copy(operatorPlanID_[:], []byte(1))
	// copy(value[:], []byte("bar"))
	// copy(key[:], []byte("foo"))
	// copy(value[:], []byte("bar"))

	tx, err := instance.CreateTravelPaymentPlan(auth, operatorPlanID_, operatorUserID_, amountPerInterval, totalIntervals, intervalLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) 
	fmt.Printf("tx data:




}


