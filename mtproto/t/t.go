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

import (
	"fmt"
	"github.com/teamgram/proto/bin"
	"github.com/teamgram/proto/errs"
	"github.com/teamgram/proto/mtproto"
)

func main() {
	var (
		message2 mtproto.Message2
	)
	x := bin.NewEncoder()
	x.End()

	message2 = &mtproto.TLMessageRawData{
		MsgId:   1,
		Seqno:   2,
		Bytes:   3,
		ClazzID: 4,
		Body:    []byte("1234567890"),
	}
	message2.Encode(x, 1)

	fmt.Println(message2.GetMessage2ClazzName())
	fmt.Printf("%s\n", x.Bytes())

	err := errs.Wrap(errs.ErrTokenNotExist)
	fmt.Printf("%+v", err)
}
