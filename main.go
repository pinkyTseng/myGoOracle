package main

import (
	"fmt"
	"log"

	"github.com/pinkyTseng/myGoOracle/oracleListener"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	var userComd string
	log.Println("main start run")
	oracleListener.TheTestFun()
	fmt.Scan(&userComd)
}
