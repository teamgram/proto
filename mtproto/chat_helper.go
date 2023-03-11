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

const (
	ChatMemberNormal  = 0
	ChatMemberCreator = 1
	ChatMemberAdmin   = 2
	ChatMemberBanned  = 4
	// ChatMemberMigratedTo = 5
)

const (
	ChatMemberStateNormal      = 0 // normal
	ChatMemberStateLeft        = 1 // left
	ChatMemberStateKicked      = 2 // kicked
	ChatMemberStateMigrated    = 3 // migrated
	ChatMemberStateChatDeleted = 4 // deleted
)

var (
// Cache
// ExportedChatInviteEmpty = mtproto.MakeTLChatInviteEmpty(&mtproto.ExportedChatInvite{}).To_ExportedChatInvite()
)

func SplitChatAndChannelIdList(idList []int64) (chatIdList, channelIdList []int64) {
	for _, id := range idList {
		if ChatIdIsChat(id) {
			chatIdList = append(chatIdList, id)
		} else {
			channelIdList = append(channelIdList, id)
		}
	}
	return
}

func MakeChatMessageService(fromId int64, chatType int32, chatId int64, slient bool, action *MessageAction) *Message {
	message := MakeTLMessageService(&Message{
		// TODO(@benqi): fill it
		Out:         true,
		Mentioned:   false,
		MediaUnread: false,
		Silent:      slient,
		Post:        false,
		Legacy:      false,
		Id:          0,
		FromId:      MakePeerUser(fromId),
		PeerId:      MakePeerUtil(chatType, chatId).ToPeer(),
		ReplyTo:     nil,
		Date:        int32(time.Now().Unix()),
		Action:      action,
		TtlPeriod:   nil,
	})
	return message.To_Message()
}

func ToSafeChats(chats []*Chat) []*Chat {
	if chats == nil {
		return []*Chat{}
	} else {
		return chats
	}
}
