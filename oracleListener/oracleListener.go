package oracleListener

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/big"
	"math/rand"
	"net/http/httptrace"
	"os"
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

type privateKeySetting struct {
	CallbackPrivateKey string `json:"CallbackPrivateKey"`
}

type contractAddresses struct {
	OracleAddrStr string `json:"oracleAddr"`
	DiceAddrStr   string `json:"diceAddr"`
}

type network struct {
	EthereumUri     string `json:"ethereumUri"`
	EthereumHttpUri string `json:"ethereumHttpUri"`
	ChainID         string `json:"chainID"`
}

//setting vars
var envSettingDirKey string = "MyGethOracleSetting"
var detailMode bool = false

var ethereumUri string = ""
var ethereumHttpUri string = ""
var oracleAddrStr string = ""
var diceAddrStr string = ""
var callbackPrivateKey string = ""
var ChainID int64

var configDir string

var oracleABI abi.ABI
var packageCommonErr error

var clientTrace *httptrace.ClientTrace
var traceCtx context.Context

var errorCount int

func init() {
	rand.Seed(time.Now().UnixNano())
	errorCount = 0
	initDynamicSettings()

	setUsedContext()
	initOracleCompoent()
	go listenToEvent()
}

func setContractAddresses() {
	f, err := os.Open(configDir + "ContractAddresses.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	v := contractAddresses{}
	json.Unmarshal(content, &v)

	oracleAddrStr = v.OracleAddrStr
	diceAddrStr = v.DiceAddrStr
}

func getNetworkSetting() string {
	f, err := os.Open(configDir + "networkSetting.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func setNetwork(networkSetting string) {
	f, err := os.Open(configDir + "network/" + networkSetting + ".json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	v := network{}
	json.Unmarshal(content, &v)

	ethereumHttpUri = v.EthereumHttpUri
	ethereumUri = v.EthereumUri
	ChainID, err = strconv.ParseInt(v.ChainID, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
}

func setCallbackPrivateKey() {
	f, err := os.Open(configDir + "privateKey.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	content, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	v := privateKeySetting{}
	json.Unmarshal(content, &v)

	callbackPrivateKey = v.CallbackPrivateKey
}

func initDynamicSettings() {
	configDir = os.Getenv(envSettingDirKey)
	if configDir == "" {
		log.Fatal("Need set the environment variable")
	}
	setContractAddresses()
	networkSetting := getNetworkSetting()
	setNetwork(networkSetting)
	setCallbackPrivateKey()
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

// func DialHttpWithCustomContext(rawurl string, customContext *context.Context) (*ethclient.Client, error) {
// 	return ethclient.DialContext(*customContext, rawurl)
// }

func genAuthOpt(privateKey *ecdsa.PrivateKey, nonce uint64, gasPrice *big.Int) *bind.TransactOpts {
	defer func() {
		if r := recover(); r != nil {
			// notify server owner credential has some error
			panic(r)
		}
	}()
	// auth := bind.NewKeyedTransactor(privateKey)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(ChainID))
	if err != nil {
		log.Panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Context = traceCtx
	return auth
}

func getAuthNoQuryData() (*ecdsa.PrivateKey, common.Address) {
	defer func() {
		if r := recover(); r != nil {
			// notify server owner credential has some error
			panic(r)
		}
	}()
	privateKey, err := crypto.HexToECDSA(callbackPrivateKey)
	if err != nil {
		log.Panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Panic("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return privateKey, fromAddress
}

func getAuthQuryData(client *ethclient.Client, fromAddress common.Address) (uint64, *big.Int) {
	nonce, err := client.PendingNonceAt(traceCtx, fromAddress)
	if err != nil {
		log.Panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(traceCtx)
	if err != nil {
		log.Panic(err)
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
		log.Panic(upPackErr)
	}

	id := event.Id
	query := event.Query

	value, err := randomIntInc(query)
	if err != nil {
		// May be notify user from server
		log.Panic(err)
	}
	log.Println("get random value: ", value)
	log.Println("get id: ", id)
	return id, value
}

func getRandom(eventObj types.Log) {
	defer func() {
		if r := recover(); r != nil {
			//if errorCount >= XX, notify server owner
			errorCount++
		}
	}()
	log.Println("getRandom trigger")
	id, value := handleEventData(eventObj)
	client, err := ethclient.Dial(ethereumHttpUri)
	if err != nil {
		log.Panic(err)
	}
	defer client.Close()

	privateKey, fromAddress := getAuthNoQuryData()
	nonce, gasPrice := getAuthQuryData(client, fromAddress)
	auth := genAuthOpt(privateKey, nonce, gasPrice)

	callbackIns, err := contractMeta.NewDice(common.HexToAddress(diceAddrStr), client)
	if err != nil {
		log.Panicf("Fail NewDice: %v", err)
	}

	intRandomStr := strconv.Itoa(value)
	tx, err := callbackIns.Callback(auth, id, intRandomStr)
	if err != nil {
		log.Panicf("Fail exec Callbacks: %v", err)
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
		return 0, err
	}

	max, err := strconv.Atoi(maxStr)
	if err != nil {
		return 0, err
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
