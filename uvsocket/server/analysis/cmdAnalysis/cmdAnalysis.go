package cmdAnalysis

import (
	"fmt"
	"GoStudy/uvsocket/serverConf"
)

func CMDAnalysis(id byte, byte13Channel <-chan []byte) {
	var (
		cacheArr = new([serverConf.CMD_ANALYSIS_CACHE_SIZE][]byte)
		cache = cacheArr[0:0]
	)
	fmt.Printf("goroutine-%v:CMDAnalysis get from byte13channel:%p,%p\n", id, &byte13Channel, byte13Channel)
	for {
		select {
		case byte13 := <-byte13Channel:
			cache = append(cache, byte13)
			if len(cache) == cap(cache) {

			}
		default:

		//fmt.Println(id, ":no data sleep 1s...")
		//	time.Sleep(time.Millisecond * 1000)
		}
	}
}

