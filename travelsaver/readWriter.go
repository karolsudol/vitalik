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

type ReadWriter struct {
	WSS             string
	HTTPS           string
	ContractAddress string
	BQ              BQ
	instanceAddress common.Address
	// clinet          ethclient.Client
	auth *bind.TransactOpts
}

func (r ReadWriter) New() error {
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

func (r ReadWriter) readTravelPlan1(ID *big.Int) (TravelSaverTravelPlan, error) {
	client, err := ethclient.Dial(r.HTTPS)
	if err != nil {
		log.Fatalf("ethClient HTTPS dial err: %v", err)
	}
	var object TravelSaverTravelPlan
	instance, err := NewTravelSaver(r.instanceAddress, client)
	if err != nil {
		return object, fmt.Errorf("instance of new TravelSaver contract err: %v", err)

	}

	object, err = instance.GetTravelPlanDetails(&bind.CallOpts{}, ID)
	if err != nil {
		return object, fmt.Errorf("instance of TravelSaver GetTravelPlanDetails err: %v", err)
	}
	return object, nil
}

func (r ReadWriter) writeCreateTravelPaymentPlan(operatorPlanID, operatorUserID, amountPerInterval, totalIntervals, intervalLength *big.Int) error {
	client, err := ethclient.Dial(r.HTTPS)
	if err != nil {
		log.Fatalf("ethClient HTTPS dial err: %v", err)
	}
	instance, err := NewTravelSaver(r.instanceAddress, client)
	if err != nil {
		return fmt.Errorf("new instance of NewTravelSaver err: %v", err)
	}

	// operatorPlanID_ := big.NewInt(operatorPlanID)
	// operatorUserID_ := big.NewInt(2)
	// amountPerInterval := big.NewInt(1)
	// totalIntervals := big.NewInt(3)
	// intervalLength := big.NewInt(10)

	tx, err := instance.CreateTravelPaymentPlan(r.auth, operatorPlanID, operatorUserID, amountPerInterval, totalIntervals, intervalLength)
	if err != nil {
		return fmt.Errorf("instance CreateTravelPaymentPlan err: %v", err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
	fmt.Println(binary.BigEndian.Uint64(tx.Data()))
	tx.Value()

	return nil

}

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

	// fmt.Println("GetTravelPlanDetails:")
	// prettyPrint(object)
	return object, nil
}

// func Write()  {

// 	client, err := ethclient.Dial("https://alfajores-forno.celo-testnet.org")
// 	if err != nil {
// 	log.Fatal(err)
// 	}

// 	privateKey, err := crypto.HexToECDSA("42497423a6d4c1542322c024b9711a35cd9e29b11adab83bd5e5ff28f468194e")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	publicKey := privateKey.Public()
// 	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
// 	if !ok {
// 		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
// 	}

// 	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
// 	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	gasPrice, err := client.SuggestGasPrice(context.Background())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	auth := bind.NewKeyedTransactor(privateKey)
// 	auth.Nonce = big.NewInt(int64(nonce))
// 	auth.Value = big.NewInt(0)     // in wei
// 	auth.GasLimit = uint64(300000) // in units
// 	auth.GasPrice = gasPrice

// 	address := common.HexToAddress("0xa883d9C6F7FC4baB52AcD2E42E51c4c528d7F7D3")
// 	instance, err := NewTravelSaver(address, client)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	operatorPlanID_ := big.NewInt(2)
// 	operatorUserID_ := big.NewInt(2)
// 	amountPerInterval := big.NewInt(1)
// 	totalIntervals := big.NewInt(3)
// 	intervalLength := big.NewInt(10)
// 	// copy(operatorPlanID_[:], []byte(1))
// 	// copy(value[:], []byte("bar"))
// 	// copy(key[:], []byte("foo"))
// 	// copy(value[:], []byte("bar"))

// 	tx, err := instance.CreateTravelPaymentPlan(auth, operatorPlanID_, operatorUserID_, amountPerInterval, totalIntervals, intervalLength)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())
// 	fmt.Println(binary.BigEndian.Uint64(tx.Data()))
// 	tx.Value()

// }
