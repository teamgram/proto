// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package main

import (
	"encoding/hex"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
	"github.com/teamgram/proto/v2/mt"
)

const (
	ClazzID_echo       = 0x05162463
	ClazzID_echo2      = 0x05162464
	ClazzID_echos_echo = 0x05162465
)

func init() {
	iface.RegisterClazzID(ClazzID_echo, func() iface.TLObject { return &TLEcho{} })
	iface.RegisterClazzID(ClazzID_echo2, func() iface.TLObject { return &TLEcho2{} })
	iface.RegisterClazzID(ClazzID_echos_echo, func() iface.TLObject { return &TLEchosEcho{} })
}

type TLEcho struct {
	Message string `json:"message"`
}

func (m *TLEcho) ClazzName() string {
	return "echo"
}

func (m *TLEcho) EchoName() string {
	return m.ClazzName()
}

func (m *TLEcho) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_echo)
	x.PutString(m.Message)
	return nil
}

func (m *TLEcho) Decode(d *bin.Decoder) (err error) {
	_, err = d.ClazzID()
	m.Message, err = d.String()
	return
}

type TLEcho2 struct {
	Message string `json:"message"`
}

func (m *TLEcho2) ClazzName() string {
	return "echo2"
}

func (m *TLEcho2) EchoName() string {
	return m.ClazzName()
}

func (m *TLEcho2) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_echo2)
	x.PutString(m.Message)
	return nil
}

func (m *TLEcho2) Decode(d *bin.Decoder) (err error) {
	_, err = d.ClazzID()
	m.Message, err = d.String()
	return
}

type Echo interface {
	EchoName() string
}

type TLEchosEcho struct {
	Message string `json:"message"`
}

func (m *TLEchosEcho) Encode(x *bin.Encoder, layer int32) error {
	x.PutClazzID(ClazzID_echos_echo)
	x.PutString(m.Message)
	return nil
}

func (m *TLEchosEcho) Decode(d *bin.Decoder) (err error) {
	_, err = d.ClazzID()
	m.Message, err = d.String()
	return
}

type RPCEchos interface {
	EchosEcho(in *TLEchosEcho) (Echo, error)
}

func main() {
	var (
		echo Echo
	)

	echo = &TLEcho{Message: "hello"}

	switch t := echo.(type) {
	case nil:
		fmt.Println("nil")
	case *TLEcho:
		fmt.Println(t.EchoName())
	case *TLEcho2:
		fmt.Println(t.EchoName())
	default:
		panic("unknown type")
	}

	reqPQ := &mt.TLReqPq{
		ClazzID: mt.ClazzID_resPQ,
		Nonce:   bin.Int128{},
	}

	x := bin.NewEncoder()
	_ = reqPQ.Encode(x, 0)
	fmt.Println(hex.EncodeToString(x.Bytes()))
}
