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
	"github.com/gogo/protobuf/types"
	"time"
)

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

func (m *Document) FixData() *Document {
	v := m.GetFixedSize()

	if m.Size2 == 0 {
		m.Size2 = v
	}
	if m.Size2_INT64 == 0 {
		m.Size2_INT64 = v
	}
	if m.Size2_INT32 == 0 {
		m.Size2_INT32 = int32(v)
	}

	for i := 0; i < len(m.Attributes); i++ {
		m.Attributes[i] = m.Attributes[i].FixData()
	}

	return m
}

func (m *DocumentAttribute) GetFixedDuration() float64 {
	if m.Duration_FLOAT64 != 0 {
		return m.Duration_FLOAT64
	}
	if m.Duration_INT32 != 0 {
		return float64(m.Duration_INT32)
	}
	if m.Duration != 0 {
		return m.Duration
	}

	return 0
}

func (m *DocumentAttribute) FixData() *DocumentAttribute {
	v := m.GetFixedDuration()

	if m.Duration == 0 {
		m.Duration = v
	}
	if m.Duration_INT32 == 0 {
		m.Duration_INT32 = int32(v)
	}
	if m.Duration_FLOAT64 == 0 {
		m.Duration_FLOAT64 = v
	}

	return m
}

func (m *PollResults) FixData() *PollResults {
	if m.Solution != nil {
		if m.SolutionEntities == nil {
			m.SolutionEntities = []*MessageEntity{}
		}
	}
	return m
}

func (m *PollAnswer) GetTextString() (text string) {
	if m.Text != "" {
		text = m.Text
	}
	if m.Text_STRING != "" {
		text = m.Text_STRING
	}
	if m.Text_TEXTWITHENTITIES != nil {
		text = m.Text_TEXTWITHENTITIES.GetText()
	}

	return
}

func (m *PollAnswer) FixData() *PollAnswer {
	if m.Text == "" {
		if m.Text_STRING != "" {
			m.Text = m.Text_STRING
		} else if m.Text_TEXTWITHENTITIES != nil {
			m.Text = m.Text_TEXTWITHENTITIES.GetText()
		}
	}
	if m.Text_STRING == "" {
		m.Text_STRING = m.Text
	}
	if m.Text_TEXTWITHENTITIES == nil {
		m.Text_TEXTWITHENTITIES = MakeTLTextWithEntities(&TextWithEntities{
			Text:     m.Text,
			Entities: make([]*MessageEntity, 0),
		}).To_TextWithEntities()
	}
	return m
}

func (m *Poll) GetQuestionString() (q string) {
	if m.Question != "" {
		q = m.Question
	}
	if m.Question_STRING != "" {
		q = m.Question_STRING
	}
	if m.Question_TEXTWITHENTITIES != nil {
		q = m.Question_TEXTWITHENTITIES.GetText()
	}

	return
}

func (m *Poll) FixData() *Poll {
	if m.Question == "" {
		if m.Question_STRING != "" {
			m.Question = m.Question_STRING
		} else if m.Question_TEXTWITHENTITIES != nil {
			m.Question = m.Question_TEXTWITHENTITIES.GetText()
		}
	}
	if m.Question_STRING == "" {
		m.Question_STRING = m.Question
	}
	if m.Question_TEXTWITHENTITIES == nil {
		m.Question_TEXTWITHENTITIES = MakeTLTextWithEntities(&TextWithEntities{
			Text:     m.Question,
			Entities: make([]*MessageEntity, 0),
		}).To_TextWithEntities()
	}

	for i := 0; i < len(m.Answers); i++ {
		m.Answers[i] = m.Answers[i].FixData()
	}

	return m
}

func (m *Update) FixData() *Update {
	if m.GetMedia() != nil {
		m.Media = m.Media.FixData()
	}
	if m.GetMessage_MESSAGE() != nil {
		m.Message_MESSAGE = m.Message_MESSAGE.FixData()
	}

	return m
}

func (m *MessageMedia) FixData() *MessageMedia {
	if m.GetDocument() != nil {
		m.Document = m.Document.FixData()
	}

	if m.GetResults() != nil {
		m.Results = m.Results.FixData()
	}

	if m.GetPoll() != nil {
		m.Poll = m.Poll.FixData()
	}

	return m
}

func (m *Message) FixData() *Message {
	if m.GetMedia() != nil {
		m.Media = m.Media.FixData()
	}

	if m.GetReactions() != nil {
		m.Reactions = m.Reactions.FixData()
	}

	if m.GetTtlPeriod() != nil {
		ttlPeriod := m.GetTtlPeriod().GetValue() + m.GetDate() - int32(time.Now().Unix())
		if ttlPeriod < 0 {
			ttlPeriod = 0
		}
		m.TtlPeriod = &types.Int32Value{Value: ttlPeriod}
	}

	if m.GetAction() != nil {
		m.Action = m.Action.FixData()
	}

	if m.GetReplyTo() != nil {
		m.ReplyTo = m.ReplyTo.FixData()
	}

	return m
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

func (m *ReactionCount) FixData() *ReactionCount {
	if m.Reaction != "" {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction
		}
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction,
			}).To_Reaction()
		}
	}

	if m.Reaction_STRING != "" {
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction_STRING,
			}).To_Reaction()
		}
	}

	if m.Reaction_REACTION != nil {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction_REACTION.GetEmoticon()
		}
	}

	return m
}

func (m *MessagePeerReaction) FixData() *MessagePeerReaction {
	if m.Reaction != "" {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction
		}
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction,
			}).To_Reaction()
		}
	}

	if m.Reaction_STRING != "" {
		if m.Reaction_REACTION == nil {
			m.Reaction_REACTION = MakeTLReactionEmoji(&Reaction{
				Emoticon: m.Reaction_STRING,
			}).To_Reaction()
		}
	}

	if m.Reaction_REACTION != nil {
		if m.Reaction_STRING == "" {
			m.Reaction_STRING = m.Reaction_REACTION.GetEmoticon()
		}
	}

	return m
}

func (m *MessageReactions) FixData() *MessageReactions {
	for _, v := range m.Results {
		v.FixData()
	}
	for _, v := range m.RecentReactions {
		v.FixData()
	}

	return m
}
