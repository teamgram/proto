// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package main

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/teamgram/proto/example/example"
	"github.com/teamgram/proto/example/example/codec"
	"github.com/teamgram/proto/example/example/echo"

	// "github.com/cloudwego/kitex-examples/codec/json"
	"github.com/cloudwego/kitex/client"
	"log"
	"time"
)

func main() {
	cCodec := codec.NewJsonCodec(true)
	client, err := echo.NewClient("echo", client.WithHostPorts("0.0.0.0:8888"), client.WithCodec(cCodec))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &example.TLExampleEcho{Message: "my request"}
		ctx := context.Background()
		ctx = metainfo.WithValue(ctx, "temp", "temp-value")       // 附加元信息到 context 里
		ctx = metainfo.WithPersistentValue(ctx, "logid", "12345") // 附加能持续透传的元信息
		resp, err := client.Echo(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
