// Copyright © 2025 The Teamgram Authors.
//  All Rights Reserved.
//
// Author: @benqi (wubenqi@gmail.com)

package mtproto

/*
*****************************************************************

	// profileTabPosts#b98cd696 = ProfileTab;
	// profileTabGifts#4d4bd46a = ProfileTab;
	// profileTabMedia#72c64955 = ProfileTab;
	// profileTabFiles#ab339c00 = ProfileTab;
	// profileTabMusic#9f27d26e = ProfileTab;
	// profileTabVoice#e477092e = ProfileTab;
	// profileTabLinks#d3656499 = ProfileTab;
	// profileTabGifs#a2c0f695 = ProfileTab;

******************************************************************
*/
const (
	ProfileTabTypeUnknown = 0
	ProfileTabTypePosts   = 1
	ProfileTabTypeGifts   = 2
	ProfileTabTypeMedia   = 3
	ProfileTabTypeFiles   = 4
	ProfileTabTypeMusic   = 5
	ProfileTabTypeVoice   = 6
	ProfileTabTypeLinks   = 7
	ProfileTabTypeGifs    = 8
)

func FromProfileTabToType(tabType *ProfileTab) int32 {
	switch tabType.GetPredicateName() {
	case Predicate_profileTabPosts:
		return ProfileTabTypePosts
	case Predicate_profileTabGifts:
		return ProfileTabTypeGifts
	case Predicate_profileTabMedia:
		return ProfileTabTypeMedia
	case Predicate_profileTabFiles:
		return ProfileTabTypeFiles
	case Predicate_profileTabMusic:
		return ProfileTabTypeMusic
	case Predicate_profileTabVoice:
		return ProfileTabTypeVoice
	case Predicate_profileTabLinks:
		return ProfileTabTypeLinks
	case Predicate_profileTabGifs:
		return ProfileTabTypeGifs
	default:
		return ProfileTabTypeUnknown
	}
}

func ToProfileTabByType(tabType int32) *ProfileTab {
	switch tabType {
	case ProfileTabTypePosts:
		return MakeTLProfileTabPosts(nil).To_ProfileTab()
	case ProfileTabTypeGifts:
		return MakeTLProfileTabGifts(nil).To_ProfileTab()
	case ProfileTabTypeMedia:
		return MakeTLProfileTabMedia(nil).To_ProfileTab()
	case ProfileTabTypeFiles:
		return MakeTLProfileTabFiles(nil).To_ProfileTab()
	case ProfileTabTypeMusic:
		return MakeTLProfileTabMusic(nil).To_ProfileTab()
	case ProfileTabTypeVoice:
		return MakeTLProfileTabVoice(nil).To_ProfileTab()
	case ProfileTabTypeLinks:
		return MakeTLProfileTabLinks(nil).To_ProfileTab()
	case ProfileTabTypeGifs:
		return MakeTLProfileTabGifs(nil).To_ProfileTab()
	default:
		return nil
	}
}
