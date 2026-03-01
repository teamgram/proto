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

package crypto

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	serverSide = 0
	clientSide = 1
)

type AuthKey struct {
	authKeyId int64
	authKey   []byte
	side      int // client or server
}

func calcAuthKeyId(keyData []byte) int64 {
	sha1 := Sha1Digest(keyData)
	// Lower 64 bits = 8 bytes of 20 byte SHA1 hash.
	return int64(binary.LittleEndian.Uint64(sha1[12:]))
}

func NewAuthKey(keyId int64, keyData []byte) *AuthKey {
	// TODO(@benqi): check keyData len

	// check keyId
	if keyId == 0 {
		keyId = calcAuthKeyId(keyData)
	}
	return &AuthKey{
		authKeyId: keyId,
		authKey:   keyData,
		side:      serverSide,
	}
}

// CreateAuthKey
/*
## android client, pushAuthKey algo:
- authKey
	if (SharedConfig.pushAuthKey == null) {
		SharedConfig.pushAuthKey = new byte[256];
		Utilities.random.nextBytes(SharedConfig.pushAuthKey);
		SharedConfig.saveConfig();
	}

- calcAuthKeyId
	SharedConfig.pushAuthKeyId = new byte[8];
	byte[] authKeyHash = Utilities.computeSHA1(SharedConfig.pushAuthKey);
	System.arraycopy(authKeyHash, authKeyHash.length - 8, SharedConfig.pushAuthKeyId, 0, 8);

*/
// for bots...
func CreateAuthKey() *AuthKey {
	key := new(AuthKey)

	key.authKey = GenerateNonce(256)
	key.authKeyId = calcAuthKeyId(key.authKey)
	key.side = serverSide

	return key
}

func NewClientAuthKey(keyId int64, keyData []byte) *AuthKey {
	// check keyId
	// check keyId
	if keyId == 0 {
		keyId = calcAuthKeyId(keyData)
	}

	return &AuthKey{
		authKeyId: keyId,
		authKey:   keyData,
		side:      clientSide,
	}
}

func (k *AuthKey) AuthKeyId() int64 {
	return k.authKeyId
}

func (k *AuthKey) AuthKey() []byte {
	return k.authKey
}

func (k *AuthKey) Equals(o *AuthKey) bool {
	return bytes.Equal(k.authKey, o.authKey)
}

//func (k *AuthKey) CalcAuthKeyId() int64 {
//	return calcAuthKeyId(k.authKey)
//}

func (k *AuthKey) calcX(incoming bool) int {
	var x = 0
	if k.side == serverSide {
		if incoming {
			x = 8
		}
	} else {
		if !incoming {
			x = 8
		}
	}
	return x
}

func (k *AuthKey) prepareAESV1(msgKey []byte, incoming bool) ([]byte, []byte) {
	x := k.calcX(incoming)

	var dataA [48]byte
	copy(dataA[:16], msgKey[:16])
	copy(dataA[16:], k.authKey[x:x+32])
	sha1A := sha1.Sum(dataA[:])

	var dataB [48]byte
	copy(dataB[:16], k.authKey[32+x:32+x+16])
	copy(dataB[16:32], msgKey[:16])
	copy(dataB[32:], k.authKey[48+x:48+x+16])
	sha1B := sha1.Sum(dataB[:])

	var dataC [48]byte
	copy(dataC[:32], k.authKey[64+x:64+x+32])
	copy(dataC[32:], msgKey[:16])
	sha1C := sha1.Sum(dataC[:])

	var dataD [48]byte
	copy(dataD[:16], msgKey[:16])
	copy(dataD[16:], k.authKey[96+x:96+x+32])
	sha1D := sha1.Sum(dataD[:])

	aesKey := make([]byte, 32)
	aesIV := make([]byte, 32)

	copy(aesKey, sha1A[:8])
	copy(aesKey[8:], sha1B[8:20])
	copy(aesKey[20:], sha1C[4:16])
	copy(aesIV, sha1A[8:20])
	copy(aesIV[12:], sha1B[:8])
	copy(aesIV[20:], sha1C[16:20])
	copy(aesIV[24:], sha1D[:8])

	return aesKey, aesIV
}

func (k *AuthKey) prepareAES(msgKey []byte, incoming bool) ([]byte, []byte) {
	x := k.calcX(incoming)

	var dataA [52]byte
	copy(dataA[:16], msgKey[:16])
	copy(dataA[16:], k.authKey[x:x+36])
	sha256A := sha256.Sum256(dataA[:])

	var dataB [52]byte
	copy(dataB[:36], k.authKey[40+x:40+x+36])
	copy(dataB[36:], msgKey[:16])
	sha256B := sha256.Sum256(dataB[:])

	aesKey := make([]byte, 32)
	aesIV := make([]byte, 32)

	copy(aesKey[:8], sha256A[:8])
	copy(aesKey[8:24], sha256B[8:24])
	copy(aesKey[24:], sha256A[24:32])
	copy(aesIV[:8], sha256B[:8])
	copy(aesIV[8:24], sha256A[8:24])
	copy(aesIV[24:], sha256B[24:32])

	return aesKey, aesIV
}

