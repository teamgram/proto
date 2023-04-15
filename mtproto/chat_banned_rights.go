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

const (
	BANNED_VIEW_MESSAGES    = 1 << 0
	BANNED_SEND_MESSAGES    = 1 << 1
	BANNED_SEND_MEDIA       = 1 << 2
	BANNED_SEND_STICKERS    = 1 << 3
	BANNED_SEND_GIFS        = 1 << 4
	BANNED_SEND_GAMES       = 1 << 5
	BANNED_SEND_INLINE      = 1 << 6
	BANNED_EMBED_LINKS      = 1 << 7
	BANNED_SEND_POLLS       = 1 << 8
	BANNED_CHANGE_INFO      = 1 << 10
	BANNED_INVITE_USERS     = 1 << 15
	BANNED_PIN_MESSAGES     = 1 << 17
	BANNED_MANAGE_TOPICS    = 1 << 18
	BANNED_SEND_PHOTOS      = 1 << 19
	BANNED_SEND_VIDEOS      = 1 << 20
	BANNED_SEND_ROUNDVIDEOS = 1 << 21
	BANNED_SEND_AUDIOS      = 1 << 22
	BANNED_SEND_VOICES      = 1 << 23
	BANNED_SEND_DOCS        = 1 << 24
	BANNED_SEND_PLAIN       = 1 << 25
)

type (
	BannedRights int64
)

func MakeDefaultBannedRights() *ChatBannedRights {
	return MakeTLChatBannedRights(&ChatBannedRights{
		ViewMessages:    false,
		SendMessages:    false,
		SendMedia:       false,
		SendStickers:    false,
		SendGifs:        false,
		SendGames:       false,
		SendInline:      false,
		EmbedLinks:      false,
		SendPolls:       false,
		ChangeInfo:      false,
		InviteUsers:     false,
		PinMessages:     false,
		ManageTopics:    false,
		SendPhotos:      false,
		SendVideos:      false,
		SendRoundvideos: false,
		SendAudios:      false,
		SendVoices:      false,
		SendDocs:        false,
		SendPlain:       false,
		UntilDate:       math.MaxInt32,
	}).To_ChatBannedRights()
}

func (m BannedRights) ToChatBannedRights() *ChatBannedRights {
	if m == 0 {
		return nil
	}

	return MakeTLChatBannedRights(&ChatBannedRights{
		ViewMessages:    (m & BANNED_VIEW_MESSAGES) != 0,
		SendMessages:    (m & BANNED_SEND_MESSAGES) != 0,
		SendMedia:       (m & BANNED_SEND_MEDIA) != 0,
		SendStickers:    (m & BANNED_SEND_STICKERS) != 0,
		SendGifs:        (m & BANNED_SEND_GIFS) != 0,
		SendGames:       (m & BANNED_SEND_GAMES) != 0,
		SendInline:      (m & BANNED_SEND_INLINE) != 0,
		EmbedLinks:      (m & BANNED_EMBED_LINKS) != 0,
		SendPolls:       (m & BANNED_SEND_POLLS) != 0,
		ChangeInfo:      (m & BANNED_CHANGE_INFO) != 0,
		InviteUsers:     (m & BANNED_INVITE_USERS) != 0,
		PinMessages:     (m & BANNED_PIN_MESSAGES) != 0,
		ManageTopics:    (m & BANNED_MANAGE_TOPICS) != 0,
		SendPhotos:      (m & BANNED_SEND_PHOTOS) != 0,
		SendVideos:      (m & BANNED_SEND_VIDEOS) != 0,
		SendRoundvideos: (m & BANNED_SEND_ROUNDVIDEOS) != 0,
		SendAudios:      (m & BANNED_SEND_AUDIOS) != 0,
		SendVoices:      (m & BANNED_SEND_VOICES) != 0,
		SendDocs:        (m & BANNED_SEND_DOCS) != 0,
		SendPlain:       (m & BANNED_SEND_PLAIN) != 0,
		UntilDate:       int32(m >> 32),
	}).To_ChatBannedRights()
}

func (m *ChatBannedRights) hasRights() bool {
	return m.ViewMessages ||
		m.SendMessages ||
		m.SendMedia ||
		m.SendStickers ||
		m.SendGifs ||
		m.SendGames ||
		m.SendInline ||
		m.EmbedLinks ||
		m.SendPolls ||
		m.ChangeInfo ||
		m.InviteUsers ||
		m.PinMessages ||
		m.ManageTopics ||
		m.SendPhotos ||
		m.SendVideos ||
		m.SendRoundvideos ||
		m.SendAudios ||
		m.SendVoices ||
		m.SendDocs ||
		m.SendPlain
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

func (m *ChatBannedRights) CanManageTopics(date int32) bool {
	return !m.GetManageTopics() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendPhotos(date int32) bool {
	return !m.GetSendPhotos() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendVideos(date int32) bool {
	return !m.GetSendVideos() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendRoundvideos(date int32) bool {
	return !m.GetSendRoundvideos() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendAudios(date int32) bool {
	return !m.GetSendAudios() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendVoices(date int32) bool {
	return !m.GetSendVoices() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendDocs(date int32) bool {
	return !m.GetSendDocs() || date >= m.UntilDate
}

func (m *ChatBannedRights) CanSendPlain(date int32) bool {
	return !m.GetSendPlain() || date >= m.UntilDate
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
	if m.GetManageTopics() {
		rights |= BANNED_MANAGE_TOPICS
	}
	if m.GetSendPhotos() {
		rights |= BANNED_SEND_PHOTOS
	}
	if m.GetSendVideos() {
		rights |= BANNED_SEND_VIDEOS
	}
	if m.GetSendRoundvideos() {
		rights |= BANNED_SEND_ROUNDVIDEOS
	}
	if m.GetSendAudios() {
		rights |= BANNED_SEND_AUDIOS
	}
	if m.GetSendVoices() {
		rights |= BANNED_SEND_VOICES
	}
	if m.GetSendDocs() {
		rights |= BANNED_SEND_DOCS
	}
	if m.GetSendPlain() {
		rights |= BANNED_SEND_PLAIN
	}

	untilDate := m.GetUntilDate()
	if untilDate == 0 {
		untilDate = math.MaxInt32
	}
	rights = BannedRights(untilDate)<<32 | rights

	return rights

}
