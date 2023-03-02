package unsafe

import (
	"reflect"
	"unsafe"
)

const (
	defaultLen = int(int32(^uint32(0)>>1)) // int32 max
)

// string => reflect.StringHeader => StringHeader.Data(字符串的指向字节数组的指针) => [defaultLen]byte => []byte
func StringToByteUnsafe(src string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&src))
	return (*(*[defaultLen]byte)(unsafe.Pointer(strHeader.Data)))[:len(src):len(src)]
}
