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
	"math"
)

const (
	ScheduledStateOK     int32 = 0
	ScheduledStateSent   int32 = 1
	ScheduledStateDelete int32 = 2
)

const (
	// TODO:
	// MaxNebulaChatUserID      = (1 << 48) - 1 // 281474976710655
	// MinNebulaChatChatID      = 1 << 48       // 281474976710656
	// MaxNebulaChatChatID      = (2 << 48) - 1 // 562949953421311
	// MinNebulaChatMegagroupID = 2 << 48       // 562949953421312
	// MaxNebulaChatMegagroupID = (3 << 48) - 1 // 844424930131967
	// MinNebulaChatChannelID   = 3 << 48       // 844424930131968
	// MaxNebulaChatChannelID   = (4 << 48) - 1 // 1125899906842623

	MinNebulaChatChatID    = 1              // 1
	MaxNebulaChatChatID    = 1073741824 - 1 // 1
	MinNebulaChatChannelID = 1073741824     // 1073801854
	MaxNebulaChatChannelID = math.MaxInt32
)

//func PeerIdIsUser(id int64) bool {
//	return id>>48 == 0
//}
//
//func PeerIdIsChat(id int64) bool {
//	return id>>48 == 1
//}
//
//func PeerIdIsMegagroup(id int64) bool {
//	return id>>48 == 2
//}

func ChatIdIsChat(id int64) bool {
	return id < MinNebulaChatChannelID
}

func ChatIdIsChannel(id int64) bool {
	return id >= MinNebulaChatChannelID
}

//
//func MakePeerUtilByPeerId(id int64) *PeerUtil {
//	t := id >> 48
//	switch t {
//	case 0:
//		return MakeUserPeerUtil(id)
//	case 1:
//		return MakeChatPeerUtil(id)
//	case 2, 3:
//		return MakeChannelPeerUtil(id)
//	default:
//		panic("invalid peer_id " + strconv.FormatInt(id, 10))
//	}
//}

func ToSafeUsers(users []*User) []*User {
	if users == nil {
		return []*User{}
	} else {
		return users
	}
}

/*
## Inline bot samples
Here are some sample inline bots, in case you’re curious to see one in action. Try any of these:
@gif – GIF search
@vid – Video search
@pic – Yandex image search
@bing – Bing image search
@wiki – Wikipedia search
@imdb – IMDB search
@bold – Make bold, italic or fixed sys text

## NEW
@youtube - Connect your account for personalized results
@music - Search and send classical music
@foursquare – Find and send venue addresses
@sticker – Find and send stickers based on emoji
*/
const (
	BotFatherId   = int64(6)
	BotFatherName = "BotFather"
)

const (
	BotGifId          = int64(101)
	BotGifName        = "gif"
	BotVidId          = int64(102)
	BotVidName        = "vid"
	BotPicId          = int64(103)
	BotPicName        = "pic"
	BotBingId         = int64(104)
	BotBingName       = "bing"
	BotWikiId         = int64(105)
	BotWikiName       = "wiki"
	BotImdbId         = int64(106)
	BotImdbName       = "imdb"
	BotBoldId         = int64(107)
	BotBoldName       = "bold"
	BotYoutubeId      = int64(108)
	BotYoutubeName    = "youtube"
	BotMusicId        = int64(109)
	BotMusicName      = "music"
	BotFoursquareId   = int64(110)
	BotFoursquareName = "foursquare"
	BotStickerId      = int64(111)
	BotStickerName    = "sticker"
)

var (
	botIdNameMap = map[int64]string{
		BotFatherId:     BotFatherName,
		BotGifId:        BotGifName,
		BotVidId:        BotVidName,
		BotPicId:        BotPicName,
		BotBingId:       BotBingName,
		BotWikiId:       BotWikiName,
		BotImdbId:       BotImdbName,
		BotBoldId:       BotBoldName,
		BotYoutubeId:    BotYoutubeName,
		BotMusicId:      BotMusicName,
		BotFoursquareId: BotFoursquareName,
		BotStickerId:    BotStickerName,
	}

	botNameIdMap = map[string]int64{
		BotFatherName:     BotFatherId,
		BotGifName:        BotGifId,
		BotVidName:        BotVidId,
		BotPicName:        BotPicId,
		BotBingName:       BotBingId,
		BotWikiName:       BotWikiId,
		BotImdbName:       BotImdbId,
		BotBoldName:       BotBoldId,
		BotYoutubeName:    BotYoutubeId,
		BotMusicName:      BotMusicId,
		BotFoursquareName: BotFoursquareId,
		BotStickerName:    BotStickerId,
	}
)

func GetBotNameById(id int64) (n string) {
	n, _ = botIdNameMap[id]
	return
}

func GetBotIdByName(n string) (id int64) {
	id, _ = botNameIdMap[n]
	return
}

func IsBotFather(id int64) bool {
	return id == BotFatherId
}

func IsBotBing(id int64) bool {
	return id == BotBingId
}

func IsBotPic(id int64) bool {
	return id == BotPicId
}

func IsBotGif(id int64) bool {
	return id == BotGifId
}

func IsBotFoursquare(id int64) bool {
	return id == BotFoursquareId
}
