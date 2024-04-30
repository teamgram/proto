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

// /////////////////////////////////////////////////////////////////////////////////////////////////
// updateShortMessage#914fbf11 flags:#
//
//	out:flags.1?true
//	mentioned:flags.4?true
//	media_unread:flags.5?true
//	silent:flags.13?true
//	id:int
//	user_id:int
//	message:string
//	pts:int
//	pts_count:int
//	date:int
//	fwd_from:flags.2?MessageFwdHeader
//	via_bot_id:flags.11?int
//	reply_to_msg_id:flags.3?int
//	entities:flags.7?Vector<MessageEntity> = Updates;
func makeMessageByUpdateShortMessage(userId int64, shortMessage *TLUpdateShortMessage) (message *Message) {
	var (
		fromId, peerId int64
	)

	if shortMessage.GetOut() {
		fromId = userId
		peerId = shortMessage.GetUserId()
	} else {
		fromId = shortMessage.GetUserId()
		peerId = userId
	}

	message = MakeTLMessage(&Message{
		Out:         shortMessage.GetOut(),
		Mentioned:   shortMessage.GetMentioned(),
		MediaUnread: shortMessage.GetMediaUnread(),
		Silent:      shortMessage.GetSilent(),
		Id:          shortMessage.GetId(),
		FromId:      MakePeerUser(fromId),
		PeerId:      MakeTLPeerUser(&Peer{UserId: peerId}).To_Peer(),
		Message:     shortMessage.GetMessage(),
		Date:        shortMessage.GetDate(),
		FwdFrom:     shortMessage.GetFwdFrom(),
		ViaBotId:    shortMessage.GetViaBotId(),
		ReplyTo:     shortMessage.GetReplyTo(),
		Entities:    shortMessage.GetEntities(),
		TtlPeriod:   shortMessage.GetTtlPeriod(),
	}).To_Message()
	return
}

// updateShortChatMessage#16812688 flags:#
//
//	out:flags.1?true
//	mentioned:flags.4?true
//	media_unread:flags.5?true
//	silent:flags.13?true
//	id:int
//	from_id:int
//	chat_id:int
//	message:string
//	pts:int
//	pts_count:int
//	date:int
//	fwd_from:flags.2?MessageFwdHeader
//	via_bot_id:flags.11?int
//	reply_to_msg_id:flags.3?int
//	entities:flags.7?Vector<MessageEntity> = Updates;
func makeMessageByUpdateShortChatMessage(shortMessage *TLUpdateShortChatMessage) *Message {
	message := MakeTLMessage(&Message{
		Out:         shortMessage.GetOut(),
		Mentioned:   shortMessage.GetMentioned(),
		MediaUnread: shortMessage.GetMediaUnread(),
		Silent:      shortMessage.GetSilent(),
		Id:          shortMessage.GetId(),
		FromId:      MakePeerUser(shortMessage.GetFromId()),
		PeerId:      MakePeerChat(shortMessage.GetChatId()),
		Message:     shortMessage.GetMessage(),
		Date:        shortMessage.GetDate(),
		FwdFrom:     shortMessage.GetFwdFrom(),
		ViaBotId:    shortMessage.GetViaBotId(),
		ReplyTo:     shortMessage.GetReplyTo(),
		Entities:    shortMessage.GetEntities(),
		TtlPeriod:   shortMessage.GetTtlPeriod(),
	})
	return message.To_Message()
}

func makeNewUpdateMessageByShortMessage(userId int64, shortMessage *TLUpdateShortMessage) *Update {
	return MakeTLUpdateNewMessage(&Update{
		Message_MESSAGE: makeMessageByUpdateShortMessage(userId, shortMessage),
		Pts_INT32:       shortMessage.GetPts(),
		PtsCount:        shortMessage.GetPtsCount(),
	}).To_Update()
}

func makeNewUpdateMessageByShortChatMessage(userId int64, shortMessage *TLUpdateShortChatMessage) *Update {
	return MakeTLUpdateNewMessage(&Update{
		Message_MESSAGE: makeMessageByUpdateShortChatMessage(shortMessage),
		Pts_INT32:       shortMessage.GetPts(),
		PtsCount:        shortMessage.GetPtsCount(),
	}).To_Update()
}

type UpdateVisitedFunc func(userId int64, update *Update, users []*User, chats []*Chat, date int32)

func VisitUpdates(userId int64, updates *Updates, handlers map[string]UpdateVisitedFunc) {
	if handlers == nil {
		// log.Warnf("VisitUpdates - handlers is nil")
		return
	}

	switch updates.PredicateName {
	case Predicate_updateAccountResetAuthorization:
		// ignore
	case Predicate_updatesTooLong:
		// ignore
	case Predicate_updateShortMessage:
		if vF, ok := handlers[Predicate_updateNewMessage]; ok {
			vF(userId,
				makeNewUpdateMessageByShortMessage(userId, updates.To_UpdateShortMessage()),
				nil,
				nil,
				updates.Date)
		}
	case Predicate_updateShortChatMessage:
		if vF, ok := handlers[Predicate_updateNewMessage]; ok {
			vF(userId,
				makeNewUpdateMessageByShortChatMessage(userId, updates.To_UpdateShortChatMessage()),
				nil,
				nil,
				updates.Date)
		}
	case Predicate_updateShort:
		if vF, ok := handlers[updates.Update.GetPredicateName()]; ok {
			vF(userId, updates.Update, nil, nil, updates.Date)
		}
	case Predicate_updatesCombined:
		for _, update := range updates.Updates {
			if update == nil {
				continue
			}
			if vF, ok := handlers[update.PredicateName]; ok {
				vF(userId, update, updates.Users, updates.Chats, updates.Date)
			}
		}
	case Predicate_updates:
		for _, update := range updates.Updates {
			if update == nil {
				continue
			}
			if vF, ok := handlers[update.PredicateName]; ok {
				vF(userId, update, updates.Users, updates.Chats, updates.Date)
			}
		}
	case Predicate_updateShortSentMessage:
		// ignore
	}
}
