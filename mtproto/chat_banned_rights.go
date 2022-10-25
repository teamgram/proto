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
	"math"
)

// chatBannedRights#9f120418 flags:#
//	view_messages:flags.0?true
//	send_messages:flags.1?true
//	send_media:flags.2?true
//	send_stickers:flags.3?true
//	send_gifs:flags.4?true
//	send_games:flags.5?true
//	send_inline:flags.6?true
//	embed_links:flags.7?true
//	send_polls:flags.8?true
//	change_info:flags.10?true
//	invite_users:flags.15?true
//	pin_messages:flags.17?true
//	until_date:int = ChatBannedRights;

const (
	BANNED_VIEW_MESSAGES = 1 << 0
	BANNED_SEND_MESSAGES = 1 << 1
	BANNED_SEND_MEDIA    = 1 << 2
	BANNED_SEND_STICKERS = 1 << 3
	BANNED_SEND_GIFS     = 1 << 4
	BANNED_SEND_GAMES    = 1 << 5
	BANNED_SEND_INLINE   = 1 << 6
	BANNED_EMBED_LINKS   = 1 << 7
	BANNED_SEND_POLLS    = 1 << 8
	BANNED_CHANGE_INFO   = 1 << 10
	BANNED_INVITE_USERS  = 1 << 15
	BANNED_PIN_MESSAGES  = 1 << 17
)

type BannedRights int64

func MakeDefaultBannedRights() *ChatBannedRights {
	return MakeTLChatBannedRights(&ChatBannedRights{
		ViewMessages: false,
		SendMessages: false,
		SendMedia:    false,
		SendStickers: false,
		SendGifs:     false,
		SendGames:    false,
		SendInline:   false,
		EmbedLinks:   false,
		SendPolls:    false,
		ChangeInfo:   false,
		InviteUsers:  false,
		PinMessages:  false,
		UntilDate:    math.MaxInt32,
	}).To_ChatBannedRights()
}

func (m BannedRights) ToChatBannedRights() *ChatBannedRights {
	if m == 0 {
		return nil
	}

	return MakeTLChatBannedRights(&ChatBannedRights{
		ViewMessages: (m & BANNED_VIEW_MESSAGES) != 0,
		SendMessages: (m & BANNED_SEND_MESSAGES) != 0,
		SendMedia:    (m & BANNED_SEND_MEDIA) != 0,
		SendStickers: (m & BANNED_SEND_STICKERS) != 0,
		SendGifs:     (m & BANNED_SEND_GIFS) != 0,
		SendGames:    (m & BANNED_SEND_GAMES) != 0,
		SendInline:   (m & BANNED_SEND_INLINE) != 0,
		EmbedLinks:   (m & BANNED_EMBED_LINKS) != 0,
		SendPolls:    (m & BANNED_SEND_POLLS) != 0,
		ChangeInfo:   (m & BANNED_CHANGE_INFO) != 0,
		InviteUsers:  (m & BANNED_INVITE_USERS) != 0,
		PinMessages:  (m & BANNED_PIN_MESSAGES) != 0,
		UntilDate:    int32(m >> 32),
	}).To_ChatBannedRights()
}

func (m *ChatBannedRights) hasRights() bool {
	return m.ViewMessages ||
		m.SendMessages ||
		m.SendStickers ||
		m.SendGifs ||
		m.SendGames ||
		m.EmbedLinks ||
		m.SendPolls ||
		m.ChangeInfo ||
		m.InviteUsers ||
		m.PinMessages
}

func (m *ChatBannedRights) NoBanRights() bool {
	return m == nil || !m.hasRights()
}

func (m *ChatBannedRights) IsKick() bool {
	return m.GetViewMessages() && m.GetUntilDate() == math.MaxInt32
}

func (m *ChatBannedRights) IsBan(date int32) bool {
	return m.hasRights() && date <= m.UntilDate
}

func (m *ChatBannedRights) IsRestrict(date int32) bool {
	return !m.IsKick() && m.hasRights() && date <= m.UntilDate
}

func (m *ChatBannedRights) CanViewMessages(date int32) bool {
	return !m.GetViewMessages() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendMessages(date int32) bool {
	return !m.GetSendMessages() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendMedia(date int32) bool {
	return !m.GetSendMedia() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendStickers(date int32) bool {
	return !m.GetSendStickers() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendGifs(date int32) bool {
	return !m.GetSendGifs() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendGames(date int32) bool {
	return !m.GetSendGames() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendInline(date int32) bool {
	return !m.GetSendInline() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanEmbedLinks(date int32) bool {
	return !m.GetEmbedLinks() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendPolls(date int32) bool {
	return !m.GetSendPolls() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanChangeInfo(date int32) bool {
	return !m.GetChangeInfo() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanInviteUsers(date int32) bool {
	return !m.GetInviteUsers() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanPinMessages(date int32) bool {
	return !m.GetPinMessages() || date >= m.UntilDate
}

func (m *ChatAdminRights) DisallowChannel() bool {
	// return m.CanPostMessages() || m.CanEditMessages()
	return false
}

func (m *ChatBannedRights) ToBannedRights() BannedRights {
	var (
		rights BannedRights = 0
	)

	if m.GetViewMessages() {
		rights |= BANNED_VIEW_MESSAGES
	}
	if m.GetSendMessages() {
		rights |= BANNED_SEND_MESSAGES
	}
	if m.GetSendMedia() {
		rights |= BANNED_SEND_MEDIA
	}
	if m.GetSendStickers() {
		rights |= BANNED_SEND_STICKERS
	}
	if m.GetSendGifs() {
		rights |= BANNED_SEND_GIFS
	}
	if m.GetSendGames() {
		rights |= BANNED_SEND_GAMES
	}
	if m.GetSendInline() {
		rights |= BANNED_SEND_INLINE
	}
	if m.GetEmbedLinks() {
		rights |= BANNED_EMBED_LINKS
	}
	if m.GetSendPolls() {
		rights |= BANNED_SEND_POLLS
	}
	if m.GetChangeInfo() {
		rights |= BANNED_CHANGE_INFO
	}
	if m.GetInviteUsers() {
		rights |= BANNED_INVITE_USERS
	}
	if m.GetPinMessages() {
		rights |= BANNED_PIN_MESSAGES
	}

	untilDate := m.GetUntilDate()
	if untilDate == 0 {
		untilDate = math.MaxInt32
	}
	rights = BannedRights(untilDate)<<32 | rights

	return rights

}
