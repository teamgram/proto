// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package example

import (
	"context"
	"github.com/teamgram/proto/iface"

	"github.com/teamgram/proto/bin"
)

const (
	ClazzID_echo         = 0x73f1f8dc
	CkazzID_example_echo = 0x3072cfa1
)

func init() {
	iface.RegisterClazzID(ClazzID_echo, func() iface.TLObject { return &TLEcho{} })
	iface.RegisterClazzID(CkazzID_example_echo, func() iface.TLObject { return &TLExampleEcho{} })
}

type Echo struct {
	ClazzName string         `json:"_name"`
	Data2     iface.TLObject `json:"_data2"`
}

// TLEcho
// echo message:string = Echo;
type TLEcho struct {
	Message string `json:"message"`
}

func (m *TLEcho) EchoName() string {
	return "Echo"
}

func (m *TLEcho) Encode(x *bin.Encoder, layer int32) {
	x.PutClazzID(ClazzID_echo)
	x.PutString(m.Message)
}

func (m *TLEcho) Decode(d *bin.Decoder) (err error) {
	m.Message, err = d.String()

	return
}

// TLExampleEcho
// example.echo message:string = Echo;
type TLExampleEcho struct {
	ClazzID uint32 `json:"_id"`
	Message string `json:"message"`
}

func (m *TLExampleEcho) Encode(x *bin.Encoder, layer int32) {
	x.PutClazzID(CkazzID_example_echo)
	x.PutString(m.Message)
}

func (m *TLExampleEcho) Decode(d *bin.Decoder) (err error) {
	m.Message, err = d.String()

	return
}

type RPCEcho interface {
	ExampleEcho(ctx context.Context, req *TLExampleEcho) (r Echo, err error)
}
