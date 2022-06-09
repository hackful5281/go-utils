package utilx

import "time"

// Now 获取当前时间戳，单位: 毫秒
func Now() int64 {
	return time.Now().UnixNano() / 1e6
}
