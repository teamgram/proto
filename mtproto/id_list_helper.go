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

type IDListHelper struct {
	UserIdList    []int64
	ChatIdList    []int64
	ChannelIdList []int64
}

func NewIDListHelper(id ...int64) *IDListHelper {
	idHelper := &IDListHelper{
		UserIdList:    []int64{},
		ChatIdList:    []int64{},
		ChannelIdList: []int64{},
	}
	if len(id) > 0 {
		idHelper.UserIdList = id
	}

	return idHelper
}

func appendIdF(idList []int64, id int64) []int64 {
	for _, i := range idList {
		if id != 0 && i == id {
			return idList
		}
	}
	return append(idList, id)
}

func (m *IDListHelper) AppendUsers(id ...int64) {
	for _, id2 := range id {
		m.UserIdList = appendIdF(m.UserIdList, id2)
	}
}

func (m *IDListHelper) AppendChats(id ...int64) {
	for _, id2 := range id {
		m.ChatIdList = appendIdF(m.ChatIdList, id2)
	}
}

func (m *IDListHelper) AppendChannels(id ...int64) {
	for _, id2 := range id {
		m.ChannelIdList = appendIdF(m.ChannelIdList, id2)
	}
}

func (m *IDListHelper) PickByPeer(peer *Peer) {
	if peer == nil {
		return
	}

	switch peer.GetPredicateName() {
	case Predicate_peerUser:
		m.UserIdList = appendIdF(m.UserIdList, peer.GetUserId())
	case Predicate_peerChat:
		m.ChatIdList = appendIdF(m.ChatIdList, peer.GetChatId())
	case Predicate_peerChannel:
		m.ChannelIdList = appendIdF(m.ChannelIdList, peer.GetChannelId())
	}
}

func (m *IDListHelper) PickByPeerUtil(peerType int32, peerId int64) {
	switch peerType {
	case PEER_USER:
		m.UserIdList = appendIdF(m.UserIdList, peerId)
	case PEER_CHAT:
		m.ChatIdList = appendIdF(m.ChatIdList, peerId)
	case PEER_CHANNEL:
		m.ChannelIdList = appendIdF(m.ChannelIdList, peerId)
	}
}

func (m *IDListHelper) PickByFwdFrom(fwd *MessageFwdHeader) {
	// from_id:flags.0?Peer
	if fwd.GetFromId() != nil {
		m.PickByPeer(fwd.GetFromId())
	}

	// saved_from_peer:flags.4?Peer
	if fwd.GetSavedFromPeer() != nil {
		m.PickByPeer(fwd.GetSavedFromPeer())
	}
}

// PickByReplyTo
// messageReplyHeader#a6d57763 flags:#
//	reply_to_msg_id:int
//	reply_to_peer_id:flags.0?Peer
//	reply_to_top_id:flags.1?int = MessageReplyHeader;
func (m *IDListHelper) PickByReplyTo(replyTo *MessageReplyHeader) {
	m.PickByPeer(replyTo.GetReplyToPeerId())
}

// PickByReplies
// messageReplies#4128faac flags:#
//	comments:flags.0?true
//	replies:int
//	replies_pts:int
//	recent_repliers:flags.1?Vector<Peer>
//	channel_id:flags.0?int
//	max_id:flags.2?int
//	read_max_id:flags.3?int = MessageReplies;
//
func (m *IDListHelper) PickByReplies(replies *MessageReplies) {
	for _, p := range replies.GetRecentRepliers() {
		m.PickByPeer(p)
	}
	if replies.GetChannelId() != nil {
		m.ChannelIdList = appendIdF(m.ChannelIdList, replies.GetChannelId().GetValue())
	}
}

