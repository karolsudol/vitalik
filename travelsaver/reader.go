package travelsaver

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func (r ReadWriter) readTravelPlan(ID *big.Int) (TravelSaverTravelPlan, error) {

	var o TravelSaverTravelPlan

	err := godotenv.Load(".env")
	if err != nil {
		return o, fmt.Errorf("load env key err: %v", err)
	}

	key := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial(r.HTTPS)
	if err != nil {
		return o, fmt.Errorf("ethClient HTTPS dial err: %v", err)

	}
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return o, fmt.Errorf("crypto.HexToECDSA err: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return o, fmt.Errorf("crypto.HexToECDSA err: %s", "cannot assert type: publicKey is not of type *ecdsa.PublicKey")

	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return o, fmt.Errorf("nonce err: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return o, fmt.Errorf("gasPrice err: %v", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(r.ContractAddress)

	instance, err := NewTravelSaver(address, client)
	if err != nil {
		return o, fmt.Errorf("instance contract err: %v", err)
	}

	object, err := instance.GetTravelPlanDetails(&bind.CallOpts{}, ID)
	if err != nil {
		return o, fmt.Errorf("object GetTravelPlanDetails err: %v", err)
	}

	return object, nil
}
