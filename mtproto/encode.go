// Copyright (c) 2019-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

import (
	"encoding/binary"
	"math"
	"math/big"
	"sync"
)

type EncodeBuf struct {
	buf []byte
}

var encodeBufPool = sync.Pool{
	New: func() interface{} { return &EncodeBuf{buf: make([]byte, 0, 512)} },
}

// GetEncodeBuf returns a pooled EncodeBuf. Call PutEncodeBuf when done.
func GetEncodeBuf() *EncodeBuf {
	e := encodeBufPool.Get().(*EncodeBuf)
	e.buf = e.buf[:0]
	return e
}

// PutEncodeBuf returns an EncodeBuf to the pool.
func PutEncodeBuf(e *EncodeBuf) {
	if cap(e.buf) <= 65536 {
		encodeBufPool.Put(e)
	}
}

func NewEncodeBuf(cap int) *EncodeBuf {
	return &EncodeBuf{make([]byte, 0, cap)}
}

func (e *EncodeBuf) GetBuf() []byte {
	return e.buf
}

func (e *EncodeBuf) GetOffset() int {
	return len(e.buf)
}

func (e *EncodeBuf) Int16(s int16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], uint16(s))
}

func (e *EncodeBuf) UInt16(s uint16) {
	e.buf = append(e.buf, 0, 0)
	binary.LittleEndian.PutUint16(e.buf[len(e.buf)-2:], s)
}

func (e *EncodeBuf) Int(s int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(s))
}

func (e *EncodeBuf) IntOffset(offset int, s int32) {
	binary.LittleEndian.PutUint32(e.buf[offset:], uint32(s))
}

func (e *EncodeBuf) UInt(s uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], s)
}

func (e *EncodeBuf) Long(s int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(s))
}

func (e *EncodeBuf) Double(s float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(s))
}

func (e *EncodeBuf) String(s string) {
	e.StringBytes([]byte(s))
}

func (e *EncodeBuf) BigInt(s *big.Int) {
	e.StringBytes(s.Bytes())
}

func (e *EncodeBuf) StringBytes(s []byte) {
	size := len(s)
	if size < 254 {
		padding := (4 - (size + 1) % 4) & 3
		e.buf = append(e.buf, byte(size))
		e.buf = append(e.buf, s...)
		for i := 0; i < padding; i++ {
			e.buf = append(e.buf, 0)
		}
	} else {
		padding := (4 - size%4) & 3
		var hdr [4]byte
		binary.LittleEndian.PutUint32(hdr[:], uint32(size<<8|254))
		e.buf = append(e.buf, hdr[:]...)
		e.buf = append(e.buf, s...)
		for i := 0; i < padding; i++ {
			e.buf = append(e.buf, 0)
		}
	}
}

func (e *EncodeBuf) Bytes(s []byte) {
	e.buf = append(e.buf, s...)
}

func (e *EncodeBuf) VectorInt(vList []int32) {
	e.UInt(uint32(CRC32_vector))
	e.Int(int32(len(vList)))
	for _, v := range vList {
		e.Int(v)
	}
}

func (e *EncodeBuf) VectorLong(vList []int64) {
	e.UInt(uint32(CRC32_vector))
	e.Int(int32(len(vList)))
	for _, v := range vList {
		e.Long(v)
	}
}

func (e *EncodeBuf) VectorString(vList []string) {
	e.UInt(uint32(CRC32_vector))
	e.Int(int32(len(vList)))
	for _, v := range vList {
		e.String(v)
	}
}

func (e *EncodeBuf) VectorBytes(vList [][]byte) {
	e.UInt(uint32(CRC32_vector))
	e.Int(int32(len(vList)))
	for _, v := range vList {
		e.StringBytes(v)
	}
}
