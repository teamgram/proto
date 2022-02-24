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
	VIEW_MESSAGES = 1 << 0
	SEND_MESSAGES = 1 << 1
	SEND_MEDIA    = 1 << 2
	SEND_STICKERS = 1 << 3
	SEND_GIFS     = 1 << 4
	SEND_GAMES    = 1 << 5
	SEND_INLINE   = 1 << 6
	EMBED_LINKS   = 1 << 7
	SEND_POLLS    = 1 << 8
	CHANGE_INFO2  = 1 << 10
	INVITE_USERS2 = 1 << 15
	PIN_MESSAGE   = 1 << 17
)

//const (
//	noBannedRights = 1 << 31
//)

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

//func MakeBannedRights(bannedRights *ChatBannedRights) BannedRights {
//	var (
//		rights BannedRights = 0
//	)
//
//	if bannedRights.GetViewMessages() {
//		rights |= VIEW_MESSAGES
//	}
//	if bannedRights.GetSendMessages() {
//		rights |= SEND_MESSAGES
//	}
//	if bannedRights.GetSendMedia() {
//		rights |= SEND_MEDIA
//	}
//	if bannedRights.GetSendStickers() {
//		rights |= SEND_STICKERS
//	}
//	if bannedRights.GetSendGifs() {
//		rights |= SEND_GIFS
//	}
//	if bannedRights.GetSendGames() {
//		rights |= SEND_GAMES
//	}
//	if bannedRights.GetSendInline() {
//		rights |= SEND_INLINE
//	}
//	if bannedRights.GetEmbedLinks() {
//		rights |= EMBED_LINKS
//	}
//	if bannedRights.GetSendPolls() {
//		rights |= SEND_POLLS
//	}
//	if bannedRights.GetChangeInfo() {
//		rights |= CHANGE_INFO2
//	}
//	if bannedRights.GetInviteUsers() {
//		rights |= INVITE_USERS2
//	}
//	if bannedRights.GetPinMessages() {
//		rights |= PIN_MESSAGE
//	}
//
//	untilDate := bannedRights.GetUntilDate()
//	if untilDate == 0 {
//		untilDate = math.MaxInt32
//	}
//	rights = BannedRights(untilDate)<<32 | rights
//
//	return rights
//}

func (m BannedRights) ToChatBannedRights() *ChatBannedRights {
	if m == 0 {
		return nil
	}

	return MakeTLChatBannedRights(&ChatBannedRights{
		ViewMessages: (m & VIEW_MESSAGES) != 0,
		SendMessages: (m & SEND_MESSAGES) != 0,
		SendMedia:    (m & SEND_MEDIA) != 0,
		SendStickers: (m & SEND_STICKERS) != 0,
		SendGifs:     (m & SEND_GIFS) != 0,
		SendGames:    (m & SEND_GAMES) != 0,
		SendInline:   (m & SEND_INLINE) != 0,
		EmbedLinks:   (m & EMBED_LINKS) != 0,
		SendPolls:    (m & SEND_POLLS) != 0,
		ChangeInfo:   (m & CHANGE_INFO2) != 0,
		InviteUsers:  (m & INVITE_USERS2) != 0,
		PinMessages:  (m & PIN_MESSAGE) != 0,
		UntilDate:    int32(m >> 32),
	}).To_ChatBannedRights()
}

