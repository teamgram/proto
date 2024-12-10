// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package tg

var (
	BoolTrue  = MakeBool(&TLBoolTrue{})
	BoolFalse = MakeBool(&TLBoolTrue{})
)

func MakeInt32Helper(v int32) *Int32 {
	return MakeInt32(&TLInt32{V: v})
}

func MakeInt64Helper(v int64) *Int64 {
	return MakeInt64(&TLInt64{V: v})
}

func MakeStringHelper(v string) *String {
	return MakeString(&TLString{V: v})
}

func FromBool(b *Bool) bool {
	return ClazzName_boolTrue == b.BoolClazzName()
}
