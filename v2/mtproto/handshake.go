// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

import (
	"github.com/teamgram/proto/v2/bin"
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzID_resPQ = 0x05162463
)

func init() {
	iface.RegisterClazzID(ClazzID_resPQ, func() iface.TLObject { return &TLResPQ{} })
}

// TLResPQ
// resPQ#05162463 nonce:int128 server_nonce:int128 pq:string server_public_key_fingerprints:Vector<long> = ResPQ;
type TLResPQ struct {
	Nonce                       int64   `json:"nonce"`
	ServerNonce                 int64   `json:"server_nonce"`
	Pq                          string  `json:"pq"`
	ServerPublicKeyFingerprints []int64 `json:"server_public_key_fingerprints"`
}

func (m *TLResPQ) ClazzName() string {
	return "message2"
}

func (m *TLResPQ) Encode(x *bin.Encoder, layer int32) error {
	_ = layer

	//x.PutInt64(m.MsgId)
	//x.PutInt32(m.Seqno)
	//x.PutInt32(m.Bytes)
	//x.Put(m.Body)

	return nil
}

func (m *TLResPQ) Decode(d *bin.Decoder) (err error) {
	return nil
}
