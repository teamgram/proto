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

func (m *Document) GetFixedSize() int64 {
	if m.Size2_INT64 != 0 {
		return m.Size2_INT64
	}
	if m.Size2_INT32 != 0 {
		return int64(m.Size2_INT32)
	}
	if m.Size2 != 0 {
		return m.Size2
	}

	return 0
}

func (m *EncryptedFile) GetFixedSize() int64 {
	if m.Size2_INT64 != 0 {
		return m.Size2_INT64
	}
	if m.Size2_INT32 != 0 {
		return int64(m.Size2_INT32)
	}

	return 0
}
