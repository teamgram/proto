// Copyright 2024 Teamgram Authors
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
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func toMTProtoErrorCod(code codes.Code) codes.Code {
	switch code {
	case ErrSeeOther:
		return code
	case ErrBadRequest:
		return code
	case ErrUnauthorized:
		return code
	case ErrForbidden:
		return code
	case ErrNotFound:
		return code
	case ErrNotAcceptable:
		return code
	case ErrFlood:
		return code
	case ErrInternal:
		return code
	case ErrNotReturnClient:
		return code
	default:
		return ErrInternal
	}
}

func NewRpcError(e error) *TLRpcError {
	if e == nil {
		return nil
	}

	if rErr, ok := status.FromError(e); ok {
		return &TLRpcError{Data2: &RpcError{
			ErrorCode:    int32(toMTProtoErrorCod(rErr.Code())),
			ErrorMessage: rErr.Message(),
		}}
	} else {
		var err *TLRpcError
		switch {
		case errors.As(e, &err):
			return err
		default:
			return &TLRpcError{Data2: &RpcError{
				ErrorCode:    int32(toMTProtoErrorCod(codes.Internal)),
				ErrorMessage: "INTERNAL_SERVER_ERROR",
			}}
		}
	}
}

func (m *TLRpcError) IsOK() bool {
	if m == nil {
		return true
	}
	return m.GetErrorCode() == int32(codes.OK)
}

func (m *TLRpcError) Error() string {
	if m == nil {
		return ""
	}

	return fmt.Sprintf("rpc(TLRpcError) error: code = %d desc = %s", m.Code(), m.Message())
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
