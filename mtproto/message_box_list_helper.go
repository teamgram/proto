// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mtproto

func (m *MessageBoxList) Visit(
	toUserId int64,
	cb1 func(messageList []*Message),
	cb2 func(userIdList []int64),
	cb3 func(chatIdList []int64),
	cb4 func(channelIdList []int64)) {
	var (
		idHelper    = NewIDListHelper(toUserId)
		messageList = make([]*Message, 0, len(m.BoxList))
	)

	for _, box := range m.BoxList {
		message := box.ToMessage(toUserId)
		messageList = append(messageList, message)
		idHelper.PickByMessage(message)
	}

	if cb1 != nil {
		cb1(messageList)
	}

	idHelper.Visit(cb2, cb3, cb4)
}

func (m *MessageBoxList) Length() int32 {
	return int32(len(m.BoxList))
}

func (m *MessageBoxList) Walk(cb func(idx int, v *MessageBox)) {
	if cb == nil {
		return
	}

	for i, v := range m.BoxList {
		cb(i, v)
	}
}
