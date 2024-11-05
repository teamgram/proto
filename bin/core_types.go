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

package bin

const (
	MTPROTO_VERSION = 2
)

// Basic TL types.
const (
	ClazzID_int       = 0xa8509bda // int = Int (0xa8509bda)
	ClazzID_long      = 0x22076cba // long = Long (0x22076cba)
	ClazzID_double    = 0x2210c154 // double = Double (0x2210c154)
	ClazzID_string    = 0xb5286e24 // string = String (0xb5286e24)
	ClazzID_vector    = 0x1cb5c415 // vector {t:Type} # [ t ] = Vector t
	ClazzID_bytes     = 0xe937bb82 // bytes#e937bb82 = Bytes
	ClazzID_boolTrue  = 0x997275b5 // boolTrue#997275b5 = Bool
	ClazzID_boolFalse = 0xbc799737 // boolFalse#bc799737 = Bool
)

// WordLen represents 4-byte sequence.
// Values in TL are generally aligned to Word.
const WordLen = 4

func nearestPaddedValueLength(l int) int {
	n := WordLen * (l / WordLen)
	if n < l {
		n += WordLen
	}
	return n
}
