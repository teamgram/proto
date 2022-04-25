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
	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
)

type DialogID struct {
	A int64
	B int64
}

func (did *DialogID) IsZero() bool {
	return did == nil || (did.A == 0 && did.B == 0)
}

func (did *DialogID) ToPeerDialogId(selfId int64) int64 {
	if did.A < 0 {
		return -did.B
	} else {
		if selfId == did.A {
			return did.B
		} else {
			return did.A
		}
	}
}

func (did *DialogID) ToPeer(selfId int64) *Peer {
	if did.A < 0 {
		switch -did.A {
		case PEER_CHAT:
			return MakePeerChat(did.B)
		case PEER_CHANNEL:
			return MakePeerChannel(did.B)
		default:
			panic(ErrPeerIdInvalid)
		}
	} else {
		if selfId == did.A {
			return MakePeerUser(did.B)
		} else if selfId == did.B {
			return MakePeerUser(did.A)
		} else {
			panic(ErrPeerIdInvalid)
		}
	}
}

func (did *DialogID) ToPeerUtil(selfId int64) *PeerUtil {
	if did.A < 0 {
		switch -did.A {
		case PEER_CHAT:
			return MakeChatPeerUtil(did.B)
		case PEER_CHANNEL:
			return MakeChannelPeerUtil(did.B)
		default:
			panic(ErrPeerIdInvalid)
		}
	} else {
		if selfId == did.A {
			return MakeUserPeerUtil(did.B)
		} else if selfId == did.B {
			return MakeUserPeerUtil(did.A)
		} else {
			panic(ErrPeerIdInvalid)
		}
	}
}

func MakeDialogId(fromId int64, peerType int32, peerId int64) (did DialogID) {
	switch peerType {
	case PEER_SELF:
		did.A = fromId
		did.B = fromId
	case PEER_USER:
		if fromId <= peerId {
			did.A = fromId
			did.B = peerId
		} else {
			did.A = peerId
			did.B = fromId
		}
	case PEER_CHAT:
		did.A = -int64(peerType)
		did.B = peerId
	case PEER_CHANNEL:
		did.A = -int64(peerType)
		did.B = peerId
	default:
		// log.Warnf("invalid peer{%d, %d, %d}", fromId, peerType, peerId)
	}
	return
}

func GetPeerIdByDialogId(userId int64, did DialogID) int64 {
	if did.A < 0 {
		return did.B
	} else {
		if userId == did.A {
			return did.B
		} else {
			return did.A
		}
	}
}

func GetPeerByDialogId(userId int64, did DialogID) *Peer {
	if did.A < 0 {
		return MakePeerChat(did.B)
	} else {
		if userId == did.A {
			return MakePeerUser(did.B)
		} else {
			return MakePeerUser(did.A)
		}
	}
}

func CheckHasMention(entities []*MessageEntity, userId int64) bool {
	for _, e := range entities {
		// check name
		switch e.PredicateName {
		case Predicate_messageEntityMentionName:
			if e.UserId_INT64 == userId {
				return true
			}
		case Predicate_messageEntityMention:
			if e.UserId_INT64 == userId {
				return true
			}
		}

		// check crc32
		switch e.GetConstructor() {
		case CRC32_messageEntityMentionName:
			if e.UserId_INT64 == userId {
				return true
			}
		case CRC32_messageEntityMention:
			if e.UserId_INT64 == userId {
				return true
			}
		}
	}
	return false
}

func CheckHasMediaUnread(m *Message) bool {
	return IsVoiceMessage(m) || IsRoundVideoMessage(m)
}

func (m *MessageBox) ToMessage(toUserId int64) *Message {
	var (
		message = m.Message
	)

	message.MediaUnread = m.MediaUnread

	switch m.PeerType {
	case PEER_SELF, PEER_USER:
	case PEER_CHAT:
		message.Mentioned = m.Mentioned
	case PEER_CHANNEL:
		message = proto.Clone(message).(*Message)

		if m.SenderUserId == toUserId {
			message.Out = true
		} else {
			message.Out = false
		}

		if message.Views == nil {
			// megagroup
			message.FromId = MakePeerUser(m.SenderUserId)
		} else {
			if message.Forwards == nil {
				message.Forwards = &types.Int32Value{Value: 0}
			}
		}
	default:
		panic("unknown peer_type")
	}

	message.Pinned = m.Pinned
	return message
}

func (m *MessageBox) IsOut(toUserId int64) bool {
	if m.GetPeerType() == PEER_CHANNEL {
		if m.GetSenderUserId() == toUserId {
			return true
		} else {
			return false
		}
	}

	return m.GetMessage().GetOut()
}

func ToSafeMessageBoxList(messages []*MessageBox) []*MessageBox {
	if messages == nil {
		return []*MessageBox{}
	} else {
		return messages
	}
}
