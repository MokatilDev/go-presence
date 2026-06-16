package utils

import "encoding/binary"

func Encode(opcode, dataLength int32) []byte {
	buffer := make([]byte, 8)

	binary.LittleEndian.PutUint32(buffer[:4], uint32(opcode))
	binary.LittleEndian.PutUint32(buffer[4:8], uint32(dataLength))

	return buffer
}

func Decode(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data[4:8])
}
