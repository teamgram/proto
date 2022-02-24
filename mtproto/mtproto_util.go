// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mtproto

import (
	"github.com/gogo/protobuf/types"
)

const (
	AuthKeyTypeUnknown   = -1
	AuthKeyTypePerm      = 0
	AuthKeyTypeTemp      = 1
	AuthKeyTypeMediaTemp = 2
)

var (
	BoolTrue       = MakeTLBoolTrue(nil).To_Bool()
	BoolFalse      = MakeTLBoolFalse(nil).To_Bool()
	UpdatesTooLong = MakeTLUpdatesTooLong(nil).To_Updates()
	EmptyVoid      = MakeTLVoid(nil).To_Void()
)

func NewAuthKeyInfo(keyId int64, key []byte, keyType int) *AuthKeyInfo {
	keyData := &AuthKeyInfo{
		AuthKeyId:          keyId,
		AuthKey:            key,
		AuthKeyType:        int32(keyType),
		PermAuthKeyId:      0,
		TempAuthKeyId:      0,
		MediaTempAuthKeyId: 0,
	}

	switch keyType {
	case AuthKeyTypePerm:
		keyData.PermAuthKeyId = keyId
	case AuthKeyTypeTemp:
		keyData.TempAuthKeyId = keyId
	case AuthKeyTypeMediaTemp:
		keyData.MediaTempAuthKeyId = keyId
	}

	return keyData
}

func ToBool(b bool) *Bool {
	if b {
		return BoolTrue
	} else {
		return BoolFalse
	}
}

func FromBool(b *Bool) bool {
	return Predicate_boolTrue == b.GetPredicateName()
}

func MakeFlagsInt32(v int32) *types.Int32Value {
	if v == 0 {
		return nil
	} else {
		return &types.Int32Value{Value: v}
	}
}

func MakeFlagsInt64(v int64) *types.Int64Value {
	if v == 0 {
		return nil
	} else {
		return &types.Int64Value{Value: v}
	}
}

func MakeFlagsString(v string) *types.StringValue {
	if v == "" {
		return nil
	} else {
		return &types.StringValue{Value: v}
	}
}

func BoolToInt8(b bool) int8 {
	if b {
		return 1
	} else {
		return 0
	}
}

func Int8ToBool(b int8) bool {
	if b == 1 {
		return true
	} else {
		return false
	}
}

//func (m *TLRpcError) ToError() *ecode.Status {
//	return ecode.Error(ecode.New(int(m.GetErrorCode())), m.GetErrorMessage())
//}

type MessageEntitySlice []*MessageEntity

func (m MessageEntitySlice) Len() int {
	return len(m)
}
func (m MessageEntitySlice) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m MessageEntitySlice) Less(i, j int) bool {
	// TODO(@benqi): if date[i] == date[j]
	return m[i].Offset < m[j].Offset
}
