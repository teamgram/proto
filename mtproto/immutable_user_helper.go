// Copyright (c) 2023-present,  Teamgram Studio (https://teamgram.io).
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
	"github.com/golang/protobuf/ptypes/wrappers"
)

func (m *ImmutableUser) Id() int64 {
	return m.User.Id
}

func (m *ImmutableUser) AccessHash() int64 {
	return m.User.AccessHash
}

func (m *ImmutableUser) FirstName() string {
	return m.User.FirstName
}

func (m *ImmutableUser) LastName() string {
	return m.User.LastName
}

func (m *ImmutableUser) SetFirstName(v string) {
	m.User.FirstName = v
}

func (m *ImmutableUser) SetLastName(v string) {
	m.User.LastName = v
}

func (m *ImmutableUser) Username() string {
	return m.User.Username
}

func (m *ImmutableUser) SetUsername(v string) {
	m.User.Username = v
}

func (m *ImmutableUser) Phone() string {
	return m.User.Phone
}

func (m *ImmutableUser) Deleted() bool {
	return m.User.Deleted
}

func (m *ImmutableUser) Restricted() bool {
	return m.User.Restricted
}

func (m *ImmutableUser) ProfilePhoto() *UserProfilePhoto {
	return MakeUserProfilePhotoByPhoto(m.User.ProfilePhoto)
}

func (m *ImmutableUser) Photo() *Photo {
	return m.User.ProfilePhoto
}

func (m *ImmutableUser) About() string {
	return m.User.About.GetValue()
}

func (m *ImmutableUser) SetAbout(v string) {
	m.User.About = MakeFlagsString(v)
}

func (m *ImmutableUser) IsBot() bool {
	return m.User.Bot != nil
}

func (m *ImmutableUser) BotChatHistory() bool {
	return m.User.Bot.GetBotChatHistory()
}

func (m *ImmutableUser) BotNochats() bool {
	return m.User.Bot.GetBotNochats()
}

func (m *ImmutableUser) BotInlineGeo() bool {
	return m.User.Bot.GetBotInlineGeo()
}

func (m *ImmutableUser) RestrictionReason() []*RestrictionReason {
	return m.User.RestrictionReason
}

func (m *ImmutableUser) BotInlinePlaceholder() *wrappers.StringValue {
	return m.User.Bot.GetBotInlinePlaceholder()
}

func (m *ImmutableUser) Verified() bool {
	return m.User.Verified
}

func (m *ImmutableUser) SetVerified(v bool) {
	m.User.Verified = v
}

func (m *ImmutableUser) Support() bool {
	return m.User.Support
}

func (m *ImmutableUser) Scam() bool {
	return m.User.Scam
}

func (m *ImmutableUser) Fake() bool {
	return m.User.Fake
}

func (m *ImmutableUser) BotInfoVersion() int32 {
	return m.User.Bot.GetBotInfoVersion()
}

func (m *ImmutableUser) BotAttachMenu() bool {
	return m.User.Bot.GetBotAttachMenu()
}

func (m *ImmutableUser) Premium() bool {
	return m.User.Premium
}

func (m *ImmutableUser) AttachMenuEnabled() bool {
	return m.User.Bot.GetBotAttachMenu()
}

func (m *ImmutableUser) BotCanEdit() bool {
	return m.User.Bot.GetBotCanEdit()
}

func (m *ImmutableUser) CheckStoriesHidden(id int64) bool {
	for _, v := range m.StoriesHiddens {
		if id == v {
			return true
		}
	}

	return false
}

func (m *ImmutableUser) StoriesUnavailable() bool {
	return m.User.StoriesUnavailable
}

func (m *ImmutableUser) EmojiStatus() *EmojiStatus {
	return m.User.EmojiStatus
}

func (m *ImmutableUser) StoriesMaxId() *wrappers.Int32Value {
	return MakeFlagsInt32(m.User.StoriesMaxId)
}

func (m *ImmutableUser) Color() *PeerColor {
	return m.User.Color
}

func (m *ImmutableUser) ProfileColor() *PeerColor {
	return m.User.ProfileColor
}

