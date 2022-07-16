package memoryAccessor

import (
	"encoding/binary"
	. "github.com/0xrawsec/golang-win32/win32"
)

func ReadDWORD(pid DWORD, pointer DWORD) DWORD {
	byteArray, _ := memoryRead(pid, pointer, 4)
	return DWORD(binary.LittleEndian.Uint32(byteArray))
}
