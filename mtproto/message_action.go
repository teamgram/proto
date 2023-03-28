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

	"github.com/gogo/protobuf/types"
)

func (m *MessageAction) FixData() *MessageAction {
	if m.GetPredicateName() != Predicate_messageActionTopicEdit {
		if m.Title_STRING == "" {
			m.Title_STRING = m.Title
		}
	}

	return m
}

// MakeMessageActionEmpty
// messageActionEmpty#b6aef7b0 = MessageAction;
func MakeMessageActionEmpty() *MessageAction {
	return MakeTLMessageActionEmpty(nil).To_MessageAction()
}

// MakeMessageActionChatCreate
// messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
func MakeMessageActionChatCreate(title string, users []int64) *MessageAction {
	return MakeTLMessageActionChatCreate(&MessageAction{
		Title:        title,
		Title_STRING: title,
		Users:        users,
	}).To_MessageAction()
}

// MakeMessageActionChatEditTitle
// messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;
func MakeMessageActionChatEditTitle(title string) *MessageAction {
	return MakeTLMessageActionChatEditTitle(&MessageAction{
		Title:        title,
		Title_STRING: title,
	}).To_MessageAction()
}

// MakeMessageActionChatEditPhoto
// messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;
func MakeMessageActionChatEditPhoto(photo *Photo) *MessageAction {
	return MakeTLMessageActionChatEditPhoto(&MessageAction{
		Photo: photo,
	}).To_MessageAction()
}

// MakeMessageActionChatDeletePhoto
// messageActionChatDeletePhoto#95e3fbef = MessageAction;
func MakeMessageActionChatDeletePhoto() *MessageAction {
	return MakeTLMessageActionChatDeletePhoto(&MessageAction{}).To_MessageAction()
}

// MakeMessageActionChatAddUser
// messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;
func MakeMessageActionChatAddUser(users ...int64) *MessageAction {
	return MakeTLMessageActionChatAddUser(&MessageAction{
		Users: users,
	}).To_MessageAction()
}

// MakeMessageActionChatDeleteUser
// messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;
func MakeMessageActionChatDeleteUser(userId int64) *MessageAction {
	return MakeTLMessageActionChatDeleteUser(&MessageAction{
		UserId: userId,
	}).To_MessageAction()
}

// MakeMessageActionChatJoinByLink
// messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;
func MakeMessageActionChatJoinByLink(inviterId int64) *MessageAction {
	return MakeTLMessageActionChatJoinedByLink(&MessageAction{
		InviterId: inviterId,
	}).To_MessageAction()
}

// MakeMessageActionChannelCreate
// messageActionChannelCreate#95d2ac92 title:string = MessageAction;
func MakeMessageActionChannelCreate(title string) *MessageAction {
	return MakeTLMessageActionChannelCreate(&MessageAction{
		Title:        title,
		Title_STRING: title,
	}).To_MessageAction()
}

// MakeMessageActionChatMigrateTo
// messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;
func MakeMessageActionChatMigrateTo(channelId int64) *MessageAction {
	return MakeTLMessageActionChatMigrateTo(&MessageAction{
		ChannelId: channelId,
	}).To_MessageAction()
}

// MakeMessageActionChannelMigrateFrom
// messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;
func MakeMessageActionChannelMigrateFrom(title string, chatId int64) *MessageAction {
	return MakeTLMessageActionChannelMigrateFrom(&MessageAction{
		Title:        title,
		Title_STRING: title,
		ChatId:       chatId,
	}).To_MessageAction()
}

// MakeMessageActionPinMessage
// messageActionPinMessage#94bd38ed = MessageAction;
func MakeMessageActionPinMessage() *MessageAction {
	return MakeTLMessageActionPinMessage(nil).To_MessageAction()
}

// MakeMessageActionHistoryClear
// messageActionHistoryClear#9fbab604 = MessageAction;
func MakeMessageActionHistoryClear() *MessageAction {
	return MakeTLMessageActionHistoryClear(nil).To_MessageAction()
}

