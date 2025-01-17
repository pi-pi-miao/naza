// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/naza
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

// Package bele 提供了大小端的转换操作
//
// be是big endian的缩写，即大端
// le是little endian的缩写，即小端
//
// assume local is `le`
//
package bele

import (
	"encoding/binary"
	"io"
	"math"
)

func BEUint16(p []byte) uint16 {
	return binary.BigEndian.Uint16(p)
}

func BEUint24(p []byte) uint32 {
	return uint32(p[2]) | uint32(p[1])<<8 | uint32(p[0])<<16
}

func BEUint32(p []byte) (ret uint32) {
	return binary.BigEndian.Uint32(p)
}

func BEFloat64(p []byte) (ret float64) {
	a := binary.BigEndian.Uint64(p)
	return math.Float64frombits(a)
}

func LEUint32(p []byte) (ret uint32) {
	return binary.LittleEndian.Uint32(p)
}

func BEPutUint24(out []byte, in uint32) {
	out[0] = byte(in >> 16)
	out[1] = byte(in >> 8)
	out[2] = byte(in)
}

func BEPutUint32(out []byte, in uint32) {
	binary.BigEndian.PutUint32(out, in)
}

func LEPutUint32(out []byte, in uint32) {
	binary.LittleEndian.PutUint32(out, in)
}

func WriteBEUint24(writer io.Writer, in uint32) error {
	_, err := writer.Write([]byte{uint8(in >> 16), uint8(in >> 8), uint8(in & 0xFF)})
	return err
}

func WriteBE(writer io.Writer, in interface{}) error {
	return binary.Write(writer, binary.BigEndian, in)
}

func WriteLE(writer io.Writer, in interface{}) error {
	return binary.Write(writer, binary.LittleEndian, in)
}
