package learn

import (
	"time"
)

func Now() int64 {
	return time.Now().Unix()
}

func Sleep(milli int64) {
	time.Sleep(time.Duration(milli) * time.Millisecond)
}
