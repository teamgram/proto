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

func MakeChatPhotoByPhoto(photo *Photo) (chatPhoto *ChatPhoto) {
	// chatPhoto#1c6e1c11 flags:#
	//	has_video:flags.0?true
	//	photo_id:long
	//	stripped_thumb:flags.1?bytes
	//	dc_id:int = ChatPhoto;
	switch photo.GetPredicateName() {
	case Predicate_photo:
		sizes := photo.GetSizes()
		if len(sizes) == 0 {
			// TODO(@benqi): log
			chatPhoto = MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
		} else {
			// TODO(@benqi): check PhotoSize is photoSizeEmpty
			chatPhoto = MakeTLChatPhoto(&ChatPhoto{
				HasVideo:      len(photo.VideoSizes) > 0,
				PhotoId:       photo.GetId(),
				StrippedThumb: nil,
				DcId:          1,
			}).To_ChatPhoto()
		}
	default:
		chatPhoto = MakeTLChatPhotoEmpty(nil).To_ChatPhoto()
	}
	return

}

func MakeUserProfilePhotoByPhoto(photo *Photo) (userProfilePhoto *UserProfilePhoto) {
	switch photo.GetPredicateName() {
	case Predicate_photo:
		sizes := photo.GetSizes()
		if len(sizes) > 0 {
			userProfilePhoto = MakeTLUserProfilePhoto(&UserProfilePhoto{
				HasVideo:      len(photo.VideoSizes) > 0,
				PhotoId:       photo.GetId(),
				StrippedThumb: nil,
				DcId:          1,
			}).To_UserProfilePhoto()
		}
	case Predicate_photoEmpty:
		userProfilePhoto = MakeTLUserProfilePhotoEmpty(nil).To_UserProfilePhoto()
	}
	return
}
