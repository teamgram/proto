// Copyright 2022 Teamgram Authors
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

func (m *TLMessageRawData) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	// x.Int(int32(CRC32_message2))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(m.Body)
	return x.buf
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

func NewTLMsgRawDataContainer() *TLMsgRawDataContainer {
	return &TLMsgRawDataContainer{
		Messages: []*TLMessageRawData{},
	}
}

func (m *TLMsgRawDataContainer) String() string {
	return "{msg_container#73f1f8dc}"
}

func (m *TLMsgRawDataContainer) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(CRC32_msg_container))

	x.Int(int32(len(m.Messages)))
	for _, m := range m.Messages {
		x.Bytes(m.Encode(layer))
	}
	return x.buf
}

func (m *TLMsgRawDataContainer) Decode(dbuf *DecodeBuf) error {
	len := dbuf.Int()
	// log.Infof("msg_container: messages size: %d", len)
	for i := 0; i < int(len); i++ {
		// log.Infof("msg_container: decode messages[%d]: ", i)
		message2 := &TLMessageRawData{}
		err := message2.Decode(dbuf)
		if err != nil {
			// log.Errorf("Decode message2 error: %v", err)
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

func (m *TLMessage2) Encode(layer int32) []byte {
	oBuf := m.Object.Encode(layer)
	m.Bytes = int32(len(oBuf))

	x := NewEncodeBuf(16 + len(oBuf))
	x.Long(m.MsgId)
	x.Int(m.Seqno)
	x.Int(m.Bytes)
	x.Bytes(oBuf)

	return x.buf
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
		err := fmt.Errorf("decode core_message error(%v): %s", dBuf2.err, hex.EncodeToString(b))
		// log.Error(err.Error())
		return err
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
	Messages []TLMessage2
	// RawMessages []TLMessageRawData

}

func (m *TLMsgContainer) String() string {
	return "{msg_container#73f1f8dc}"
}

func (m *TLMsgContainer) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(CRC32_msg_container))

	x.Int(int32(len(m.Messages)))
	for _, m := range m.Messages {
		x.Bytes(m.Encode(layer))
	}
	return x.buf
}

func (m *TLMsgContainer) Decode(dbuf *DecodeBuf) error {
	len2 := dbuf.Int()
	// log.Debugf("msg_container: messages size: %d", len)
	for i := 0; i < int(len2); i++ {
		// log.Debugf("msg_container: decode messages[%d]: ", i)
		message2 := &TLMessage2{}
		err := message2.Decode(dbuf)
		if err != nil {
			// log.Errorf("Decode message2 error: %v", err)
			return err
		}
		m.Messages = append(m.Messages, *message2)
	}
	return dbuf.err
}

func (m *TLMsgContainer) DebugString() string {
	return fmt.Sprintf(`{"msg_container#73f1f8dc": []}`)
}

// TLMsgCopy
// msg_copy#e06046b2 orig_message:Message2 = MessageCopy; // parsed manually, not used - use msg_container
type TLMsgCopy struct {
	OrigMessage TLMessage2
}

func (m *TLMsgCopy) String() string {
	return "{msg_copy#e06046b2}"
}

func (m *TLMsgCopy) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(CRC32_msg_copy))
	x.Bytes(m.OrigMessage.Encode(layer))
	return x.buf
}

func (m *TLMsgCopy) Decode(dbuf *DecodeBuf) error {
	o := dbuf.Object()
	message2, _ := o.(*TLMessage2)
	m.OrigMessage = *message2
	return dbuf.err
}

func (m *TLMsgCopy) DebugString() string {
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

func (m *TLGzipPacked) Encode(layer int32) []byte {
	if len(m.PackedData) == 0 {
		return m.PackedData
	}

	var (
		err error
		b   = new(bytes.Buffer)
	)
	gz := gzip.NewWriter(b)
	// _, err = io.Copy(gz, bytes.NewBuffer(m.PackedData))
	_, err = gz.Write(m.PackedData)
	gz.Flush()
	clErr := gz.Close()

	if err != nil {
		// log.Errorf("gzip write: %v", err)
		return m.PackedData
	}
	if clErr != nil {
		// log.Errorf("gzip write: %v", err)
		return m.PackedData
	}

	x := NewEncodeBuf(512)
	x.Int(int32(CRC32_gzip_packed))
	x.StringBytes(b.Bytes())
	return x.buf
}

func (m *TLGzipPacked) Decode(dbuf *DecodeBuf) error {
	data := dbuf.StringBytes()
	if dbuf.err != nil {
		return dbuf.err
	}

	var (
		gz  io.ReadCloser
		err error
	)
	if gz, err = gzip.NewReader(bytes.NewBuffer(data)); err != nil {
		if gz, err = zlib.NewReader(bytes.NewBuffer(data)); err != nil {
			dbuf.err = err
			return fmt.Errorf("gzip read1: %v", err)
		}
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		dbuf.err = err
		return fmt.Errorf("gzip read2: %v", err)
	}
	if clErr != nil {
		dbuf.err = clErr
		return clErr
	}

	m.PackedData = buf.Bytes()

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
	return "{rpc_result#f35c6d01: req_msg_id:" + string(m.ReqMsgId) + "}"
}

func (m *TLRpcResult) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(CRC32_rpc_result))
	x.Long(m.ReqMsgId)
	rawBody := m.Result.Encode(layer)
	if len(rawBody) > 256 {
		switch m.Result.(type) {
		case *Upload_WebFile:
		case *Upload_CdnFile:
		case *Upload_File:
		default:
			gzipPacked := &TLGzipPacked{
				PackedData: rawBody,
			}
			rawBody = gzipPacked.Encode(layer)
		}
	}
	x.Bytes(rawBody)
	return x.buf
}

func (m *TLRpcResult) Decode(dbuf *DecodeBuf) error {
	m.ReqMsgId = dbuf.Long()
	m.Result = dbuf.Object()
	return dbuf.err
}

func (m *TLRpcResult) DebugString() string {
	return fmt.Sprintf(`{"rpc_result#f35c6d01": {"req_msg_id": %d}}`, m.ReqMsgId)
}
