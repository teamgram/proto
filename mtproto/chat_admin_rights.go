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

const (
	ADMIN_CHANGE_INFO     int32 = 1 << 0
	ADMIN_POST_MESSAGES   int32 = 1 << 1
	ADMIN_EDIT_MESSAGES   int32 = 1 << 2
	ADMIN_DELETE_MESSAGES int32 = 1 << 3
	ADMIN_BAN_USERS       int32 = 1 << 4
	ADMIN_INVITE_USERS    int32 = 1 << 5
	ADMIN_PIN_MESSAGES    int32 = 1 << 7
	ADMIN_ADD_ADMINS      int32 = 1 << 9
	ADMIN_ANONYMOUS       int32 = 1 << 10
	ADMIN_MANAGE_CALL     int32 = 1 << 11
	ADMIN_OTHER           int32 = 1 << 12
	ADMIN_POST_STORIES    int32 = 1 << 14
	ADMIN_EDIT_STORIES    int32 = 1 << 15
	ADMIN_DELETE_STORIES  int32 = 1 << 16
)

type AdminRights int32

func MakeChatAdminRightsHelper(adminRights *ChatAdminRights) AdminRights {
	var rights = int32(0)

	if adminRights.GetChangeInfo() {
		rights |= ADMIN_CHANGE_INFO
	}
	if adminRights.GetPostMessages() {
		rights |= ADMIN_POST_MESSAGES
	}
	if adminRights.GetEditMessages() {
		rights |= ADMIN_EDIT_MESSAGES
	}
	if adminRights.GetDeleteMessages() {
		rights |= ADMIN_DELETE_MESSAGES
	}
	if adminRights.GetBanUsers() {
		rights |= ADMIN_BAN_USERS
	}
	if adminRights.GetInviteUsers() {
		rights |= ADMIN_INVITE_USERS
	}
	if adminRights.GetPinMessages() {
		rights |= ADMIN_PIN_MESSAGES
	}
	if adminRights.GetAddAdmins() {
		rights |= ADMIN_ADD_ADMINS
	}
	if adminRights.GetAnonymous() {
		rights |= ADMIN_ANONYMOUS
	}
	if adminRights.GetManageCall() {
		rights |= ADMIN_MANAGE_CALL
	}
	if adminRights.GetOther() {
		rights |= ADMIN_OTHER
	}
	if adminRights.GetManageTopics() {
		rights |= ADMIN_MANAGE_TOPICS
	}
	if adminRights.GetPostStories() {
		rights |= ADMIN_POST_STORIES
	}
	if adminRights.GetEditStories() {
		rights |= ADMIN_EDIT_STORIES
	}
	if adminRights.GetDeleteStories() {
		rights |= ADMIN_DELETE_STORIES
	}

	return AdminRights(rights)
}

func (m AdminRights) ToChatAdminRights() *ChatAdminRights {
	if int32(m) == 0 {
		return nil
	} else {
		return MakeTLChatAdminRights(&ChatAdminRights{
			ChangeInfo:     int32(m)&ADMIN_CHANGE_INFO != 0,
			PostMessages:   int32(m)&ADMIN_POST_MESSAGES != 0,
			EditMessages:   int32(m)&ADMIN_EDIT_MESSAGES != 0,
			DeleteMessages: int32(m)&ADMIN_DELETE_MESSAGES != 0,
			BanUsers:       int32(m)&ADMIN_BAN_USERS != 0,
			InviteUsers:    int32(m)&ADMIN_INVITE_USERS != 0,
			PinMessages:    int32(m)&ADMIN_PIN_MESSAGES != 0,
			AddAdmins:      int32(m)&ADMIN_ADD_ADMINS != 0,
			Anonymous:      int32(m)&ADMIN_ANONYMOUS != 0,
			ManageCall:     int32(m)&ADMIN_MANAGE_CALL != 0,
			Other:          int32(m)&ADMIN_OTHER != 0,
			ManageTopics:   int32(m)&ADMIN_MANAGE_TOPICS != 0,
			PostStories:    int32(m)&ADMIN_POST_STORIES != 0,
			EditStories:    int32(m)&ADMIN_EDIT_STORIES != 0,
			DeleteStories:  int32(m)&ADMIN_DELETE_STORIES != 0,
		}).To_ChatAdminRights()
	}
}

func (m AdminRights) HasAdminRights() bool {
	return m != 0
}

func (m AdminRights) CanChangeInfo() bool {
	return int32(m)&ADMIN_CHANGE_INFO != 0
}

