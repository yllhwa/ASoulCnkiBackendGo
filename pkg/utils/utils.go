package utils

import (
	"encoding/binary"
)

func IntToBytes() {

}
func BytesToInt(data []byte) int64 {
	data = append([]byte{0, 0}, data...)
	return int64(binary.BigEndian.Uint64(data))
}