// MakeMessageActionGameScore
// messageActionGameScore#92a72876 game_id:long score:int = MessageAction;
func MakeMessageActionGameScore(gameId int64, score int32) *MessageAction {
	return MakeTLMessageActionGameScore(&MessageAction{
		GameId: gameId,
		Score:  score,
	}).To_MessageAction()
}

// MakeMessageActionPaymentSentMe
// messageActionPaymentSentMe#8f31b327 flags:# currency:string total_amount:long payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string charge:PaymentCharge = MessageAction;
func MakeMessageActionPaymentSentMe(currency string, totalAmount int64, payload []byte, info *PaymentRequestedInfo, shippingOptionId string, charge *PaymentCharge) *MessageAction {
	action := MakeTLMessageActionGameScore(&MessageAction{
		Currency:         currency,
		TotalAmount:      totalAmount,
		Payload:          payload,
		Info:             info,
		ShippingOptionId: nil,
		Charge:           charge,
	}).To_MessageAction()

	if shippingOptionId != "" {
		action.ShippingOptionId = &types.StringValue{Value: shippingOptionId}
	}
	return action
}

// MakeMessageActionPaymentSent
// messageActionPaymentSent#40699cd0 currency:string total_amount:long = MessageAction;
func MakeMessageActionPaymentSent(currency string, totalAmount int64) *MessageAction {
	return MakeTLMessageActionPaymentSent(&MessageAction{
		Currency:    currency,
		TotalAmount: totalAmount,
	}).To_MessageAction()
}

// MakeMessageActionPhoneCall
// messageActionPhoneCall#80e11a7f flags:# video:flags.2?true call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;
func MakeMessageActionPhoneCall(video bool, callId int64, reason *PhoneCallDiscardReason, duration int32) *MessageAction {
	action := MakeTLMessageActionPhoneCall(&MessageAction{
		Video:    video,
		CallId:   callId,
		Reason:   reason,
		Duration: nil,
	}).To_MessageAction()
	if duration > 0 {
		action.Duration = &types.Int32Value{Value: duration}
	}
	return action
}

// MakeMessageActionScreenshotTaken
// messageActionScreenshotTaken#4792929b = MessageAction;
func MakeMessageActionScreenshotTaken() *MessageAction {
	return MakeTLMessageActionScreenshotTaken(nil).To_MessageAction()
}

// MakeMessageActionCustomAction
// messageActionCustomAction#fae69f56 message:string = MessageAction;
func MakeMessageActionCustomAction(message string) *MessageAction {
	return MakeTLMessageActionCustomAction(&MessageAction{
		Message: message,
	}).To_MessageAction()
}

// MakeMessageActionBotAllowed
// messageActionBotAllowed#abe9affe domain:string = MessageAction;
// messageActionBotAllowed#c516d679 flags:# attach_menu:flags.1?true domain:flags.0?string app:flags.2?BotApp = MessageAction;
func MakeMessageActionBotAllowed(domain string) *MessageAction {
	return MakeTLMessageActionBotAllowed(&MessageAction{
		AttachMenu:        false,
		Domain_STRING:     domain,
		Domain_FLAGSTRING: MakeFlagsString(domain),
		App:               nil,
	}).To_MessageAction()
}

// MakeMessageActionSecureValuesSentMe
// messageActionSecureValuesSentMe#1b287353 values:Vector<SecureValue> credentials:SecureCredentialsEncrypted = MessageAction;
func MakeMessageActionSecureValuesSentMe(values []*SecureValue, credentials *SecureCredentialsEncrypted) *MessageAction {
	return MakeTLMessageActionSecureValuesSentMe(&MessageAction{
		Values:      values,
		Credentials: credentials,
	}).To_MessageAction()
}

