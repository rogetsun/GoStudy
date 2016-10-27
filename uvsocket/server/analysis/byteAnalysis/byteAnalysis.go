package byteAnalysis

import (
	"net"
	"fmt"
	"GoStudy/uvsocket/server/channel"
	. "GoStudy/uvsocket/serverConf"
)

func ByteRead(con net.Conn) {
	var (
		lastDataArr = new([FRAME_LEN]byte)
		lastData []byte
		lastDataLen int
		//若上次读取半截，抛去本次的半截
		readIdx int
	)

	READ_LESS:for {

		//重建接收缓存
		data := make([]byte, SOCKET_BYTE_BUFFER_LEN)
		//读取数据
		length, err := con.Read(data)
		if err != nil {
			fmt.Printf("Client %v quit.\n%v\n", con.RemoteAddr(), err)
			con.Close()
			break READ_LESS
		}
		fmt.Printf("%v send %v bytes: %X\n", con.RemoteAddr(), length, data[0:length])

		readIdx = 0

		//如果上次存在只收了一半的情况
		if lastDataLen = len(lastData); lastDataLen > 0 {
			fmt.Printf("上次剩余字节：%v\n", lastData)
			//上次只收了一半，但这次可能没有上次剩余后半部分的情况处理
			if data[FRAME_LEN - lastDataLen] == FRAME_FIRST_BYTE || lastDataLen + length == FRAME_LEN {
				//接的发的情况
				lastData = lastDataArr[lastDataLen:]
				copy(lastData, data[0:])
				lastData = lastDataArr[0:]
				fmt.Printf("两次拼接:%X\n", lastData)
				channel.FrameByteChan <- lastData
				lastDataArr = new([FRAME_LEN]byte)
				lastData = lastDataArr[0:0]
				readIdx = FRAME_LEN - lastDataLen
			} else {
				//可能非接的发的情况
				lastDataArr = new([FRAME_LEN]byte)
				lastData = lastDataArr[0:0]
			}
		}

		ANALYSIS_FRAME:for i := readIdx; i < length; {

			if data[i] == FRAME_FIRST_BYTE {
				//找到头
				if i + FRAME_LEN <= length {
					//没有读到末尾不够一个frame长度的情况
					fmt.Println("byte_analysis get ", FRAME_LEN, "bytes:", data[i:i + FRAME_LEN])
					channel.FrameByteChan <- data[i:i + FRAME_LEN]
					i += FRAME_LEN
				} else {
					fmt.Println("deal half of ", FRAME_LEN, " bytes")
					//只剩 半个frame长度字节处理
					lastData = lastDataArr[:length - i]
					copy(lastData, data[i:])
					fmt.Println(lastData)
					break ANALYSIS_FRAME
				}
			}
		}
		//res = "You said:" + res
		//con.Write([]byte(res))
	}
}
