package logs

import (
	"os"
	"bufio"
	"io"
	"time"
	"adumonitor/config"
)

const (
	ZERO int64 = 0
	ONE  int   = 2
)

func Read(readCh chan []byte) {
	logPath := config.Config.Logs
	file, err := os.Open(logPath.Path)
	if err != nil {
		panic(err)
	}

	file.Seek(ZERO, ONE)
	reader := bufio.NewReader(file)

	for {
		lineString, err := reader.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(time.Second)
			continue
		}
		if err != nil {
			panic(err)
		}
		readCh <- lineString
	}

}
