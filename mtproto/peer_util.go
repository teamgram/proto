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
	"fmt"
)

const (
	PEER_EMPTY           = 0
	PEER_SELF            = 1
	PEER_USER            = 2
	PEER_CHAT            = 3
	PEER_CHANNEL         = 4
	PEER_USERS           = 5
	PEER_CHATS           = 6
	PEER_ENCRYPTED_CHAT  = 7
	PEER_BROADCASTS      = 8
	PEER_USER_MESSAGE    = 10
	PEER_CHANNEL_MESSAGE = 11
	PEER_UNKNOWN         = -1
)

var (
// EmptyPeer
)

//type PeerUtil struct {
//	selfId     int64
//	PeerType   int32
//	PeerId     int64
//	AccessHash int64
//}

func (p *PeerUtil) ToString() (s string) {
	switch p.PeerType {
	case PEER_EMPTY:
		return fmt.Sprintf("PEER_EMPTY: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_SELF:
		return fmt.Sprintf("PEER_SELF: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USER:
		return fmt.Sprintf("PEER_USER: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHAT:
		return fmt.Sprintf("PEER_CHAT: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHANNEL:
		return fmt.Sprintf("PEER_CHANNEL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_USERS:
		return fmt.Sprintf("PEER_USERS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	case PEER_CHATS:
		return fmt.Sprintf("PEER_CHATS: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	//case PEER_ALL:
	//	return fmt.Sprintf("PEER_ALL: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	default:
		return fmt.Sprintf("PEER_UNKNOWN: {peer_id: %d, access_hash: %d", p.PeerId, p.AccessHash)
	}
	// return
}

func (p *PeerUtil) CanDoSendMessage() bool {
	switch p.PeerType {
	case PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		return true
	default:
		return false
	}
}

func FromInputUser(selfId int64, user *InputUser) (p *PeerUtil) {
	p = &PeerUtil{}
	switch user.PredicateName {
	case Predicate_inputUserEmpty:
		p.PeerType = PEER_EMPTY
	case Predicate_inputUserSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
	case Predicate_inputUser:
		p.PeerType = PEER_USER
		p.PeerId = user.UserId
		p.AccessHash = user.GetAccessHash()
	case Predicate_inputUserFromMessage:
		//p.PeerType = PEER_USER
		//p.PeerId = user.UserId
		//p.AccessHash = user.GetAccessHash()
	default:
		p.PeerType = PEER_EMPTY
		// panic(fmt.Sprintf("FromInputUser(%v) error!", user))
	}
	return
}

func FromInputPeer(peer *InputPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.PredicateName {
	case Predicate_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case Predicate_inputPeerSelf:
		p.PeerType = PEER_SELF
	case Predicate_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.UserId
		p.AccessHash = peer.GetAccessHash()
	case Predicate_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.ChatId
	case Predicate_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.ChannelId
		p.AccessHash = peer.AccessHash
	default:
		panic(fmt.Sprintf("FromInputPeer(%v) error!", peer))
	}
	return
}

func FromInputPeer2(selfId int64, peer *InputPeer) (p *PeerUtil) {
	p = &PeerUtil{
		SelfId: selfId,
	}
	switch peer.PredicateName {
	case Predicate_inputPeerEmpty:
		p.PeerType = PEER_EMPTY
	case Predicate_inputPeerSelf:
		p.PeerType = PEER_SELF
		p.PeerId = selfId
	case Predicate_inputPeerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.UserId
		p.AccessHash = peer.AccessHash
	case Predicate_inputPeerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.ChatId
	case Predicate_inputPeerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.ChannelId
		p.AccessHash = peer.AccessHash
	default:
		p.PeerType = PEER_UNKNOWN
	}
	return
}

func FromInputEncryptedChat(peer *InputEncryptedChat) (p *PeerUtil) {
	p = &PeerUtil{}

	switch peer.GetPredicateName() {
	case Predicate_inputEncryptedChat:
		p.PeerType = PEER_ENCRYPTED_CHAT
		p.PeerId = int64(peer.GetChatId())
		p.AccessHash = peer.GetAccessHash()
	default:
		p.PeerType = PEER_UNKNOWN
	}

	return
}

//
//func (p *PeerUtil) IsSelf() bool {
//	switch p.PeerType {
//	case PEER_SELF:
//		return true
//	case PEER_USER:
//		return p.selfId == p.PeerId
//	}
//	return false
//}

func (p *PeerUtil) ToInputPeer() (peer *InputPeer) {
	switch p.PeerType {
	case PEER_EMPTY:
		peer = MakeTLInputPeerEmpty(nil).To_InputPeer()
	case PEER_SELF:
		peer = MakeTLInputPeerSelf(nil).To_InputPeer()
	case PEER_USER:
		peer = MakeTLInputPeerUser(&InputPeer{
			UserId:     p.PeerId,
			AccessHash: p.AccessHash,
		}).To_InputPeer()
	case PEER_CHAT:
		peer = MakeTLInputPeerChat(&InputPeer{
			ChatId: p.PeerId,
		}).To_InputPeer()
	case PEER_CHANNEL:
		peer = MakeTLInputPeerChannel(&InputPeer{
			ChannelId:  p.PeerId,
			AccessHash: p.AccessHash,
		}).To_InputPeer()
	default:
		panic(fmt.Sprintf("ToInputPeer(%v) error!", p))
	}
	return
}

func FromPeer(peer *Peer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.PredicateName {
	case Predicate_peerUser:
		p.PeerType = PEER_USER
		p.PeerId = peer.UserId
	case Predicate_peerChat:
		p.PeerType = PEER_CHAT
		p.PeerId = peer.ChatId
	case Predicate_peerChannel:
		p.PeerType = PEER_CHANNEL
		p.PeerId = peer.ChannelId
	default:
		panic(fmt.Sprintf("FromPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) ToPeer() (peer *Peer) {
	switch p.PeerType {
	case PEER_SELF:
		if p.PeerId != 0 {
			peer = MakeTLPeerUser(&Peer{
				UserId: p.PeerId,
			}).To_Peer()
		} else if p.SelfId != 0 {
			peer = MakeTLPeerUser(&Peer{
				// Constructor: CRC32_peerUser,
				UserId: p.SelfId,
			}).To_Peer()
		} else {
			panic(fmt.Sprintf("ToPeer(%v) error!", p))
		}
	case PEER_USER:
		peer = MakeTLPeerUser(&Peer{
			// Constructor: CRC32_peerUser,
			UserId: p.PeerId,
		}).To_Peer()
	case PEER_CHAT:
		peer = MakeTLPeerChat(&Peer{
			// Constructor: CRC32_peerChat,
			ChatId: p.PeerId,
		}).To_Peer()
	case PEER_CHANNEL:
		peer = MakeTLPeerChannel(&Peer{
			// Constructor: CRC32_peerChannel,
			ChannelId: p.PeerId,
		}).To_Peer()
	default:
		peer = nil
		// panic(fmt.Sprintf("ToPeer(%v) error!", p))
	}
	return
}

func (p *PeerUtil) IsEmpty() bool {
	return p.PeerType == PEER_EMPTY
}

func (p *PeerUtil) IsSelf() bool {
	return p.PeerType == PEER_SELF
}

func (p *PeerUtil) IsUser() bool {
	return p.PeerType == PEER_USER || p.PeerType == PEER_SELF
}

func (p *PeerUtil) IsChat() bool {
	return p.PeerType == PEER_CHAT
}

func (p *PeerUtil) IsChatOrChannel() bool {
	return p.IsChat() || p.IsChannel()
}

func (p *PeerUtil) IsUserOrChatOrChannel() bool {
	return p.IsUser() || p.IsChat() || p.IsChannel()
}

func (p *PeerUtil) IsChatOrUser() bool {
	return p.IsUser() || p.IsChat()
}

func (p *PeerUtil) IsChannel() bool {
	return p.PeerType == PEER_CHANNEL
}

func (p *PeerUtil) IsEncryptedChat() bool {
	return p.PeerType == PEER_ENCRYPTED_CHAT
}

func (p *PeerUtil) IsSelfUser(id int64) bool {
	return p.PeerType == PEER_SELF || p.PeerType == PEER_USER && p.PeerId == id
}

func FromInputNotifyPeer(selfId int64, peer *InputNotifyPeer) (p *PeerUtil) {
	p = &PeerUtil{}
	switch peer.PredicateName {
	case Predicate_inputNotifyPeer:
		p = FromInputPeer2(selfId, peer.Peer)
	case Predicate_inputNotifyUsers:
		p.PeerType = PEER_USERS
	case Predicate_inputNotifyChats:
		p.PeerType = PEER_CHATS
	case Predicate_inputNotifyBroadcasts:
		p.PeerType = PEER_BROADCASTS
	default:
		// log.Errorf("fromInputNotifyPeer: invalid peer - %v", peer)
		p.PeerType = PEER_UNKNOWN
	}
	return
}

func (p *PeerUtil) ToInputNotifyPeer(peer *InputNotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = MakeTLInputNotifyPeer(&InputNotifyPeer{
			Constructor: CRC32_inputNotifyPeer,
			Peer:        p.ToInputPeer(),
		}).To_InputNotifyPeer()
	case PEER_USERS:
		peer = MakeTLInputNotifyUsers(&InputNotifyPeer{
			Constructor: CRC32_inputNotifyUsers,
		}).To_InputNotifyPeer()
	case PEER_CHATS:
		peer = MakeTLInputNotifyChats(&InputNotifyPeer{
			Constructor: CRC32_inputNotifyChats,
		}).To_InputNotifyPeer()
	//case PEER_ALL:
	//	peer = &InputNotifyPeer{
	//		Constructor: TLConstructor_CRC32_inputNotifyAll,
	//		Data2:       &InputNotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToInputNotifyPeer(%v) error!", p))
	}
	return
}

//func FromNotifyPeer(peer *NotifyPeer) (p *PeerUtil) {
//	p = &PeerUtil{}
//	switch peer.GetConstructor() {
//	case CRC32_notifyPeer:
//		p = FromPeer(peer.GetPeer())
//	case CRC32_notifyUsers:
//		p.PeerType = PEER_USERS
//	case CRC32_notifyChats:
//		p.PeerType = PEER_CHATS
//	case CRC32_notifyAll:
//		p.PeerType = PEER_ALL
//	default:
//		panic(fmt.Sprintf("FromNotifyPeer(%v) error!", p))
//	}
//	return
//}

func (p *PeerUtil) ToNotifyPeer() (peer *NotifyPeer) {
	switch p.PeerType {
	case PEER_EMPTY, PEER_SELF, PEER_USER, PEER_CHAT, PEER_CHANNEL:
		peer = MakeTLNotifyPeer(&NotifyPeer{
			// Constructor: CRC32_notifyPeer,
			Peer: p.ToPeer(),
		}).To_NotifyPeer()
	case PEER_USERS:
		peer = MakeTLNotifyUsers(&NotifyPeer{
			// Constructor: CRC32_notifyUsers,
		}).To_NotifyPeer()
	case PEER_CHATS:
		peer = MakeTLNotifyChats(&NotifyPeer{
			// Constructor: CRC32_notifyChats,
		}).To_NotifyPeer()
	case PEER_BROADCASTS:
		peer = MakeTLNotifyBroadcasts(&NotifyPeer{
			// Constructor: CRC32_notifyBroadcasts,
		}).To_NotifyPeer()
	//case PEER_ALL:
	//	peer = &NotifyPeer{
	//		Constructor: TLConstructor_CRC32_notifyAll,
	//		Data2:       &NotifyPeer_Data{},
	//	}
	default:
		panic(fmt.Sprintf("ToNotifyPeer(%v) error!", p))
	}
	return
}

func ToPeerByTypeAndID(peerType int32, peerId int64) (peer *Peer) {
	switch peerType {
	case PEER_USER:
		peer = MakeTLPeerUser(&Peer{
			UserId: peerId,
		}).To_Peer()
	case PEER_CHAT:
		peer = MakeTLPeerChat(&Peer{
			ChatId: peerId,
		}).To_Peer()
	case PEER_CHANNEL:
		peer = MakeTLPeerChannel(&Peer{
			ChannelId: peerId,
		}).To_Peer()
	default:
		panic(fmt.Sprintf("ToPeerByTypeAndID(%d, %d) error!", peerType, peerId))
	}
	return
}

func PickAllIdListByPeers(peers []*Peer) (idList, chatIdList, channelIdList []int64) {
	for _, p := range peers {
		switch p.PredicateName {
		case Predicate_peerUser:
			idList = append(idList, p.UserId)
		case Predicate_peerChat:
			chatIdList = append(chatIdList, p.ChatId)
		case Predicate_peerChannel:
			channelIdList = append(channelIdList, p.ChannelId)
		}
	}
	if idList == nil {
		idList = []int64{}
	}
	if chatIdList == nil {
		chatIdList = []int64{}
	}
	if channelIdList == nil {
		channelIdList = []int64{}
	}
	return
}

func MakePeerUser(peerId int64) *Peer {
	return MakeTLPeerUser(&Peer{
		UserId: peerId,
	}).To_Peer()
}

func MakePeerChat(peerId int64) *Peer {
	return MakeTLPeerChat(&Peer{
		ChatId: peerId,
	}).To_Peer()
}

func MakePeerChannel(peerId int64) *Peer {
	return MakeTLPeerChannel(&Peer{
		ChannelId: peerId,
	}).To_Peer()
}

func MakePeer(peerType int32, peerId int64) *Peer {
	switch peerType {
	case PEER_CHAT:
		return MakePeerChat(peerId)
	case PEER_CHANNEL:
		return MakePeerChannel(peerId)
	default:
		return MakePeerUser(peerId)
	}
}

func MakePeerUtil(peerType int32, peerId int64) *PeerUtil {
	return &PeerUtil{
		PeerType: peerType,
		PeerId:   peerId,
	}
}

func MakeUserPeerUtil(peerId int64) *PeerUtil {
	return &PeerUtil{
		PeerType: PEER_USER,
		PeerId:   peerId,
	}
}

func MakeChatPeerUtil(peerId int64) *PeerUtil {
	return &PeerUtil{
		PeerType: PEER_CHAT,
		PeerId:   peerId,
	}
}

func MakeChannelPeerUtil(peerId int64) *PeerUtil {
	return &PeerUtil{
		PeerType: PEER_CHANNEL,
		PeerId:   peerId,
	}
}

func MakeInputPeerChat(peerId int64) *InputPeer {
	return MakeTLInputPeerChat(&InputPeer{
		ChatId: peerId,
	}).To_InputPeer()
}

func MakeInputPeerChannel(peerId int64) *InputPeer {
	return MakeTLInputPeerChannel(&InputPeer{
		ChannelId: peerId,
	}).To_InputPeer()
}

func MakePeerDialogId(peerType int32, peerId int64) int64 {
	var (
		id int64 = 0
	)

	switch peerType {
	case PEER_SELF, PEER_USER:
		id = peerId
	case PEER_CHAT, PEER_CHANNEL:
		id = -peerId
	default:
		// LOGO
	}

	return id
}

func GetPeerUtilByPeerDialogId(id int64) (int32, int64) {
	if id > 0 {
		return PEER_USER, id
	} else {
		if ChatIdIsChat(-id) {
			return PEER_CHAT, -id
		} else {
			return PEER_CHANNEL, -id
		}
	}
}

func IsChannelInputPeer(peer *InputPeer) bool {
	return peer.GetPredicateName() == Predicate_inputPeerChannel
}

func IsChatInputPeer(peer *InputPeer) bool {
	return peer.GetPredicateName() == Predicate_inputPeerChat
}

func IsUserInputPeer(peer *InputPeer) bool {
	return peer.GetPredicateName() == Predicate_inputPeerUser ||
		peer.GetPredicateName() == Predicate_inputPeerSelf
}

func PeerIsChannel(peer *Peer) bool {
	return peer.GetPredicateName() == Predicate_peerChannel
}
