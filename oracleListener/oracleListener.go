package oracleListener

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"strings"

	"log"

	// logger "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/pinkyTseng/myGoOracle/contractMeta"
)

var ethereumUri string = "ws://localhost:8546"
var callbackAddrStr string = "0x20faBc5Ea9baFD923E8B9fB4abD09Eef54dC06cC"
var oracleAddrStr string = "0x702cEa0Ce3b7cEE1080be8bD89899C9a06B171D9"
var diceAddrStr string = "0xb37Bdc74Bc8A4D22d8BE376E610BcB1113996dc8"

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
	log.Println("get random value: %v", value)
	log.Println("get id: %v", id)

	// demoIns, err := NewCalldemo(common.HexToAddress("0x6D4e7f39E8cAA4aB8b4917B82f1E69a924712906"), conn)
	// if err != nil {
	// 	log.Fatalf("Fail NewCalldemo: %v", err)
	// }

	// val, err := demoIns.Message(nil)
	// if err != nil {
	// 	log.Fatalf("Fail get Message: %v", err)
	// }

	// // must to unlock the cbAddress before firing
	// web3.personal.unlockAccount(web3.eth.defaultAccount, password);
	// console.log(`fire _callback(${id[0]}, ${value}) back to Dice`);
	// let txid = dice.instance._callback(id[0], '' + value, {gas: 200000});
	// console.log(`txid = ${txid}`)

	// eventObj.

	// // get query id
	// var id = _.at(result, 'args.id');
	// // get query string
	// var query = _.split(_.at(result, 'args.query'), '-', 2);
	// // generate a random number according to query string
	// var value = _randomIntInc(query[0], query[1]);

	// // must to unlock the cbAddress before firing
	// web3.personal.unlockAccount(web3.eth.defaultAccount, password);
	// console.log(`fire _callback(${id[0]}, ${value}) back to Dice`);
	// let txid = dice.instance._callback(id[0], '' + value, {gas: 200000});
	// console.log(`txid = ${txid}`)

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
