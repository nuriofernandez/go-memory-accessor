package memoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
)

func ReadInt(pid DWORD, pointer DWORD) int {
	ba, _ := memoryRead(pid, pointer, 4)
	var value int32
	value |= int32(ba[0])
	value |= int32(ba[1]) << 8
	value |= int32(ba[2]) << 16
	value |= int32(ba[3]) << 24
	return int(value)
}
