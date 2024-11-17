// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package echoclient

import (
	"context"

	api "github.com/teamgram/proto/v2/rpc/examples/echo/echo"
	"github.com/teamgram/proto/v2/rpc/examples/echo/echo/echo"
)

type EchoClient interface {
	EchosEcho(ctx context.Context, req *api.TLEchosEcho) (r *api.Echo, err error)
}

type defaultEchoClient struct {
	cli echo.Client
}

func NewEchoClient(cli echo.Client) EchoClient {
	return &defaultEchoClient{
		cli: cli,
	}
}

// EchosEcho
// echos.echo message:string = Echo;
func (m *defaultEchoClient) EchosEcho(ctx context.Context, in *api.TLEchosEcho) (*api.Echo, error) {
	return m.cli.EchosEcho(ctx, in)
}
