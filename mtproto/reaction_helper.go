// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

import (
	"strconv"
)

var (
	reactionEmpty = MakeTLReactionEmpty(nil).To_Reaction()
)

func (m *Reaction) ToString() string {
	v := m.Emoticon
	if m.DocumentId != 0 {
		v = strconv.FormatInt(m.DocumentId, 10)
	}

	return v
}

func FromReaction(s string) *Reaction {
	if s == "" {
		return reactionEmpty
	}
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return MakeTLReactionEmoji(&Reaction{
			Emoticon: s,
		}).To_Reaction()
	}

	return MakeTLReactionCustomEmoji(&Reaction{
		DocumentId: id,
	}).To_Reaction()
}
