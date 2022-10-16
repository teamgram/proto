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

// chatReactionsNone#eafc32bc = ChatReactions;
// chatReactionsAll#52928bca flags:# allow_custom:flags.0?true = ChatReactions;
// chatReactionsSome#661d4037 reactions:Vector<Reaction> = ChatReactions;

const (
	ChatReactionsTypeNotDefined       int32 = 0
	ChatReactionsTypeNone             int32 = 1
	ChatReactionsTypeAllNoAllowCustom int32 = 2
	ChatReactionsTypeAllAllowCustom   int32 = 3
	ChatReactionsTypeSome             int32 = 4
)

var (
	defaultChatReactionsNone              = MakeTLChatReactionsNone(nil).To_ChatReactions()
	defaultChatReactionsAllNotAllowCustom = MakeTLChatReactionsAll(&ChatReactions{
		AllowCustom: false,
	}).To_ChatReactions()
	defaultChatReactionsAllAllowCustom = MakeTLChatReactionsAll(&ChatReactions{
		AllowCustom: true,
	}).To_ChatReactions()
)

func (m *ChatReactions) ToChatReactions() (int32, []string) {
	switch m.GetPredicateName() {
	case Predicate_chatReactionsNone:
		return ChatReactionsTypeNone, nil
	case Predicate_chatReactionsAll:
		if !m.AllowCustom {
			return ChatReactionsTypeAllNoAllowCustom, nil
		} else {
			return ChatReactionsTypeAllAllowCustom, nil
		}
	case Predicate_chatReactionsSome:
		reactions := make([]string, 0, len(m.Reactions))
		for _, r := range m.Reactions {
			reactions = append(reactions, r.ToString())
		}

		return ChatReactionsTypeSome, reactions
	}

	return ChatReactionsTypeNotDefined, nil
}

func FromChatReactions(t int32, reactions []string) *ChatReactions {
	switch t {
	case ChatReactionsTypeNone:
		return defaultChatReactionsNone
	case ChatReactionsTypeAllNoAllowCustom:
		return defaultChatReactionsAllNotAllowCustom
	case ChatReactionsTypeAllAllowCustom:
		return defaultChatReactionsAllAllowCustom
	case ChatReactionsTypeSome:
		chatReactions := MakeTLChatReactionsSome(&ChatReactions{
			Reactions: make([]*Reaction, 0, len(reactions)),
		}).To_ChatReactions()
		for _, r := range reactions {
			chatReactions.Reactions = append(chatReactions.Reactions, FromReaction(r))
		}
		return chatReactions
	}

	return nil
}