// MakeMessageActionSecureValuesSent
// messageActionSecureValuesSent#d95c6154 types:Vector<SecureValueType> = MessageAction;
func MakeMessageActionSecureValuesSent(types []*SecureValueType) *MessageAction {
	return MakeTLMessageActionSecureValuesSent(&MessageAction{
		Types: types,
	}).To_MessageAction()
}

// MakeMessageActionContactSignUp
// messageActionContactSignUp#f3f25f76 = MessageAction;
func MakeMessageActionContactSignUp() *MessageAction {
	return MakeTLMessageActionContactSignUp(nil).To_MessageAction()
}

// MakeMessageActionGeoProximityReached
// messageActionGeoProximityReached#98e0d697 from_id:Peer to_id:Peer distance:int = MessageAction;
func MakeMessageActionGeoProximityReached(fromId, toId *Peer, distance int32) *MessageAction {
	return MakeTLMessageActionGeoProximityReached(&MessageAction{
		FromId:   fromId,
		ToId:     toId,
		Distance: distance,
	}).To_MessageAction()
}

// MakeMessageActionGroupCall
// messageActionGroupCall#7a0d7f42 flags:# call:InputGroupCall duration:flags.0?int = MessageAction;
func MakeMessageActionGroupCall(call *InputGroupCall, duration int32) *MessageAction {
	return MakeTLMessageActionGroupCall(&MessageAction{
		Call:     call,
		Duration: MakeFlagsInt32(duration),
	}).To_MessageAction()
}

// MakeMessageActionInviteToGroupCall
// messageActionInviteToGroupCall#502f92f7 call:InputGroupCall users:Vector<long> = MessageAction;
func MakeMessageActionInviteToGroupCall(call *InputGroupCall, users ...int64) *MessageAction {
	return MakeTLMessageActionInviteToGroupCall(&MessageAction{
		Call:  call,
		Users: users,
	}).To_MessageAction()
}

// MakeMessageActionSetMessagesTTL
// messageActionSetMessagesTTL#aa1afbfd period:int = MessageAction;
func MakeMessageActionSetMessagesTTL(period int32) *MessageAction {
	return MakeTLMessageActionSetMessagesTTL(&MessageAction{
		Period: period,
	}).To_MessageAction()
}

// MakeMessageActionGroupCallScheduled
// messageActionGroupCallScheduled#b3a07661 call:InputGroupCall schedule_date:int = MessageAction;
func MakeMessageActionGroupCallScheduled(call *InputGroupCall, scheduleDate int32) *MessageAction {
	return MakeTLMessageActionGroupCallScheduled(&MessageAction{
		Call:         call,
		ScheduleDate: scheduleDate,
	}).To_MessageAction()
}

// MakeMessageActionSetChatTheme
// messageActionSetChatTheme#aa786345 emoticon:string = MessageAction;
func MakeMessageActionSetChatTheme(emoticon string) *MessageAction {
	return MakeTLMessageActionSetChatTheme(&MessageAction{
		Emoticon: emoticon,
	}).To_MessageAction()
}

// MakeMessageActionChatJoinedByRequest
// messageActionChatJoinedByRequest#ebbca3cb = MessageAction;
func MakeMessageActionChatJoinedByRequest() *MessageAction {
	return MakeTLMessageActionChatJoinedByRequest(nil).To_MessageAction()
}

// MakeMessageActionWebViewDataSentMe
// messageActionWebViewDataSentMe#47dd8079 text:string data:string = MessageAction;
func MakeMessageActionWebViewDataSentMe(text, data string) *MessageAction {
	return MakeTLMessageActionWebViewDataSentMe(&MessageAction{
		Text: text,
		Data: data,
	}).To_MessageAction()
}

// MakeMessageActionWebViewDataSent
// messageActionWebViewDataSent#b4c38cb5 text:string = MessageAction;
func MakeMessageActionWebViewDataSent(text string) *MessageAction {
	return MakeTLMessageActionWebViewDataSent(&MessageAction{
		Text: text,
	}).To_MessageAction()
}

