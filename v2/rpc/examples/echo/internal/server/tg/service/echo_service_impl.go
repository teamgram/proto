// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package service

import (
	"context"

	"github.com/teamgram/proto/v2/rpc/examples/echo/echo"
	"github.com/teamgram/proto/v2/rpc/examples/echo/internal/core"

	"github.com/cloudwego/kitex/pkg/klog"
)

// EchosEcho
// echos.echo message:string = Echo;
func (s *Service) EchosEcho(ctx context.Context, request *echo.TLEchosEcho) (*echo.Echo, error) {
	c := core.New(ctx, s.svcCtx)
	klog.Infof("echos.echo - metadata: {}, request: %v", request)

	r, err := c.EchosEcho(request)
	if err != nil {
		return nil, err
	}

	klog.Infof("echos.echo - reply: %v", r)
	return r, err
}
