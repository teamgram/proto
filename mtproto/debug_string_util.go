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
	"encoding/hex"
	"strconv"
	"strings"
	"sync"
)

func DebugStringVectorInt(v []int32) string {
	return joinInt32s(v, ",")
}

func DebugStringVectorLong(v []int64) string {
	return joinInt64s(v, ",")
}

func DebugStringVectorString(v []string) string {
	if len(v) == 0 {
		return ""
	} else {
		return "\"" + strings.Join(v, "\",") + "\""
	}
}

func DebugStringBytes(v []byte) string {
	if len(v) <= 64 {
		return hex.EncodeToString(v)
	} else {
		return hex.EncodeToString(v[:64])
	}
}

func DebugStringString(v string) string {
	if len(v) <= 64 {
		return v
	} else {
		return v[:64]
	}
}

var (
	bfPool = sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer([]byte{})
		},
	}
)

// joinInt32s format int32 slice like:n1,n2,n3.
func joinInt32s(is []int32, p string) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(int64(is[0]), 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(int64(i), 10))
		buf.WriteString(p)
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}

// joinInt64s format int64 slice like:n1,n2,n3.
func joinInt64s(is []int64, p string) string {
	if len(is) == 0 {
		return ""
	}
	if len(is) == 1 {
		return strconv.FormatInt(is[0], 10)
	}
	buf := bfPool.Get().(*bytes.Buffer)
	for _, i := range is {
		buf.WriteString(strconv.FormatInt(i, 10))
		buf.WriteString(p)
	}
	if buf.Len() > 0 {
		buf.Truncate(buf.Len() - 1)
	}
	s := buf.String()
	buf.Reset()
	bfPool.Put(buf)
	return s
}
