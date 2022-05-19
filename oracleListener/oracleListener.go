package oracleListener

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"math/rand"
	"net/http/httptrace"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/pinkyTseng/myGoOracle/contractMeta"
)

//setting vars
var ethereumUri string = "ws://localhost:8545"
var ethereumHttpUri string = "http://localhost:8545"
var oracleAddrStr string = "0xD02053727A6a5879cfB07EC2d1223E6D8635218d"
var diceAddrStr string = "0x35a59391ee5A30D70DFa04A88c0924793408ED0b"
var callbackPrivateKey string = "1d02daeddc5b76af157a41def1704b900b902359afc906ac973e299f419d70f3"
var ChainID int64 = 1337

var detailMode bool = false

// package vars
var oracleABI abi.ABI
var packageCommonErr error

var clientTrace *httptrace.ClientTrace
var traceCtx context.Context

func init() {
	// log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	rand.Seed(time.Now().UnixNano())

	setUsedContext()
	initOracleCompoent()
	go listenToEvent()
}

func setUsedContext() {
	clientTrace = &httptrace.ClientTrace{
		GotConn:      func(info httptrace.GotConnInfo) { log.Printf("!!!conn was reused: %t", info.Reused) },
		ConnectStart: func(network, addr string) { log.Println("!!!ConnectStart") },
		ConnectDone:  func(network, addr string, err error) { log.Println("!!!ConnectDone") },
	}
	if detailMode {
		traceCtx = httptrace.WithClientTrace(context.Background(), clientTrace)
	} else {
		traceCtx = context.Background()
	}
}

func listenToEvent() {

	client, err := ethclient.Dial(ethereumUri)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	oracleAddress := common.HexToAddress(oracleAddrStr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{oracleAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			getRandom(vLog)
		}
	}
}

func initOracleCompoent() {
	oracleABI, packageCommonErr = abi.JSON(strings.NewReader(string(contractMeta.OracleABI)))
	if packageCommonErr != nil {
		log.Fatal(packageCommonErr)
	}
}

func DialHttpWithCustomContext(rawurl string, customContext *context.Context) (*ethclient.Client, error) {
	return ethclient.DialContext(*customContext, rawurl)
}

func genAuthOpt(privateKey *ecdsa.PrivateKey, nonce uint64, gasPrice *big.Int) *bind.TransactOpts {
	// auth := bind.NewKeyedTransactor(privateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(ChainID))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Context = traceCtx
	return auth
}

func getAuthNoQuryData() (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(callbackPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, fromAddress
}

func getAuthQuryData(client *ethclient.Client, fromAddress common.Address) (uint64, *big.Int) {
	// nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	nonce, err := client.PendingNonceAt(traceCtx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	gasPrice, err := client.SuggestGasPrice(traceCtx)
	if err != nil {
		log.Fatal(err)
	}
	return nonce, gasPrice
}

func handleEventData(eventObj types.Log) ([32]byte, int) {
	event := struct {
		Id    [32]byte
		Query string
	}{}
	upPackErr := oracleABI.UnpackIntoInterface(&event, "QueryEvent", eventObj.Data)
	if upPackErr != nil {
		log.Fatal(upPackErr)
	}

	id := event.Id
	query := event.Query

	value, err := randomIntInc(query)
	if err != nil {
		log.Println(err)
	}
	log.Println("get random value: ", value)
	log.Println("get id: ", id)
	return id, value
}

func getRandom(eventObj types.Log) {
	log.Println("getRandom trigger")
	id, value := handleEventData(eventObj)
	client, err := ethclient.Dial(ethereumHttpUri)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	privateKey, fromAddress := getAuthNoQuryData()
	nonce, gasPrice := getAuthQuryData(client, fromAddress)
	auth := genAuthOpt(privateKey, nonce, gasPrice)

	callbackIns, err := contractMeta.NewDice(common.HexToAddress(diceAddrStr), client)
	if err != nil {
		log.Fatalf("Fail NewDice: %v", err)
	}

	intRandomStr := strconv.Itoa(value)
	tx, err := callbackIns.Callback(auth, id, intRandomStr)
	if err != nil {
		log.Fatalf("Fail exec Callbacks: %v", err)
	}
	_ = tx
	//log.Println("tx is: ", tx)
}

func randomIntInc(randomRange string) (int, error) {
	randomRangeArr := strings.Split(randomRange, "-")
	if len(randomRangeArr) != 2 {
		return 0, errors.New("event query param formmat wrong")
	}

	minStr := randomRangeArr[0]
	maxStr := randomRangeArr[1]

	min, err := strconv.Atoi(minStr)
	if err != nil {
		log.Fatal(err)
	}

	max, err := strconv.Atoi(maxStr)
	if err != nil {
		log.Fatal(err)
	}

	if min < 0 {
		return 0, errors.New("min is negative")
	}

	if max < 0 {
		return 0, errors.New("max is negative")
	}

	diff := max - min
	if diff >= 0 {
		return rand.Intn(diff) + min, nil
	} else {
		return 0, errors.New("min > max")
	}
}

func TheTestFun() {
	log.Println("theTestFun info")
}