func (m *ImmutableUser) Birthday() *Birthday {
	return FromBirthdayString(m.User.Birthday)
}

func (m *ImmutableUser) CheckContact(cId int64) (bool, bool) {
	// TODO: sort.Search
	for _, c := range m.Contacts {
		if cId == c.ContactUserId {
			return true, c.MutualContact
		}
	}

	return false, false
}

func (m *ImmutableUser) GetContactData(cId int64) *ContactData {
	// TODO: sort.Search
	for _, c := range m.Contacts {
		if cId == c.ContactUserId {
			return c
		}
	}

	return nil
}

func (m *ImmutableUser) CheckReverseContact(cId int64) (bool, bool) {
	// TODO: sort.Search
	for _, c := range m.ReverseContacts {
		if cId == c.UserId {
			return true, c.MutualContact
		}
	}

	return false, false
}

func (m *ImmutableUser) GetReverseContactData(cId int64) *ContactData {
	// TODO: sort.Search
	for _, c := range m.ReverseContacts {
		if cId == c.UserId {
			return c
		}
	}

	return nil
}

func (m *ImmutableUser) CheckPrivacy(keyType int, id int64) bool {
	var (
		rules *PrivacyKeyRules
	)

	if m.Id() == id {
		return true
	}

	for _, v := range m.KeysPrivacyRules {
		if v.Key == int32(keyType) {
			rules = v
			break
		}
	}

	if rules == nil {
		return true
	}

	isContact, _ := m.CheckContact(id)
	allow := privacyIsAllow(rules.Rules, id, isContact)

	return allow
}

func (m *ImmutableUser) ToUnsafeUser(selfUser *ImmutableUser) *User {
	if m.Deleted() {
		return m.ToDeletedUser()
	}

	if m.Id() == selfUser.Id() {
		return m.ToSelfUser()
	}

	user := MakeTLUser(&User{
		Self:                 false,
		Contact:              false,
		MutualContact:        false,
		Deleted:              false,
		Bot:                  m.IsBot(),
		BotChatHistory:       m.BotChatHistory(),
		BotNochats:           m.BotNochats(),
		Verified:             m.Verified(),
		Restricted:           m.Restricted(),
		Min:                  false,
		BotInlineGeo:         m.BotInlineGeo(),
		Support:              m.Support(),
		Scam:                 m.Scam(),
		ApplyMinPhoto:        false,
		Fake:                 m.Fake(),
		BotAttachMenu:        m.BotAttachMenu(),
		Premium:              m.Premium(),
		AttachMenuEnabled:    m.AttachMenuEnabled(),
		BotCanEdit:           m.BotCanEdit(),
		CloseFriend:          false,
		StoriesHidden:        false,
		StoriesUnavailable:   m.StoriesUnavailable(),
		Id:                   m.Id(),
		AccessHash:           MakeFlagsInt64(m.AccessHash()),
		FirstName:            MakeFlagsString(m.FirstName()),
		LastName:             MakeFlagsString(m.LastName()),
		Username:             MakeFlagsString(m.Username()),
		Phone:                nil,
		Photo:                nil,
		Status:               MakeUserStatus(m.LastSeenAt, true),
		BotInfoVersion:       MakeFlagsInt32(m.BotInfoVersion()),
		RestrictionReason:    m.RestrictionReason(),
		BotInlinePlaceholder: m.BotInlinePlaceholder(),
		LangCode:             nil,
		EmojiStatus:          m.EmojiStatus(),
		Usernames:            nil,
		StoriesMaxId:         m.StoriesMaxId(),
		Color_FLAGPEERCOLOR:  m.Color(),
		Color:                m.Color().GetColor(),
		Color_FLAGINT32:      m.Color().GetColor(),
		BackgroundEmojiId:    m.Color().GetBackgroundEmojiId(),
		ProfileColor:         m.ProfileColor(),
	}).To_User()

	if m.IsBot() {
		return user
	}

	// contact
	contact := selfUser.GetContactData(m.Id())
	if contact != nil {
		user.Contact = true
		user.MutualContact = contact.MutualContact
		user.FirstName = contact.FirstName
		user.LastName = contact.LastName
		user.Phone = contact.Phone
		user.CloseFriend = contact.CloseFriend
	} else {
		// reverse contact
		reverseContact := m.GetReverseContactData(selfUser.Id())
		if reverseContact != nil {
			user.Contact = true
			user.MutualContact = reverseContact.MutualContact
			user.FirstName = reverseContact.FirstName
			user.LastName = reverseContact.LastName
			user.Phone = reverseContact.Phone
			user.CloseFriend = reverseContact.CloseFriend
		}
	}

	// phone
	if m.CheckPrivacy(PHONE_NUMBER, selfUser.Id()) {
		user.Phone = MakeFlagsString(m.Phone())
	}

	// photo
	if m.CheckPrivacy(PROFILE_PHOTO, selfUser.Id()) {
		user.Photo = m.ProfilePhoto()
	}

	// status
	allowTimestamp := selfUser.CheckPrivacy(STATUS_TIMESTAMP, m.Id())
	if allowTimestamp {
		allowTimestamp = m.CheckPrivacy(STATUS_TIMESTAMP, selfUser.Id())
	}
	user.Status = MakeUserStatus(m.LastSeenAt, allowTimestamp)

	// TODO
	// StoriesHidden
	for _, v := range selfUser.StoriesHiddens {
		if v == m.Id() {
			user.StoriesHidden = true
			break
		}
	}

	return user
}

