package main

import (
	"GoStudy/uvsocket/server"
	"GoStudy/uvsocket/server/analysis/cmdAnalysis"
	"GoStudy/uvsocket/server/channel"
	"GoStudy/uvsocket/serverConf"
	//"GoStudy/uvsocket/server/db"
	//fmt "log"
	"fmt"
	"time"
)

func main() {
	//err := db.DBTest()
	//if err != nil {
	//	fmt.Fatal(err)
	//	return
	//}
	for i := 0; i < serverConf.GOROUTINE_CMD_ANALYSIS_SIZE; i++ {
		go cmdAnalysis.CMDAnalysis(i+1, channel.FrameByteChan)
	}
	go func() {
		for {
			fmt.Printf("ChannelWatcher:FrameByteChan[%+v]\n", len(channel.FrameByteChan))
			time.Sleep(time.Duration(1) * time.Second)
		}
	}()
	server.Listen(serverConf.SERVER_IP, serverConf.SERVER_PORT)
}
