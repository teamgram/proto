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

package metadata

import (
	"encoding/base64"

	"github.com/teamgram/proto/mtproto"

	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	headerRpcError = "rpc_error"
)

// RpcErrorFromMD Server To Client
func RpcErrorFromMD(md metadata.MD) (rpcErr *mtproto.TLRpcError) {
	val := NiceMD(md).Get(headerRpcError)
	if val == "" {
		// TODO(@benqi): 未设置rpc_error
		rpcErr = mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
		// log.Errorf("%v", rpcErr)
		return
	}

	// proto.Marshal()
	buf, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		rpcErr = mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
		// log.Errorf("%v", rpcErr)
		return
	}

	rpcErr = &mtproto.TLRpcError{}
	err = proto.Unmarshal(buf, rpcErr)
	if err != nil {
		rpcErr = mtproto.NewRpcError(status.Convert(mtproto.ErrInternelServerError))
		// log.Errorf("%v", rpcErr)
		return
	}

	// log.Errorf("%v", rpcErr)
	return rpcErr
}

func RpcErrorToMD(md *mtproto.TLRpcError) (metadata.MD, error) {
	buf, err := proto.Marshal(md)
	if err != nil {
		// log.Errorf("Marshal rpc_metadata error: %v", err)
		return nil, err
	}

	return metadata.Pairs(headerRpcError, base64.StdEncoding.EncodeToString(buf)), nil
}