// PickByMessageAction PickByMessageAction
func (m *IDListHelper) PickByMessageAction(action *MessageAction) {
	switch action.GetPredicateName() {
	case Predicate_messageActionChatCreate:
		// messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
		m.AppendUsers(action.Users...)
	case Predicate_messageActionChatAddUser:
		//messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;
		m.AppendUsers(action.Users...)
	case Predicate_messageActionChatDeleteUser:
		//messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;
		m.AppendUsers(action.UserId)
	case Predicate_messageActionChatJoinedByLink:
		// messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;
		m.AppendUsers(action.InviterId)
	case Predicate_messageActionChatMigrateTo:
		// messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;
		m.AppendChannels(action.ChannelId)
	case Predicate_messageActionChannelMigrateFrom:
		// messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;
		m.AppendChats(action.ChatId)
	case Predicate_messageActionGeoProximityReached:
		// messageActionGeoProximityReached#98e0d697 from_id:Peer to_id:Peer distance:int = MessageAction;
		m.PickByPeer(action.GetFromId())
		m.PickByPeer(action.GetToId())
	case Predicate_messageActionInviteToGroupCall:
		// messageActionInviteToGroupCall#76b9f11a call:InputGroupCall users:Vector<int> = MessageAction;
		m.AppendUsers(action.Users...)
	}
}

func (m *IDListHelper) PickByMessageMedia(media *MessageMedia) {
	switch media.GetPredicateName() {
	case Predicate_messageMediaContact:
		// messageMediaContact#cbf24940 phone_number:string first_name:string last_name:string vcard:string user_id:int = MessageMedia;
		m.AppendUsers(media.UserId)
	}
}

// PickByMessage PickByMessage
func (m *IDListHelper) PickByMessage(message *Message) {
	// message#bce383d2 flags:#
	//	out:flags.1?true
	//	mentioned:flags.4?true
	//	media_unread:flags.5?true
	//	silent:flags.13?true
	//	post:flags.14?true
	//	from_scheduled:flags.18?true
	//	legacy:flags.19?true
	//	edit_hide:flags.21?true
	//	pinned:flags.24?true
	//	id:int
	//	from_id:flags.8?Peer
	//	peer_id:Peer
	//	fwd_from:flags.2?MessageFwdHeader
	//	via_bot_id:flags.11?int
	//	reply_to:flags.3?MessageReplyHeader
	//	date:int
	//	message:string
	//	media:flags.9?MessageMedia
	//	reply_markup:flags.6?ReplyMarkup
	//	entities:flags.7?Vector<MessageEntity>
	//	views:flags.10?int
	//	forwards:flags.10?int
	//	replies:flags.23?MessageReplies
	//	edit_date:flags.15?int
	//	post_author:flags.16?string
	//	grouped_id:flags.17?long
	//	restriction_reason:flags.22?Vector<RestrictionReason>
	//	ttl_period:flags.25?int = Message;

	if message == nil {
		return
	}

	// from_id:flags.8?Peer
	if message.GetFromId() != nil {
		m.PickByPeer(message.GetFromId())
	}

	// peer_id:Peer
	m.PickByPeer(message.PeerId)

	// fwd_from:flags.2?MessageFwdHeader
	if message.GetFwdFrom() != nil {
		m.PickByFwdFrom(message.GetFwdFrom())
	}

	// reply_markup:flags.6?ReplyMarkup}
	if message.GetReplyTo() != nil {
		m.PickByReplyTo(message.GetReplyTo())
	}

	if message.GetReplies() != nil {
		m.PickByReplies(message.GetReplies())
	}

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
	if message.GetAction() != nil {
		m.PickByMessageAction(message.GetAction())
	}

	if message.GetMedia() != nil {
		m.PickByMessageMedia(message.Media)
	}
}

func (m *IDListHelper) PickByMessages(messages ...*Message) {
	for _, message := range messages {
		m.PickByMessage(message)
	}
}

func (m *IDListHelper) PickByUpdates(update ...*Update) {
	// TODO: handle other
	for _, upd := range update {
		m.PickByMessage(upd.GetMessage_MESSAGE())
		if upd.GetChannelId() != 0 {
			m.ChannelIdList = appendIdF(m.ChannelIdList, upd.GetChannelId())
		}
	}
}

