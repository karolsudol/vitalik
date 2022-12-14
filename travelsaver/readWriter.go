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

type ReadWriter struct {
	WSS             string
	HTTPS           string
	ContractAddress string
	BQ              BQ
	instanceAddress common.Address
	SubID           string
	auth            *bind.TransactOpts
}

func (r *ReadWriter) New() error {
	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("load env key err: %v", err)
	}

	key := os.Getenv("PRIVATE_KEY")

	client, err := ethclient.Dial(r.HTTPS)
	if err != nil {
		return fmt.Errorf("ethClient HTTPS dial err: %v", err)
	}
	// r.clinet = *client
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

	r.auth = auth

	r.instanceAddress = common.HexToAddress(r.ContractAddress)

	return nil
}
