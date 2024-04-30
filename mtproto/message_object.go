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
	//MessageText     = 1

	MessageEmpty     = 0
	MessageTypePhoto = 1
	MessageTypePoll  = 17
)

/*
	public void setType() {
	    int oldType = type;
	    isRoundVideoCached = 0;
	    if (messageOwner instanceof TLRPC.TL_message || messageOwner instanceof TLRPC.TL_messageForwarded_old2) {
	        if (isMediaEmpty()) {
	            type = 0;
	            if (TextUtils.isEmpty(messageText) && eventId == 0) {
	                messageText = "Empty message";
	            }
	        } else if (messageOwner.media.ttl_seconds != 0 && (messageOwner.media.photo instanceof TLRPC.TL_photoEmpty || messageOwner.media.document instanceof TLRPC.TL_documentEmpty)) {
	            contentType = 1;
	            type = 10;	// EncryptedPhoto
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaPhoto) {
	            type = 1;	// Photo
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaGeo || messageOwner.media instanceof TLRPC.TL_messageMediaVenue || messageOwner.media instanceof TLRPC.TL_messageMediaGeoLive) {
	            type = 4;	// Geo: messageMediaGeo || messageMediaVenue || messageMediaGeoLive
	        } else if (isRoundVideo()) {
	            type = 5;	// RoundVideo
	        } else if (isVideo()) {
	            type = 3;	// Video
	        } else if (isVoice()) {
	            type = 2;	// Voice
	        } else if (isMusic()) {
	            type = 14;	// Music
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaContact) {
	            type = 12;	// Contact
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaPoll) {
	            type = TYPE_POLL; // Poll
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaUnsupported) {
	            type = 0;	// Unsupported
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaDocument) {
	            if (messageOwner.media.document != null && messageOwner.media.document.mime_type != null) {
	                if (isGifDocument(messageOwner.media.document)) {
	                    type = 8;	// Gif
	                } else if (messageOwner.media.document.mime_type.equals("image/webp") && isSticker()) {
	                    type = 13;	// Sticker
	                } else {
	                    type = 9;	// Doc
	                }
	            } else {
	                type = 9;		// Doc
	            }
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaGame) {
	            type = 0;			//
	        } else if (messageOwner.media instanceof TLRPC.TL_messageMediaInvoice) {
	            type = 0;			//
	        }
	    } else if (messageOwner instanceof TLRPC.TL_messageService) {
	        if (messageOwner.action instanceof TLRPC.TL_messageActionLoginUnknownLocation) {
	            type = 0;			//
	        } else if (messageOwner.action instanceof TLRPC.TL_messageActionChatEditPhoto || messageOwner.action instanceof TLRPC.TL_messageActionUserUpdatedPhoto) {
	            contentType = 1;
	            type = 11;			// messageActionChatEditPhoto || messageActionUserUpdatedPhoto
	        } else if (messageOwner.action instanceof TLRPC.TL_messageEncryptedAction) {
	            if (messageOwner.action.encryptedAction instanceof TLRPC.TL_decryptedMessageActionScreenshotMessages || messageOwner.action.encryptedAction instanceof TLRPC.TL_decryptedMessageActionSetMessageTTL) {
	                contentType = 1;
	                type = 10;
	            } else {
	                contentType = -1;
	                type = -1;
	            }
	        } else if (messageOwner.action instanceof TLRPC.TL_messageActionHistoryClear) {
	            contentType = -1;
	            type = -1;
	        } else if (messageOwner.action instanceof TLRPC.TL_messageActionPhoneCall) {
	            type = 16;
	        } else {
	            contentType = 1;
	            type = 10;
	        }
	    }
	    if (oldType != 1000 && oldType != type) {
	        generateThumbs(false);
	    }
	}
*/
func getMessageType(m *Message) (mType int) {
	switch m.PredicateName {
	case Predicate_message:
		// message := m.To_Message()
		if IsMediaEmpty(m) {
			mType = 0
			// mType = MessageEmpty
		} else if m.Media.TtlSeconds != nil &&
			(m.Media.Photo_FLAGPHOTO.PredicateName == Predicate_photoEmpty ||
				m.Media.Document.PredicateName == Predicate_document) {
			mType = 10 // EncryptedPhoto
		} else if m.Media.PredicateName == Predicate_messageMediaPhoto {
			mType = 1 // Photo
		} else if m.Media.PredicateName == Predicate_messageMediaGeo ||
			m.Media.PredicateName == Predicate_messageMediaVenue ||
			m.Media.PredicateName == Predicate_messageMediaGeoLive {
			mType = 4 // Geo: messageMediaGeo || messageMediaVenue || messageMediaGeoLive
		} else if IsRoundVideoMessage(m) {
			mType = 5 // RoundVideo
		} else if IsVideoMessage(m) {
			mType = 3 // RoundVideo
		} else if IsVoiceMessage(m) {
			mType = 2 // Voice
		} else if IsMusicMessage(m) {
			mType = 14 // Music
		} else if m.Media.PredicateName == Predicate_messageMediaContact {
			mType = 12 // Contact
		} else if m.Media.PredicateName == Predicate_messageMediaPoll {
			mType = 17 // Poll
		} else if m.Media.PredicateName == Predicate_messageMediaUnsupported {
			mType = 0 // Unsupported
		} else if m.Media.PredicateName == Predicate_messageMediaDocument {
			if m.Media.Document != nil && m.Media.Document.MimeType != "" {
				if IsGifDocument(m.Media.Document) {
					mType = 8 //
				} else if m.Media.Document.MimeType == "image/webp" && IsStickerMessage(m) {
					mType = 13
				} else {
					mType = 9
				}
			} else {
				mType = 9
			}
		} else if m.Media.PredicateName == Predicate_messageMediaGame {
			mType = 0
		} else if m.Media.PredicateName == Predicate_messageMediaInvoice {
			mType = 0
		}
	case Predicate_messageService:
		// FIXME
		// if m.Action.PredicateName == mtproto.Predicate_messageActionLoginUnknownLocation {
		//
		// }
		//
		// FIXME
		// m.Action.PredicateName == mtproto.Predicate_messageActionUserUpdatedPhoto {
		//

		if m.Action.PredicateName == Predicate_messageActionChatEditPhoto {
			mType = 11
		} else if m.Action.PredicateName == Predicate_messageActionHistoryClear {
			mType = -1
		} else if m.Action.PredicateName == Predicate_messageActionPhoneCall {
			mType = 16
		} else {
			mType = 10
		}
	case Predicate_messageEmpty:
	default:
	}
	return
}

