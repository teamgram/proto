// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mtproto

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"sync"
)

// TLMessageRawData
// message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessageRawData struct {
	MsgId   int64
	Seqno   int32
	Bytes   int32
	ClassId int32
	Body    []byte
}

func (m *TLMessageRawData) String() string {
	return fmt.Sprintf("{message2#5bb8e511 - msg_id: %d, seq_no: %d, bytes: %d, class_id: %d}",
		m.MsgId,
		m.Seqno,
		m.Bytes,
		m.ClassId)
}

func (m *TLMessageRawData) Encode(x *EncodeBuf, layer int32) error {
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(m.Body)

	return nil
}

func (m *TLMessageRawData) Decode(dBuf *DecodeBuf) error {
	m.MsgId = dBuf.Long()
	m.Seqno = dBuf.Int()
	m.Bytes = dBuf.Int()
	b := dBuf.Bytes(int(m.Bytes))
	if dBuf.err != nil {
		return dBuf.err
	}
	if len(b) < 4 {
		return fmt.Errorf("decode error")
	}
	m.ClassId = int32(binary.LittleEndian.Uint32(b))
	if !CheckClassID(m.ClassId) {
		return fmt.Errorf("not register classId: %d", m.ClassId)
	}
	m.Body = b
	return nil
}

func (m *TLMessageRawData) DebugString() string {
	return fmt.Sprintf(`{"message2#5bb8e511": {"msg_id": %d, "seq_no": %d, "bytes": %d, "class_id": "0x%x"}}`,
		m.MsgId,
		m.Seqno,
		m.Bytes,
		uint32(m.ClassId))
}

// TLMsgRawDataContainer
// msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
type TLMsgRawDataContainer struct {
	Messages []*TLMessageRawData
}

// gzipBufPool 用于 TLGzipPacked Encode/Decode 复用 bytes.Buffer，减少分配
var gzipBufPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

func NewTLMsgRawDataContainer() *TLMsgRawDataContainer {
	return &TLMsgRawDataContainer{
		Messages: nil, // Decode 时会预分配，无需分配空 slice
	}
}

func (m *TLMsgRawDataContainer) String() string {
	return "{msg_container#73f1f8dc}"
}

func (m *TLMsgRawDataContainer) Encode(x *EncodeBuf, layer int32) error {
	x.Int(int32(CRC32_msg_container))

	x.Int(int32(len(m.Messages)))
	for _, v := range m.Messages {
		v.Encode(x, layer)
	}
	return nil
}

func (m *TLMsgRawDataContainer) Decode(dbuf *DecodeBuf) error {
	n := dbuf.Int()
	if n > 0 {
		m.Messages = make([]*TLMessageRawData, 0, int(n))
	}
	for i := 0; i < int(n); i++ {
		message2 := &TLMessageRawData{}
		err := message2.Decode(dbuf)
		if err != nil {
			return err
		}
		m.Messages = append(m.Messages, message2)
	}
	return dbuf.err
}

func (m *TLMsgRawDataContainer) DebugString() string {
	return fmt.Sprintf(`{"msg_container#73f1f8dc": []}`)
}

// TLMessage2
// message2 msg_id:long seqno:int bytes:int body:Object = Message2; // parsed manually
type TLMessage2 struct {
	MsgId  int64
	Seqno  int32
	Bytes  int32
	Object TLObject
}

func (m *TLMessage2) String() string {
	return fmt.Sprintf("{message2#5bb8e511 - msg_id: %d, seq_no: %d, object: {%v}}",
		m.MsgId,
		m.Seqno,
		m.Object)
}

func (m *TLMessage2) Encode(x *EncodeBuf, layer int32) error {
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	offset := x.GetOffset()
	x.Int(m.Bytes)
	m.Object.Encode(x, layer)
	x.IntOffset(offset, int32(x.GetOffset()-offset-4))

	return nil
}

func (m *TLMessage2) Decode(dBuf *DecodeBuf) error {
	m.MsgId = dBuf.Long()
	m.Seqno = dBuf.Int()
	m.Bytes = dBuf.Int()
	// log.Debugf("message2: {msg_id: %d, seqno: %d, bytes: %d}", m.MsgId, m.Seqno, m.Bytes)
	b := dBuf.Bytes(int(m.Bytes))

	dBuf2 := NewDecodeBuf(b)
	m.Object = dBuf2.Object()
	if m.Object == nil {
		hexLen := len(b)
		if hexLen > 128 {
			hexLen = 128
		}
		return fmt.Errorf("decode core_message error(%v): %s", dBuf2.err, hex.EncodeToString(b[:hexLen]))
	}

	// log.Info("Sucess decoded core_message: ", m.Object.String())
	return dBuf2.err
}

func (m *TLMessage2) DebugString() string {
	return fmt.Sprintf(`{"message2#5bb8e511": {"msg_id": %d, "seq_no": %d, "object": "%s"}`,
		m.MsgId,
		m.Seqno,
		reflect.TypeOf(m.Object))
}

// TLMsgContainer
// msg_container#73f1f8dc messages:vector<message2> = MessageContainer; // parsed manually
type TLMsgContainer struct {
	Messages []*TLMessage2
}

func (m *TLMsgContainer) String() string {
	return "{msg_container#73f1f8dc}"
}

func (m *TLMsgContainer) Encode(x *EncodeBuf, layer int32) error {
	// x := NewEncodeBuf(512)
	x.Int(int32(CRC32_msg_container))

	x.Int(int32(len(m.Messages)))
	for _, v := range m.Messages {
		v.Encode(x, layer)
	}
	return nil
}

