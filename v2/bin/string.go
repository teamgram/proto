// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package bin

import (
	"io"
)

const (
	// If L <= 253, the serialization contains one byte with the value of L,
	// then L bytes of the string followed by 0 to 3 characters containing 0,
	// such that the overall length of the value be divisible by 4,
	// whereupon all of this is interpreted as a sequence of int(L/4)+1 32-bit numbers.
	maxSmallStringLength = 253
	// If L >= 254, the serialization contains byte 254, followed by 3 bytes with
	// the string length L, followed by L bytes of the string, further followed
	// by 0 to 3 null padding bytes.
	firstLongStringByte = 254
)

func (e *Encoder) encodeString(v string) {
	l := len(v)
	if l <= maxSmallStringLength {
		_ = e.w.WriteByte(byte(l))
		_, _ = e.w.WriteString(v)
		currentLen := l + 1
		_, _ = e.w.Write(make([]byte, nearestPaddedValueLength(currentLen)-currentLen))
		return
	} else {
		_, _ = e.w.Write([]byte{firstLongStringByte,
			byte(l),
			byte(l >> 8),
			byte(l >> 16),
		})
		_, _ = e.w.WriteString(v)
		currentLen := l + 4
		_, _ = e.w.Write(make([]byte, nearestPaddedValueLength(currentLen)-currentLen))
	}
}

func decodeString(b []byte) (padding int, v string, err error) {
	if len(b) == 0 {
		return 0, "", io.ErrUnexpectedEOF
	}
	if b[0] == firstLongStringByte {
		if len(b) < 4 {
			return 0, "", io.ErrUnexpectedEOF
		}
		strLen := uint32(b[1]) | uint32(b[2])<<8 | uint32(b[3])<<16
		if len(b) < (int(strLen) + 4) {
			return 0, "", io.ErrUnexpectedEOF
		}
		return nearestPaddedValueLength(int(strLen) + 4), string(b[4 : strLen+4]), nil
	}
	strLen := int(b[0])
	if len(b) < (strLen + 1) {
		return 0, "", io.ErrUnexpectedEOF
	}
	if strLen > maxSmallStringLength {
		return 0, "", &InvalidLengthError{
			Length: strLen,
			Where:  "string",
		}
	}
	return nearestPaddedValueLength(strLen + 1), string(b[1 : strLen+1]), nil
}
