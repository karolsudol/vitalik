package travelsaver

import (
	"context"
	"crypto/ecdsa"
	"encoding/binary"
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

func (r ReadWriter) writeCreateTravelPaymentPlan(operatorPlanID, operatorUserID, amountPerInterval, totalIntervals, intervalLength *big.Int) error {
	client, err := ethclient.Dial(r.HTTPS)
	if err != nil {
		log.Fatalf("ethClient HTTPS dial err: %v", err)
	}
	instance, err := NewTravelSaver(r.instanceAddress, client)
	if err != nil {
		return fmt.Errorf("new instance of NewTravelSaver err: %v", err)
	}

	tx, err := instance.CreateTravelPaymentPlan(r.auth, operatorPlanID, operatorUserID, amountPerInterval, totalIntervals, intervalLength)
	if err != nil {
		return fmt.Errorf("instance CreateTravelPaymentPlan err: %v", err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Println(binary.BigEndian.Uint64(tx.Data()))
	tx.Value()

	return nil
}

func writeCreateTravelPaymentPlan(m vitaliksMsg) error {

	ctx := context.Background()

	err := godotenv.Load(".env")
	if err != nil {
		return fmt.Errorf("load env key err: %v", err)
	}

	key := os.Getenv("PRIVATE_KEY")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		return fmt.Errorf("crypto.HexToECDSA err: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("crypto.HexToECDSA err: %s", "cannot assert type: publicKey is not of type *ecdsa.PublicKey")

	}

	client, err := ethclient.Dial(m.HTTPS)
	if err != nil {
		log.Fatalf("ethClient HTTPS dial err: %v", err)
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return fmt.Errorf("nonce err: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("gasPrice err: %v", err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	instance, err := NewTravelSaver(common.HexToAddress(m.ADDR), client)
	if err != nil {
		return fmt.Errorf("new instance of NewTravelSaver err: %v", err)
	}

	paymentPlanID_ := big.NewInt(int64(m.ID))
	opts := &bind.TransactOpts{
		Nonce:    auth.Nonce,
		GasPrice: auth.GasPrice,
		From:     auth.From,
		Value:    auth.Value,
		Signer:   auth.Signer,
		GasLimit: auth.GasLimit,
	}
	tx, err := instance.RunInterval(opts, paymentPlanID_)
	if err != nil {
		return fmt.Errorf("instance RunInterval err: %v", err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Println(binary.BigEndian.Uint64(tx.Data()))
	tx.Value()

	return nil

}