func (m *TLMsgContainer) Decode(dbuf *DecodeBuf) error {
	n := dbuf.Int()
	if n > 0 {
		m.Messages = make([]*TLMessage2, 0, int(n))
	}
	for i := 0; i < int(n); i++ {
		message2 := new(TLMessage2)
		err := message2.Decode(dbuf)
		if err != nil {
			return err
		}
		m.Messages = append(m.Messages, message2)
	}
	return dbuf.err
}

func (m *TLMsgContainer) DebugString() string {
	return fmt.Sprintf(`{"msg_container#73f1f8dc": []}`)
}

// TLMsgCopy
// msg_copy#e06046b2 orig_message:Message2 = MessageCopy; // parsed manually, not used - use msg_container
type TLMsgCopy struct {
	OrigMessage *TLMessage2
}

func (m *TLMsgCopy) String() string {
	return "{msg_copy#e06046b2}"
}

func (m *TLMsgCopy) Encode(x *EncodeBuf, layer int32) error {
	x.Int(int32(CRC32_msg_copy))
	if m.OrigMessage != nil {
		m.OrigMessage.Encode(x, layer)
	}
	return nil
}

func (m *TLMsgCopy) Decode(dbuf *DecodeBuf) error {
	o := dbuf.Object()
	if o == nil {
		return dbuf.err
	}
	message2, ok := o.(*TLMessage2)
	if !ok {
		return fmt.Errorf("msg_copy: expected *TLMessage2, got %T", o)
	}
	m.OrigMessage = message2
	return dbuf.err
}

func (m *TLMsgCopy) DebugString() string {
	if m.OrigMessage == nil {
		return `{"msg_copy#e06046b2": {"orig_message": null}}`
	}
	return fmt.Sprintf(`{"msg_copy#e06046b2": {"orig_message": %s}}`, m.OrigMessage.DebugString())
}

// TLGzipPacked
// gzip_packed#3072cfa1 packed_data:string = Object; // parsed manually
type TLGzipPacked struct {
	PackedData []byte
	Obj        TLObject
}

func (m *TLGzipPacked) String() string {
	return "{gzip_packed#3072cfa1}"
}

func (m *TLGzipPacked) Encode(x *EncodeBuf, layer int32) error {
	if len(m.PackedData) == 0 {
		return nil
	}
	b := gzipBufPool.Get().(*bytes.Buffer)
	b.Reset()
	defer func() {
		b.Reset()
		gzipBufPool.Put(b)
	}()
	gz := gzip.NewWriter(b)
	_, err := gz.Write(m.PackedData)
	gz.Flush()
	clErr := gz.Close()
	if err != nil || clErr != nil {
		x.Bytes(m.PackedData)
		return nil
	}
	x.Int(int32(CRC32_gzip_packed))
	x.StringBytes(b.Bytes())
	return nil
}

func (m *TLGzipPacked) Decode(dbuf *DecodeBuf) error {
	data := dbuf.StringBytes()
	if dbuf.err != nil {
		return dbuf.err
	}
	r := bytes.NewReader(data)
	var gz io.ReadCloser
	gz, err := gzip.NewReader(r)
	if err != nil {
		r.Reset(data)
		gz, err = zlib.NewReader(r)
		if err != nil {
			dbuf.err = err
			return fmt.Errorf("gzip read1: %v", err)
		}
	}
	buf := gzipBufPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer func() {
		buf.Reset()
		gzipBufPool.Put(buf)
	}()
	_, err = io.Copy(buf, gz)
	clErr := gz.Close()
	if err != nil {
		dbuf.err = err
		return fmt.Errorf("gzip read2: %v", err)
	}
	if clErr != nil {
		dbuf.err = clErr
		return clErr
	}
	m.PackedData = append([]byte(nil), buf.Bytes()...)
	dbuf2 := NewDecodeBuf(m.PackedData)
	m.Obj = dbuf2.Object()
	dbuf.err = dbuf2.err
	return dbuf.err
}

func (m *TLGzipPacked) DebugString() string {
	return fmt.Sprintf(`{"gzip_packed#3072cfa1": {}}`)
}

// TLRpcResult
// rpc_result#f35c6d01 req_msg_id:long result:Object = RpcResult; // parsed manually
type TLRpcResult struct {
	ReqMsgId int64
	Result   TLObject
}

func (m *TLRpcResult) String() string {
	return fmt.Sprintf("{rpc_result#f35c6d01: req_msg_id: %d, result: %s}", m.ReqMsgId, reflect.TypeOf(m.Result))
}

func (m *TLRpcResult) Encode(x *EncodeBuf, layer int32) error {
	x.Int(int32(CRC32_rpc_result))
	x.Long(m.ReqMsgId)

	rawBody, offset := func() ([]byte, int) {
		x2 := GetEncodeBuf()
		defer PutEncodeBuf(x2)
		m.Result.Encode(x2, layer)
		return append([]byte(nil), x2.GetBuf()...), x2.GetOffset()
	}()

	if offset > 256 {
		switch m.Result.(type) {
		case *Upload_WebFile:
			x.Bytes(rawBody)
		case *Upload_CdnFile:
			x.Bytes(rawBody)
		case *Upload_File:
			x.Bytes(rawBody)
		default:
			gzipPacked := &TLGzipPacked{
				PackedData: rawBody,
			}
			gzipPacked.Encode(x, layer)
		}
	} else {
		x.Bytes(rawBody)
	}

	return nil

}

func (m *TLRpcResult) Decode(dbuf *DecodeBuf) error {
	m.ReqMsgId = dbuf.Long()
	m.Result = dbuf.Object()
	return dbuf.err
}

func (m *TLRpcResult) DebugString() string {
	return fmt.Sprintf(`{"rpc_result#f35c6d01": {"req_msg_id": %d}}`, m.ReqMsgId)
}