// PickByRules
// TODO(@benqi): pick chat and channel
func (m *IDListHelper) PickByRules(rules []*PrivacyRule) {
	for _, r := range rules {
		for _, id := range r.Users {
			m.UserIdList = appendIdF(m.UserIdList, id)
		}

		for _, id := range r.Chats {
			if id >= MinNebulaChatChannelID {
				m.ChannelIdList = appendIdF(m.ChannelIdList, id)
			} else {
				m.ChatIdList = appendIdF(m.ChatIdList, id)
			}
		}
	}
}

func (m *IDListHelper) PickByChannelParticipant(p *ChannelParticipant) {
	switch p.GetPredicateName() {
	case Predicate_channelParticipant:
		// channelParticipant#c00c07c0 user_id:long date:int = ChannelParticipant;

		m.UserIdList = appendIdF(m.UserIdList, p.UserId)
	case Predicate_channelParticipantSelf:
		// channelParticipantSelf#35a8bfa7 flags:# via_request:flags.0?true user_id:long inviter_id:long date:int = ChannelParticipant;

		m.UserIdList = appendIdF(m.UserIdList, p.UserId)
		m.UserIdList = appendIdF(m.UserIdList, p.InviterId_INT64)
	case Predicate_channelParticipantCreator:
		// channelParticipantCreator#2fe601d3 flags:# user_id:long admin_rights:ChatAdminRights rank:flags.0?string = ChannelParticipant;

		m.UserIdList = appendIdF(m.UserIdList, p.UserId)
	case Predicate_channelParticipantAdmin:
		// channelParticipantAdmin#34c3bb53 flags:# can_edit:flags.0?true self:flags.1?true user_id:long inviter_id:flags.1?long promoted_by:long date:int admin_rights:ChatAdminRights rank:flags.2?string = ChannelParticipant;

		m.UserIdList = appendIdF(m.UserIdList, p.UserId)
		if p.InviterId_FLAGINT64 != nil {
			m.UserIdList = appendIdF(m.UserIdList, p.GetInviterId_FLAGINT64().GetValue())
		}
		m.UserIdList = appendIdF(m.UserIdList, p.PromotedBy)
	case Predicate_channelParticipantBanned:
		// channelParticipantBanned#6df8014e flags:# left:flags.0?true peer:Peer kicked_by:long date:int banned_rights:ChatBannedRights = ChannelParticipant;

		m.PickByPeer(p.Peer)
		m.UserIdList = appendIdF(m.UserIdList, p.KickedBy)
	case Predicate_channelParticipantLeft:
		//channelParticipantLeft#1b03f006 peer:Peer = ChannelParticipant;

		m.PickByPeer(p.Peer)
	}
}

func (m *IDListHelper) Visit(cb1 func(userIdList []int64), cb2 func(chatIdList []int64), cb3 func(channelIdList []int64)) {
	if cb1 != nil && len(m.UserIdList) > 0 {
		cb1(m.UserIdList)
	}
	if cb2 != nil && len(m.ChatIdList) > 0 {
		cb2(m.ChatIdList)
	}
	if cb3 != nil && len(m.ChannelIdList) > 0 {
		cb3(m.ChannelIdList)
	}
}

func (m *IDListHelper) ToUsersAndChats(
	cb1 func(idList []int64) []*User,
	cb2 func(idList []int64) []*Chat,
	cb3 func(idList []int64) []*Chat) (users []*User, chats []*Chat) {
	if cb1 != nil && len(m.ChatIdList) > 0 {
		users = cb1(m.UserIdList)
	}
	if users == nil {
		users = []*User{}
	}

	if cb2 != nil && len(m.ChatIdList) > 0 {
		chats = cb2(m.ChatIdList)
	}

	if cb3 != nil && len(m.ChannelIdList) > 0 {
		chats = append(chats, cb3(m.ChannelIdList)...)
	}

	if chats == nil {
		chats = []*Chat{}
	}

	return
}
