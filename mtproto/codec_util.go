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
	"crypto/sha1"
	"crypto/sha256"
	"time"
)

//int64_t ConnectionsManager::generateMessageId() {
//int64_t messageId = (int64_t) ((((double) getCurrentTimeMillis() + ((double) timeDifference) * 1000) * 4294967296.0) / 1000.0);
//if (messageId <= lastOutgoingMessageId) {
//messageId = lastOutgoingMessageId + 1;
//}
//while (messageId % 4 != 0) {
//messageId++;
//}
//lastOutgoingMessageId = messageId;
//return messageId;
//}

func GenerateMessageId() int64 {
	const nano = 1000 * 1000 * 1000
	unixnano := time.Now().UnixNano()

	messageId := ((unixnano / nano) << 32) | ((unixnano % nano) & -4)
	for {
		//rpc_response
		if (messageId % 4) != 1 {
			messageId += 1
		} else {
			break
		}

		/****************************
		 * // rpc_request
		 * if (messageId % 4) != 3 {
		 * 	messageId += 1
		 * } else {
		 * 	break
		 * }
		 */
	}

	return messageId
}

/*
uint32_t x = incoming ? 8 : 0;
static uint8_t sha[68];

SHA256_Init(&sha256Ctx);
SHA256_Update(&sha256Ctx, messageKey, 16);
SHA256_Update(&sha256Ctx, authKey + x, 36);
SHA256_Final(sha, &sha256Ctx);

SHA256_Init(&sha256Ctx);
SHA256_Update(&sha256Ctx, authKey + 40 + x, 36);
SHA256_Update(&sha256Ctx, messageKey, 16);
SHA256_Final(sha + 32, &sha256Ctx);

memcpy(result, sha, 8);
memcpy(result + 8, sha + 32 + 8, 16);
memcpy(result + 8 + 16, sha + 24, 8);

memcpy(result + 32, sha + 32, 8);
memcpy(result + 32 + 8, sha + 8, 16);
memcpy(result + 32 + 8 + 16, sha + 32 + 24, 8);
*/
func generateMessageKey(msgKey, authKey []byte, incoming bool) (aesKey, aesIV []byte) {
	var x = 0
	if incoming {
		x = 8
	}

	switch MTPROTO_VERSION {
	case 2:
		var tA [52]byte
		copy(tA[:16], msgKey[:16])
		copy(tA[16:], authKey[x:x+36])
		sha256A := sha256.Sum256(tA[:])

		var tB [52]byte
		copy(tB[:36], authKey[40+x:40+x+36])
		copy(tB[36:], msgKey[:16])
		sha256B := sha256.Sum256(tB[:])

		var aesKeyArr [32]byte
		copy(aesKeyArr[:8], sha256A[:8])
		copy(aesKeyArr[8:24], sha256B[8:24])
		copy(aesKeyArr[24:], sha256A[24:32])

		var aesIVArr [32]byte
		copy(aesIVArr[:8], sha256B[:8])
		copy(aesIVArr[8:24], sha256A[8:24])
		copy(aesIVArr[24:], sha256B[24:32])

		aesKey = aesKeyArr[:]
		aesIV = aesIVArr[:]

	default:
		var tA [48]byte
		copy(tA[:16], msgKey[:16])
		copy(tA[16:], authKey[x:x+32])
		sha1A := sha1.Sum(tA[:])

		var tB [48]byte
		copy(tB[:16], authKey[32+x:32+x+16])
		copy(tB[16:32], msgKey[:16])
		copy(tB[32:], authKey[48+x:48+x+16])
		sha1B := sha1.Sum(tB[:])

		var tC [48]byte
		copy(tC[:32], authKey[64+x:64+x+32])
		copy(tC[32:], msgKey[:16])
		sha1C := sha1.Sum(tC[:])

		var tD [48]byte
		copy(tD[:16], msgKey[:16])
		copy(tD[16:], authKey[96+x:96+x+32])
		sha1D := sha1.Sum(tD[:])

		var aesKeyArr [32]byte
		copy(aesKeyArr[:8], sha1A[:8])
		copy(aesKeyArr[8:20], sha1B[8:20])
		copy(aesKeyArr[20:], sha1C[4:16])

		var aesIVArr [32]byte
		copy(aesIVArr[:12], sha1A[8:20])
		copy(aesIVArr[12:20], sha1B[:8])
		copy(aesIVArr[20:24], sha1C[16:20])
		copy(aesIVArr[24:], sha1D[:8])

		aesKey = aesKeyArr[:]
		aesIV = aesIVArr[:]
	}

	return
}
