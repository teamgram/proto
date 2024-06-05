// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mtproto

func (m *MessageBoxList) Visit(
	toUserId int64,
	cb1 func(messageList []*Message),
	cb2 func(userIdList []int64),
	cb3 func(chatIdList []int64),
	cb4 func(channelIdList []int64)) {
	var (
		idHelper    = NewIDListHelper(toUserId)
		messageList = make([]*Message, 0, len(m.BoxList))
	)

	for _, box := range m.BoxList {
		message := box.ToMessage(toUserId)
		messageList = append(messageList, message)
		idHelper.PickByMessage(message)
	}

	if cb1 != nil {
		cb1(messageList)
	}

	idHelper.Visit(cb2, cb3, cb4)
}

func (m *MessageBoxList) Length() int32 {
	return int32(len(m.BoxList))
}

func (m *MessageBoxList) Walk(cb func(idx int, v *MessageBox)) {
	if cb == nil {
		return
	}

	for i, v := range m.BoxList {
		cb(i, v)
	}
}

type ChannelClazz interface {
	ToUnsafeChat(selfUserId int64) *Chat
	Id() int64
}

type MessageBoxListHelper struct {
	BoxList  []*MessageBox
	Users    *MutableUsers
	Chats    []*MutableChat
	Channels []ChannelClazz
}

func MakeBoxListByBoxList(boxList []*MessageBox) *MessageBoxListHelper {
	return &MessageBoxListHelper{
		BoxList:  boxList,
		Users:    nil,
		Chats:    nil,
		Channels: nil,
	}
}

func MakeBoxListByBoxListUsers(boxList []*MessageBox, users []*ImmutableUser) *MessageBoxListHelper {
	return &MessageBoxListHelper{
		BoxList:  boxList,
		Users:    MakeTLMutableUsers(&MutableUsers{Users: users}).To_MutableUsers(),
		Chats:    nil,
		Channels: nil,
	}
}

func MakeBoxListByBoxListUsersChats(boxList []*MessageBox, users []*ImmutableUser, chat ...*MutableChat) *MessageBoxListHelper {
	return &MessageBoxListHelper{
		BoxList:  boxList,
		Users:    MakeTLMutableUsers(&MutableUsers{Users: users}).To_MutableUsers(),
		Chats:    chat,
		Channels: nil,
	}
}

func MakeBoxListByBoxListUsersChannels(boxList []*MessageBox, users []*ImmutableUser, chat ...ChannelClazz) *MessageBoxListHelper {
	return &MessageBoxListHelper{
		BoxList:  boxList,
		Users:    MakeTLMutableUsers(&MutableUsers{Users: users}).To_MutableUsers(),
		Chats:    nil,
		Channels: chat,
	}
}

func (m *MessageBoxListHelper) Visit(
	toUserId int64,
	cb1 func(messageList []*Message),
	cb2 func(users []*User, rawIdList []int64),
	cb3 func(chats []*Chat, rawIdList []int64),
	cb4 func(chats []*Chat, rawIdList []int64),
) {
	var (
		idHelper    = NewIDListHelper(toUserId)
		messageList = make([]*Message, 0, len(m.BoxList))
	)

	for _, box := range m.BoxList {
		message := box.ToMessage(toUserId)
		messageList = append(messageList, message)
		idHelper.PickByMessage(message)
	}

	if cb1 != nil {
		cb1(messageList)
	}

	idHelper.Visit(
		func(userIdList []int64) {
			if cb2 != nil {
				rawIdList := make([]int64, 0, len(userIdList))
				users := m.Users.GetUserListByIdList(toUserId, userIdList...)
				for _, id := range userIdList {
					found := false
					for _, user := range users {
						if user.Id == id {
							found = true
							break
						}
					}
					if !found {
						rawIdList = appendIdF(rawIdList, id)
					}
				}
				cb2(users, rawIdList)
			}
		},
		func(chatIdList []int64) {
			if cb3 != nil {
				rawIdList := make([]int64, 0, len(chatIdList))
				chats := make([]*Chat, 0, len(chatIdList))
				for _, id := range chatIdList {
					var (
						foundChat *Chat
					)
					for _, chat := range m.Chats {
						if chat.Id() == id {
							foundChat = chat.ToUnsafeChat(toUserId)
							break
						}
					}
					if foundChat == nil {
						rawIdList = appendIdF(rawIdList, id)
					} else {
						chats = append(chats, foundChat)
					}
				}
				cb3(chats, rawIdList)
			}
		},
		func(channelIdList []int64) {
			if cb4 != nil {
				rawIdList := make([]int64, 0, len(channelIdList))
				chats := make([]*Chat, 0, len(channelIdList))
				for _, id := range channelIdList {
					var (
						foundChat *Chat
					)
					for _, chat := range m.Channels {
						if chat.Id() == id {
							foundChat = chat.ToUnsafeChat(toUserId)
							break
						}
					}
					if foundChat == nil {
						rawIdList = appendIdF(rawIdList, id)
					} else {
						chats = append(chats, foundChat)
					}
				}
				cb4(chats, rawIdList)
			}
		})
}
