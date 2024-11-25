/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package mt

import (
	"fmt"

	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
)

var (
	_ iface.TLObject
)

// BindAuthKeyInnerClazz <--
//   - TL_BindAuthKeyInner
type BindAuthKeyInnerClazz interface {
	iface.TLObject
	BindAuthKeyInnerClazzName() string
}

func DecodeBindAuthKeyInnerClazz(d *bin.Decoder) (BindAuthKeyInnerClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_bind_auth_key_inner:
		x := &TLBindAuthKeyInner{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeBindAuthKeyInner - unexpected clazzId: %d", id)
	}
}

// TLBindAuthKeyInner <--
type TLBindAuthKeyInner struct {
	ClazzID       uint32 `json:"_id"`
	Nonce         int64  `json:"nonce"`
	TempAuthKeyId int64  `json:"temp_auth_key_id"`
	PermAuthKeyId int64  `json:"perm_auth_key_id"`
	TempSessionId int64  `json:"temp_session_id"`
	ExpiresAt     int32  `json:"expires_at"`
}

func (m *TLBindAuthKeyInner) BindAuthKeyInnerClazzName() string {
	return ClazzName_bind_auth_key_inner
}

func (m *TLBindAuthKeyInner) ClazzName() string {
	return ClazzName_bind_auth_key_inner
}

func (m *TLBindAuthKeyInner) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x75a3f765: func() error {
			x.PutClazzID(0x75a3f765)

			x.PutInt64(m.Nonce)
			x.PutInt64(m.TempAuthKeyId)
			x.PutInt64(m.PermAuthKeyId)
			x.PutInt64(m.TempSessionId)
			x.PutInt32(m.ExpiresAt)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_bind_auth_key_inner, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_bind_auth_key_inner, layer)
	}
}

