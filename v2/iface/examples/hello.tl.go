// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package main

import (
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzID_worldEmpty     = 0x05162466
	ClazzID_world          = 0x05162467
	ClazzID_hello_sayHello = 0x05162469
)

const (
	ClazzName_worldEmpty = "worldEmpty"
	ClazzName_world      = "world"
)

func init() {
	iface.RegisterClazzID(ClazzID_worldEmpty, func() iface.TLObject { return &TLWorldEmpty{} })
	iface.RegisterClazzID(ClazzID_world, func() iface.TLObject { return &TLWorld{} })
	iface.RegisterClazzID(ClazzID_hello_sayHello, func() iface.TLObject { return &TLHelloSayHello{ClazzID: ClazzID_hello_sayHello} })

	iface.RegisterClazzName(ClazzName_worldEmpty, 0, 0x05162466) // 5162463
	iface.RegisterClazzName(ClazzName_world, 0, 0x05162467)      // 83c95aec
}

// World <--
//   - TLWorldEmpty
//   - TLWorld
type World struct {
	ClazzID   int32  `json:"_id"`
	ClazzName string `json:"_name"`
	Message   string `json:"message"`
	Message2  string `json:"message2"`
}

func (m *World) Encode(x *bin.Encoder, layer int32) error {
	clazzName := m.ClazzName
	if clazzName == "" {
		clazzName = iface.GetClazzNameByID(m.ClazzID)
	}

	switch clazzName {
	case ClazzName_worldEmpty:
		t := m.To_WorldEmpty()
		return t.Encode(x, layer)
	case ClazzName_world:
		t := m.To_World()
		return t.Encode(x, layer)
	default:
		// err := fmt.Errorf("invalid predicate error: %s", m.PredicateName)
		// log.Errorf(err.Error())
		return nil
	}
}

func (m *World) Decode(d *bin.Decoder) (err error) {
	m.ClazzID, err = d.ClazzID()
	if err != nil {
		return
	}

	switch uint32(m.ClazzID) {
	case ClazzID_worldEmpty:
		m2 := MaleTLWorldEmpty(m)
		err = m2.Decode(d)
	case ClazzID_world:
		m2 := MaleTLWorld(m)
		err = m2.Decode(d)
	default:
		err = fmt.Errorf("invalid constructorId: 0x%x", uint32(m.ClazzID))
	}

	return
}

func (m *World) To_WorldEmpty() *TLWorldEmpty {
	m.ClazzName = ClazzName_worldEmpty
	return &TLWorldEmpty{
		Data2: m,
	}
}

func (m *World) To_World() *TLWorld {
	m.ClazzName = ClazzName_world
	return &TLWorld{
		Data2: m,
	}
}

// TLWorldEmpty
// worldEmpty message:string = World;
type TLWorldEmpty struct {
	Data2 *World `json:"_data2"`
}

func MaleTLWorldEmpty(data2 *World) *TLWorldEmpty {
	if data2 == nil {
		return &TLWorldEmpty{Data2: &World{
			ClazzName: ClazzName_worldEmpty,
		}}
	} else {
		data2.ClazzName = ClazzName_worldEmpty
		return &TLWorldEmpty{Data2: data2}
	}
}

func (m *TLWorldEmpty) ToClazzWorld() *World {
	m.Data2.ClazzName = ClazzName_worldEmpty
	return m.Data2
}

func (m *TLWorldEmpty) SetMessage(v string) { m.Data2.Message = v }
func (m *TLWorldEmpty) GetMessage() string  { return m.Data2.Message }

func (m *TLWorldEmpty) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		ClazzID_worldEmpty: func() error {
			x.PutClazzID(ClazzID_worldEmpty)

			x.PutString(m.GetMessage())
			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_worldEmpty, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_bind_auth_key_inner, layer)
		return nil
	}
}

func (m *TLWorldEmpty) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		ClazzID_worldEmpty: func() error {
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

// TLWorld
// world message2:string = World;
type TLWorld struct {
	Data2 *World `json:"_data2"`
}

func MaleTLWorld(data2 *World) *TLWorld {
	if data2 == nil {
		return &TLWorld{Data2: &World{
			ClazzName: ClazzName_world,
		}}
	} else {
		data2.ClazzName = ClazzName_world
		return &TLWorld{Data2: data2}
	}
}

func (m *TLWorld) ClazzName() string {
	return "world"
}

func (m *TLWorld) ToClazzWorld() *World {
	m.Data2.ClazzName = ClazzName_world
	return m.Data2
}

func (m *TLWorld) SetMessage2(v string) { m.Data2.Message2 = v }
func (m *TLWorld) GetMessage2() string  { return m.Data2.Message2 }

func (m *TLWorld) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		ClazzID_world: func() error {
			x.PutClazzID(ClazzID_world)

			x.PutString(m.GetMessage2())
			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_world, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		// log.Errorf("not found clazzId by (%s, %d)", Predicate_bind_auth_key_inner, layer)
		return nil
	}
}

func (m *TLWorld) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		ClazzID_world: func() error {
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

type TLHelloSayHello struct {
	ClazzID int32  `json:"_id"`
	Message string `json:"message"`
}

func (m *TLHelloSayHello) Encode(x *bin.Encoder, layer int32) error {
	switch m.ClazzID {
	case ClazzID_hello_sayHello:
		x.PutClazzID(ClazzID_hello_sayHello)
		x.PutString(m.Message)

		return nil
	default:
		return fmt.Errorf("invalid clazz_id error: 0x%x", m.ClazzID)
	}
}

func (m *TLHelloSayHello) Decode(d *bin.Decoder) (err error) {
	switch m.ClazzID {
	case ClazzID_hello_sayHello:
		m.Message, err = d.String()
	default:
		err = fmt.Errorf("invalid clazz_id error: 0x%x", m.ClazzID)
	}

	return
}
