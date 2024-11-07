// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mtproto

import (
	"context"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/teamgram/proto/example/example"
)

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// ExampleEcho implements the Echo interface.
func (s *EchoImpl) ExampleEcho(ctx context.Context, req *example.TLExampleEcho) (resp *example.Echo, err error) {
	temp, ok1 := metainfo.GetValue(ctx, "temp")
	logid, ok2 := metainfo.GetPersistentValue(ctx, "logid")

	if !(ok1 && ok2) {
		// panic("It looks like the protocol does not support transmitting meta information")
	}
	klog.Debug(temp)  // "temp-value"
	klog.Debug(logid) // "12345"
	klog.Debug("echo called")
	return &example.Echo{Clazz: &example.TLEcho{Message: req.Message}}, nil
}
