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
	"time"
)

func MakeEmptyUpdates() *Updates {
	return MakeTLUpdates(&Updates{
		Updates: []*Update{},
		Users:   []*User{},
		Chats:   []*Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

func MakeUpdatesByUpdates(updates ...*Update) *Updates {
	return MakeTLUpdates(&Updates{
		Updates: updates,
		Users:   []*User{},
		Chats:   []*Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

func MakeUpdatesByUpdatesUsers(users []*User, updates ...*Update) *Updates {
	if users == nil {
		users = []*User{}
	}
	return MakeTLUpdates(&Updates{
		Updates: updates,
		Users:   users,
		Chats:   []*Chat{},
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

func MakeUpdatesByUpdatesChats(chats []*Chat, updates ...*Update) *Updates {
	if chats == nil {
		chats = []*Chat{}
	}
	return MakeTLUpdates(&Updates{
		Updates: updates,
		Users:   []*User{},
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

func MakeUpdatesByUpdatesUsersChats(users []*User, chats []*Chat, updates ...*Update) *Updates {
	if users == nil {
		users = []*User{}
	}
	if chats == nil {
		chats = []*Chat{}
	}

	return MakeTLUpdates(&Updates{
		Updates: updates,
		Users:   users,
		Chats:   chats,
		Date:    int32(time.Now().Unix()),
		Seq:     0,
	}).To_Updates()
}

func (m *Updates) SetUsers(users []*User) {
	m.Users = users
}

func (m *Updates) AddSafeUser(user *User) {
	if user != nil {
		m.Users = append(m.Users, user)
	}
}

func (m *Updates) PushUser(user ...*User) {
	for _, v := range user {
		if v != nil {
			m.Users = append(m.Users, v)
		}
	}
}

func (m *Updates) PushBackUpdate(update ...*Update) {
	m.Updates = append(m.Updates, update...)
}

func (m *Updates) PushFrontUpdate(update ...*Update) {
	m.Updates = append(update, m.Updates...)
}

func (m *Updates) AddSafeChat(chat *Chat) {
	if chat != nil {
		m.Chats = append(m.Chats, chat)
	}
}

func (m *Updates) PushChat(chat ...*Chat) {
	for _, v := range chat {
		if v != nil {
			m.Chats = append(m.Chats, v)
		}
	}
}

func (m *Updates) PushUsersAndChats(users []*User, chats []*Chat) {
	m.Users = append(m.Users, users...)
	m.Chats = append(m.Chats, chats...)
}

func MakeReplyUpdates(
	cb1 func(idList []int64) []*User,
	cb2 func(idList []int64) []*Chat,
	cb3 func(idList []int64) []*Chat,
	update ...*Update) *Updates {

	var (
		updates  = make([]*Update, 0, len(update))
		idHelper = NewIDListHelper()
	)

	for _, upd := range update {
		switch upd.PredicateName {
		case Predicate_updateNewChannelMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case Predicate_messageActionChannelCreate:
					updates = append(updates, MakeUpdateChannel(upd.Message_MESSAGE.GetPeerId().GetChannelId()))
				}
			}
		}
	}

	// updateList = append(updateList, m.updates...)
	for _, upd := range update {
		switch upd.PredicateName {
		case Predicate_updateNewMessage:
			updates = append(updates, MakeTLUpdateMessageID(&Update{
				Id_INT32: upd.Message_MESSAGE.Id,
				RandomId: upd.RandomId,
			}).To_Update())
			switch upd.Message_MESSAGE.PredicateName {
			case Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case Predicate_messageActionChatCreate:
					chatParticipants := MakeTLChatParticipants(&ChatParticipants{
						ChatId:       upd.Message_MESSAGE.GetPeerId().GetChatId(),
						Participants: make([]*ChatParticipant, 0, len(upd.Message_MESSAGE.Action.Users)),
						Version:      1,
					}).To_ChatParticipants()
					chatParticipants.Participants = append(chatParticipants.Participants,
						MakeTLChatParticipantCreator(&ChatParticipant{
							UserId: upd.Message_MESSAGE.FromId.GetUserId(),
						}).To_ChatParticipant())
					for _, id := range upd.Message_MESSAGE.Action.Users {
						chatParticipants.Participants = append(chatParticipants.Participants,
							MakeTLChatParticipant(&ChatParticipant{
								UserId:    id,
								InviterId: upd.Message_MESSAGE.FromId.GetUserId(),
								Date:      upd.Message_MESSAGE.Date,
							}).To_ChatParticipant())
					}
					updates = append(updates, MakeTLUpdateChatParticipants(&Update{
						Participants_CHATPARTICIPANTS: chatParticipants,
					}).To_Update())
				}
			}
		case Predicate_updateNewChannelMessage:
			updates = append(updates, MakeTLUpdateMessageID(&Update{
				Id_INT32: upd.Message_MESSAGE.Id,
				RandomId: upd.RandomId,
			}).To_Update())
		case Predicate_updateNewScheduledMessage:
			updates = append(updates, MakeTLUpdateMessageID(&Update{
				Id_INT32: upd.Message_MESSAGE.Id,
				RandomId: upd.RandomId,
			}).To_Update())
		}
	}

	for _, upd := range update {
		switch upd.PredicateName {
		case Predicate_updateNewChannelMessage:
			updates = append(updates, MakeTLUpdateReadChannelInbox(&Update{
				FolderId:         nil,
				ChannelId:        upd.Message_MESSAGE.GetPeerId().GetChannelId(),
				MaxId:            upd.Message_MESSAGE.Id,
				StillUnreadCount: 0,
				Pts_INT32:        upd.Pts_INT32,
			}).To_Update())
		}
		updates = append(updates, upd)
	}

	rUpdates := MakeUpdatesByUpdates(updates...)

	idHelper.PickByUpdates(update...)
	if cb1 != nil {
		rUpdates.PushUser(cb1(idHelper.UserIdList)...)
	}
	if cb2 != nil {
		rUpdates.PushChat(cb2(idHelper.ChatIdList)...)
	}
	if cb3 != nil {
		rUpdates.PushChat(cb3(idHelper.ChannelIdList)...)
	}

	return rUpdates
}

func MakeSyncNotMeUpdates(
	cb1 func(idList []int64) []*User,
	cb2 func(idList []int64) []*Chat,
	cb3 func(idList []int64) []*Chat,
	update ...*Update) *Updates {

	var (
		updates  = make([]*Update, 0, len(update))
		idHelper = NewIDListHelper()
	)

	for _, upd := range update {
		switch upd.PredicateName {
		case Predicate_updateNewMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case Predicate_messageActionChatCreate:
					chatParticipants := MakeTLChatParticipants(&ChatParticipants{
						ChatId:       upd.Message_MESSAGE.GetPeerId().GetChatId(),
						Participants: make([]*ChatParticipant, 0, len(upd.Message_MESSAGE.Action.Users)),
						Version:      1,
					}).To_ChatParticipants()
					chatParticipants.Participants = append(chatParticipants.Participants,
						MakeTLChatParticipantCreator(&ChatParticipant{
							UserId: upd.Message_MESSAGE.FromId.GetUserId(),
						}).To_ChatParticipant())
					for _, id := range upd.Message_MESSAGE.Action.Users {
						chatParticipants.Participants = append(chatParticipants.Participants,
							MakeTLChatParticipant(&ChatParticipant{
								UserId:    id,
								InviterId: upd.Message_MESSAGE.FromId.GetUserId(),
								Date:      upd.Message_MESSAGE.Date,
							}).To_ChatParticipant())
					}
					updates = append(updates, MakeTLUpdateChatParticipants(&Update{
						Participants_CHATPARTICIPANTS: chatParticipants,
					}).To_Update())
				}
			}
		case Predicate_updateNewChannelMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case Predicate_messageActionChannelCreate:
					updates = append(updates, MakeUpdateChannel(upd.Message_MESSAGE.GetPeerId().GetChannelId()))
				}
			}
		}
	}

	for _, upd := range update {
		switch upd.PredicateName {
		case Predicate_updateNewChannelMessage:
			updates = append(updates, MakeTLUpdateReadChannelInbox(&Update{
				FolderId:         nil,
				ChannelId:        upd.Message_MESSAGE.GetPeerId().GetChannelId(),
				MaxId:            upd.Message_MESSAGE.Id,
				StillUnreadCount: 0,
				Pts_INT32:        upd.Pts_INT32,
			}).To_Update())
		}
		updates = append(updates, upd)
	}

	rUpdates := MakeUpdatesByUpdates(updates...)

	idHelper.PickByUpdates(update...)
	if cb1 != nil {
		rUpdates.PushUser(cb1(idHelper.UserIdList)...)
	}
	if cb2 != nil {
		rUpdates.PushChat(cb2(idHelper.ChatIdList)...)
	}
	if cb3 != nil {
		rUpdates.PushChat(cb3(idHelper.ChannelIdList)...)
	}

	return rUpdates
}

func MakePushUpdates(
	cb1 func(idList []int64) []*User,
	cb2 func(idList []int64) []*Chat,
	cb3 func(idList []int64) []*Chat,
	update ...*Update) *Updates {

	var (
		updates  = make([]*Update, 0, len(update))
		idHelper = NewIDListHelper()
	)

	for _, upd := range update {
		switch upd.PredicateName {
		case Predicate_updateNewMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case Predicate_messageActionChatCreate:
					chatParticipants := MakeTLChatParticipants(&ChatParticipants{
						ChatId:       upd.Message_MESSAGE.GetPeerId().GetChatId(),
						Participants: make([]*ChatParticipant, 0, len(upd.Message_MESSAGE.Action.Users)),
						Version:      1,
					}).To_ChatParticipants()
					chatParticipants.Participants = append(chatParticipants.Participants,
						MakeTLChatParticipantCreator(&ChatParticipant{
							UserId: upd.Message_MESSAGE.FromId.GetUserId(),
						}).To_ChatParticipant())
					for _, id := range upd.Message_MESSAGE.Action.Users {
						chatParticipants.Participants = append(chatParticipants.Participants,
							MakeTLChatParticipant(&ChatParticipant{
								UserId:    id,
								InviterId: upd.Message_MESSAGE.FromId.GetUserId(),
								Date:      upd.Message_MESSAGE.Date,
							}).To_ChatParticipant())
					}
					updates = append(updates, MakeTLUpdateChatParticipants(&Update{
						Participants_CHATPARTICIPANTS: chatParticipants,
					}).To_Update())
				}
			}
		case Predicate_updateNewChannelMessage:
			switch upd.Message_MESSAGE.PredicateName {
			case Predicate_messageService:
				switch upd.Message_MESSAGE.Action.PredicateName {
				case Predicate_messageActionChatAddUser:
					updates = append(updates, MakeUpdateChannel(upd.Message_MESSAGE.GetPeerId().GetChannelId()))
				}
			}
		case Predicate_updateChannel:
			// channelIdList = append(channelIdList, upd.ChannelId)
		}
	}

	for _, upd := range update {
		updates = append(updates, upd)
	}

	rUpdates := MakeUpdatesByUpdates(updates...)

	idHelper.PickByUpdates(update...)
	if cb1 != nil {
		rUpdates.PushUser(cb1(idHelper.UserIdList)...)
	}
	if cb2 != nil {
		rUpdates.PushChat(cb2(idHelper.ChatIdList)...)
	}
	if cb3 != nil {
		rUpdates.PushChat(cb3(idHelper.ChannelIdList)...)
	}

	return rUpdates
}

func MakeUpdateChannel(channelId int64) *Update {
	return MakeTLUpdateChannel(&Update{
		ChannelId: channelId,
	}).To_Update()
}

func MakeUpdateChat(chatId int64) *Update {
	return MakeTLUpdateChat(&Update{
		ChatId_INT64: chatId,
	}).To_Update()
}
