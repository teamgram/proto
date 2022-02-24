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

//const (
//	/*
//		void BoxController::loadMoreRows() {
//			if (_loadRequestId || _allLoaded) {
//				return;
//			}
//
//			_loadRequestId = _api.request(MTPmessages_Search(
//				MTP_flags(0),
//				MTP_inputPeerEmpty(),
//				MTP_string(),
//				MTP_inputUserEmpty(),
//				MTP_inputMessagesFilterPhoneCalls(MTP_flags(0)),
//				MTP_int(0),
//				MTP_int(0),
//				MTP_int(_offsetId),
//				MTP_int(0),
//				MTP_int(_offsetId ? kFirstPageCount : kPerPageCount),
//				MTP_int(0),
//				MTP_int(0),
//				MTP_int(0)
//			)).done([this](const MTPmessages_Messages &result) {
//
//		   inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;
//	*/
//	/*
//		const auto filter = [&] {
//			using Type = Storage::SharedMediaType;
//			switch (type) {
//			case Type::Photo:
//				return MTP_inputMessagesFilterPhotos();
//			case Type::Video:
//				return MTP_inputMessagesFilterVideo();
//			case Type::PhotoVideo:
//				return MTP_inputMessagesFilterPhotoVideo();
//			case Type::MusicFile:
//				return MTP_inputMessagesFilterMusic();
//			case Type::File:
//				return MTP_inputMessagesFilterDocument();
//			case Type::VoiceFile:
//				return MTP_inputMessagesFilterVoice();
//			case Type::RoundVoiceFile:
//				return MTP_inputMessagesFilterRoundVoice();
//			case Type::RoundFile:
//				return MTP_inputMessagesFilterRoundVideo();
//			case Type::GIF:
//				return MTP_inputMessagesFilterGif();
//			case Type::Link:
//				return MTP_inputMessagesFilterUrl();
//			case Type::ChatPhoto:
//				return MTP_inputMessagesFilterChatPhotos();
//			}
//			return MTP_inputMessagesFilterEmpty();
//		}();
//	*/
//	/*
//	   inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
//
//	   inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
//	   inputMessagesFilterGeo#e7026d0d = MessagesFilter;
//	   inputMessagesFilterContacts#e062db83 = MessagesFilter;
//	*/
//	// Allow forward declarations.
//	Photo          = 1  // inputMessagesFilterPhotos#9609a51c = MessagesFilter;
//	Video          = 2  // inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
//	PhotoVideo     = 3  // inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
//	MusicFile      = 4  // inputMessagesFilterMusic#3751b49e = MessagesFilter;
//	File           = 5  // inputMessagesFilterDocument#9eddf188 = MessagesFilter;
//	VoiceFile      = 6  // inputMessagesFilterVoice#50f5c392 = MessagesFilter;
//	Link           = 7  // inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
//	ChatPhoto      = 8  // inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;
//	RoundVoiceFile = 9  // inputMessagesFilterRoundVoice#7a7c17a4 = MessagesFilter;
//	GIF            = 10 // inputMessagesFilterGif#ffc86587 = MessagesFilter;
//	RoundFile      = 11 // inputMessagesFilterRoundVideo#b549da53 = MessagesFilter;
//)

const (
	MEDIA_EMPTY      = -1
	MEDIA_PHOTOVIDEO = 0
	MEDIA_FILE       = 1
	MEDIA_AUDIO      = 2
	MEDIA_URL        = 3
	MEDIA_MUSIC      = 4
	MEDIA_PHONE_CALL = 5
	MEDIA_GIF        = 6
)

// GetMediaType
/*
   public final static int MEDIA_PHOTOVIDEO = 0;
   public final static int MEDIA_FILE = 1;
   public final static int MEDIA_AUDIO = 2;
   public final static int MEDIA_URL = 3;
   public final static int MEDIA_MUSIC = 4;
   public final static int MEDIA_TYPES_COUNT = 5;

   public static int getMediaType(TLRPC.Message message) {
       if (message == null) {
           return -1;
       }
       if (message.media instanceof TLRPC.TL_messageMediaPhoto) {
           return MEDIA_PHOTOVIDEO;
       } else if (message.media instanceof TLRPC.TL_messageMediaDocument) {
           if (MessageObject.isVoiceMessage(message) || MessageObject.isRoundVideoMessage(message)) {
               return MEDIA_AUDIO;
           } else if (MessageObject.isVideoMessage(message)) {
               return MEDIA_PHOTOVIDEO;
           } else if (MessageObject.isStickerMessage(message) || MessageObject.isAnimatedStickerMessage(message)) {
               return -1;
           } else if (MessageObject.isNewGifMessage(message)) {
               return -1;
           } else if (MessageObject.isMusicMessage(message)) {
               return MEDIA_MUSIC;
           } else {
               return MEDIA_FILE;
           }
       } else if (!message.entities.isEmpty()) {
           for (int a = 0; a < message.entities.size(); a++) {
               TLRPC.MessageEntity entity = message.entities.get(a);
               if (entity instanceof TLRPC.TL_messageEntityUrl || entity instanceof TLRPC.TL_messageEntityTextUrl || entity instanceof TLRPC.TL_messageEntityEmail) {
                   return MEDIA_URL;
               }
           }
       }
       return -1;
   }
*/
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

