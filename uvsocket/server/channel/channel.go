package channel

import "GoStudy/uvsocket/serverConf"

var FrameByteChan = make(chan []byte, serverConf.FRAME_BYTE_CHAN_SIZE)