//func (m BannedRights) NoBanRights() bool {
//	return m.Rights == 0 && m.UntilDate == math.MaxInt32
//}
//
//func (m ChatBannedRightsHelper) IsKick() bool {
//	return m.Rights&VIEW_MESSAGES != 0 && m.UntilDate == math.MaxInt32
//}
//
//func (m ChatBannedRightsHelper) IsBan(date int32) bool {
//	return m.Rights != 0 && date <= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) IsRestrict(date int32) bool {
//	return !m.IsKick() && m.Rights != 0 && date <= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanViewMessages(date int32) bool {
//	return m.Rights&VIEW_MESSAGES == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanSendMessages(date int32) bool {
//	return m.Rights&SEND_MESSAGES == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanSendMedia(date int32) bool {
//	return m.Rights&SEND_MEDIA == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanSendStickers(date int32) bool {
//	return m.Rights&SEND_STICKERS == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanSendGifs(date int32) bool {
//	return m.Rights&SEND_GIFS == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanSendGames(date int32) bool {
//	return m.Rights&SEND_GAMES == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanSendInline(date int32) bool {
//	return m.Rights&SEND_INLINE == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanEmbedLinks(date int32) bool {
//	return m.Rights&EMBED_LINKS == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanSendPolls(date int32) bool {
//	return m.Rights&SEND_POLLS == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanChangeInfo(date int32) bool {
//	return m.Rights&CHANGE_INFO2 == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanInviteUsers(date int32) bool {
//	return m.Rights&INVITE_USERS2 == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) CanPinMessage(date int32) bool {
//	return m.Rights&PIN_MESSAGE == 0 || date >= m.UntilDate
//}
//
//func (m ChatBannedRightsHelper) DisallowChannel() bool {
//	// return m.CanPostMessages() || m.CanEditMessages()
//	return false
//}
//
//func FromChatBannedRights(bannedRights *TLChatBannedRights) (int32, int32) {
//	var (
//		rights    int32 = 0
//		untilDate int32 = 0
//	)
//
//	if bannedRights.GetViewMessages() {
//		rights |= VIEW_MESSAGES
//	}
//	if bannedRights.GetSendMessages() {
//		rights |= SEND_MESSAGES
//	}
//	if bannedRights.GetSendMedia() {
//		rights |= SEND_MEDIA
//	}
//	if bannedRights.GetSendStickers() {
//		rights |= SEND_STICKERS
//	}
//	if bannedRights.GetSendGifs() {
//		rights |= SEND_GIFS
//	}
//	if bannedRights.GetSendGames() {
//		rights |= SEND_GAMES
//	}
//	if bannedRights.GetSendInline() {
//		rights |= SEND_INLINE
//	}
//	if bannedRights.GetEmbedLinks() {
//		rights |= EMBED_LINKS
//	}
//	if bannedRights.GetSendPolls() {
//		rights |= SEND_POLLS
//	}
//	if bannedRights.GetChangeInfo() {
//		rights |= CHANGE_INFO2
//	}
//	if bannedRights.GetInviteUsers() {
//		rights |= INVITE_USERS2
//	}
//	if bannedRights.GetPinMessages() {
//		rights |= PIN_MESSAGE
//	}
//
//	untilDate = bannedRights.GetUntilDate()
//	if rights == 0 {
//		untilDate = 0
//	} else if untilDate == 0 {
//		untilDate = math.MaxInt32
//	}
//
//	return rights, untilDate
//}
//
//func ToChatBannedRights(rights, untilDate int32) *ChatBannedRights {
//	if rights == 0 && untilDate == 0 {
//		return nil
//	}
//
//	bannedRights := MakeTLChatBannedRights(nil)
//
//	if (rights & VIEW_MESSAGES) != 0 {
//		bannedRights.SetViewMessages(true)
//	}
//	if (rights & SEND_MESSAGES) != 0 {
//		bannedRights.SetSendMessages(true)
//	}
//	if (rights & SEND_MEDIA) != 0 {
//		bannedRights.SetSendMedia(true)
//	}
//	if (rights & SEND_STICKERS) != 0 {
//		bannedRights.SetSendStickers(true)
//	}
//	if (rights & SEND_GIFS) != 0 {
//		bannedRights.SetSendGifs(true)
//	}
//	if (rights & SEND_GAMES) != 0 {
//		bannedRights.SetSendGames(true)
//	}
//	if (rights & SEND_INLINE) != 0 {
//		bannedRights.SetSendInline(true)
//	}
//	if (rights & EMBED_LINKS) != 0 {
//		bannedRights.SetEmbedLinks(true)
//	}
//	if (rights & SEND_POLLS) != 0 {
//		bannedRights.SetSendPolls(true)
//	}
//	if (rights & CHANGE_INFO2) != 0 {
//		bannedRights.SetChangeInfo(true)
//	}
//	if (rights & INVITE_USERS2) != 0 {
//		bannedRights.SetInviteUsers(true)
//	}
//	if (rights & PIN_MESSAGE) != 0 {
//		bannedRights.SetPinMessages(true)
//	}
//
//	bannedRights.SetUntilDate(untilDate)
//	return bannedRights.To_ChatBannedRights()
//}

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
		rights |= VIEW_MESSAGES
	}
	if m.GetSendMessages() {
		rights |= SEND_MESSAGES
	}
	if m.GetSendMedia() {
		rights |= SEND_MEDIA
	}
	if m.GetSendStickers() {
		rights |= SEND_STICKERS
	}
	if m.GetSendGifs() {
		rights |= SEND_GIFS
	}
	if m.GetSendGames() {
		rights |= SEND_GAMES
	}
	if m.GetSendInline() {
		rights |= SEND_INLINE
	}
	if m.GetEmbedLinks() {
		rights |= EMBED_LINKS
	}
	if m.GetSendPolls() {
		rights |= SEND_POLLS
	}
	if m.GetChangeInfo() {
		rights |= CHANGE_INFO2
	}
	if m.GetInviteUsers() {
		rights |= INVITE_USERS2
	}
	if m.GetPinMessages() {
		rights |= PIN_MESSAGE
	}

	untilDate := m.GetUntilDate()
	if untilDate == 0 {
		untilDate = math.MaxInt32
	}
	rights = BannedRights(untilDate)<<32 | rights

	return rights

}