/*
##

```
inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
inputMessagesFilterPhotos#9609a51c = MessagesFilter;
inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
inputMessagesFilterDocument#9eddf188 = MessagesFilter;
inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
inputMessagesFilterGif#ffc86587 = MessagesFilter;
inputMessagesFilterVoice#50f5c392 = MessagesFilter;
inputMessagesFilterMusic#3751b49e = MessagesFilter;
inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;
inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;
inputMessagesFilterRoundVoice#7a7c17a4 = MessagesFilter;
inputMessagesFilterRoundVideo#b549da53 = MessagesFilter;
inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
inputMessagesFilterGeo#e7026d0d = MessagesFilter;
inputMessagesFilterContacts#e062db83 = MessagesFilter;
inputMessagesFilterPinned#1bb00451 = MessagesFilter;
```

## android

- 已使用
	- inputMessagesFilterEmpty
	- inputMessagesFilterPhotoVideo	---- MEDIA_PHOTOVIDEO	//
	- inputMessagesFilterDocument	---- MEDIA_FILE			//
	- inputMessagesFilterUrl		---- MEDIA_URL			// 只要有url的，都可以
	- inputMessagesFilterMusic		---- MEDIA_MUSIC
	- inputMessagesFilterChatPhotos

		```
			public void loadDialogPhotos(final int did, final int count, final long max_id, final boolean fromCache, final int classGuid) {
				if (fromCache) {
					getMessagesStorage().getDialogPhotos(did, count, max_id, classGuid);
				} else {
					if (did > 0) {
						TLRPC.User user = getUser(did);
						if (user == null) {
							return;
						}
						TLRPC.TL_photos_getUserPhotos req = new TLRPC.TL_photos_getUserPhotos();
						req.limit = count;
						req.offset = 0;
						req.max_id = (int) max_id;
						req.user_id = getInputUser(user);
						int reqId = getConnectionsManager().sendRequest(req, (response, error) -> {
							if (error == null) {
								TLRPC.photos_Photos res = (TLRPC.photos_Photos) response;
								processLoadedUserPhotos(res, did, count, max_id, false, classGuid);
							}
						});
						getConnectionsManager().bindRequestToGuid(reqId, classGuid);
					} else if (did < 0) {
						TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
						req.filter = new TLRPC.TL_inputMessagesFilterChatPhotos();
						req.limit = count;
						req.offset_id = (int) max_id;
						req.q = "";
						req.peer = getInputPeer(did);
						int reqId = getConnectionsManager().sendRequest(req, (response, error) -> {
							if (error == null) {
								TLRPC.messages_Messages messages = (TLRPC.messages_Messages) response;
								TLRPC.TL_photos_photos res = new TLRPC.TL_photos_photos();
								res.count = messages.count;
								res.users.addAll(messages.users);
								for (int a = 0; a < messages.messages.size(); a++) {
									TLRPC.Message message = messages.messages.get(a);
									if (message.action == null || message.action.photo == null) {
										continue;
									}
									res.photos.add(message.action.photo);
								}
								processLoadedUserPhotos(res, did, count, max_id, false, classGuid);
							}
						});
						getConnectionsManager().bindRequestToGuid(reqId, classGuid);
					}
				}
			}

		```

	- inputMessagesFilterPhoneCalls

		```
			private void getCalls(int max_id, final int count) {
				if (loading) {
					return;
				}
				loading = true;
				if (emptyView != null && !firstLoaded) {
					emptyView.showProgress();
				}
				if (listViewAdapter != null) {
					listViewAdapter.notifyDataSetChanged();
				}
				TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
				req.limit = count;
				req.peer = new TLRPC.TL_inputPeerEmpty();
				req.filter = new TLRPC.TL_inputMessagesFilterPhoneCalls();
				req.q = "";
				req.offset_id = max_id;
				int reqId = ConnectionsManager.getInstance(currentAccount).sendRequest(req, (response, error) -> AndroidUtilities.runOnUIThread(() -> {

		```

	- inputMessagesFilterRoundVoice	---- MEDIA_AUDIO

- 未使用
	- inputMessagesFilterPhotos
	- inputMessagesFilterVideo
	- inputMessagesFilterGif
	- inputMessagesFilterVoice
	- inputMessagesFilterRoundVideo
	- inputMessagesFilterMyMentions
	- inputMessagesFilterGeo
	- inputMessagesFilterContacts

## 规则

```
	TLRPC.TL_messages_search req = new TLRPC.TL_messages_search();
	req.limit = count;
	req.offset_id = max_id;
	if (type == MEDIA_PHOTOVIDEO) {
		req.filter = new TLRPC.TL_inputMessagesFilterPhotoVideo();
	} else if (type == MEDIA_FILE) {
		req.filter = new TLRPC.TL_inputMessagesFilterDocument();
	} else if (type == MEDIA_AUDIO) {
		req.filter = new TLRPC.TL_inputMessagesFilterRoundVoice();
	} else if (type == MEDIA_URL) {
		req.filter = new TLRPC.TL_inputMessagesFilterUrl();
	} else if (type == MEDIA_MUSIC) {
		req.filter = new TLRPC.TL_inputMessagesFilterMusic();
	}
```

*/

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

