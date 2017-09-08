package main

import "fmt"

/**
连接器接口
*/
type Connector interface {
	Connect()
}

/**
usb接口
*/
type USB interface {
	GetName() string
	Connector
}

/**
PhoneConnector
*/
type PhoneConnector struct {
	Name string
}

//注意实现GetName的是*PhoneConnector,而不是PhoneConnector
func (pc *PhoneConnector) GetName() string {
	return pc.Name
}

//注意实现Connect的是*PhoneConnector,而不是PhoneConnector
func (pc *PhoneConnector) Connect() {
	fmt.Printf("%v connected\n", pc.GetName())
}

func main() {
	var pc = new(PhoneConnector)
	pc.Name = "IPhone6 S"
	pc.Connect()
	fmt.Println(pc.GetName())
	disconnect(pc) //pc并没有实现接口方法，因此此处需要给入指针，用new构建对象或者加&地址符号
}

func disconnect(c Connector) {
	usb, ok := c.(USB) //断言类型转换
	if ok {
		fmt.Printf("%v disconnected \n", usb.GetName())
	} else {
		fmt.Println("unknown decive")
	}

	/**
	类型switch
	*/
	switch v := c.(type) {
	case USB:
		fmt.Printf("%v disconnected \n", v.GetName())
	case Connector:
		fmt.Println("%v disconnected \n", v)
	default:
		fmt.Println("unknown decive")
	}

}
