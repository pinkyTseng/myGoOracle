package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/pinkyTseng/myGoOracle/oracleListener"
)

func init() {

	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   false,
	}
	log.SetFormatter(formatter)

	//log輸出為json格式
	//log.SetFormatter(&log.JSONFormatter{})
	//輸出設定為標準輸出(預設為stderr)
	log.SetOutput(os.Stdout)
	//設定要輸出的log等級
	log.SetLevel(log.DebugLevel)
}

func main() {

	//first start by the command -> run truffle migration -> change config at listener -> run the go app
	// geth --datadir /Users/pinkytseng/Documents/pinkyGethData --networkid 18 --port 30303 --http --http.addr 0.0.0.0 --http.vhosts "*"  --http.port 8545 --http.api 'db,net,eth,web3,personal' --http.corsdomain "*"  --ws --dev --dev.period 1 console 2> 1.log

	var userComd string

	// fmt.Println("aaa")
	log.Info("test info log")
	log.Info("test info log2")
	oracleListener.TheTestFun()
	fmt.Scan(&userComd)

}
