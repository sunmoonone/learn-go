package learn

import (
	"fmt"
)

func StringToByteSlice(s string) ([]byte, error) {
	bs := make([]byte, len(s))
	copy(bs, s)
	return bs, nil
}

func StringToByteSlice2(s string) []byte {
	return []byte(s)
}

func ByteSlice2String(bs []byte) string {
	return string(bs)
}

func String(v interface{}) string {
	return fmt.Sprintf("%v", v)
}
