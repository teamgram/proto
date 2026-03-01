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
	"crypto/aes"
	"encoding/binary"
	"errors"
)

type AES256IGECryptor struct {
	aesKey []byte
	aesIV  []byte
}

func NewAES256IGECryptor(aesKey, aesIV []byte) *AES256IGECryptor {
	// guard conditions
	if (len(aesIV)) < aes.BlockSize {
		return nil
	}

	k := len(aesKey)
	switch k {
	default:
		return nil
	case 16, 24, 32:
		break
	}
	return &AES256IGECryptor{
		aesKey: aesKey,
		aesIV:  aesIV,
	}
}

// xorBlocks performs XOR on 16-byte AES blocks using uint64-aligned operations.
// This is ~4-8x faster than byte-by-byte XOR for fixed 16-byte blocks.
func xorBlocks(dst, src []byte) {
	_ = dst[15] // bounds check hint
	_ = src[15]
	a := binary.LittleEndian.Uint64(dst[:8])
	b := binary.LittleEndian.Uint64(src[:8])
	binary.LittleEndian.PutUint64(dst[:8], a^b)
	a = binary.LittleEndian.Uint64(dst[8:16])
	b = binary.LittleEndian.Uint64(src[8:16])
	binary.LittleEndian.PutUint64(dst[8:16], a^b)
}

// Encrypt
// data长度必须是aes.BlockSize(16)的倍数，如果不是请调用者补齐
func (c *AES256IGECryptor) Encrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to encrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	var t, x, y [aes.BlockSize]byte
	copy(x[:], c.aesIV[:aes.BlockSize])
	copy(y[:], c.aesIV[aes.BlockSize:])
	encrypted := make([]byte, len(data))

	for i := 0; i < len(data); i += aes.BlockSize {
		// x ^= plaintext
		xorBlocks(x[:], data[i:i+aes.BlockSize])
		// t = E(x)
		block.Encrypt(t[:], x[:])
		// t ^= y
		xorBlocks(t[:], y[:])
		// x = t (ciphertext), y = plaintext
		copy(x[:], t[:])
		copy(y[:], data[i:i+aes.BlockSize])
		copy(encrypted[i:], t[:])
	}

	return encrypted, nil
}

func (c *AES256IGECryptor) Decrypt(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.aesKey)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("AES256IGE: data too small to decrypt")
	}
	if len(data)%aes.BlockSize != 0 {
		return nil, errors.New("AES256IGE: data not divisible by block size")
	}

	var t, x, y [aes.BlockSize]byte
	copy(x[:], c.aesIV[:aes.BlockSize])
	copy(y[:], c.aesIV[aes.BlockSize:])
	decrypted := make([]byte, len(data))

	for i := 0; i < len(data); i += aes.BlockSize {
		// y ^= ciphertext
		xorBlocks(y[:], data[i:i+aes.BlockSize])
		// t = D(y)
		block.Decrypt(t[:], y[:])
		// t ^= x
		xorBlocks(t[:], x[:])
		// y = t (plaintext), x = ciphertext
		copy(y[:], t[:])
		copy(x[:], data[i:i+aes.BlockSize])
		copy(decrypted[i:], t[:])
	}

	return decrypted, nil
}
