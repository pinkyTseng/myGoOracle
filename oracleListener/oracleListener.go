package oracleListener

import (
	"context"
	"encoding/json"
	"errors"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"crypto/ecdsa"
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
)

var ethereumUri string = "ws://localhost:8545"
var ethereumHttpUri string = "http://localhost:8545"
var callbackAddrStr string = "0x9c66139681c918a41e8B05Caf5653aff0aA410b8"
var oracleAddrStr string = "0x8058Cd17a879a23dC515b3B152ABCa1A5eB6a0af"
var diceAddrStr string = "0x2E26C409DbbA7A335Fd207189b33Ea8bcd20A199"

var callbackPrivateKey string = "85dec634b5fd256afaab46c39f2be809a6bcda4dcbc1ae961584d24c87377a5a"

// var client *ethclient.Client

//var password string = ""; //替換成可以unlock callbackAddr的密碼

// type config struct {
// 	ethereumUri string
// 	callbackAddr string
// 	oracleAddr string
// 	diceAddr string
// 	// password string
// }

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	rand.Seed(time.Now().UnixNano())
	go listenToEvent()
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
			log.Println(vLog)
			//log.Info(vLog) // pointer to event log
			getRandom(vLog)
		}
	}
}

func getRandom(eventObj types.Log) {
	//log.Info("event trigger")
	log.Println("event trigger")
	oracleABI, err := abi.JSON(strings.NewReader(string(contractMeta.OracleABI)))
	if err != nil {
		log.Fatal(err)
	}

	//for _, vLog := range logs {
	// event := struct {
	//   Key   [32]byte
	//   Value [32]byte
	// }{}

	// unpackerr := oracleABI.UnpackIntoInterface(&event, "QueryEvent", eventObj.Data)
	// if unpackerr != nil {
	//   log.Fatal(unpackerr)
	// }

	// fmt.Println(string(event.Key[:]))   // foo
	// fmt.Println(string(event.Value[:])) // bar
	// }

	result, err := oracleABI.Unpack("QueryEvent", eventObj.Data)
	if err != nil {
		//log.Println(err)
		log.Fatal(err)
	}

	//event QueryEvent(bytes32 id, string query);
	id := result[0]
	query := result[1]
	// randomRange := strings.Split(query.(string), "-")
	// var value = _randomIntInc(randomRange[0], randomRange[1])
	value, err := randomIntInc(query.(string))
	if err != nil {
		//log.Error(err)
		log.Println(err)
	}
	log.Println("get random value: ", value)
	log.Println("get id: ", id)

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

	callbackIns, err := contractMeta.NewDice(common.HexToAddress(diceAddrStr), client)
	if err != nil {
		log.Fatalf("Fail NewDice: %v", err)
	}

	//~~~~~~~~~~~~~~~~~~~~

	// keyFile := "/Users/pinkytseng/Documents/pinkyGethData/keystore/UTC--2022-05-18T06-07-40.701761000Z--0405326679ad037df5c2e28868e98e6bea71e8ad"
	// reader, _ := os.Open(keyFile)
	// opts, err := bind.NewTransactorWithChainID(reader, "", big.NewInt(18))
	// if err != nil {
	// 	log.Fatalf("Fail NewTransactor: %v", err)
	// }

	// gasPrice, err := client.SuggestGasPrice(context.Background())
	// log.Println("gasPrice: ", gasPrice)
	// opts.GasPrice = big.NewInt(8000)

	clientChainid, err := client.NetworkID(context.Background())
	log.Println("clientChainid: ", clientChainid)

	// []byte([]uint8{uint8(2), uint8(3)})

	idByteSlice, err := json.Marshal(id)
	if err != nil {
		log.Fatalf("json.Marshal(id) fail: %v", err)
	}
	var idByteArr [32]byte
	copy(idByteArr[:], idByteSlice[:32])

	intRandomStr := strconv.Itoa(value)

	log.Println("idByteArr: ", idByteArr)
	log.Println("intRandomStr: ", intRandomStr)

	tx, err := callbackIns.Callback(auth, idByteArr, intRandomStr)
	if err != nil {
		log.Fatalf("Fail exec Callbacks: %v", err)
	}
	log.Println("tx is: ", tx)

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