func (m *TLBindAuthKeyInner) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x75a3f765: func() (err error) {
			m.Nonce, err = d.Int64()
			m.TempAuthKeyId, err = d.Int64()
			m.PermAuthKeyId, err = d.Int64()
			m.TempSessionId, err = d.Int64()
			m.ExpiresAt, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// BindAuthKeyInner <--
type BindAuthKeyInner struct {
	// ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	BindAuthKeyInnerClazz
}

func MakeBindAuthKeyInner(c BindAuthKeyInnerClazz) *BindAuthKeyInner {
	return &BindAuthKeyInner{
		// ClazzID:   c.ClazzID(),
		ClazzName:             c.BindAuthKeyInnerClazzName(),
		BindAuthKeyInnerClazz: c,
	}
}

func (m *BindAuthKeyInner) Encode(x *bin.Encoder, layer int32) error {
	if m.BindAuthKeyInnerClazz != nil {
		return m.BindAuthKeyInnerClazz.Encode(x, layer)
	}

	return fmt.Errorf("BindAuthKeyInner - invalid Clazz")
}

func (m *BindAuthKeyInner) Decode(d *bin.Decoder) (err error) {
	m.BindAuthKeyInnerClazz, err = DecodeBindAuthKeyInnerClazz(d)
	return
}

func (m *BindAuthKeyInner) Match(f ...interface{}) {
	switch c := m.BindAuthKeyInnerClazz.(type) {
	case *TLBindAuthKeyInner:
		for _, v := range f {
			if f1, ok := v.(func(c *TLBindAuthKeyInner) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ClientDHInnerDataClazz <--
//   - TL_ClientDHInnerData
type ClientDHInnerDataClazz interface {
	iface.TLObject
	ClientDHInnerDataClazzName() string
}

func DecodeClientDHInnerDataClazz(d *bin.Decoder) (ClientDHInnerDataClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_client_DH_inner_data:
		x := &TLClientDHInnerData{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeClientDHInnerData - unexpected clazzId: %d", id)
	}
}

// TLClientDHInnerData <--
type TLClientDHInnerData struct {
	ClazzID     uint32     `json:"_id"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	RetryId     int64      `json:"retry_id"`
	GB          string     `json:"g_b"`
}

func (m *TLClientDHInnerData) ClientDHInnerDataClazzName() string {
	return ClazzName_client_DH_inner_data
}

func (m *TLClientDHInnerData) ClazzName() string {
	return ClazzName_client_DH_inner_data
}

func (m *TLClientDHInnerData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6643b654: func() error {
			x.PutClazzID(0x6643b654)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt64(m.RetryId)
			x.PutString(m.GB)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_client_DH_inner_data, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_client_DH_inner_data, layer)
	}
}

func (m *TLClientDHInnerData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6643b654: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			m.RetryId, err = d.Int64()
			m.GB, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ClientDHInnerData <--
type ClientDHInnerData struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	ClientDHInnerDataClazz
}

func MakeClientDHInnerData(c ClientDHInnerDataClazz) *ClientDHInnerData {
	return &ClientDHInnerData{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		ClientDHInnerDataClazz: c,
	}
}

func (m *ClientDHInnerData) Encode(x *bin.Encoder, layer int32) error {
	if m.ClientDHInnerDataClazz != nil {
		return m.ClientDHInnerDataClazz.Encode(x, layer)
	}

	return fmt.Errorf("ClientDHInnerData - invalid Clazz")
}

func (m *ClientDHInnerData) Decode(d *bin.Decoder) (err error) {
	m.ClientDHInnerDataClazz, err = DecodeClientDHInnerDataClazz(d)
	return
}

func (m *ClientDHInnerData) Match(f ...interface{}) {
	switch c := m.ClientDHInnerDataClazz.(type) {
	case *TLClientDHInnerData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLClientDHInnerData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// DestroyAuthKeyResClazz <--
//   - TL_DestroyAuthKeyOk
//   - TL_DestroyAuthKeyNone
//   - TL_DestroyAuthKeyFail
type DestroyAuthKeyResClazz interface {
	iface.TLObject
	DestroyAuthKeyResClazzName() string
}

func DecodeDestroyAuthKeyResClazz(d *bin.Decoder) (DestroyAuthKeyResClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_destroy_auth_key_ok:
		x := &TLDestroyAuthKeyOk{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_destroy_auth_key_none:
		x := &TLDestroyAuthKeyNone{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_destroy_auth_key_fail:
		x := &TLDestroyAuthKeyFail{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDestroyAuthKeyRes - unexpected clazzId: %d", id)
	}
}

// TLDestroyAuthKeyOk <--
type TLDestroyAuthKeyOk struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLDestroyAuthKeyOk) DestroyAuthKeyResClazzName() string {
	return ClazzName_destroy_auth_key_ok
}

func (m *TLDestroyAuthKeyOk) ClazzName() string {
	return ClazzName_destroy_auth_key_ok
}

func (m *TLDestroyAuthKeyOk) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xf660e1d4: func() error {
			x.PutClazzID(0xf660e1d4)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key_ok, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_auth_key_ok, layer)
	}
}

func (m *TLDestroyAuthKeyOk) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xf660e1d4: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroyAuthKeyNone <--
type TLDestroyAuthKeyNone struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLDestroyAuthKeyNone) DestroyAuthKeyResClazzName() string {
	return ClazzName_destroy_auth_key_none
}

func (m *TLDestroyAuthKeyNone) ClazzName() string {
	return ClazzName_destroy_auth_key_none
}

func (m *TLDestroyAuthKeyNone) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa9f2259: func() error {
			x.PutClazzID(0xa9f2259)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key_none, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_auth_key_none, layer)
	}
}

func (m *TLDestroyAuthKeyNone) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa9f2259: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroyAuthKeyFail <--
type TLDestroyAuthKeyFail struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLDestroyAuthKeyFail) DestroyAuthKeyResClazzName() string {
	return ClazzName_destroy_auth_key_fail
}

func (m *TLDestroyAuthKeyFail) ClazzName() string {
	return ClazzName_destroy_auth_key_fail
}

func (m *TLDestroyAuthKeyFail) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xea109b13: func() error {
			x.PutClazzID(0xea109b13)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_destroy_auth_key_fail, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_auth_key_fail, layer)
	}
}

func (m *TLDestroyAuthKeyFail) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xea109b13: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DestroyAuthKeyRes <--
type DestroyAuthKeyRes struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	DestroyAuthKeyResClazz
}

func MakeDestroyAuthKeyRes(c DestroyAuthKeyResClazz) *DestroyAuthKeyRes {
	return &DestroyAuthKeyRes{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		DestroyAuthKeyResClazz: c,
	}
}

func (m *DestroyAuthKeyRes) Encode(x *bin.Encoder, layer int32) error {
	if m.DestroyAuthKeyResClazz != nil {
		return m.DestroyAuthKeyResClazz.Encode(x, layer)
	}

	return fmt.Errorf("DestroyAuthKeyRes - invalid Clazz")
}

func (m *DestroyAuthKeyRes) Decode(d *bin.Decoder) (err error) {
	m.DestroyAuthKeyResClazz, err = DecodeDestroyAuthKeyResClazz(d)
	return
}

func (m *DestroyAuthKeyRes) Match(f ...interface{}) {
	switch c := m.DestroyAuthKeyResClazz.(type) {
	case *TLDestroyAuthKeyOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroyAuthKeyOk) interface{}); ok {
				f1(c)
			}
		}
	case *TLDestroyAuthKeyNone:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroyAuthKeyNone) interface{}); ok {
				f1(c)
			}
		}
	case *TLDestroyAuthKeyFail:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroyAuthKeyFail) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// PQInnerDataClazz <--
//   - TL_PQInnerData
//   - TL_PQInnerDataDc
//   - TL_PQInnerDataTemp
//   - TL_PQInnerDataTempDc
type PQInnerDataClazz interface {
	iface.TLObject
	PQInnerDataClazzName() string
}

func DecodePQInnerDataClazz(d *bin.Decoder) (PQInnerDataClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_p_q_inner_data:
		x := &TLPQInnerData{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_p_q_inner_data_dc:
		x := &TLPQInnerDataDc{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_p_q_inner_data_temp:
		x := &TLPQInnerDataTemp{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_p_q_inner_data_temp_dc:
		x := &TLPQInnerDataTempDc{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePQInnerData - unexpected clazzId: %d", id)
	}
}

// TLPQInnerData <--
type TLPQInnerData struct {
	ClazzID     uint32     `json:"_id"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
}

func (m *TLPQInnerData) PQInnerDataClazzName() string {
	return ClazzName_p_q_inner_data
}

func (m *TLPQInnerData) ClazzName() string {
	return ClazzName_p_q_inner_data
}

func (m *TLPQInnerData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x83c95aec: func() error {
			x.PutClazzID(0x83c95aec)

			x.PutString(m.Pq)
			x.PutString(m.P)
			x.PutString(m.Q)
			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt256(m.NewNonce)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data, layer)
	}
}

func (m *TLPQInnerData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x83c95aec: func() (err error) {
			m.Pq, err = d.String()
			m.P, err = d.String()
			m.Q, err = d.String()
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonce.Decode(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPQInnerDataDc <--
type TLPQInnerDataDc struct {
	ClazzID     uint32     `json:"_id"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
	Dc          int32      `json:"dc"`
}

func (m *TLPQInnerDataDc) PQInnerDataClazzName() string {
	return ClazzName_p_q_inner_data_dc
}

func (m *TLPQInnerDataDc) ClazzName() string {
	return ClazzName_p_q_inner_data_dc
}

func (m *TLPQInnerDataDc) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa9f55f95: func() error {
			x.PutClazzID(0xa9f55f95)

			x.PutString(m.Pq)
			x.PutString(m.P)
			x.PutString(m.Q)
			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt256(m.NewNonce)
			x.PutInt32(m.Dc)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data_dc, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data_dc, layer)
	}
}

func (m *TLPQInnerDataDc) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa9f55f95: func() (err error) {
			m.Pq, err = d.String()
			m.P, err = d.String()
			m.Q, err = d.String()
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonce.Decode(d)
			m.Dc, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPQInnerDataTemp <--
type TLPQInnerDataTemp struct {
	ClazzID     uint32     `json:"_id"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
	ExpiresIn   int32      `json:"expires_in"`
}

func (m *TLPQInnerDataTemp) PQInnerDataClazzName() string {
	return ClazzName_p_q_inner_data_temp
}

func (m *TLPQInnerDataTemp) ClazzName() string {
	return ClazzName_p_q_inner_data_temp
}

func (m *TLPQInnerDataTemp) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3c6a84d4: func() error {
			x.PutClazzID(0x3c6a84d4)

			x.PutString(m.Pq)
			x.PutString(m.P)
			x.PutString(m.Q)
			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt256(m.NewNonce)
			x.PutInt32(m.ExpiresIn)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data_temp, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data_temp, layer)
	}
}

func (m *TLPQInnerDataTemp) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3c6a84d4: func() (err error) {
			m.Pq, err = d.String()
			m.P, err = d.String()
			m.Q, err = d.String()
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonce.Decode(d)
			m.ExpiresIn, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLPQInnerDataTempDc <--
type TLPQInnerDataTempDc struct {
	ClazzID     uint32     `json:"_id"`
	Pq          string     `json:"pq"`
	P           string     `json:"p"`
	Q           string     `json:"q"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	NewNonce    bin.Int256 `json:"new_nonce"`
	Dc          int32      `json:"dc"`
	ExpiresIn   int32      `json:"expires_in"`
}

func (m *TLPQInnerDataTempDc) PQInnerDataClazzName() string {
	return ClazzName_p_q_inner_data_temp_dc
}

func (m *TLPQInnerDataTempDc) ClazzName() string {
	return ClazzName_p_q_inner_data_temp_dc
}

func (m *TLPQInnerDataTempDc) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x56fddf88: func() error {
			x.PutClazzID(0x56fddf88)

			x.PutString(m.Pq)
			x.PutString(m.P)
			x.PutString(m.Q)
			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt256(m.NewNonce)
			x.PutInt32(m.Dc)
			x.PutInt32(m.ExpiresIn)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_p_q_inner_data_temp_dc, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_p_q_inner_data_temp_dc, layer)
	}
}

func (m *TLPQInnerDataTempDc) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x56fddf88: func() (err error) {
			m.Pq, err = d.String()
			m.P, err = d.String()
			m.Q, err = d.String()
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonce.Decode(d)
			m.Dc, err = d.Int32()
			m.ExpiresIn, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// PQInnerData <--
type PQInnerData struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	PQInnerDataClazz
}

func MakePQInnerData(c PQInnerDataClazz) *PQInnerData {
	return &PQInnerData{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		PQInnerDataClazz: c,
	}
}

func (m *PQInnerData) Encode(x *bin.Encoder, layer int32) error {
	if m.PQInnerDataClazz != nil {
		return m.PQInnerDataClazz.Encode(x, layer)
	}

	return fmt.Errorf("PQInnerData - invalid Clazz")
}

func (m *PQInnerData) Decode(d *bin.Decoder) (err error) {
	m.PQInnerDataClazz, err = DecodePQInnerDataClazz(d)
	return
}

func (m *PQInnerData) Match(f ...interface{}) {
	switch c := m.PQInnerDataClazz.(type) {
	case *TLPQInnerData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerData) interface{}); ok {
				f1(c)
			}
		}
	case *TLPQInnerDataDc:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerDataDc) interface{}); ok {
				f1(c)
			}
		}
	case *TLPQInnerDataTemp:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerDataTemp) interface{}); ok {
				f1(c)
			}
		}
	case *TLPQInnerDataTempDc:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPQInnerDataTempDc) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ResPQClazz <--
//   - TL_ResPQ
type ResPQClazz interface {
	iface.TLObject
	ResPQClazzName() string
}

func DecodeResPQClazz(d *bin.Decoder) (ResPQClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_resPQ:
		x := &TLResPQ{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeResPQ - unexpected clazzId: %d", id)
	}
}

// TLResPQ <--
type TLResPQ struct {
	ClazzID                     uint32     `json:"_id"`
	Nonce                       bin.Int128 `json:"nonce"`
	ServerNonce                 bin.Int128 `json:"server_nonce"`
	Pq                          string     `json:"pq"`
	ServerPublicKeyFingerprints []int64    `json:"server_public_key_fingerprints"`
}

func (m *TLResPQ) ResPQClazzName() string {
	return ClazzName_resPQ
}

func (m *TLResPQ) ClazzName() string {
	return ClazzName_resPQ
}

func (m *TLResPQ) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5162463: func() error {
			x.PutClazzID(0x5162463)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutString(m.Pq)

			iface.EncodeInt64List(x, m.ServerPublicKeyFingerprints)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_resPQ, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_resPQ, layer)
	}
}

func (m *TLResPQ) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5162463: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			m.Pq, err = d.String()

			m.ServerPublicKeyFingerprints, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ResPQ <--
type ResPQ struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	ResPQClazz
}

func MakeResPQ(c ResPQClazz) *ResPQ {
	return &ResPQ{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		ResPQClazz: c,
	}
}

func (m *ResPQ) Encode(x *bin.Encoder, layer int32) error {
	if m.ResPQClazz != nil {
		return m.ResPQClazz.Encode(x, layer)
	}

	return fmt.Errorf("ResPQ - invalid Clazz")
}

func (m *ResPQ) Decode(d *bin.Decoder) (err error) {
	m.ResPQClazz, err = DecodeResPQClazz(d)
	return
}

func (m *ResPQ) Match(f ...interface{}) {
	switch c := m.ResPQClazz.(type) {
	case *TLResPQ:
		for _, v := range f {
			if f1, ok := v.(func(c *TLResPQ) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ServerDHInnerDataClazz <--
//   - TL_ServerDHInnerData
type ServerDHInnerDataClazz interface {
	iface.TLObject
	ServerDHInnerDataClazzName() string
}

func DecodeServerDHInnerDataClazz(d *bin.Decoder) (ServerDHInnerDataClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_server_DH_inner_data:
		x := &TLServerDHInnerData{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeServerDHInnerData - unexpected clazzId: %d", id)
	}
}

// TLServerDHInnerData <--
type TLServerDHInnerData struct {
	ClazzID     uint32     `json:"_id"`
	Nonce       bin.Int128 `json:"nonce"`
	ServerNonce bin.Int128 `json:"server_nonce"`
	G           int32      `json:"g"`
	DhPrime     string     `json:"dh_prime"`
	GA          string     `json:"g_a"`
	ServerTime  int32      `json:"server_time"`
}

func (m *TLServerDHInnerData) ServerDHInnerDataClazzName() string {
	return ClazzName_server_DH_inner_data
}

func (m *TLServerDHInnerData) ClazzName() string {
	return ClazzName_server_DH_inner_data
}

func (m *TLServerDHInnerData) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xb5890dba: func() error {
			x.PutClazzID(0xb5890dba)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt32(m.G)
			x.PutString(m.DhPrime)
			x.PutString(m.GA)
			x.PutInt32(m.ServerTime)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_server_DH_inner_data, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_server_DH_inner_data, layer)
	}
}

func (m *TLServerDHInnerData) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xb5890dba: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			m.G, err = d.Int32()
			m.DhPrime, err = d.String()
			m.GA, err = d.String()
			m.ServerTime, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ServerDHInnerData <--
type ServerDHInnerData struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	ServerDHInnerDataClazz
}

func MakeServerDHInnerData(c ServerDHInnerDataClazz) *ServerDHInnerData {
	return &ServerDHInnerData{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		ServerDHInnerDataClazz: c,
	}
}

func (m *ServerDHInnerData) Encode(x *bin.Encoder, layer int32) error {
	if m.ServerDHInnerDataClazz != nil {
		return m.ServerDHInnerDataClazz.Encode(x, layer)
	}

	return fmt.Errorf("ServerDHInnerData - invalid Clazz")
}

func (m *ServerDHInnerData) Decode(d *bin.Decoder) (err error) {
	m.ServerDHInnerDataClazz, err = DecodeServerDHInnerDataClazz(d)
	return
}

func (m *ServerDHInnerData) Match(f ...interface{}) {
	switch c := m.ServerDHInnerDataClazz.(type) {
	case *TLServerDHInnerData:
		for _, v := range f {
			if f1, ok := v.(func(c *TLServerDHInnerData) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// ServerDHParamsClazz <--
//   - TL_ServerDHParamsFail
//   - TL_ServerDHParamsOk
type ServerDHParamsClazz interface {
	iface.TLObject
	ServerDHParamsClazzName() string
}

func DecodeServerDHParamsClazz(d *bin.Decoder) (ServerDHParamsClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_server_DH_params_fail:
		x := &TLServerDHParamsFail{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_server_DH_params_ok:
		x := &TLServerDHParamsOk{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeServerDHParams - unexpected clazzId: %d", id)
	}
}

// TLServerDHParamsFail <--
type TLServerDHParamsFail struct {
	ClazzID      uint32     `json:"_id"`
	Nonce        bin.Int128 `json:"nonce"`
	ServerNonce  bin.Int128 `json:"server_nonce"`
	NewNonceHash bin.Int128 `json:"new_nonce_hash"`
}

func (m *TLServerDHParamsFail) ServerDHParamsClazzName() string {
	return ClazzName_server_DH_params_fail
}

func (m *TLServerDHParamsFail) ClazzName() string {
	return ClazzName_server_DH_params_fail
}

func (m *TLServerDHParamsFail) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x79cb045d: func() error {
			x.PutClazzID(0x79cb045d)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt128(m.NewNonceHash)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_server_DH_params_fail, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_server_DH_params_fail, layer)
	}
}

func (m *TLServerDHParamsFail) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x79cb045d: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonceHash.Decode(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLServerDHParamsOk <--
type TLServerDHParamsOk struct {
	ClazzID         uint32     `json:"_id"`
	Nonce           bin.Int128 `json:"nonce"`
	ServerNonce     bin.Int128 `json:"server_nonce"`
	EncryptedAnswer string     `json:"encrypted_answer"`
}

func (m *TLServerDHParamsOk) ServerDHParamsClazzName() string {
	return ClazzName_server_DH_params_ok
}

func (m *TLServerDHParamsOk) ClazzName() string {
	return ClazzName_server_DH_params_ok
}

func (m *TLServerDHParamsOk) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd0e8075c: func() error {
			x.PutClazzID(0xd0e8075c)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutString(m.EncryptedAnswer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_server_DH_params_ok, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_server_DH_params_ok, layer)
	}
}

func (m *TLServerDHParamsOk) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd0e8075c: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			m.EncryptedAnswer, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// ServerDHParams <--
type ServerDHParams struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	ServerDHParamsClazz
}

func MakeServerDHParams(c ServerDHParamsClazz) *ServerDHParams {
	return &ServerDHParams{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		ServerDHParamsClazz: c,
	}
}

func (m *ServerDHParams) Encode(x *bin.Encoder, layer int32) error {
	if m.ServerDHParamsClazz != nil {
		return m.ServerDHParamsClazz.Encode(x, layer)
	}

	return fmt.Errorf("ServerDHParams - invalid Clazz")
}

func (m *ServerDHParams) Decode(d *bin.Decoder) (err error) {
	m.ServerDHParamsClazz, err = DecodeServerDHParamsClazz(d)
	return
}

func (m *ServerDHParams) Match(f ...interface{}) {
	switch c := m.ServerDHParamsClazz.(type) {
	case *TLServerDHParamsFail:
		for _, v := range f {
			if f1, ok := v.(func(c *TLServerDHParamsFail) interface{}); ok {
				f1(c)
			}
		}
	case *TLServerDHParamsOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLServerDHParamsOk) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// SetClientDHParamsAnswerClazz <--
//   - TL_DhGenOk
//   - TL_DhGenRetry
//   - TL_DhGenFail
type SetClientDHParamsAnswerClazz interface {
	iface.TLObject
	SetClientDHParamsAnswerClazzName() string
}

func DecodeSetClientDHParamsAnswerClazz(d *bin.Decoder) (SetClientDHParamsAnswerClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_dh_gen_ok:
		x := &TLDhGenOk{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_dh_gen_retry:
		x := &TLDhGenRetry{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_dh_gen_fail:
		x := &TLDhGenFail{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeSetClientDHParamsAnswer - unexpected clazzId: %d", id)
	}
}

// TLDhGenOk <--
type TLDhGenOk struct {
	ClazzID       uint32     `json:"_id"`
	Nonce         bin.Int128 `json:"nonce"`
	ServerNonce   bin.Int128 `json:"server_nonce"`
	NewNonceHash1 bin.Int128 `json:"new_nonce_hash1"`
}

func (m *TLDhGenOk) SetClientDHParamsAnswerClazzName() string {
	return ClazzName_dh_gen_ok
}

func (m *TLDhGenOk) ClazzName() string {
	return ClazzName_dh_gen_ok
}

func (m *TLDhGenOk) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x3bcbf734: func() error {
			x.PutClazzID(0x3bcbf734)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt128(m.NewNonceHash1)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dh_gen_ok, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dh_gen_ok, layer)
	}
}

func (m *TLDhGenOk) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x3bcbf734: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonceHash1.Decode(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDhGenRetry <--
type TLDhGenRetry struct {
	ClazzID       uint32     `json:"_id"`
	Nonce         bin.Int128 `json:"nonce"`
	ServerNonce   bin.Int128 `json:"server_nonce"`
	NewNonceHash2 bin.Int128 `json:"new_nonce_hash2"`
}

func (m *TLDhGenRetry) SetClientDHParamsAnswerClazzName() string {
	return ClazzName_dh_gen_retry
}

func (m *TLDhGenRetry) ClazzName() string {
	return ClazzName_dh_gen_retry
}

func (m *TLDhGenRetry) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x46dc1fb9: func() error {
			x.PutClazzID(0x46dc1fb9)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt128(m.NewNonceHash2)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dh_gen_retry, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dh_gen_retry, layer)
	}
}

func (m *TLDhGenRetry) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x46dc1fb9: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonceHash2.Decode(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDhGenFail <--
type TLDhGenFail struct {
	ClazzID       uint32     `json:"_id"`
	Nonce         bin.Int128 `json:"nonce"`
	ServerNonce   bin.Int128 `json:"server_nonce"`
	NewNonceHash3 bin.Int128 `json:"new_nonce_hash3"`
}

func (m *TLDhGenFail) SetClientDHParamsAnswerClazzName() string {
	return ClazzName_dh_gen_fail
}

func (m *TLDhGenFail) ClazzName() string {
	return ClazzName_dh_gen_fail
}

func (m *TLDhGenFail) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa69dae02: func() error {
			x.PutClazzID(0xa69dae02)

			x.PutInt128(m.Nonce)
			x.PutInt128(m.ServerNonce)
			x.PutInt128(m.NewNonceHash3)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_dh_gen_fail, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_dh_gen_fail, layer)
	}
}

func (m *TLDhGenFail) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa69dae02: func() (err error) {
			err = m.Nonce.Decode(d)
			err = m.ServerNonce.Decode(d)
			err = m.NewNonceHash3.Decode(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// SetClientDHParamsAnswer <--
type SetClientDHParamsAnswer struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	SetClientDHParamsAnswerClazz
}

func MakeSetClientDHParamsAnswer(c SetClientDHParamsAnswerClazz) *SetClientDHParamsAnswer {
	return &SetClientDHParamsAnswer{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		SetClientDHParamsAnswerClazz: c,
	}
}

func (m *SetClientDHParamsAnswer) Encode(x *bin.Encoder, layer int32) error {
	if m.SetClientDHParamsAnswerClazz != nil {
		return m.SetClientDHParamsAnswerClazz.Encode(x, layer)
	}

	return fmt.Errorf("SetClientDHParamsAnswer - invalid Clazz")
}

func (m *SetClientDHParamsAnswer) Decode(d *bin.Decoder) (err error) {
	m.SetClientDHParamsAnswerClazz, err = DecodeSetClientDHParamsAnswerClazz(d)
	return
}

func (m *SetClientDHParamsAnswer) Match(f ...interface{}) {
	switch c := m.SetClientDHParamsAnswerClazz.(type) {
	case *TLDhGenOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDhGenOk) interface{}); ok {
				f1(c)
			}
		}
	case *TLDhGenRetry:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDhGenRetry) interface{}); ok {
				f1(c)
			}
		}
	case *TLDhGenFail:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDhGenFail) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// AccessPointRuleClazz <--
//   - TL_AccessPointRule
type AccessPointRuleClazz interface {
	iface.TLObject
	AccessPointRuleClazzName() string
}

func DecodeAccessPointRuleClazz(d *bin.Decoder) (AccessPointRuleClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_accessPointRule:
		x := &TLAccessPointRule{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeAccessPointRule - unexpected clazzId: %d", id)
	}
}

// TLAccessPointRule <--
type TLAccessPointRule struct {
	ClazzID          uint32    `json:"_id"`
	PhonePrefixRules string    `json:"phone_prefix_rules"`
	DcId             int32     `json:"dc_id"`
	Ips              []*IpPort `json:"ips"`
}

func (m *TLAccessPointRule) AccessPointRuleClazzName() string {
	return ClazzName_accessPointRule
}

func (m *TLAccessPointRule) ClazzName() string {
	return ClazzName_accessPointRule
}

func (m *TLAccessPointRule) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4679b65f: func() error {
			x.PutClazzID(0x4679b65f)

			x.PutString(m.PhonePrefixRules)
			x.PutInt32(m.DcId)
			// x.PutClazzID(iface.ClazzID_vector)
			x.PutInt(len(m.Ips))
			for _, v := range m.Ips {
				_ = v.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_accessPointRule, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_accessPointRule, layer)
	}
}

func (m *TLAccessPointRule) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4679b65f: func() (err error) {
			m.PhonePrefixRules, err = d.String()
			m.DcId, err = d.Int32()
			// c2, err2 := d.ClazzID()
			// if c2 != int32(iface.ClazzID_vector) {
			//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
			//     return err2
			// }
			l2, err2 := d.Int()
			_ = err2
			v2 := make([]*IpPort, l2)
			for i := 0; i < l2; i++ {
				vv := new(IpPort)
				err2 = vv.Decode(d)
				_ = err2
				v2[i] = vv
			}
			m.Ips = v2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// AccessPointRule <--
type AccessPointRule struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	AccessPointRuleClazz
}

func MakeAccessPointRule(c AccessPointRuleClazz) *AccessPointRule {
	return &AccessPointRule{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		AccessPointRuleClazz: c,
	}
}

func (m *AccessPointRule) Encode(x *bin.Encoder, layer int32) error {
	if m.AccessPointRuleClazz != nil {
		return m.AccessPointRuleClazz.Encode(x, layer)
	}

	return fmt.Errorf("AccessPointRule - invalid Clazz")
}

func (m *AccessPointRule) Decode(d *bin.Decoder) (err error) {
	m.AccessPointRuleClazz, err = DecodeAccessPointRuleClazz(d)
	return
}

func (m *AccessPointRule) Match(f ...interface{}) {
	switch c := m.AccessPointRuleClazz.(type) {
	case *TLAccessPointRule:
		for _, v := range f {
			if f1, ok := v.(func(c *TLAccessPointRule) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// BadMsgNotificationClazz <--
//   - TL_BadMsgNotification
//   - TL_BadServerSalt
type BadMsgNotificationClazz interface {
	iface.TLObject
	BadMsgNotificationClazzName() string
}

func DecodeBadMsgNotificationClazz(d *bin.Decoder) (BadMsgNotificationClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_bad_msg_notification:
		x := &TLBadMsgNotification{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_bad_server_salt:
		x := &TLBadServerSalt{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeBadMsgNotification - unexpected clazzId: %d", id)
	}
}

// TLBadMsgNotification <--
type TLBadMsgNotification struct {
	ClazzID     uint32 `json:"_id"`
	BadMsgId    int64  `json:"bad_msg_id"`
	BadMsgSeqno int32  `json:"bad_msg_seqno"`
	ErrorCode   int32  `json:"error_code"`
}

func (m *TLBadMsgNotification) BadMsgNotificationClazzName() string {
	return ClazzName_bad_msg_notification
}

func (m *TLBadMsgNotification) ClazzName() string {
	return ClazzName_bad_msg_notification
}

func (m *TLBadMsgNotification) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa7eff811: func() error {
			x.PutClazzID(0xa7eff811)

			x.PutInt64(m.BadMsgId)
			x.PutInt32(m.BadMsgSeqno)
			x.PutInt32(m.ErrorCode)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_bad_msg_notification, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_bad_msg_notification, layer)
	}
}

func (m *TLBadMsgNotification) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa7eff811: func() (err error) {
			m.BadMsgId, err = d.Int64()
			m.BadMsgSeqno, err = d.Int32()
			m.ErrorCode, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLBadServerSalt <--
type TLBadServerSalt struct {
	ClazzID       uint32 `json:"_id"`
	BadMsgId      int64  `json:"bad_msg_id"`
	BadMsgSeqno   int32  `json:"bad_msg_seqno"`
	ErrorCode     int32  `json:"error_code"`
	NewServerSalt int64  `json:"new_server_salt"`
}

func (m *TLBadServerSalt) BadMsgNotificationClazzName() string {
	return ClazzName_bad_server_salt
}

func (m *TLBadServerSalt) ClazzName() string {
	return ClazzName_bad_server_salt
}

func (m *TLBadServerSalt) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xedab447b: func() error {
			x.PutClazzID(0xedab447b)

			x.PutInt64(m.BadMsgId)
			x.PutInt32(m.BadMsgSeqno)
			x.PutInt32(m.ErrorCode)
			x.PutInt64(m.NewServerSalt)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_bad_server_salt, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_bad_server_salt, layer)
	}
}

func (m *TLBadServerSalt) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xedab447b: func() (err error) {
			m.BadMsgId, err = d.Int64()
			m.BadMsgSeqno, err = d.Int32()
			m.ErrorCode, err = d.Int32()
			m.NewServerSalt, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// BadMsgNotification <--
type BadMsgNotification struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	BadMsgNotificationClazz
}

func MakeBadMsgNotification(c BadMsgNotificationClazz) *BadMsgNotification {
	return &BadMsgNotification{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		BadMsgNotificationClazz: c,
	}
}

func (m *BadMsgNotification) Encode(x *bin.Encoder, layer int32) error {
	if m.BadMsgNotificationClazz != nil {
		return m.BadMsgNotificationClazz.Encode(x, layer)
	}

	return fmt.Errorf("BadMsgNotification - invalid Clazz")
}

func (m *BadMsgNotification) Decode(d *bin.Decoder) (err error) {
	m.BadMsgNotificationClazz, err = DecodeBadMsgNotificationClazz(d)
	return
}

func (m *BadMsgNotification) Match(f ...interface{}) {
	switch c := m.BadMsgNotificationClazz.(type) {
	case *TLBadMsgNotification:
		for _, v := range f {
			if f1, ok := v.(func(c *TLBadMsgNotification) interface{}); ok {
				f1(c)
			}
		}
	case *TLBadServerSalt:
		for _, v := range f {
			if f1, ok := v.(func(c *TLBadServerSalt) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// DestroySessionResClazz <--
//   - TL_DestroySessionOk
//   - TL_DestroySessionNone
type DestroySessionResClazz interface {
	iface.TLObject
	DestroySessionResClazzName() string
}

func DecodeDestroySessionResClazz(d *bin.Decoder) (DestroySessionResClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_destroy_session_ok:
		x := &TLDestroySessionOk{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_destroy_session_none:
		x := &TLDestroySessionNone{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeDestroySessionRes - unexpected clazzId: %d", id)
	}
}

// TLDestroySessionOk <--
type TLDestroySessionOk struct {
	ClazzID   uint32 `json:"_id"`
	SessionId int64  `json:"session_id"`
}

func (m *TLDestroySessionOk) DestroySessionResClazzName() string {
	return ClazzName_destroy_session_ok
}

func (m *TLDestroySessionOk) ClazzName() string {
	return ClazzName_destroy_session_ok
}

func (m *TLDestroySessionOk) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe22045fc: func() error {
			x.PutClazzID(0xe22045fc)

			x.PutInt64(m.SessionId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_destroy_session_ok, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_session_ok, layer)
	}
}

func (m *TLDestroySessionOk) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe22045fc: func() (err error) {
			m.SessionId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLDestroySessionNone <--
type TLDestroySessionNone struct {
	ClazzID   uint32 `json:"_id"`
	SessionId int64  `json:"session_id"`
}

func (m *TLDestroySessionNone) DestroySessionResClazzName() string {
	return ClazzName_destroy_session_none
}

func (m *TLDestroySessionNone) ClazzName() string {
	return ClazzName_destroy_session_none
}

func (m *TLDestroySessionNone) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x62d350c9: func() error {
			x.PutClazzID(0x62d350c9)

			x.PutInt64(m.SessionId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_destroy_session_none, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_destroy_session_none, layer)
	}
}

func (m *TLDestroySessionNone) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x62d350c9: func() (err error) {
			m.SessionId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// DestroySessionRes <--
type DestroySessionRes struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	DestroySessionResClazz
}

func MakeDestroySessionRes(c DestroySessionResClazz) *DestroySessionRes {
	return &DestroySessionRes{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		DestroySessionResClazz: c,
	}
}

func (m *DestroySessionRes) Encode(x *bin.Encoder, layer int32) error {
	if m.DestroySessionResClazz != nil {
		return m.DestroySessionResClazz.Encode(x, layer)
	}

	return fmt.Errorf("DestroySessionRes - invalid Clazz")
}

func (m *DestroySessionRes) Decode(d *bin.Decoder) (err error) {
	m.DestroySessionResClazz, err = DecodeDestroySessionResClazz(d)
	return
}

func (m *DestroySessionRes) Match(f ...interface{}) {
	switch c := m.DestroySessionResClazz.(type) {
	case *TLDestroySessionOk:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroySessionOk) interface{}); ok {
				f1(c)
			}
		}
	case *TLDestroySessionNone:
		for _, v := range f {
			if f1, ok := v.(func(c *TLDestroySessionNone) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// FutureSaltClazz <--
//   - TL_FutureSalt
type FutureSaltClazz interface {
	iface.TLObject
	FutureSaltClazzName() string
}

func DecodeFutureSaltClazz(d *bin.Decoder) (FutureSaltClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_future_salt:
		x := &TLFutureSalt{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeFutureSalt - unexpected clazzId: %d", id)
	}
}

// TLFutureSalt <--
type TLFutureSalt struct {
	ClazzID    uint32 `json:"_id"`
	ValidSince int32  `json:"valid_since"`
	ValidUntil int32  `json:"valid_until"`
	Salt       int64  `json:"salt"`
}

func (m *TLFutureSalt) FutureSaltClazzName() string {
	return ClazzName_future_salt
}

func (m *TLFutureSalt) ClazzName() string {
	return ClazzName_future_salt
}

func (m *TLFutureSalt) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x949d9dc: func() error {
			x.PutClazzID(0x949d9dc)

			x.PutInt32(m.ValidSince)
			x.PutInt32(m.ValidUntil)
			x.PutInt64(m.Salt)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_future_salt, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_future_salt, layer)
	}
}

func (m *TLFutureSalt) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x949d9dc: func() (err error) {
			m.ValidSince, err = d.Int32()
			m.ValidUntil, err = d.Int32()
			m.Salt, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// FutureSalt <--
type FutureSalt struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	FutureSaltClazz
}

func MakeFutureSalt(c FutureSaltClazz) *FutureSalt {
	return &FutureSalt{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		FutureSaltClazz: c,
	}
}

func (m *FutureSalt) Encode(x *bin.Encoder, layer int32) error {
	if m.FutureSaltClazz != nil {
		return m.FutureSaltClazz.Encode(x, layer)
	}

	return fmt.Errorf("FutureSalt - invalid Clazz")
}

func (m *FutureSalt) Decode(d *bin.Decoder) (err error) {
	m.FutureSaltClazz, err = DecodeFutureSaltClazz(d)
	return
}

func (m *FutureSalt) Match(f ...interface{}) {
	switch c := m.FutureSaltClazz.(type) {
	case *TLFutureSalt:
		for _, v := range f {
			if f1, ok := v.(func(c *TLFutureSalt) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// FutureSaltsClazz <--
//   - TL_FutureSalts
type FutureSaltsClazz interface {
	iface.TLObject
	FutureSaltsClazzName() string
}

func DecodeFutureSaltsClazz(d *bin.Decoder) (FutureSaltsClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_future_salts:
		x := &TLFutureSalts{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeFutureSalts - unexpected clazzId: %d", id)
	}
}

// TLFutureSalts <--
type TLFutureSalts struct {
	ClazzID  uint32          `json:"_id"`
	ReqMsgId int64           `json:"req_msg_id"`
	Now      int32           `json:"now"`
	Salts    []*TLFutureSalt `json:"salts"`
}

func (m *TLFutureSalts) FutureSaltsClazzName() string {
	return ClazzName_future_salts
}

func (m *TLFutureSalts) ClazzName() string {
	return ClazzName_future_salts
}

func (m *TLFutureSalts) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xae500895: func() error {
			x.PutClazzID(0xae500895)

			x.PutInt64(m.ReqMsgId)
			x.PutInt32(m.Now)
			// x.PutClazzID(iface.ClazzID_vector)
			x.PutInt(len(m.Salts))
			for _, v := range m.Salts {
				_ = v.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_future_salts, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_future_salts, layer)
	}
}

func (m *TLFutureSalts) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xae500895: func() (err error) {
			m.ReqMsgId, err = d.Int64()
			m.Now, err = d.Int32()
			// c2, err2 := d.ClazzID()
			// if c2 != int32(iface.ClazzID_vector) {
			//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
			//     return err2
			// }
			l2, err2 := d.Int()
			_ = err2
			v2 := make([]*TLFutureSalt, l2)
			for i := 0; i < l2; i++ {
				v2[i] = &TLFutureSalt{ClazzID: ClazzID_future_salts}
				err2 = v2[i].Decode(d)
				_ = err2
			}
			m.Salts = v2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// FutureSalts <--
type FutureSalts struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	FutureSaltsClazz
}

func MakeFutureSalts(c FutureSaltsClazz) *FutureSalts {
	return &FutureSalts{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		FutureSaltsClazz: c,
	}
}

func (m *FutureSalts) Encode(x *bin.Encoder, layer int32) error {
	if m.FutureSaltsClazz != nil {
		return m.FutureSaltsClazz.Encode(x, layer)
	}

	return fmt.Errorf("FutureSalts - invalid Clazz")
}

func (m *FutureSalts) Decode(d *bin.Decoder) (err error) {
	m.FutureSaltsClazz, err = DecodeFutureSaltsClazz(d)
	return
}

func (m *FutureSalts) Match(f ...interface{}) {
	switch c := m.FutureSaltsClazz.(type) {
	case *TLFutureSalts:
		for _, v := range f {
			if f1, ok := v.(func(c *TLFutureSalts) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// HelpConfigSimpleClazz <--
//   - TL_HelpConfigSimple
type HelpConfigSimpleClazz interface {
	iface.TLObject
	HelpConfigSimpleClazzName() string
}

func DecodeHelpConfigSimpleClazz(d *bin.Decoder) (HelpConfigSimpleClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_help_configSimple:
		x := &TLHelpConfigSimple{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeHelpConfigSimple - unexpected clazzId: %d", id)
	}
}

// TLHelpConfigSimple <--
type TLHelpConfigSimple struct {
	ClazzID uint32             `json:"_id"`
	Date    int32              `json:"date"`
	Expires int32              `json:"expires"`
	Rules   []*AccessPointRule `json:"rules"`
}

func (m *TLHelpConfigSimple) HelpConfigSimpleClazzName() string {
	return ClazzName_help_configSimple
}

func (m *TLHelpConfigSimple) ClazzName() string {
	return ClazzName_help_configSimple
}

func (m *TLHelpConfigSimple) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5a592a6c: func() error {
			x.PutClazzID(0x5a592a6c)

			x.PutInt32(m.Date)
			x.PutInt32(m.Expires)
			// x.PutClazzID(iface.ClazzID_vector)
			x.PutInt(len(m.Rules))
			for _, v := range m.Rules {
				_ = v.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_help_configSimple, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_help_configSimple, layer)
	}
}

func (m *TLHelpConfigSimple) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5a592a6c: func() (err error) {
			m.Date, err = d.Int32()
			m.Expires, err = d.Int32()
			// c2, err2 := d.ClazzID()
			// if c2 != int32(iface.ClazzID_vector) {
			//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 2, c2)
			//     return err2
			// }
			l2, err2 := d.Int()
			_ = err2
			v2 := make([]*AccessPointRule, l2)
			for i := 0; i < l2; i++ {
				vv := new(AccessPointRule)
				err2 = vv.Decode(d)
				_ = err2
				v2[i] = vv
			}
			m.Rules = v2

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// HelpConfigSimple <--
type HelpConfigSimple struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	HelpConfigSimpleClazz
}

func MakeHelpConfigSimple(c HelpConfigSimpleClazz) *HelpConfigSimple {
	return &HelpConfigSimple{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		HelpConfigSimpleClazz: c,
	}
}

func (m *HelpConfigSimple) Encode(x *bin.Encoder, layer int32) error {
	if m.HelpConfigSimpleClazz != nil {
		return m.HelpConfigSimpleClazz.Encode(x, layer)
	}

	return fmt.Errorf("HelpConfigSimple - invalid Clazz")
}

func (m *HelpConfigSimple) Decode(d *bin.Decoder) (err error) {
	m.HelpConfigSimpleClazz, err = DecodeHelpConfigSimpleClazz(d)
	return
}

func (m *HelpConfigSimple) Match(f ...interface{}) {
	switch c := m.HelpConfigSimpleClazz.(type) {
	case *TLHelpConfigSimple:
		for _, v := range f {
			if f1, ok := v.(func(c *TLHelpConfigSimple) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// HttpWaitClazz <--
//   - TL_HttpWait
type HttpWaitClazz interface {
	iface.TLObject
	HttpWaitClazzName() string
}

func DecodeHttpWaitClazz(d *bin.Decoder) (HttpWaitClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_http_wait:
		x := &TLHttpWait{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeHttpWait - unexpected clazzId: %d", id)
	}
}

// TLHttpWait <--
type TLHttpWait struct {
	ClazzID   uint32 `json:"_id"`
	MaxDelay  int32  `json:"max_delay"`
	WaitAfter int32  `json:"wait_after"`
	MaxWait   int32  `json:"max_wait"`
}

func (m *TLHttpWait) HttpWaitClazzName() string {
	return ClazzName_http_wait
}

func (m *TLHttpWait) ClazzName() string {
	return ClazzName_http_wait
}

func (m *TLHttpWait) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9299359f: func() error {
			x.PutClazzID(0x9299359f)

			x.PutInt32(m.MaxDelay)
			x.PutInt32(m.WaitAfter)
			x.PutInt32(m.MaxWait)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_http_wait, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_http_wait, layer)
	}
}

func (m *TLHttpWait) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9299359f: func() (err error) {
			m.MaxDelay, err = d.Int32()
			m.WaitAfter, err = d.Int32()
			m.MaxWait, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// HttpWait <--
type HttpWait struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	HttpWaitClazz
}

func MakeHttpWait(c HttpWaitClazz) *HttpWait {
	return &HttpWait{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		HttpWaitClazz: c,
	}
}

func (m *HttpWait) Encode(x *bin.Encoder, layer int32) error {
	if m.HttpWaitClazz != nil {
		return m.HttpWaitClazz.Encode(x, layer)
	}

	return fmt.Errorf("HttpWait - invalid Clazz")
}

func (m *HttpWait) Decode(d *bin.Decoder) (err error) {
	m.HttpWaitClazz, err = DecodeHttpWaitClazz(d)
	return
}

func (m *HttpWait) Match(f ...interface{}) {
	switch c := m.HttpWaitClazz.(type) {
	case *TLHttpWait:
		for _, v := range f {
			if f1, ok := v.(func(c *TLHttpWait) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// IpPortClazz <--
//   - TL_IpPort
//   - TL_IpPortSecret
type IpPortClazz interface {
	iface.TLObject
	IpPortClazzName() string
}

func DecodeIpPortClazz(d *bin.Decoder) (IpPortClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_ipPort:
		x := &TLIpPort{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_ipPortSecret:
		x := &TLIpPortSecret{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeIpPort - unexpected clazzId: %d", id)
	}
}

// TLIpPort <--
type TLIpPort struct {
	ClazzID uint32 `json:"_id"`
	Ipv4    int32  `json:"ipv4"`
	Port    int32  `json:"port"`
}

func (m *TLIpPort) IpPortClazzName() string {
	return ClazzName_ipPort
}

func (m *TLIpPort) ClazzName() string {
	return ClazzName_ipPort
}

func (m *TLIpPort) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xd433ad73: func() error {
			x.PutClazzID(0xd433ad73)

			x.PutInt32(m.Ipv4)
			x.PutInt32(m.Port)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_ipPort, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_ipPort, layer)
	}
}

func (m *TLIpPort) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xd433ad73: func() (err error) {
			m.Ipv4, err = d.Int32()
			m.Port, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLIpPortSecret <--
type TLIpPortSecret struct {
	ClazzID uint32 `json:"_id"`
	Ipv4    int32  `json:"ipv4"`
	Port    int32  `json:"port"`
	Secret  []byte `json:"secret"`
}

func (m *TLIpPortSecret) IpPortClazzName() string {
	return ClazzName_ipPortSecret
}

func (m *TLIpPortSecret) ClazzName() string {
	return ClazzName_ipPortSecret
}

func (m *TLIpPortSecret) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x37982646: func() error {
			x.PutClazzID(0x37982646)

			x.PutInt32(m.Ipv4)
			x.PutInt32(m.Port)
			x.PutBytes(m.Secret)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_ipPortSecret, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_ipPortSecret, layer)
	}
}

func (m *TLIpPortSecret) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x37982646: func() (err error) {
			m.Ipv4, err = d.Int32()
			m.Port, err = d.Int32()
			m.Secret, err = d.Bytes()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// IpPort <--
type IpPort struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	IpPortClazz
}

func MakeIpPort(c IpPortClazz) *IpPort {
	return &IpPort{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		IpPortClazz: c,
	}
}

func (m *IpPort) Encode(x *bin.Encoder, layer int32) error {
	if m.IpPortClazz != nil {
		return m.IpPortClazz.Encode(x, layer)
	}

	return fmt.Errorf("IpPort - invalid Clazz")
}

func (m *IpPort) Decode(d *bin.Decoder) (err error) {
	m.IpPortClazz, err = DecodeIpPortClazz(d)
	return
}

func (m *IpPort) Match(f ...interface{}) {
	switch c := m.IpPortClazz.(type) {
	case *TLIpPort:
		for _, v := range f {
			if f1, ok := v.(func(c *TLIpPort) interface{}); ok {
				f1(c)
			}
		}
	case *TLIpPortSecret:
		for _, v := range f {
			if f1, ok := v.(func(c *TLIpPortSecret) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// MsgDetailedInfoClazz <--
//   - TL_MsgDetailedInfo
//   - TL_MsgNewDetailedInfo
type MsgDetailedInfoClazz interface {
	iface.TLObject
	MsgDetailedInfoClazzName() string
}

func DecodeMsgDetailedInfoClazz(d *bin.Decoder) (MsgDetailedInfoClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_msg_detailed_info:
		x := &TLMsgDetailedInfo{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_msg_new_detailed_info:
		x := &TLMsgNewDetailedInfo{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgDetailedInfo - unexpected clazzId: %d", id)
	}
}

// TLMsgDetailedInfo <--
type TLMsgDetailedInfo struct {
	ClazzID     uint32 `json:"_id"`
	MsgId       int64  `json:"msg_id"`
	AnswerMsgId int64  `json:"answer_msg_id"`
	Bytes       int32  `json:"bytes"`
	Status      int32  `json:"status"`
}

func (m *TLMsgDetailedInfo) MsgDetailedInfoClazzName() string {
	return ClazzName_msg_detailed_info
}

func (m *TLMsgDetailedInfo) ClazzName() string {
	return ClazzName_msg_detailed_info
}

func (m *TLMsgDetailedInfo) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x276d3ec6: func() error {
			x.PutClazzID(0x276d3ec6)

			x.PutInt64(m.MsgId)
			x.PutInt64(m.AnswerMsgId)
			x.PutInt32(m.Bytes)
			x.PutInt32(m.Status)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_detailed_info, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_detailed_info, layer)
	}
}

func (m *TLMsgDetailedInfo) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x276d3ec6: func() (err error) {
			m.MsgId, err = d.Int64()
			m.AnswerMsgId, err = d.Int64()
			m.Bytes, err = d.Int32()
			m.Status, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLMsgNewDetailedInfo <--
type TLMsgNewDetailedInfo struct {
	ClazzID     uint32 `json:"_id"`
	AnswerMsgId int64  `json:"answer_msg_id"`
	Bytes       int32  `json:"bytes"`
	Status      int32  `json:"status"`
}

func (m *TLMsgNewDetailedInfo) MsgDetailedInfoClazzName() string {
	return ClazzName_msg_new_detailed_info
}

func (m *TLMsgNewDetailedInfo) ClazzName() string {
	return ClazzName_msg_new_detailed_info
}

func (m *TLMsgNewDetailedInfo) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x809db6df: func() error {
			x.PutClazzID(0x809db6df)

			x.PutInt64(m.AnswerMsgId)
			x.PutInt32(m.Bytes)
			x.PutInt32(m.Status)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_new_detailed_info, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_new_detailed_info, layer)
	}
}

func (m *TLMsgNewDetailedInfo) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x809db6df: func() (err error) {
			m.AnswerMsgId, err = d.Int64()
			m.Bytes, err = d.Int32()
			m.Status, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgDetailedInfo <--
type MsgDetailedInfo struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	MsgDetailedInfoClazz
}

func MakeMsgDetailedInfo(c MsgDetailedInfoClazz) *MsgDetailedInfo {
	return &MsgDetailedInfo{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		MsgDetailedInfoClazz: c,
	}
}

func (m *MsgDetailedInfo) Encode(x *bin.Encoder, layer int32) error {
	if m.MsgDetailedInfoClazz != nil {
		return m.MsgDetailedInfoClazz.Encode(x, layer)
	}

	return fmt.Errorf("MsgDetailedInfo - invalid Clazz")
}

func (m *MsgDetailedInfo) Decode(d *bin.Decoder) (err error) {
	m.MsgDetailedInfoClazz, err = DecodeMsgDetailedInfoClazz(d)
	return
}

func (m *MsgDetailedInfo) Match(f ...interface{}) {
	switch c := m.MsgDetailedInfoClazz.(type) {
	case *TLMsgDetailedInfo:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgDetailedInfo) interface{}); ok {
				f1(c)
			}
		}
	case *TLMsgNewDetailedInfo:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgNewDetailedInfo) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// MsgResendReqClazz <--
//   - TL_MsgResendReq
type MsgResendReqClazz interface {
	iface.TLObject
	MsgResendReqClazzName() string
}

func DecodeMsgResendReqClazz(d *bin.Decoder) (MsgResendReqClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_msg_resend_req:
		x := &TLMsgResendReq{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgResendReq - unexpected clazzId: %d", id)
	}
}

// TLMsgResendReq <--
type TLMsgResendReq struct {
	ClazzID uint32  `json:"_id"`
	MsgIds  []int64 `json:"msg_ids"`
}

func (m *TLMsgResendReq) MsgResendReqClazzName() string {
	return ClazzName_msg_resend_req
}

func (m *TLMsgResendReq) ClazzName() string {
	return ClazzName_msg_resend_req
}

func (m *TLMsgResendReq) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x7d861a08: func() error {
			x.PutClazzID(0x7d861a08)

			iface.EncodeInt64List(x, m.MsgIds)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msg_resend_req, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msg_resend_req, layer)
	}
}

func (m *TLMsgResendReq) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x7d861a08: func() (err error) {

			m.MsgIds, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgResendReq <--
type MsgResendReq struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	MsgResendReqClazz
}

func MakeMsgResendReq(c MsgResendReqClazz) *MsgResendReq {
	return &MsgResendReq{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		MsgResendReqClazz: c,
	}
}

func (m *MsgResendReq) Encode(x *bin.Encoder, layer int32) error {
	if m.MsgResendReqClazz != nil {
		return m.MsgResendReqClazz.Encode(x, layer)
	}

	return fmt.Errorf("MsgResendReq - invalid Clazz")
}

func (m *MsgResendReq) Decode(d *bin.Decoder) (err error) {
	m.MsgResendReqClazz, err = DecodeMsgResendReqClazz(d)
	return
}

func (m *MsgResendReq) Match(f ...interface{}) {
	switch c := m.MsgResendReqClazz.(type) {
	case *TLMsgResendReq:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgResendReq) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// MsgsAckClazz <--
//   - TL_MsgsAck
type MsgsAckClazz interface {
	iface.TLObject
	MsgsAckClazzName() string
}

func DecodeMsgsAckClazz(d *bin.Decoder) (MsgsAckClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_msgs_ack:
		x := &TLMsgsAck{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsAck - unexpected clazzId: %d", id)
	}
}

// TLMsgsAck <--
type TLMsgsAck struct {
	ClazzID uint32  `json:"_id"`
	MsgIds  []int64 `json:"msg_ids"`
}

func (m *TLMsgsAck) MsgsAckClazzName() string {
	return ClazzName_msgs_ack
}

func (m *TLMsgsAck) ClazzName() string {
	return ClazzName_msgs_ack
}

func (m *TLMsgsAck) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x62d6b459: func() error {
			x.PutClazzID(0x62d6b459)

			iface.EncodeInt64List(x, m.MsgIds)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msgs_ack, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_ack, layer)
	}
}

func (m *TLMsgsAck) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x62d6b459: func() (err error) {

			m.MsgIds, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsAck <--
type MsgsAck struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	MsgsAckClazz
}

func MakeMsgsAck(c MsgsAckClazz) *MsgsAck {
	return &MsgsAck{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		MsgsAckClazz: c,
	}
}

func (m *MsgsAck) Encode(x *bin.Encoder, layer int32) error {
	if m.MsgsAckClazz != nil {
		return m.MsgsAckClazz.Encode(x, layer)
	}

	return fmt.Errorf("MsgsAck - invalid Clazz")
}

func (m *MsgsAck) Decode(d *bin.Decoder) (err error) {
	m.MsgsAckClazz, err = DecodeMsgsAckClazz(d)
	return
}

func (m *MsgsAck) Match(f ...interface{}) {
	switch c := m.MsgsAckClazz.(type) {
	case *TLMsgsAck:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgsAck) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// MsgsAllInfoClazz <--
//   - TL_MsgsAllInfo
type MsgsAllInfoClazz interface {
	iface.TLObject
	MsgsAllInfoClazzName() string
}

func DecodeMsgsAllInfoClazz(d *bin.Decoder) (MsgsAllInfoClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_msgs_all_info:
		x := &TLMsgsAllInfo{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsAllInfo - unexpected clazzId: %d", id)
	}
}

// TLMsgsAllInfo <--
type TLMsgsAllInfo struct {
	ClazzID uint32  `json:"_id"`
	MsgIds  []int64 `json:"msg_ids"`
	Info    string  `json:"info"`
}

func (m *TLMsgsAllInfo) MsgsAllInfoClazzName() string {
	return ClazzName_msgs_all_info
}

func (m *TLMsgsAllInfo) ClazzName() string {
	return ClazzName_msgs_all_info
}

func (m *TLMsgsAllInfo) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x8cc0d131: func() error {
			x.PutClazzID(0x8cc0d131)

			iface.EncodeInt64List(x, m.MsgIds)

			x.PutString(m.Info)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msgs_all_info, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_all_info, layer)
	}
}

func (m *TLMsgsAllInfo) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x8cc0d131: func() (err error) {

			m.MsgIds, err = iface.DecodeInt64List(d)

			m.Info, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsAllInfo <--
type MsgsAllInfo struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	MsgsAllInfoClazz
}

func MakeMsgsAllInfo(c MsgsAllInfoClazz) *MsgsAllInfo {
	return &MsgsAllInfo{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		MsgsAllInfoClazz: c,
	}
}

func (m *MsgsAllInfo) Encode(x *bin.Encoder, layer int32) error {
	if m.MsgsAllInfoClazz != nil {
		return m.MsgsAllInfoClazz.Encode(x, layer)
	}

	return fmt.Errorf("MsgsAllInfo - invalid Clazz")
}

func (m *MsgsAllInfo) Decode(d *bin.Decoder) (err error) {
	m.MsgsAllInfoClazz, err = DecodeMsgsAllInfoClazz(d)
	return
}

func (m *MsgsAllInfo) Match(f ...interface{}) {
	switch c := m.MsgsAllInfoClazz.(type) {
	case *TLMsgsAllInfo:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgsAllInfo) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// MsgsStateInfoClazz <--
//   - TL_MsgsStateInfo
type MsgsStateInfoClazz interface {
	iface.TLObject
	MsgsStateInfoClazzName() string
}

func DecodeMsgsStateInfoClazz(d *bin.Decoder) (MsgsStateInfoClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_msgs_state_info:
		x := &TLMsgsStateInfo{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsStateInfo - unexpected clazzId: %d", id)
	}
}

// TLMsgsStateInfo <--
type TLMsgsStateInfo struct {
	ClazzID  uint32 `json:"_id"`
	ReqMsgId int64  `json:"req_msg_id"`
	Info     string `json:"info"`
}

func (m *TLMsgsStateInfo) MsgsStateInfoClazzName() string {
	return ClazzName_msgs_state_info
}

func (m *TLMsgsStateInfo) ClazzName() string {
	return ClazzName_msgs_state_info
}

func (m *TLMsgsStateInfo) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4deb57d: func() error {
			x.PutClazzID(0x4deb57d)

			x.PutInt64(m.ReqMsgId)
			x.PutString(m.Info)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msgs_state_info, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_state_info, layer)
	}
}

func (m *TLMsgsStateInfo) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4deb57d: func() (err error) {
			m.ReqMsgId, err = d.Int64()
			m.Info, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsStateInfo <--
type MsgsStateInfo struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	MsgsStateInfoClazz
}

func MakeMsgsStateInfo(c MsgsStateInfoClazz) *MsgsStateInfo {
	return &MsgsStateInfo{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		MsgsStateInfoClazz: c,
	}
}

func (m *MsgsStateInfo) Encode(x *bin.Encoder, layer int32) error {
	if m.MsgsStateInfoClazz != nil {
		return m.MsgsStateInfoClazz.Encode(x, layer)
	}

	return fmt.Errorf("MsgsStateInfo - invalid Clazz")
}

func (m *MsgsStateInfo) Decode(d *bin.Decoder) (err error) {
	m.MsgsStateInfoClazz, err = DecodeMsgsStateInfoClazz(d)
	return
}

func (m *MsgsStateInfo) Match(f ...interface{}) {
	switch c := m.MsgsStateInfoClazz.(type) {
	case *TLMsgsStateInfo:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgsStateInfo) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// MsgsStateReqClazz <--
//   - TL_MsgsStateReq
type MsgsStateReqClazz interface {
	iface.TLObject
	MsgsStateReqClazzName() string
}

func DecodeMsgsStateReqClazz(d *bin.Decoder) (MsgsStateReqClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_msgs_state_req:
		x := &TLMsgsStateReq{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeMsgsStateReq - unexpected clazzId: %d", id)
	}
}

// TLMsgsStateReq <--
type TLMsgsStateReq struct {
	ClazzID uint32  `json:"_id"`
	MsgIds  []int64 `json:"msg_ids"`
}

func (m *TLMsgsStateReq) MsgsStateReqClazzName() string {
	return ClazzName_msgs_state_req
}

func (m *TLMsgsStateReq) ClazzName() string {
	return ClazzName_msgs_state_req
}

func (m *TLMsgsStateReq) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xda69fb52: func() error {
			x.PutClazzID(0xda69fb52)

			iface.EncodeInt64List(x, m.MsgIds)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_msgs_state_req, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_msgs_state_req, layer)
	}
}

func (m *TLMsgsStateReq) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xda69fb52: func() (err error) {

			m.MsgIds, err = iface.DecodeInt64List(d)

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// MsgsStateReq <--
type MsgsStateReq struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	MsgsStateReqClazz
}

func MakeMsgsStateReq(c MsgsStateReqClazz) *MsgsStateReq {
	return &MsgsStateReq{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		MsgsStateReqClazz: c,
	}
}

func (m *MsgsStateReq) Encode(x *bin.Encoder, layer int32) error {
	if m.MsgsStateReqClazz != nil {
		return m.MsgsStateReqClazz.Encode(x, layer)
	}

	return fmt.Errorf("MsgsStateReq - invalid Clazz")
}

func (m *MsgsStateReq) Decode(d *bin.Decoder) (err error) {
	m.MsgsStateReqClazz, err = DecodeMsgsStateReqClazz(d)
	return
}

func (m *MsgsStateReq) Match(f ...interface{}) {
	switch c := m.MsgsStateReqClazz.(type) {
	case *TLMsgsStateReq:
		for _, v := range f {
			if f1, ok := v.(func(c *TLMsgsStateReq) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// NewSessionClazz <--
//   - TL_NewSessionCreated
type NewSessionClazz interface {
	iface.TLObject
	NewSessionClazzName() string
}

func DecodeNewSessionClazz(d *bin.Decoder) (NewSessionClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_new_session_created:
		x := &TLNewSessionCreated{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeNewSession - unexpected clazzId: %d", id)
	}
}

// TLNewSessionCreated <--
type TLNewSessionCreated struct {
	ClazzID    uint32 `json:"_id"`
	FirstMsgId int64  `json:"first_msg_id"`
	UniqueId   int64  `json:"unique_id"`
	ServerSalt int64  `json:"server_salt"`
}

func (m *TLNewSessionCreated) NewSessionClazzName() string {
	return ClazzName_new_session_created
}

func (m *TLNewSessionCreated) ClazzName() string {
	return ClazzName_new_session_created
}

func (m *TLNewSessionCreated) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9ec20908: func() error {
			x.PutClazzID(0x9ec20908)

			x.PutInt64(m.FirstMsgId)
			x.PutInt64(m.UniqueId)
			x.PutInt64(m.ServerSalt)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_new_session_created, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_new_session_created, layer)
	}
}

func (m *TLNewSessionCreated) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9ec20908: func() (err error) {
			m.FirstMsgId, err = d.Int64()
			m.UniqueId, err = d.Int64()
			m.ServerSalt, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// NewSession <--
type NewSession struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	NewSessionClazz
}

func MakeNewSession(c NewSessionClazz) *NewSession {
	return &NewSession{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		NewSessionClazz: c,
	}
}

func (m *NewSession) Encode(x *bin.Encoder, layer int32) error {
	if m.NewSessionClazz != nil {
		return m.NewSessionClazz.Encode(x, layer)
	}

	return fmt.Errorf("NewSession - invalid Clazz")
}

func (m *NewSession) Decode(d *bin.Decoder) (err error) {
	m.NewSessionClazz, err = DecodeNewSessionClazz(d)
	return
}

func (m *NewSession) Match(f ...interface{}) {
	switch c := m.NewSessionClazz.(type) {
	case *TLNewSessionCreated:
		for _, v := range f {
			if f1, ok := v.(func(c *TLNewSessionCreated) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// PongClazz <--
//   - TL_Pong
type PongClazz interface {
	iface.TLObject
	PongClazzName() string
}

func DecodePongClazz(d *bin.Decoder) (PongClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_pong:
		x := &TLPong{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodePong - unexpected clazzId: %d", id)
	}
}

// TLPong <--
type TLPong struct {
	ClazzID uint32 `json:"_id"`
	MsgId   int64  `json:"msg_id"`
	PingId  int64  `json:"ping_id"`
}

func (m *TLPong) PongClazzName() string {
	return ClazzName_pong
}

func (m *TLPong) ClazzName() string {
	return ClazzName_pong
}

func (m *TLPong) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x347773c5: func() error {
			x.PutClazzID(0x347773c5)

			x.PutInt64(m.MsgId)
			x.PutInt64(m.PingId)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_pong, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_pong, layer)
	}
}

func (m *TLPong) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x347773c5: func() (err error) {
			m.MsgId, err = d.Int64()
			m.PingId, err = d.Int64()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// Pong <--
type Pong struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	PongClazz
}

func MakePong(c PongClazz) *Pong {
	return &Pong{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		PongClazz: c,
	}
}

func (m *Pong) Encode(x *bin.Encoder, layer int32) error {
	if m.PongClazz != nil {
		return m.PongClazz.Encode(x, layer)
	}

	return fmt.Errorf("Pong - invalid Clazz")
}

func (m *Pong) Decode(d *bin.Decoder) (err error) {
	m.PongClazz, err = DecodePongClazz(d)
	return
}

func (m *Pong) Match(f ...interface{}) {
	switch c := m.PongClazz.(type) {
	case *TLPong:
		for _, v := range f {
			if f1, ok := v.(func(c *TLPong) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// RpcDropAnswerClazz <--
//   - TL_RpcAnswerUnknown
//   - TL_RpcAnswerDroppedRunning
//   - TL_RpcAnswerDropped
type RpcDropAnswerClazz interface {
	iface.TLObject
	RpcDropAnswerClazzName() string
}

func DecodeRpcDropAnswerClazz(d *bin.Decoder) (RpcDropAnswerClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_rpc_answer_unknown:
		x := &TLRpcAnswerUnknown{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_rpc_answer_dropped_running:
		x := &TLRpcAnswerDroppedRunning{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_rpc_answer_dropped:
		x := &TLRpcAnswerDropped{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeRpcDropAnswer - unexpected clazzId: %d", id)
	}
}

// TLRpcAnswerUnknown <--
type TLRpcAnswerUnknown struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLRpcAnswerUnknown) RpcDropAnswerClazzName() string {
	return ClazzName_rpc_answer_unknown
}

func (m *TLRpcAnswerUnknown) ClazzName() string {
	return ClazzName_rpc_answer_unknown
}

func (m *TLRpcAnswerUnknown) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x5e2ad36e: func() error {
			x.PutClazzID(0x5e2ad36e)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_rpc_answer_unknown, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_answer_unknown, layer)
	}
}

func (m *TLRpcAnswerUnknown) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x5e2ad36e: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLRpcAnswerDroppedRunning <--
type TLRpcAnswerDroppedRunning struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLRpcAnswerDroppedRunning) RpcDropAnswerClazzName() string {
	return ClazzName_rpc_answer_dropped_running
}

func (m *TLRpcAnswerDroppedRunning) ClazzName() string {
	return ClazzName_rpc_answer_dropped_running
}

func (m *TLRpcAnswerDroppedRunning) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xcd78e586: func() error {
			x.PutClazzID(0xcd78e586)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_rpc_answer_dropped_running, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_answer_dropped_running, layer)
	}
}

func (m *TLRpcAnswerDroppedRunning) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xcd78e586: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLRpcAnswerDropped <--
type TLRpcAnswerDropped struct {
	ClazzID uint32 `json:"_id"`
	MsgId   int64  `json:"msg_id"`
	SeqNo   int32  `json:"seq_no"`
	Bytes   int32  `json:"bytes"`
}

func (m *TLRpcAnswerDropped) RpcDropAnswerClazzName() string {
	return ClazzName_rpc_answer_dropped
}

func (m *TLRpcAnswerDropped) ClazzName() string {
	return ClazzName_rpc_answer_dropped
}

func (m *TLRpcAnswerDropped) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xa43ad8b7: func() error {
			x.PutClazzID(0xa43ad8b7)

			x.PutInt64(m.MsgId)
			x.PutInt32(m.SeqNo)
			x.PutInt32(m.Bytes)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_rpc_answer_dropped, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_answer_dropped, layer)
	}
}

func (m *TLRpcAnswerDropped) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xa43ad8b7: func() (err error) {
			m.MsgId, err = d.Int64()
			m.SeqNo, err = d.Int32()
			m.Bytes, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// RpcDropAnswer <--
type RpcDropAnswer struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	RpcDropAnswerClazz
}

func MakeRpcDropAnswer(c RpcDropAnswerClazz) *RpcDropAnswer {
	return &RpcDropAnswer{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		RpcDropAnswerClazz: c,
	}
}

func (m *RpcDropAnswer) Encode(x *bin.Encoder, layer int32) error {
	if m.RpcDropAnswerClazz != nil {
		return m.RpcDropAnswerClazz.Encode(x, layer)
	}

	return fmt.Errorf("RpcDropAnswer - invalid Clazz")
}

func (m *RpcDropAnswer) Decode(d *bin.Decoder) (err error) {
	m.RpcDropAnswerClazz, err = DecodeRpcDropAnswerClazz(d)
	return
}

func (m *RpcDropAnswer) Match(f ...interface{}) {
	switch c := m.RpcDropAnswerClazz.(type) {
	case *TLRpcAnswerUnknown:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRpcAnswerUnknown) interface{}); ok {
				f1(c)
			}
		}
	case *TLRpcAnswerDroppedRunning:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRpcAnswerDroppedRunning) interface{}); ok {
				f1(c)
			}
		}
	case *TLRpcAnswerDropped:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRpcAnswerDropped) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// RpcErrorClazz <--
//   - TL_RpcError
type RpcErrorClazz interface {
	iface.TLObject
	RpcErrorClazzName() string
}

func DecodeRpcErrorClazz(d *bin.Decoder) (RpcErrorClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_rpc_error:
		x := &TLRpcError{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeRpcError - unexpected clazzId: %d", id)
	}
}

// TLRpcError <--
type TLRpcError struct {
	ClazzID      uint32 `json:"_id"`
	ErrorCode    int32  `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

func (m *TLRpcError) RpcErrorClazzName() string {
	return ClazzName_rpc_error
}

func (m *TLRpcError) ClazzName() string {
	return ClazzName_rpc_error
}

func (m *TLRpcError) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x2144ca19: func() error {
			x.PutClazzID(0x2144ca19)

			x.PutInt32(m.ErrorCode)
			x.PutString(m.ErrorMessage)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_rpc_error, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_rpc_error, layer)
	}
}

func (m *TLRpcError) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x2144ca19: func() (err error) {
			m.ErrorCode, err = d.Int32()
			m.ErrorMessage, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// RpcError <--
type RpcError struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	RpcErrorClazz
}

func MakeRpcError(c RpcErrorClazz) *RpcError {
	return &RpcError{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		RpcErrorClazz: c,
	}
}

func (m *RpcError) Encode(x *bin.Encoder, layer int32) error {
	if m.RpcErrorClazz != nil {
		return m.RpcErrorClazz.Encode(x, layer)
	}

	return fmt.Errorf("RpcError - invalid Clazz")
}

func (m *RpcError) Decode(d *bin.Decoder) (err error) {
	m.RpcErrorClazz, err = DecodeRpcErrorClazz(d)
	return
}

func (m *RpcError) Match(f ...interface{}) {
	switch c := m.RpcErrorClazz.(type) {
	case *TLRpcError:
		for _, v := range f {
			if f1, ok := v.(func(c *TLRpcError) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// TlsBlockClazz <--
//   - TL_TlsBlockString
//   - TL_TlsBlockRandom
//   - TL_TlsBlockZero
//   - TL_TlsBlockDomain
//   - TL_TlsBlockGrease
//   - TL_TlsBlockPublicKey
//   - TL_TlsBlockScope
type TlsBlockClazz interface {
	iface.TLObject
	TlsBlockClazzName() string
}

func DecodeTlsBlockClazz(d *bin.Decoder) (TlsBlockClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_tlsBlockString:
		x := &TLTlsBlockString{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_tlsBlockRandom:
		x := &TLTlsBlockRandom{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_tlsBlockZero:
		x := &TLTlsBlockZero{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_tlsBlockDomain:
		x := &TLTlsBlockDomain{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_tlsBlockGrease:
		x := &TLTlsBlockGrease{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_tlsBlockPublicKey:
		x := &TLTlsBlockPublicKey{}
		_ = x.Decode(d)
		return x, nil
	case ClazzName_tlsBlockScope:
		x := &TLTlsBlockScope{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeTlsBlock - unexpected clazzId: %d", id)
	}
}

// TLTlsBlockString <--
type TLTlsBlockString struct {
	ClazzID uint32 `json:"_id"`
	Data    string `json:"data"`
}

func (m *TLTlsBlockString) TlsBlockClazzName() string {
	return ClazzName_tlsBlockString
}

func (m *TLTlsBlockString) ClazzName() string {
	return ClazzName_tlsBlockString
}

func (m *TLTlsBlockString) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4218a164: func() error {
			x.PutClazzID(0x4218a164)

			x.PutString(m.Data)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockString, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockString, layer)
	}
}

func (m *TLTlsBlockString) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4218a164: func() (err error) {
			m.Data, err = d.String()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockRandom <--
type TLTlsBlockRandom struct {
	ClazzID uint32 `json:"_id"`
	Length  int32  `json:"length"`
}

func (m *TLTlsBlockRandom) TlsBlockClazzName() string {
	return ClazzName_tlsBlockRandom
}

func (m *TLTlsBlockRandom) ClazzName() string {
	return ClazzName_tlsBlockRandom
}

func (m *TLTlsBlockRandom) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x4d4dc41e: func() error {
			x.PutClazzID(0x4d4dc41e)

			x.PutInt32(m.Length)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockRandom, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockRandom, layer)
	}
}

func (m *TLTlsBlockRandom) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x4d4dc41e: func() (err error) {
			m.Length, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockZero <--
type TLTlsBlockZero struct {
	ClazzID uint32 `json:"_id"`
	Length  int32  `json:"length"`
}

func (m *TLTlsBlockZero) TlsBlockClazzName() string {
	return ClazzName_tlsBlockZero
}

func (m *TLTlsBlockZero) ClazzName() string {
	return ClazzName_tlsBlockZero
}

func (m *TLTlsBlockZero) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9333afb: func() error {
			x.PutClazzID(0x9333afb)

			x.PutInt32(m.Length)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockZero, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockZero, layer)
	}
}

func (m *TLTlsBlockZero) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9333afb: func() (err error) {
			m.Length, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockDomain <--
type TLTlsBlockDomain struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLTlsBlockDomain) TlsBlockClazzName() string {
	return ClazzName_tlsBlockDomain
}

func (m *TLTlsBlockDomain) ClazzName() string {
	return ClazzName_tlsBlockDomain
}

func (m *TLTlsBlockDomain) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x10e8636f: func() error {
			x.PutClazzID(0x10e8636f)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockDomain, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockDomain, layer)
	}
}

func (m *TLTlsBlockDomain) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x10e8636f: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockGrease <--
type TLTlsBlockGrease struct {
	ClazzID uint32 `json:"_id"`
	Seed    int32  `json:"seed"`
}

func (m *TLTlsBlockGrease) TlsBlockClazzName() string {
	return ClazzName_tlsBlockGrease
}

func (m *TLTlsBlockGrease) ClazzName() string {
	return ClazzName_tlsBlockGrease
}

func (m *TLTlsBlockGrease) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe675a1c1: func() error {
			x.PutClazzID(0xe675a1c1)

			x.PutInt32(m.Seed)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockGrease, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockGrease, layer)
	}
}

func (m *TLTlsBlockGrease) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe675a1c1: func() (err error) {
			m.Seed, err = d.Int32()

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockPublicKey <--
type TLTlsBlockPublicKey struct {
	ClazzID uint32 `json:"_id"`
}

func (m *TLTlsBlockPublicKey) TlsBlockClazzName() string {
	return ClazzName_tlsBlockPublicKey
}

func (m *TLTlsBlockPublicKey) ClazzName() string {
	return ClazzName_tlsBlockPublicKey
}

func (m *TLTlsBlockPublicKey) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x9eb95b5c: func() error {
			x.PutClazzID(0x9eb95b5c)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockPublicKey, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockPublicKey, layer)
	}
}

func (m *TLTlsBlockPublicKey) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x9eb95b5c: func() (err error) {

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TLTlsBlockScope <--
type TLTlsBlockScope struct {
	ClazzID uint32      `json:"_id"`
	Entries []*TlsBlock `json:"entries"`
}

func (m *TLTlsBlockScope) TlsBlockClazzName() string {
	return ClazzName_tlsBlockScope
}

func (m *TLTlsBlockScope) ClazzName() string {
	return ClazzName_tlsBlockScope
}

func (m *TLTlsBlockScope) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0xe725d44f: func() error {
			x.PutClazzID(0xe725d44f)

			_ = iface.EncodeObjectList(x, m.Entries, layer)

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsBlockScope, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsBlockScope, layer)
	}
}

func (m *TLTlsBlockScope) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0xe725d44f: func() (err error) {
			c3, err2 := d.ClazzID()
			if c3 != iface.ClazzID_vector {
				// dBuf.err = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 3, c3)
				return err2
			}
			l3, err3 := d.Int()
			v3 := make([]*TlsBlock, l3)
			for i := 0; i < l3; i++ {
				vv := new(TlsBlock)
				err3 = vv.Decode(d)
				_ = err3
				v3[i] = vv
			}
			m.Entries = v3

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TlsBlock <--
type TlsBlock struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	TlsBlockClazz
}

func MakeTlsBlock(c TlsBlockClazz) *TlsBlock {
	return &TlsBlock{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		TlsBlockClazz: c,
	}
}

func (m *TlsBlock) Encode(x *bin.Encoder, layer int32) error {
	if m.TlsBlockClazz != nil {
		return m.TlsBlockClazz.Encode(x, layer)
	}

	return fmt.Errorf("TlsBlock - invalid Clazz")
}

func (m *TlsBlock) Decode(d *bin.Decoder) (err error) {
	m.TlsBlockClazz, err = DecodeTlsBlockClazz(d)
	return
}

func (m *TlsBlock) Match(f ...interface{}) {
	switch c := m.TlsBlockClazz.(type) {
	case *TLTlsBlockString:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockString) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockRandom:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockRandom) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockZero:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockZero) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockDomain:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockDomain) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockGrease:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockGrease) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockPublicKey:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockPublicKey) interface{}); ok {
				f1(c)
			}
		}
	case *TLTlsBlockScope:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsBlockScope) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}

