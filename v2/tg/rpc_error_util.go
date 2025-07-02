// Copyright (c) 2024 The Teamgram Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)

package tg

import (
	"errors"

	"github.com/teamgram/proto/v2/iface/ecode"
	"github.com/teamgram/proto/v2/mt"
)

type (
	TLRpcError = mt.TLRpcError
	RpcError   = mt.RpcError
)

var (
	DecodeRpcErrorClazz = mt.DecodeRpcErrorClazz
)

func NewRpcError(e error) *TLRpcError {
	var (
		err ecode.CodeError
	)
	ok := errors.As(e, &err)
	if ok {
		return &TLRpcError{
			ErrorCode:    int32(err.Code()),
			ErrorMessage: err.Msg(),
		}
	} else {
		return &TLRpcError{
			ErrorCode:    ErrInternal,
			ErrorMessage: e.Error(),
		}
	}
}
