// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package echo

import (
	"context"
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzID_echo       = 0x05162463
	ClazzID_echo2      = 0x05162464
	ClazzID_echos_echo = 0x05162465
)

const (
	ClazzName_echo  = "echo"
	ClazzName_echo2 = "echo2"
)

func init() {
	iface.RegisterClazzID(ClazzID_echo, func() iface.TLObject { return &TLEcho{} })
	iface.RegisterClazzID(ClazzID_echo2, func() iface.TLObject { return &TLEcho2{} })
	iface.RegisterClazzID(ClazzID_echos_echo, func() iface.TLObject { return &TLEchosEcho{ClazzID: ClazzID_echos_echo} })

	iface.RegisterClazzName(ClazzName_echo, 0, ClazzID_echo)   // 5162463
	iface.RegisterClazzName(ClazzName_echo2, 0, ClazzID_echo2) // 83c95aec
}

// Echo <--
//   - TLWorldEmpty
//   - TLWorld
type Echo struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	Message   string `json:"message"`
	Message2  string `json:"message2"`
}

func (m *Echo) Encode(x *bin.Encoder, layer int32) error {
	clazzName := m.ClazzName
	if clazzName == "" {
		clazzName = iface.GetClazzNameByID(m.ClazzID)
	}

	switch clazzName {
	case ClazzName_echo:
		t := m.ToClazzEcho()
		return t.Encode(x, layer)
	case ClazzName_echo2:
		t := m.ToClazzEcho2()
		return t.Encode(x, layer)
	default:
		// err := fmt.Errorf("invalid predicate error: %s", m.PredicateName)
		// log.Errorf(err.Error())
		return nil
	}
}

func (m *Echo) Decode(d *bin.Decoder) (err error) {
	m.ClazzID, err = d.ClazzID()
	if err != nil {
		return
	}

	switch uint32(m.ClazzID) {
	case ClazzID_echo:
		m2 := MakeTLEcho(m)
		err = m2.Decode(d)
	case ClazzID_echo2:
		m2 := MakeTLEcho(m)
		err = m2.Decode(d)
	default:
		err = fmt.Errorf("invalid constructorId: 0x%x", uint32(m.ClazzID))
	}

	return
}

func (m *Echo) ToClazzEcho() *TLEcho {
	m.ClazzName = ClazzName_echo
	return &TLEcho{
		Data2: m,
	}
}

func (m *Echo) ToClazzEcho2() *TLEcho2 {
	m.ClazzName = ClazzName_echo2
	return &TLEcho2{
		Data2: m,
	}
}

// TLEcho
// echo message:string = Echo;
type TLEcho struct {
	Data2 *Echo `json:"_data2"`
}

func MakeTLEcho(data2 *Echo) *TLEcho {
	if data2 == nil {
		return &TLEcho{Data2: &Echo{
			ClazzName: ClazzName_echo,
		}}
	} else {
		data2.ClazzName = ClazzName_echo
		return &TLEcho{Data2: data2}
	}
}

func (m *TLEcho) ToClazzEcho() *Echo {
	m.Data2.ClazzName = ClazzName_echo
	return m.Data2
}

func (m *TLEcho) SetMessage(v string) { m.Data2.Message = v }
func (m *TLEcho) GetMessage() string  { return m.Data2.Message }

func (m *TLEcho) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		ClazzID_echo: func() error {
			x.PutClazzID(ClazzID_echo)

			x.PutString(m.GetMessage())
			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_echo, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_bind_auth_key_inner, layer)
		return nil
	}
}

func (m *TLEcho) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		ClazzID_echo: func() error {
			v, err2 := d.String()
			m.SetMessage(v)
			return err2
		},
	}

	if f, ok := decodeF[uint32(m.Data2.ClazzID)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.ClazzID))
	}
}

// TLEcho2
// echo2 message:string = Echo;
type TLEcho2 struct {
	Data2 *Echo `json:"_data2"`
}

func MaleTLEcho2(data2 *Echo) *TLEcho2 {
	if data2 == nil {
		return &TLEcho2{Data2: &Echo{
			ClazzName: ClazzName_echo2,
		}}
	} else {
		data2.ClazzName = ClazzName_echo2
		return &TLEcho2{Data2: data2}
	}
}

func (m *TLEcho2) ClazzName() string {
	return "echo2"
}

func (m *TLEcho2) ToClazzEcho() *Echo {
	m.Data2.ClazzName = ClazzName_echo2
	return m.Data2
}

func (m *TLEcho2) SetMessage2(v string) { m.Data2.Message2 = v }
func (m *TLEcho2) GetMessage2() string  { return m.Data2.Message2 }

func (m *TLEcho2) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		ClazzID_echo2: func() error {
			x.PutClazzID(ClazzID_echo2)

			x.PutString(m.GetMessage2())
			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_echo2, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_bind_auth_key_inner, layer)
		return nil
	}
}

func (m *TLEcho2) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		ClazzID_echo2: func() error {
			v, err2 := d.String()
			m.SetMessage2(v)
			return err2
		},
	}

	if f, ok := decodeF[uint32(m.Data2.ClazzID)]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", uint32(m.Data2.ClazzID))
	}
}

// TLEchosEcho
// echos.echo message:string = Echo;
type TLEchosEcho struct {
	ClazzID uint32 `json:"_id"`
	Message string `json:"message"`
}

func (m *TLEchosEcho) Encode(x *bin.Encoder, layer int32) error {
	switch m.ClazzID {
	case ClazzID_echos_echo:
		x.PutClazzID(ClazzID_echos_echo)
		x.PutString(m.Message)

		return nil
	default:
		return fmt.Errorf("invalid clazz_id error: 0x%x", m.ClazzID)
	}
}

func (m *TLEchosEcho) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case ClazzID_echos_echo:
		m.Message, err = d.String()
	default:
		err = fmt.Errorf("invalid clazz_id error: 0x%x", m.ClazzID)
	}

	return
}

type RPCEchos interface {
	EchosEcho(ctx context.Context, in *TLEchosEcho) (*Echo, error)
}
