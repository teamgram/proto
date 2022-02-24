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
	PTS_UPDATE_TYPE_UNKNOWN = 0

	// pts
	PTS_UPDATE_NEW_MESSAGE          = 1
	PTS_UPDATE_DELETE_MESSAGES      = 2
	PTS_UPDATE_READ_HISTORY_OUTBOX  = 3
	PTS_UPDATE_READ_HISTORY_INBOX   = 4
	PTS_UPDATE_WEBPAGE              = 5
	PTS_UPDATE_READ_MESSAGE_CONENTS = 6
	PTS_UPDATE_EDIT_MESSAGE         = 7

	// qts
	QTS_UPDATE_NEW_ENCRYPTED_MESSAGE = 8

	// channel pts
	PTS_UPDATE_NEW_CHANNEL_MESSAGE     = 9
	PTS_UPDATE_DELETE_CHANNEL_MESSAGES = 10
	PTS_UPDATE_EDIT_CHANNEL_MESSAGE    = 11
	PTS_UPDATE_EDIT_CHANNEL_WEBPAGE    = 12

	// added
	PTS_UPDATE_FOLDER_PEERS            = 13
	PTS_UPDATE_PINNED_MESSAGES         = 14
	PTS_UPDATE_PINNED_CHANNEL_MESSAGES = 15
	QTS_UPDATE_MMESSAGE_POLL_VOTE      = 16
	QTS_UPDATE_BOT_STOPPED             = 17
)

func GetUpdateType(update *Update) int32 {
	switch update.PredicateName {
	case Predicate_updateNewMessage:
		return PTS_UPDATE_NEW_MESSAGE
	case Predicate_updateDeleteMessages:
		return PTS_UPDATE_DELETE_MESSAGES
	case Predicate_updateReadHistoryOutbox:
		return PTS_UPDATE_READ_HISTORY_OUTBOX
	case Predicate_updateReadHistoryInbox:
		return PTS_UPDATE_READ_HISTORY_INBOX
	case Predicate_updateWebPage:
		return PTS_UPDATE_WEBPAGE
	case Predicate_updateReadMessagesContents:
		return PTS_UPDATE_READ_MESSAGE_CONENTS
	case Predicate_updateEditMessage:
		return PTS_UPDATE_EDIT_MESSAGE

	case Predicate_updateNewEncryptedMessage:
		return QTS_UPDATE_NEW_ENCRYPTED_MESSAGE

	case Predicate_updateNewChannelMessage:
		return PTS_UPDATE_NEW_CHANNEL_MESSAGE
	case Predicate_updateDeleteChannelMessages:
		return PTS_UPDATE_DELETE_CHANNEL_MESSAGES
	case Predicate_updateEditChannelMessage:
		return PTS_UPDATE_EDIT_CHANNEL_MESSAGE
	case Predicate_updateChannelWebPage:
		return PTS_UPDATE_EDIT_CHANNEL_WEBPAGE

	case Predicate_updateFolderPeers:
		return PTS_UPDATE_FOLDER_PEERS
	case Predicate_updatePinnedMessages:
		return PTS_UPDATE_PINNED_MESSAGES
	case Predicate_updatePinnedChannelMessages:
		return PTS_UPDATE_PINNED_CHANNEL_MESSAGES
	case Predicate_updateMessagePollVote:
		return QTS_UPDATE_MMESSAGE_POLL_VOTE
	case "updateBotStopped": // mtproto.Predicate_updateBotStopped:
		return QTS_UPDATE_BOT_STOPPED
	}
	return PTS_UPDATE_TYPE_UNKNOWN
}
