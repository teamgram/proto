// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: @benqi (wubenqi@gmail.com)

package mtproto

import (
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (x *PeerColor) GetBackgroundEmojiId() *wrapperspb.Int64Value {
	return x.GetBackgroundEmojiId_FLAGINT64()
}
