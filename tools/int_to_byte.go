package tool

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func IntTiByte() {
	k := uint32(2048)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, k)
	fmt.Println(bytesBuffer.Bytes())

	bytesBuffer.Reset()
	binary.Write(bytesBuffer, binary.LittleEndian, k)
	fmt.Println("littleEndian:", bytesBuffer.Bytes())
}
