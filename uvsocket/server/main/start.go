package main

import (
	"GoStudy/uvsocket/server/analysis/cmdAnalysis"
	"GoStudy/uvsocket/server/channel"
	"GoStudy/uvsocket/server"
	"GoStudy/uvsocket/serverConf"
)

func main() {
	for i := 0; i < serverConf.GOROUTINE_CMD_ANALYSIS_SIZE; i++ {
		go cmdAnalysis.CMDAnalysis(i + 1, channel.FrameByteChan)
	}
	server.Listen(serverConf.SERVER_IP, serverConf.SERVER_PORT)
}