// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package iface

import (
	"fmt"

	"github.com/teamgram/proto/v2/bin"
)

//const (
//	MTPROTO_VERSION = 2
//)

// Basic TL types.
const (
	ClazzID_int32     = -1932527041 // 0x8ccffa3f
	ClazzID_long      = 1253220205  // 0x4ab29f6d
	ClazzID_int64     = -1568590240 // 0xa2813660
	ClazzID_double    = 1431132616  // 0x554d59c8
	ClazzID_string    = 194458693   // 0xb973445
	ClazzID_void      = 470303800   // 0x1c084438
	ClazzID_boolFalse = -1132882121 // bc799737
	ClazzID_boolTrue  = -1720552011 // 997275b5
	CClazzID_true     = 1072550713  // 3fedd339
	ClazzID_vector    = 481674261
)

// Size32 represents 4-byte sequence.
// Values in TL are generally aligned to Word.
const Size32 = 4

type TLObject interface {
	Encode(x *bin.Encoder, layer int32) error
	Decode(d *bin.Decoder) error
	// EncodeBare(x *bin.Encoder, layer int32) error
	// DecodeBare(d *bin.Decoder) error
	// ClazzID() uint32
	// ClazzName() string
	// String() string
}

func DecodeObject(d *bin.Decoder) (TLObject, error) {
	clazzID, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	r := NewTLObjectByClazzID(clazzID)
	if r == nil {
		err = fmt.Errorf("can't find registered classId: %x", clazzID)
		return nil, err
	}

	// log.Infof("newTLObjectByClassID, classID: %x", uint32(classID))
	err = r.Decode(d)
	if err != nil {
		err = fmt.Errorf("object(%x) decode error: %v", uint32(clazzID), err)
		return nil, err
	}

	return r, nil
}

func EncodeObject(obj TLObject, layer int32) ([]byte, error) {
	x := bin.NewEncoder()

	err := obj.Encode(x, layer)
	if err != nil {
		return nil, err
	}

	return x.Bytes(), nil
}

func EncodeObjectList[T TLObject](x *bin.Encoder, vList []T, layer int32) error {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, obj := range vList {
		_ = obj.Encode(x, layer)
	}

	return nil
}

// EncodeBool serializes bare boolean.
func EncodeBool(x *bin.Encoder, v bool) {
	var (
		c int32
	)

	switch v {
	case true:
		c = int32(ClazzID_boolTrue)
		x.PutClazzID(uint32(c))
	case false:
		c = int32(ClazzID_boolFalse)
		x.PutClazzID(uint32(c))
	}
}

// DecodeBool decodes bare boolean from Buffer.
func DecodeBool(d *bin.Decoder) (bool, error) {
	v, err := d.PeekClazzID()
	if err != nil {
		return false, err
	}
	switch v {
	case ClazzID_boolTrue:
		_ = d.ConsumeClazzID(ClazzID_boolTrue)
		return true, nil
	case ClazzID_boolFalse:
		_ = d.ConsumeClazzID(ClazzID_boolTrue)
		return false, nil
	default:
		return false, bin.NewUnexpectedClazzID(v)
	}
}

func EncodeInt32List(x *bin.Encoder, vList []int32) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutInt32(v)
	}
}

func DecodeInt32List(d *bin.Decoder) ([]int32, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Length: n,
			Where:  "vector",
		}
	}

	vList := make([]int32, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.Int32()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}

func EncodeInt64List(x *bin.Encoder, vList []int64) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutInt64(v)
	}
}

func DecodeInt64List(d *bin.Decoder) ([]int64, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Length: n,
			Where:  "vector",
		}
	}

	vList := make([]int64, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.Int64()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}

func EncodeStringList(x *bin.Encoder, vList []string) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutString(v)
	}
}

func DecodeStringList(d *bin.Decoder) ([]string, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Length: n,
			Where:  "vector",
		}
	}

	vList := make([]string, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.String()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}

func EncodeBytesList(x *bin.Encoder, vList [][]byte) {
	x.PutClazzID(ClazzID_vector)
	x.PutInt(len(vList))
	for _, v := range vList {
		x.PutBytes(v)
	}
}

func DecodeBytesList(d *bin.Decoder) ([][]byte, error) {
	if err := d.ConsumeClazzID(ClazzID_vector); err != nil {
		return nil, err
	}
	n, err := d.Int()
	if err != nil {
		return nil, err
	}
	if n < 0 {
		return nil, &bin.InvalidLengthError{
			Length: n,
			Where:  "vector",
		}
	}

	vList := make([][]byte, n)
	for i := 0; i < n; i++ {
		vList[i], err = d.Bytes()
		if err != nil {
			return nil, err
		}
	}

	return vList, nil
}