//- inputMessagesFilterPhotos
//- inputMessagesFilterVideo
//- inputMessagesFilterGif
//- inputMessagesFilterVoice
//- inputMessagesFilterRoundVideo
//- inputMessagesFilterMyMentions
//- inputMessagesFilterGeo
//- inputMessagesFilterContacts
//
//func GetMediaTypeByMessagesFilter(filter *mtproto.MessagesFilter) MediaType {
//	r := MEDIA_EMPTY
//	switch filter.PredicateName {
//	case mtproto.Predicate_inputMessagesFilterEmpty:
//		r = MEDIA_EMPTY
//	case mtproto.Predicate_inputMessagesFilterPhotos:
//		r = MEDIA_EMPTY
//	case mtproto.Predicate_inputMessagesFilterVideo:
//		r = MEDIA_EMPTY
//	case mtproto.Predicate_inputMessagesFilterPhotoVideo:
//		r = MEDIA_PHOTOVIDEO
//	case mtproto.Predicate_inputMessagesFilterDocument:
//		r = MEDIA_FILE
//	case mtproto.Predicate_inputMessagesFilterUrl:
//		r = MEDIA_URL
//	case mtproto.Predicate_inputMessagesFilterGif:
//		r = MEDIA_MUSIC
//	case mtproto.Predicate_inputMessagesFilterVoice:
//		r = FilterVoice
//	case mtproto.Predicate_inputMessagesFilterMusic:
//		r = FilterMusic
//	case mtproto.Predicate_inputMessagesFilterChatPhotos:
//		r = FilterChatPhotos
//	case mtproto.Predicate_inputMessagesFilterPhoneCalls:
//		r = FilterPhoneCalls
//	case mtproto.Predicate_inputMessagesFilterRoundVoice:
//		r = FilterRoundVoice
//	case mtproto.Predicate_inputMessagesFilterRoundVideo:
//		r = FilterRoundVideo
//	case mtproto.Predicate_inputMessagesFilterMyMentions:
//		r = FilterMyMentions
//	case mtproto.Predicate_inputMessagesFilterGeo:
//		r = FilterGeo
//	case mtproto.Predicate_inputMessagesFilterContacts:
//		r = FilterContacts
//	}
//
//	return r
//}

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

/*
package chat.channel.utils

import proto.gramchat.tl

object MessageTypeExUtil {
	enum class MessageTypeEx(val value: Int) {
		Unknown(0),
		Voice(1),
		Music(2),
		Url(3),
		Document(4),
		PhoneCalls(5),
		ChatPhotos(6),
		PhotoVideo(7),
		RoundVoice(8),
	}

	fun parseMessageTypeEx(media: tl.TLMessageMedia): MessageTypeEx {
		try {
			if (media is tl.TL_messageMediaPhoto) {
				return MessageTypeEx.PhotoVideo
			} else if (media is tl.TL_messageMediaDocument) {
				val document = media.document
				if (document is tl.TL_documentEmpty) {
					return MessageTypeEx.Unknown
				}

				document as tl.TL_document
				val attributes = document.attributes
				if (attributes == null || attributes.size == 0) {
					return MessageTypeEx.Unknown
				}

				var hasAnimated = false
				for (attribute in attributes) {
					// 有任意一个attribute设置round message为true，认为不是photoVideo类型
					if (attribute is tl.TL_documentAttributeVideo && attribute.round_message) {
						return MessageTypeEx.Unknown
					}

					if (attribute is tl.TL_documentAttributeAnimated) {
						hasAnimated = true
					}
				}

				if (hasAnimated) {
					for (attribute in attributes) {
						if (attribute is tl.TL_documentAttributeVideo && (attribute.w <= 1280 || attribute.h <= 1280)) {
							return MessageTypeEx.Unknown
						}
					}
				}

				return MessageTypeEx.PhotoVideo
			} else {
				return MessageTypeEx.Unknown
			}
		} catch (e: Exception) {
			return MessageTypeEx.Unknown
		}
	}
}
*/