func (m *ImmutableUser) ToSelfUser() *User {
	return MakeTLUser(&User{
		Self:                 true,
		Contact:              true,
		MutualContact:        true,
		Deleted:              false,
		Bot:                  m.IsBot(),
		BotChatHistory:       m.BotChatHistory(),
		BotNochats:           m.BotNochats(),
		Verified:             m.Verified(),
		Restricted:           m.Restricted(),
		Min:                  false,
		BotInlineGeo:         m.BotInlineGeo(),
		Support:              m.Support(),
		Scam:                 m.Scam(),
		ApplyMinPhoto:        false,
		Fake:                 m.Fake(),
		BotAttachMenu:        m.BotAttachMenu(),
		Premium:              m.Premium(),
		AttachMenuEnabled:    m.AttachMenuEnabled(),
		BotCanEdit:           m.BotCanEdit(),
		CloseFriend:          false,
		StoriesHidden:        false,
		StoriesUnavailable:   false,
		Id:                   m.Id(),
		AccessHash:           MakeFlagsInt64(m.AccessHash()),
		FirstName:            MakeFlagsString(m.FirstName()),
		LastName:             MakeFlagsString(m.LastName()),
		Username:             MakeFlagsString(m.Username()),
		Phone:                MakeFlagsString(m.Phone()),
		Photo:                m.ProfilePhoto(),
		Status:               MakeUserStatus(m.LastSeenAt, true),
		BotInfoVersion:       MakeFlagsInt32(m.BotInfoVersion()),
		RestrictionReason:    m.RestrictionReason(),
		BotInlinePlaceholder: m.BotInlinePlaceholder(),
		LangCode:             nil,
		EmojiStatus:          m.EmojiStatus(),
		Usernames:            nil,
		StoriesMaxId:         m.StoriesMaxId(),
		Color_FLAGPEERCOLOR:  m.Color(),
		Color:                m.Color().GetColor(),
		Color_FLAGINT32:      m.Color().GetColor(),
		BackgroundEmojiId:    m.Color().GetBackgroundEmojiId(),
		ProfileColor:         m.ProfileColor(),
	}).To_User()
}

