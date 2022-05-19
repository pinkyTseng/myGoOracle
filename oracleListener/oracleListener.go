package oracleListener

import (
	"context"
	// "encoding/json"
	"errors"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"crypto/ecdsa"
	// "fmt"
	"strings"

	"log"
	// logger "github.com/sirupsen/logrus"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/pinkyTseng/myGoOracle/contractMeta"

	"net/http/httptrace"
)

//setting vars
var ethereumUri string = "ws://localhost:8545"
var ethereumHttpUri string = "http://localhost:8545"
var callbackAddrStr string = "0x7dEcf302aBd2E4B31322eC0F724b663889d705F0"
var oracleAddrStr string = "0xD02053727A6a5879cfB07EC2d1223E6D8635218d"
var diceAddrStr string = "0x35a59391ee5A30D70DFa04A88c0924793408ED0b"

var callbackPrivateKey string = "1d02daeddc5b76af157a41def1704b900b902359afc906ac973e299f419d70f3"

var detailMode bool = false

// package vars
var oracleABI abi.ABI
var packageCommonErr error
var callbackClient *ethclient.Client

var clientTrace *httptrace.ClientTrace
var traceCtx context.Context

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
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

// var url string

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
			//log.Println(err)
			log.Fatal(err)
		case vLog := <-logs:
			// log.Println(vLog)
			//log.Info(vLog) // pointer to event log
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

func initCallbackCompoent() {
	callbackClient, packageCommonErr = ethclient.Dial(ethereumHttpUri)
	if packageCommonErr != nil {
		log.Fatal(packageCommonErr)
	}
	//callbackClient.Timeout
}

func DialHttpWithCustomContext(rawurl string, customContext *context.Context) (*ethclient.Client, error) {
	return ethclient.DialContext(*customContext, rawurl)
}

func getRandom(eventObj types.Log) {
	log.Println("event trigger")

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

	// //!!
	// clientTrace = &httptrace.ClientTrace{
	// 	GotConn:      func(info httptrace.GotConnInfo) { log.Printf("!!!conn was reused: %t", info.Reused) },
	// 	ConnectStart: func(network, addr string) { log.Println("!!!ConnectStart") },
	// 	ConnectDone:  func(network, addr string, err error) { log.Println("!!!ConnectDone") },
	// }
	// traceCtx = httptrace.WithClientTrace(context.Background(), clientTrace)
	// // client, err := DialHttpWithCustomContext(ethereumHttpUri, &traceCtx)
	// //!!

	client, err := ethclient.Dial(ethereumHttpUri)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	//~~~~~~~~~~~~

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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Context = traceCtx

	callbackIns, err := contractMeta.NewDice(common.HexToAddress(diceAddrStr), client)
	if err != nil {
		log.Fatalf("Fail NewDice: %v", err)
	}

	//~~~~~~~~~~~~~~~~~~~~

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// log.Println("gasPrice: ", gasPrice)
	// opts.GasPrice = big.NewInt(8000)

	intRandomStr := strconv.Itoa(value)

	// log.Println("idByteArr: ", id)
	// log.Println("intRandomStr: ", intRandomStr)

	// tx, err := callbackIns.Callback(auth, idByteArr, intRandomStr)
	tx, err := callbackIns.Callback(auth, id, intRandomStr)
	if err != nil {
		log.Fatalf("Fail exec Callbacks: %v", err)
	}
	_ = tx
	//log.Println("tx is: ", tx)

}

func randomIntInc(randomRange string) (int, error) {

	randomRangeArr := strings.Split(randomRange, "-")
	//var value = _randomIntInc(randomRange[0], randomRange[1])

	if len(randomRangeArr) != 2 {
		return 0, errors.New("event query param formmat wrong")
	}

	minStr := randomRangeArr[0]
	maxStr := randomRangeArr[1]

	// var min int = 0
	// var max int = 0
	// var resultNum int
	// var err error

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
	// n := rand.Intn(10) + 10

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
