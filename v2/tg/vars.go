// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package tg

import (
	"math"
)

const (
	// TODO:
	// MaxNebulaChatUserID      = (1 << 48) - 1 // 281474976710655
	// MinNebulaChatChatID      = 1 << 48       // 281474976710656
	// MaxNebulaChatChatID      = (2 << 48) - 1 // 562949953421311
	// MinNebulaChatMegagroupID = 2 << 48       // 562949953421312
	// MaxNebulaChatMegagroupID = (3 << 48) - 1 // 844424930131967
	// MinNebulaChatChannelID   = 3 << 48       // 844424930131968
	// MaxNebulaChatChannelID   = (4 << 48) - 1 // 1125899906842623

	MinNebulaChatChatID    = 1              // 1
	MaxNebulaChatChatID    = 1073741824 - 1 // 1
	MinNebulaChatChannelID = 1073741824     // 1073801854
	MaxNebulaChatChannelID = math.MaxInt32
)

//func PeerIdIsUser(id int64) bool {
//	return id>>48 == 0
//}
//
//func PeerIdIsChat(id int64) bool {
//	return id>>48 == 1
//}
//
//func PeerIdIsMegagroup(id int64) bool {
//	return id>>48 == 2
//}

func ChatIdIsChat(id int64) bool {
	return id < MinNebulaChatChannelID
}

func ChatIdIsChannel(id int64) bool {
	return id >= MinNebulaChatChannelID
}

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

func ToBool(b bool) *Bool {
	if b {
		return BoolTrue
	} else {
		return BoolFalse
	}
}

func FromBool(b *Bool) bool {
	return ClazzName_boolTrue == b.BoolClazzName()
}

func MakeFlagsString(i string) (v *string) {
	if i != "" {
		v = new(string)
		*v = i
	}

	return
}

func MakeFlagsZeroString() (v *string) {
	v = new(string)
	*v = ""

	return
}

func GetFlagsString(i *string) string {
	if i != nil {
		return *i
	}

	return ""
}

func MakeFlagsInt32(i int32) (v *int32) {
	if i != 0 {
		v = new(int32)
		*v = i
	}

	return
}

func MakeFlagsZeroInt32() (v *int32) {
	v = new(int32)
	*v = 0

	return
}

func GetFlagsInt32(i *int32) int32 {
	if i != nil {
		return *i
	}

	return 0
}

func MakeFlagsInt64(i int64) (v *int64) {
	if i != 0 {
		v = new(int64)
		*v = i
	}

	return
}

func MakeFlagsZeroInt64() (v *int64) {
	v = new(int64)
	*v = 0

	return
}

func GetFlagsInt64(i *int64) int64 {
	if i != nil {
		return *i
	}

	return 0
}