// TlsClientHelloClazz <--
//   - TL_TlsClientHello
type TlsClientHelloClazz interface {
	iface.TLObject
	TlsClientHelloClazzName() string
}

func DecodeTlsClientHelloClazz(d *bin.Decoder) (TlsClientHelloClazz, error) {
	id, err := d.PeekClazzID()
	if err != nil {
		return nil, err
	}

	clazzName := iface.GetClazzNameByID(id)
	switch clazzName {
	case ClazzName_tlsClientHello:
		x := &TLTlsClientHello{}
		_ = x.Decode(d)
		return x, nil
	default:
		return nil, fmt.Errorf("DecodeTlsClientHello - unexpected clazzId: %d", id)
	}
}

// TLTlsClientHello <--
type TLTlsClientHello struct {
	ClazzID uint32      `json:"_id"`
	Blocks  []*TlsBlock `json:"blocks"`
}

func (m *TLTlsClientHello) TlsClientHelloClazzName() string {
	return ClazzName_tlsClientHello
}

func (m *TLTlsClientHello) ClazzName() string {
	return ClazzName_tlsClientHello
}

func (m *TLTlsClientHello) Encode(x *bin.Encoder, layer int32) error {
	var encodeF = map[uint32]func() error{
		0x6c52c484: func() error {
			x.PutClazzID(0x6c52c484)

			// x.PutClazzID(iface.ClazzID_vector)
			x.PutInt(len(m.Blocks))
			for _, v := range m.Blocks {
				_ = v.Encode(x, layer)
			}

			return nil
		},
	}

	clazzId := iface.GetClazzIDByName(ClazzName_tlsClientHello, int(layer))
	if f, ok := encodeF[clazzId]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		return fmt.Errorf("not found clazzId by (%s, %d)", ClazzName_tlsClientHello, layer)
	}
}

