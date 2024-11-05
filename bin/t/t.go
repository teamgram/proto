// Copyright © 2024 Teamgram Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

//import (
//	"fmt"
//
//	"github.com/teamgram/proto/mtproto"
//)
//
//type TLObject interface {
//	Encode(e *mtproto.Encoder, layer int32)
//	Decode(d *mtproto.Decoder) error
//}
//
//func WriteObjectList[T TLObject](e *mtproto.Encoder, vList []T) {
//	e.WriteInt(int32(len(vList)))
//	for _, obj := range vList {
//		obj.Encode(e, 0)
//	}
//}
//
//type TLPing struct {
//	PingId int64
//}
//
//func (m *TLPing) Encode(e *mtproto.Encoder, layer int32) {
//	e.WriteLong(m.PingId)
//}
//
//func (m *TLPing) Decode(d *mtproto.Decoder) error {
//	// m.PingId = d.Long()
//	// return d.GetError()
//
//	return nil
//}
//
//func main() {
//	enc := mtproto.NewEncoder()
//	enc.WriteBytes([]byte("Hello world!"))
//	b := enc.Bytes()
//	enc.End()
//
//	enc = mtproto.NewEncoder()
//	enc.WriteBytes([]byte("oh, my god"))
//	ping := &TLPing{
//		PingId: 12345,
//	}
//	WriteObjectList(enc, []*TLPing{ping})
//	b2 := enc.Bytes()
//	enc.End()
//
//	fmt.Println(string(b))
//	fmt.Println(string(b2))
//}
