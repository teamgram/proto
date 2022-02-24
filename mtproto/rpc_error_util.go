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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRpcError(e *status.Status) (err *TLRpcError) {
	return &TLRpcError{Data2: &RpcError{
		ErrorCode:    int32(e.Code()),
		ErrorMessage: e.Message(),
	}}
}

func (m *TLRpcError) IsOK() bool {
	if m == nil {
		return true
	}
	return m.GetErrorCode() == int32(codes.OK)
}

func (m *TLRpcError) Error() string {
	return m.DebugString()
}

func (m *TLRpcError) Code() int {
	return int(m.GetErrorCode())
}

func (m *TLRpcError) Message() string {
	return m.GetErrorMessage()
}

func (m *TLRpcError) Details() []interface{} {
	return nil
}
