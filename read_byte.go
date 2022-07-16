package memoryAccessor

import (
	. "github.com/0xrawsec/golang-win32/win32"
)

func ReadByte(pid DWORD, pointer DWORD) byte {
	read, _ := memoryRead(pid, pointer, 1)
	return read[0]
}