func (m AdminRights) CanPostMessages() bool {
	return int32(m)&ADMIN_POST_MESSAGES != 0
}

func (m AdminRights) CanEditMessages() bool {
	return int32(m)&ADMIN_EDIT_MESSAGES != 0
}

func (m AdminRights) CanDeleteMessages() bool {
	return int32(m)&ADMIN_DELETE_MESSAGES != 0
}

func (m AdminRights) CanBanUsers() bool {
	return int32(m)&ADMIN_BAN_USERS != 0
}

func (m AdminRights) CanInviteUsers() bool {
	return int32(m)&ADMIN_INVITE_USERS != 0
}

func (m AdminRights) CanPinMessages() bool {
	return int32(m)&ADMIN_PIN_MESSAGES != 0
}

func (m AdminRights) CanAddAdmins() bool {
	return int32(m)&ADMIN_ADD_ADMINS != 0
}

func (m AdminRights) CanAnonymous() bool {
	return int32(m)&ADMIN_ANONYMOUS != 0
}

func (m AdminRights) CanManageCall() bool {
	return int32(m)&ADMIN_MANAGE_CALL != 0
}

func (m AdminRights) CanOther() bool {
	return int32(m)&ADMIN_OTHER != 0
}

func (m AdminRights) CanManageTopics() bool {
	return int32(m)&ADMIN_MANAGE_TOPICS != 0
}

func (m AdminRights) CanPostStories() bool {
	return int32(m)&ADMIN_POST_STORIES != 0
}

func (m AdminRights) CanEditStories() bool {
	return int32(m)&ADMIN_EDIT_STORIES != 0
}

func (m AdminRights) CanDeleteStories() bool {
	return int32(m)&ADMIN_DELETE_STORIES != 0
}

// DisallowMegagroup
// ////////////////////////////////////////////////////////////////////////////////////////
func (m AdminRights) DisallowMegagroup() bool {
	return m.CanPostMessages() || m.CanEditMessages()
}

func (m AdminRights) DisallowChat() bool {
	return m.CanPostMessages() || m.CanEditMessages()
}

func MakeDefaultChatAdminRights() *ChatAdminRights {
	return MakeTLChatAdminRights(&ChatAdminRights{
		ChangeInfo:     true,
		PostMessages:   false,
		EditMessages:   false,
		DeleteMessages: true,
		BanUsers:       true,
		InviteUsers:    true,
		PinMessages:    true,
		AddAdmins:      false,
		Anonymous:      false,
		ManageCall:     true,
		Other:          true,
		ManageTopics:   false,
		PostStories:    false,
		EditStories:    false,
		DeleteStories:  false,
	}).To_ChatAdminRights()
}

func (m *ChatAdminRights) HasAdminRights() bool {
	return m != nil
}

func (m *ChatAdminRights) CanChangeInfo() bool {
	return m.GetChangeInfo()
}

func (m *ChatAdminRights) CanPostMessages() bool {
	return m.GetPostMessages()
}

func (m *ChatAdminRights) CanEditMessages() bool {
	return m.GetEditMessages()
}

func (m *ChatAdminRights) CanDeleteMessages() bool {
	return m.GetDeleteMessages()
}

func (m *ChatAdminRights) CanBanUsers() bool {
	return m.GetBanUsers()
}

func (m *ChatAdminRights) CanInviteUsers() bool {
	return m.GetInviteUsers()
}

func (m *ChatAdminRights) CanPinMessages() bool {
	return m.GetPinMessages()
}

func (m *ChatAdminRights) CanAddAdmins() bool {
	return m.GetAddAdmins()
}

func (m *ChatAdminRights) CanAnonymous() bool {
	return m.GetAnonymous()
}

func (m *ChatAdminRights) CanManageCall() bool {
	return m.GetManageCall()
}

func (m *ChatAdminRights) CanOther() bool {
	return m.GetOther()
}

func (m *ChatAdminRights) CanManageTopics() bool {
	return m.GetManageTopics()
}

func (m *ChatAdminRights) CanPostStories() bool {
	return m.GetPostStories()
}

func (m *ChatAdminRights) CanEditStories() bool {
	return m.GetEditStories()
}

func (m *ChatAdminRights) CanDeleteStories() bool {
	return m.GetDeleteStories()
}

// DisallowMegagroup
// ////////////////////////////////////////////////////////////////////////////////////////
func (m *ChatAdminRights) DisallowMegagroup() bool {
	return m.CanPostMessages() || m.CanEditMessages()
}

func (m *ChatAdminRights) DisallowChat() bool {
	return m.CanPostMessages() || m.CanEditMessages()
}
