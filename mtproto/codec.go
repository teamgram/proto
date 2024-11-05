// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/remote"
)

var _ remote.Codec = (*MTProtoCodec)(nil)

// MTProtoCodec NewMTProtoCodec creates the mtproto codec.
type MTProtoCodec struct {
	opt *Options
}

func NewMTProtoCodec(opt *Options) *MTProtoCodec {
	return &MTProtoCodec{
		opt: opt,
	}
}

// Name codec name
func (c *MTProtoCodec) Name() string {
	return "mtproto"
}

func (c *MTProtoCodec) Encode(ctx context.Context, message remote.Message, out remote.ByteBuffer) error {
	var (
		err error
	)

	msgType := message.MessageType()
	switch msgType {
	case remote.Call, remote.Oneway:
		// payload, err = m.encodeRequestPayload(ctx, message)
	case remote.Exception:
		// payload, err = m.encodeExceptionPayload(ctx, message)
		// todo(DMwangnima): refer to exception processing logic of dubbo-java, use status to determine if this exception
		// is in outside layer.(eg. non-exist InterfaceName)
		// for now, use StatusOK by default, regardless of whether it is in outside layer.
		// status = dubbo_spec.StatusOK
	case remote.Reply:
		// payload, err = m.encodeResponsePayload(ctx, message)
		// status = dubbo_spec.StatusOK
	case remote.Heartbeat:
		// payload, err = m.encodeHeartbeatPayload(ctx, message)
		// status = dubbo_spec.StatusOK
		// eventFlag = true
	default:
		return fmt.Errorf("unsupported MessageType: %v", msgType)
	}

	err = fmt.Errorf("")
	if err != nil {
		return err
	}

	return nil
}

func (c *MTProtoCodec) Decode(ctx context.Context, message remote.Message, in remote.ByteBuffer) error {
	return nil
}
