// Copyright 2021 CloudWeGo Authors
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

package main

import (
	"context"
	"log"

	"github.com/teamgram/proto/v2/rpc/codec"
	api "github.com/teamgram/proto/v2/rpc/echo"
	"github.com/teamgram/proto/v2/rpc/echo/echo"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
)

var _ api.RPCEchos = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// EchosEcho implements the Echo interface.
func (s *EchoImpl) EchosEcho(ctx context.Context, req *api.TLEchosEcho) (resp *api.Echo, err error) {
	klog.Info("echo called")
	return api.MaleTLEcho(&api.Echo{Message: req.Message}).ToClazzEcho(), nil
}

func main() {
	zCodec := codec.NewZRpcCodec(true)
	svr := echo.NewServer(new(EchoImpl), server.WithCodec(zCodec))
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
