package main

import (
	"adumonitor/logic/logs"
	"fmt"
)

func main() {
	var buf int64 = 200
	readCh := make(chan []byte, buf)
	writeCh := make(chan *logs.Message, buf)
	go logs.Read(readCh)
	go logs.Analysis(readCh, writeCh)
	go logs.Write(writeCh)

	for {
		fmt.Println("writing...")
	}
}