func IsMediaEmpty(message *Message) bool {
	return message == nil ||
		message.GetMedia() == nil ||
		message.Media.PredicateName == Predicate_messageMediaEmpty ||
		message.Media.PredicateName == Predicate_messageMediaWebPage
}

func IsMediaEmptyWebpage(message *Message) bool {
	return message.GetMedia().GetPredicateName() == Predicate_messageMediaEmpty
}

func IsStickerMessage(message *Message) bool {
	return IsStickerDocument(message.GetMedia().GetDocument())
}

func IsLocationMessage(message *Message) bool {
	return message.GetMedia().GetPredicateName() == Predicate_messageMediaGeo ||
		message.GetMedia().GetPredicateName() == Predicate_messageMediaGeoLive ||
		message.GetMedia().GetPredicateName() == Predicate_messageMediaVenue
}

func IsMaskMessage(message *Message) bool {
	return IsMaskDocument(message.GetMedia().GetDocument())
}

func IsMusicMessage(message *Message) bool {
	if message.GetMedia().GetPredicateName() == Predicate_messageMediaWebPage {
		return IsMusicDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsMusicDocument(message.GetMedia().GetDocument())
}

func IsGifMessage(message *Message) bool {
	return IsGifDocument(message.GetMedia().GetDocument())
}

func IsRoundVideoMessage(message *Message) bool {
	if message.GetMedia().GetPredicateName() == Predicate_messageMediaWebPage {
		return IsRoundVideoDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsRoundVideoDocument(message.GetMedia().GetDocument())
}

func IsPhoto(message *Message) bool {
	if message.GetMedia().GetPredicateName() == Predicate_messageMediaWebPage {
		return message.GetMedia().GetWebpage().GetPhoto().GetPredicateName() == Predicate_photo &&
			message.GetMedia().GetWebpage().GetDocument().GetPredicateName() == Predicate_document
	}
	return message.GetMedia().GetPredicateName() == Predicate_messageMediaPhoto
}

func IsVoiceMessage(message *Message) bool {
	if message.GetMedia().GetPredicateName() == Predicate_messageMediaWebPage {
		return IsVoiceDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsVoiceDocument(message.GetMedia().GetDocument())
}

func IsNewGifMessage(message *Message) bool {
	if message.GetMedia().GetPredicateName() == Predicate_messageMediaWebPage {
		return IsNewGifDocument(message.GetMedia().GetWebpage().GetDocument())
	}
	return IsNewGifDocument(message.GetMedia().GetDocument())
}

func IsLiveLocationMessage(message *Message) bool {
	return message.GetMedia().GetPredicateName() == Predicate_messageMediaGeoLive
}

func IsVideoMessage(message *Message) bool {
	if message.GetMedia().GetPredicateName() == Predicate_messageMediaWebPage {
		return IsVideoDocument(message.Media.Webpage.Document)
	}

	return IsVideoDocument(message.GetMedia().GetDocument())
}

func IsGameMessage(message *Message) bool {
	return message.GetMedia().GetPredicateName() == Predicate_messageMediaGame
}

func IsInvoiceMessage(message *Message) bool {
	return message.GetMedia().GetPredicateName() == Predicate_messageMediaInvoice
}

/*
   public static TLRPC.InputStickerSet getInputStickerSet(TLRPC.Message message) {
       if (message.media != null && message.media.document != null) {
           for (TLRPC.DocumentAttribute attribute : message.media.document.attributes) {
               if (attribute instanceof TLRPC.TL_documentAttributeSticker) {
                   if (attribute.stickerset instanceof TLRPC.TL_inputStickerSetEmpty) {
                       return null;
                   }
                   return attribute.stickerset;
               }
           }
       }
       return null;
   }



*/

func IsForwardedMessage(message *Message) bool {
	return message.GetFwdFrom() != nil
}

func IsAnimatedStickerMessage(message *Message) bool {
	// return message.media != null && isAnimatedStickerDocument(message.media.document, !DialogObject.isSecretDialogId(message.dialog_id))
	return false
}

func IsBroadcastMessage(message *Message) bool {
	return PeerIsChannel(message.GetPeerId()) && message.GetPost()
}

func IsMegagroupMessage(message *Message) bool {
	return PeerIsChannel(message.GetPeerId()) && !message.GetPost()
}
