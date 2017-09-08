package cmdAnalysis

import (
	"GoStudy/uvsocket/server/db/receiveDao"
	"GoStudy/uvsocket/serverConf"
	"fmt"
	"time"
)

func CMDAnalysis(id int, byte13Channel <-chan []byte) {
	defer func() {
		err := recover()
		fmt.Printf("goroutine-CMDAnalysis- %v: end err:%v", id, err)
	}()

	fmt.Printf("goroutine-CMDAnalysis- %v: get from byte13channel:%p,%p\n", id, &byte13Channel, byte13Channel)
	var (
		cacheArr *[serverConf.CMD_ANALYSIS_CACHE_SIZE][]byte
		cache    [][]byte
	)

SAVE:

	if cache != nil {
		fmt.Println("goroutine-CMDAnalysis-", id, ": will insert")
		receiveDao.Insert(cache)
	}
	//初始化缓存
	cacheArr = new([serverConf.CMD_ANALYSIS_CACHE_SIZE][]byte)
	cache = cacheArr[0:0]
	fmt.Println("goroutine-CMDAnalysis-", id, ":begin <-byte13Channel...")
	for {
		select {
		case byte13 := <-byte13Channel:
			fmt.Println("goroutine-CMDAnalysis-", id, ":<-byte13Channel=", byte13)
			cache = append(cache, byte13)
			if len(cache) == cap(cache) {
				goto SAVE
			}
		default:
			if len(cache) > 0 {
				goto SAVE
			} else {
				time.Sleep(time.Millisecond * time.Duration(1000))
				fmt.Println("goroutine-CMDAnalysis-", id, ": sleep 1s")
			}
		}
	}
}
