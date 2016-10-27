package main

import (
	log "github.com/skoo87/log4go"
	"io/ioutil"
	"fmt"
	"os"
)

func SetLog() {
	w := log.NewConsoleWriter()
	w.SetColor(true)

	log.Register(w)
	log.SetLevel(log.DEBUG)
	log.SetLayout("2006-01-02 15:04:05")
}
func main() {
	fmt.Println(os.Args[0])
	fmt.Println("*****")
	cnt, err := ioutil.ReadFile("log.json")
	fmt.Println(cnt, err)
	//SetLog()
	if err := log.SetupLogWithConf("./log.json"); err != nil {
		panic(err)
	}
	defer log.Close()

	log.Info("fdsfds")
}