func (m *ImmutableUser) ToDeletedUser() *User {
	return MakeTLUser(&User{
		Self:                 false,
		Contact:              false,
		MutualContact:        false,
		Deleted:              true,
		Bot:                  false,
		BotChatHistory:       false,
		BotNochats:           false,
		Verified:             false,
		Restricted:           false,
		Min:                  false,
		BotInlineGeo:         false,
		Support:              false,
		Scam:                 false,
		ApplyMinPhoto:        false,
		Fake:                 false,
		BotAttachMenu:        false,
		Premium:              false,
		AttachMenuEnabled:    false,
		BotCanEdit:           false,
		CloseFriend:          false,
		StoriesHidden:        false,
		StoriesUnavailable:   false,
		Id:                   m.Id(),
		AccessHash:           MakeFlagsInt64(m.AccessHash()),
		FirstName:            nil,
		LastName:             nil,
		Username:             nil,
		Phone:                nil,
		Photo:                nil,
		Status:               nil,
		BotInfoVersion:       nil,
		RestrictionReason:    nil,
		BotInlinePlaceholder: nil,
		LangCode:             nil,
		EmojiStatus:          nil,
		Usernames:            nil,
		StoriesMaxId:         nil,
		Color_FLAGPEERCOLOR:  nil,
		Color:                nil,
		Color_FLAGINT32:      nil,
		BackgroundEmojiId:    nil,
		ProfileColor:         nil,
	}).To_User()
}

func (m *ImmutableUser) ToUser(selfUserId int64) *User {
	if m.Deleted() {
		return m.ToDeletedUser()
	}

	if m.Id() == selfUserId {
		return m.ToSelfUser()
	}

	user := MakeTLUser(&User{
		Self:                 false,
		Contact:              false,
		MutualContact:        false,
		Deleted:              false,
		Bot:                  m.IsBot(),
		BotChatHistory:       m.BotChatHistory(),
		BotNochats:           m.BotNochats(),
		Verified:             m.Verified(),
		Restricted:           m.Restricted(),
		Min:                  false,
		BotInlineGeo:         m.BotInlineGeo(),
		Support:              m.Support(),
		Scam:                 m.Scam(),
		ApplyMinPhoto:        false,
		Fake:                 m.Fake(),
		BotAttachMenu:        m.BotAttachMenu(),
		Premium:              m.Premium(),
		AttachMenuEnabled:    m.AttachMenuEnabled(),
		BotCanEdit:           m.BotCanEdit(),
		CloseFriend:          false,
		StoriesHidden:        false,
		StoriesUnavailable:   m.StoriesUnavailable(),
		Id:                   m.Id(),
		AccessHash:           MakeFlagsInt64(m.AccessHash()),
		FirstName:            MakeFlagsString(m.FirstName()),
		LastName:             MakeFlagsString(m.LastName()),
		Username:             MakeFlagsString(m.Username()),
		Phone:                nil,
		Photo:                nil,
		Status:               MakeUserStatus(m.LastSeenAt, false),
		BotInfoVersion:       MakeFlagsInt32(m.BotInfoVersion()),
		RestrictionReason:    m.RestrictionReason(),
		BotInlinePlaceholder: m.BotInlinePlaceholder(),
		LangCode:             nil,
		EmojiStatus:          m.EmojiStatus(),
		Usernames:            nil,
		StoriesMaxId:         m.StoriesMaxId(),
		Color_FLAGPEERCOLOR:  m.Color(),
		Color:                m.Color().GetColor(),
		Color_FLAGINT32:      m.Color().GetColor(),
		BackgroundEmojiId:    m.Color().GetBackgroundEmojiId(),
		ProfileColor:         m.ProfileColor(),
	}).To_User()

	if m.IsBot() {
		return user
	}

	reverseContact := m.GetReverseContactData(selfUserId)
	if reverseContact != nil {
		user.Contact = true
		user.MutualContact = reverseContact.MutualContact
		user.FirstName = reverseContact.FirstName
		user.LastName = reverseContact.LastName
		user.Phone = reverseContact.Phone
		user.CloseFriend = reverseContact.CloseFriend
	}

	// phone
	if m.CheckPrivacy(PHONE_NUMBER, selfUserId) {
		user.Phone = MakeFlagsString(m.Phone())
	}

	// photo
	if m.CheckPrivacy(PROFILE_PHOTO, selfUserId) {
		user.Photo = m.ProfilePhoto()
	}

	// TODO: check if we need to check privacy for status
	// status
	allowTimestamp := m.CheckPrivacy(STATUS_TIMESTAMP, selfUserId)
	user.Status = MakeUserStatus(m.LastSeenAt, allowTimestamp)

	// TODO
	// StoriesHidden

	return user
}
