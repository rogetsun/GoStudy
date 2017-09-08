package serverConf

const (
	SERVER_IP   = "0.0.0.0"
	SERVER_PORT = "2222"

	FRAME_FIRST_BYTE       = 0x88
	FRAME_LEN              = 13
	SOCKET_BYTE_BUFFER_LEN = 13 * 1000

	//
	CMD_ANALYSIS_CACHE_SIZE = 1000
)

/*
channel config
*/
const (
	FRAME_BYTE_CHAN_SIZE = 10
)

/**
goroutine 数配置
*/
const (
	GOROUTINE_CMD_ANALYSIS_SIZE = 5
)
