package tool

import (
	"unsafe"
)

// 判断系统是大端存储还是小端存储。大小端存储是由CPU和操作系统决定的。
// 在操作系统中，x86和一般的操作系统（比如Windows、FreeBSD、Linux）都是小端模式;
// Mac OS是大端模式。
func IsLittleEndian() bool {
	var i int32 = 0x01020304
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)
	b := *pb
	return (b == 0x04)
}