// MakeMessageActionGiftPremium
// messageActionGiftPremium#aba0f5c6 currency:string amount:long months:int = MessageAction;
func MakeMessageActionGiftPremium(currency string, amount int64, months int32) *MessageAction {
	return MakeTLMessageActionGiftPremium(&MessageAction{
		Currency: currency,
		Amount:   amount,
		Months:   months,
	}).To_MessageAction()
}

// MakeMessageActionTopicCreate
// messageActionTopicCreate#d999256 flags:# title:string icon_color:int icon_emoji_id:flags.0?long = MessageAction;
func MakeMessageActionTopicCreate(title string, iconColor int32, iconEmojiId int64) *MessageAction {
	return MakeTLMessageActionTopicCreate(&MessageAction{
		Title_STRING: title,
		IconColor:    iconColor,
		IconEmojiId:  MakeFlagsInt64(iconEmojiId),
	}).To_MessageAction()
}

// MakeMessageActionTopicEdit
// messageActionTopicEdit#b18a431c flags:# title:flags.0?string icon_emoji_id:flags.1?long closed:flags.2?Bool = MessageAction;
func MakeMessageActionTopicEdit(title string, iconEmojiId int64, closed *Bool) *MessageAction {
	return MakeTLMessageActionTopicEdit(&MessageAction{
		Title_FLAGSTRING: MakeFlagsString(title),
		IconEmojiId:      MakeFlagsInt64(iconEmojiId),
		Closed:           closed,
	}).To_MessageAction()
}

func MakeContactSignUpMessage(fromId, toId int64) *Message {
	return MakeTLMessageService(&Message{
		// TODO(@benqi): fill it
		Out:         true,
		Mentioned:   false,
		MediaUnread: false,
		Silent:      false,
		Post:        false,
		Legacy:      false,
		Id:          0,
		FromId:      MakePeerUser(fromId),
		PeerId:      MakePeerUser(toId),
		ReplyTo:     nil,
		Date:        int32(time.Now().Unix()),
		Action:      MakeMessageActionContactSignUp(),
		TtlPeriod:   nil,
	}).To_Message()
}

// MakePinnedMessageService
// messageService#2b085862 flags:#
//	out:flags.1?true
//	mentioned:flags.4?true
//	media_unread:flags.5?true
//	silent:flags.13?true
//	post:flags.14?true
//	legacy:flags.19?true
//	id:int
//	from_id:flags.8?Peer
//	peer_id:Peer
//	reply_to:flags.3?MessageReplyHeader
//	date:int
//	action:MessageAction
//	ttl_period:flags.25?int = Message;
func MakePinnedMessageService(slient bool, fromId int64, peer *PeerUtil, pinnedId int32) *Message {
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
		PeerId:      peer.ToPeer(),
		ReplyTo: MakeTLMessageReplyHeader(&MessageReplyHeader{
			ReplyToMsgId:  pinnedId,
			ReplyToPeerId: nil,
			ReplyToTopId:  nil,
		}).To_MessageReplyHeader(),
		Date:      int32(time.Now().Unix()),
		Action:    MakeMessageActionPinMessage(),
		TtlPeriod: nil,
	})
	return message.To_Message()
}

func MakeSetChatThemeService(fromId int64, peer *PeerUtil, emoticon string) *Message {
	message := MakeTLMessageService(&Message{
		// TODO(@benqi): fill it
		Out:         true,
		Mentioned:   false,
		MediaUnread: false,
		Silent:      true,
		Post:        false,
		Legacy:      false,
		Id:          0,
		FromId:      MakePeerUser(fromId),
		PeerId:      peer.ToPeer(),
		ReplyTo:     nil,
		Date:        int32(time.Now().Unix()),
		Action:      MakeMessageActionSetChatTheme(emoticon),
		TtlPeriod:   nil,
	})
	return message.To_Message()
}
