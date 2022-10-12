package travelsaver

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Write()  {

	client, err := ethclient.Dial("https://alfajores-forno.celo-testnet.org")
	if err != nil {
	log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("42497423a6d4c1542322c024b9711a35cd9e29b11adab83bd5e5ff28f468194e")
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

	address := common.HexToAddress("0xa883d9C6F7FC4baB52AcD2E42E51c4c528d7F7D3")
	instance, err := NewTravelSaver(address, client)
	if err != nil {
		log.Fatal(err)
	}

	operatorPlanID_ := big.NewInt(2)
	operatorUserID_ := big.NewInt(2)
	amountPerInterval := big.NewInt(1)
	totalIntervals := big.NewInt(3)
	intervalLength := big.NewInt(10)
	// copy(operatorPlanID_[:], []byte(1))
	// copy(value[:], []byte("bar"))
	// copy(key[:], []byte("foo"))
	// copy(value[:], []byte("bar"))

	tx, err := instance.CreateTravelPaymentPlan(auth, operatorPlanID_, operatorUserID_, amountPerInterval, totalIntervals, intervalLength)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) 
	fmt.Println(binary.BigEndian.Uint64(tx.Data()))
	tx.Value()




}


