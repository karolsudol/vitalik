package travelsaver

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

type reader struct {
	address common.Address
	rw      ReadWriter
}

func (r reader) load() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("load env key err: %v", err)
	}

	key := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial(r.rw.HTTPS)
	if err != nil {
		return fmt.Errorf("ethClient HTTPS dial err: %v", err)

	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return fmt.Errorf("private key ECDSA err: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")

	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("PendingNonceAt err: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("SuggestGasPrice err: %v", err)

	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	r.address = common.HexToAddress(r.rw.ContractAddress)

	return nil

}

func (r ReadWriter) readTravelPlan(ID *big.Int) {

	client, err := ethclient.Dial(r.HTTPS)
	if err != nil {
		log.Fatalf("ethClient HTTPS dial err: %v", err)
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

	address := common.HexToAddress(r.ContractAddress)

	instance, err := NewTravelSaver(address, client)
	if err != nil {
		log.Fatalf("instance contract err: %v", err)
	}

	object, err := instance.GetTravelPlanDetails(&bind.CallOpts{}, ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("GetTravelPlanDetails:")

	prettyPrint(object)

}