func (m *TLTlsClientHello) Decode(d *bin.Decoder) (err error) {
	var decodeF = map[uint32]func() error{
		0x6c52c484: func() (err error) {
			// c0, err2 := d.ClazzID()
			// if c0 != int32(iface.ClazzID_vector) {
			//     err2 = fmt.Errorf("invalid ClazzID_vector, c%d: %d", 0, c0)
			//     return err2
			// }
			l0, err2 := d.Int()
			_ = err2
			v0 := make([]*TlsBlock, l0)
			for i := 0; i < l0; i++ {
				vv := new(TlsBlock)
				err2 = vv.Decode(d)
				_ = err2
				v0[i] = vv
			}
			m.Blocks = v0

			return nil
		},
	}

	if f, ok := decodeF[m.ClazzID]; ok {
		return f()
	} else {
		return fmt.Errorf("invalid constructor: %x", m.ClazzID)
	}
}

// TlsClientHello <--
type TlsClientHello struct {
	ClazzID   uint32 `json:"_id"`
	ClazzName string `json:"_name"`
	TlsClientHelloClazz
}

func MakeTlsClientHello(c TlsClientHelloClazz) *TlsClientHello {
	return &TlsClientHello{
		// ClazzID:   c.ClazzID(),
		// ClazzName: c.ClazzName(),
		TlsClientHelloClazz: c,
	}
}

func (m *TlsClientHello) Encode(x *bin.Encoder, layer int32) error {
	if m.TlsClientHelloClazz != nil {
		return m.TlsClientHelloClazz.Encode(x, layer)
	}

	return fmt.Errorf("TlsClientHello - invalid Clazz")
}

func (m *TlsClientHello) Decode(d *bin.Decoder) (err error) {
	m.TlsClientHelloClazz, err = DecodeTlsClientHelloClazz(d)
	return
}

func (m *TlsClientHello) Match(f ...interface{}) {
	switch c := m.TlsClientHelloClazz.(type) {
	case *TLTlsClientHello:
		for _, v := range f {
			if f1, ok := v.(func(c *TLTlsClientHello) interface{}); ok {
				f1(c)
			}
		}
	default:
		//
	}
}