func (k *AuthKey) partForMsgKey(incoming bool) []byte {
	x := k.calcX(incoming)
	return k.authKey[88+x : 88+x+32]
}

// AesIgeEncryptV1
/*
| salt <br> int64	| `session_id` <br> int64 | `message_id` <br> int64 | `seq_no` <br> int32 |`message_data_length` <br> int32	| `message_data` <br> bytes | padding12..1024 <br> bytes|
|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
*/
func (k *AuthKey) AesIgeEncryptV1(rawData []byte) ([]byte, []byte, error) {
	var additionalSize = len(rawData) % 16
	if additionalSize != 0 {
		additionalSize = 16 - additionalSize
	}

	tmpData := make([]byte, 0, len(rawData)+additionalSize)
	tmpData = append(tmpData, rawData...)
	if additionalSize > 0 {
		tmpData = append(tmpData, GenerateNonce(additionalSize)...)
	}

	// calc msg_key
	var msgKeyBuf [32]byte
	sha1Sum := sha1.Sum(rawData)
	copy(msgKeyBuf[4:], sha1Sum[:])

	aesKey, aesIV := k.prepareAESV1(msgKeyBuf[8:8+16], true)
	e := NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(tmpData)
	if err != nil {
		return nil, nil, err
	}

	result := make([]byte, 16)
	copy(result, msgKeyBuf[8:8+16])
	return result, x, nil
}

func (k *AuthKey) AesIgeEncrypt(rawData []byte) ([]byte, []byte, error) {
	var additionalSize = len(rawData) % 16
	if additionalSize != 0 {
		additionalSize = 16 - additionalSize
	}

	if additionalSize < 12 {
		additionalSize += 16
	}

	tmpData := make([]byte, 0, len(rawData)+additionalSize)
	tmpData = append(tmpData, rawData...)
	if additionalSize > 0 {
		tmpData = append(tmpData, GenerateNonce(additionalSize)...)
	}

	// calc msg_key
	sha256Dig := sha256.New()
	sha256Dig.Write(k.partForMsgKey(true))
	sha256Dig.Write(tmpData)
	var msgKeyBuf [32]byte
	sha256Dig.Sum(msgKeyBuf[:0])

	aesKey, aesIV := k.prepareAES(msgKeyBuf[8:8+16], true)
	e := NewAES256IGECryptor(aesKey, aesIV)

	x, err := e.Encrypt(tmpData)
	if err != nil {
		return nil, nil, err
	}

	result := make([]byte, 16)
	copy(result, msgKeyBuf[8:8+16])
	return result, x, nil
}

func (k *AuthKey) AesIgeDecryptV1(msgKey, rawData []byte) ([]byte, error) {
	aesKey, aesIV := k.prepareAESV1(msgKey, false)
	d := NewAES256IGECryptor(aesKey, aesIV)
	x, err := d.Decrypt(rawData)
	if err != nil {
		return nil, err
	}

	var dataLen = uint32(len(rawData))
	messageLen := binary.LittleEndian.Uint32(x[28:])
	if messageLen+32 > dataLen {
		err = fmt.Errorf("aesIgeDecrypt data(%d) error - Wrong message length %d", dataLen, messageLen)
		return nil, err
	}

	var calcMsgKey [32]byte
	sha1Sum := sha1.Sum(x[:32+messageLen])
	copy(calcMsgKey[4:], sha1Sum[:])

	if !bytes.Equal(calcMsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("aesIgeDecrypt data error - (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(rawData[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			k.authKeyId,
			hex.EncodeToString(k.authKey[88:88+32]),
			hex.EncodeToString(calcMsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		return nil, err
	}

	return x, nil
}

func (k *AuthKey) AesIgeDecrypt(msgKey, rawData []byte) ([]byte, error) {
	aesKey, aesIV := k.prepareAES(msgKey, false)
	d := NewAES256IGECryptor(aesKey, aesIV)
	x, err := d.Decrypt(rawData)
	if err != nil {
		return nil, err
	}

	var dataLen = uint32(len(rawData))
	messageLen := binary.LittleEndian.Uint32(x[28:])

	if messageLen+32 > dataLen {
		err = fmt.Errorf("aesIgeDecrypt data(%d) error - Wrong message length %d", dataLen, messageLen)
		return nil, err
	}

	sha256Dig := sha256.New()
	sha256Dig.Write(k.partForMsgKey(false))
	sha256Dig.Write(x[:dataLen])
	var calcMsgKey [32]byte
	sha256Dig.Sum(calcMsgKey[:0])

	if !bytes.Equal(calcMsgKey[8:8+16], msgKey[:16]) {
		err = fmt.Errorf("aesIgeDecrypt data error - (data: %s, aesKey: %s, aseIV: %s, authKeyId: %d, authKey: %s), msgKey verify error, sign: %s, msgKey: %s",
			hex.EncodeToString(rawData[:64]),
			hex.EncodeToString(aesKey),
			hex.EncodeToString(aesIV),
			k.authKeyId,
			hex.EncodeToString(k.authKey[88:88+32]),
			hex.EncodeToString(calcMsgKey[8:8+16]),
			hex.EncodeToString(msgKey[:16]))
		return nil, err
	}

	return x, nil
}
