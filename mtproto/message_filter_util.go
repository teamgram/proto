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

//public final static int MEDIA_PHOTOVIDEO = 0;
//public final static int MEDIA_FILE = 1;
//public final static int MEDIA_AUDIO = 2;
//public final static int MEDIA_URL = 3;
//public final static int MEDIA_MUSIC = 4;
//public final static int MEDIA_GIF = 5;
//public final static int MEDIA_PHOTOS_ONLY = 6;
//public final static int MEDIA_VIDEOS_ONLY = 7;
//public final static int MEDIA_TYPES_COUNT = 8;

// Predicate_inputMessagesFilterPhotoVideo
// inputMessagesFilterPhotoVideo
const (
	MEDIA_EMPTY       = -1 //
	MEDIA_PHOTOVIDEO  = 0  // PhotoVideo -> inputMessagesFilterPhotoVideo
	MEDIA_FILE        = 1  // File -> inputMessagesFilterDocument
	MEDIA_AUDIO       = 2  // Audio(RoundVoiceFile) -> inputMessagesFilterRoundVoice
	MEDIA_URL         = 3  // Link -> inputMessagesFilterUrl
	MEDIA_MUSIC       = 4  // MusicFile -> inputMessagesFilterMusic
	MEDIA_PHONE_CALL  = 5  // inputMessagesFilterPhoneCalls
	MEDIA_GIF         = 6  // Gif -> inputMessagesFilterGif
	MEDIA_PHOTOS_ONLY = 7  // Photo -> inputMessagesFilterPhotos
	MEDIA_VIDEOS_ONLY = 8  // Video -> inputMessagesFilterVideo
	//MEDIA_VOICE_FILE = 9  // VoiceFile -> inputMessagesFilterVoice
	//MEDIA_CHAT_PHOTO = 10 // ChatPhoto -> inputMessagesFilterChatPhotos
	//MEDIA_ROUND_FILE = 11 // RoundFile ->inputMessagesFilterRoundVideo,
	//MEDIA_PINNED     = 12 // Pinned -> inputMessagesFilterPinned
)

func GetMediaType(message *Message) int32 {
	if message == nil {
		return MEDIA_EMPTY
	}

	switch message.GetPredicateName() {
	case Predicate_messageService:
		switch message.GetAction().GetPredicateName() {
		case Predicate_messageActionPhoneCall:
			return MEDIA_PHONE_CALL
		}
	case Predicate_message:
		switch message.GetMedia().GetPredicateName() {
		case Predicate_messageMediaPhoto:
			return MEDIA_PHOTOVIDEO
		case Predicate_messageMediaDocument:
			if IsVoiceMessage(message) || IsRoundVideoMessage(message) {
				return MEDIA_AUDIO
			} else if IsVideoMessage(message) {
				return MEDIA_PHOTOVIDEO
			} else if IsStickerMessage(message) || IsAnimatedStickerMessage(message) {
				return MEDIA_EMPTY
			} else if IsNewGifMessage(message) {
				return MEDIA_EMPTY
			} else if IsMusicMessage(message) {
				return MEDIA_MUSIC
			} else {
				return MEDIA_FILE
			}
		}

		for _, entity := range message.GetEntities() {
			switch entity.PredicateName {
			case Predicate_messageEntityUrl,
				Predicate_messageEntityTextUrl,
				Predicate_messageEntityEmail:
				return MEDIA_URL
			case Predicate_messageActionPhoneCall:
				return MEDIA_PHONE_CALL
			}
		}
	}
	return MEDIA_EMPTY
}

type MessagesFilterType int8

const (
	FilterEmpty      MessagesFilterType = 0
	FilterPhotos     MessagesFilterType = 1
	FilterVideo      MessagesFilterType = 2
	FilterPhotoVideo MessagesFilterType = 3
	FilterDocument   MessagesFilterType = 4
	FilterUrl        MessagesFilterType = 5
	FilterGif        MessagesFilterType = 6
	FilterVoice      MessagesFilterType = 7
	FilterMusic      MessagesFilterType = 8
	FilterChatPhotos MessagesFilterType = 9
	FilterPhoneCalls MessagesFilterType = 10
	FilterRoundVoice MessagesFilterType = 11
	FilterRoundVideo MessagesFilterType = 12
	FilterMyMentions MessagesFilterType = 13
	FilterGeo        MessagesFilterType = 14
	FilterContacts   MessagesFilterType = 15
	FilterPinned     MessagesFilterType = 16
)

func FromMessagesFilter(filter *MessagesFilter) MessagesFilterType {
	r := FilterEmpty
	switch filter.PredicateName {
	case Predicate_inputMessagesFilterEmpty:
		r = FilterEmpty
	case Predicate_inputMessagesFilterPhotos:
		r = FilterPhotos
	case Predicate_inputMessagesFilterVideo:
		r = FilterVideo
	case Predicate_inputMessagesFilterPhotoVideo:
		r = FilterPhotoVideo
	case Predicate_inputMessagesFilterDocument:
		r = FilterDocument
	case Predicate_inputMessagesFilterUrl:
		r = FilterUrl
	case Predicate_inputMessagesFilterGif:
		r = FilterGif
	case Predicate_inputMessagesFilterVoice:
		r = FilterVoice
	case Predicate_inputMessagesFilterMusic:
		r = FilterMusic
	case Predicate_inputMessagesFilterChatPhotos:
		r = FilterChatPhotos
	case Predicate_inputMessagesFilterPhoneCalls:
		r = FilterPhoneCalls
	case Predicate_inputMessagesFilterRoundVoice:
		r = FilterRoundVoice
	case Predicate_inputMessagesFilterRoundVideo:
		r = FilterRoundVideo
	case Predicate_inputMessagesFilterMyMentions:
		r = FilterMyMentions
	case Predicate_inputMessagesFilterGeo:
		r = FilterGeo
	case Predicate_inputMessagesFilterContacts:
		r = FilterContacts
	case Predicate_inputMessagesFilterPinned:
		r = FilterPinned
	}

	return r
}

// GetMessagesFilterType
// TODO(@benqi): other
func GetMessagesFilterType(msg *Message) MessagesFilterType {
	r := FilterEmpty

	switch msg.PredicateName {
	case Predicate_message:
	case Predicate_messageService:
		action := msg.Action
		switch action.PredicateName {
		case Predicate_messageActionPhoneCall:
			r = FilterPhoneCalls
		}
	}
	return r
}
