// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

func (m *ReplyMarkup) ToForwardMessage() (replyMarkup *ReplyMarkup) {
	if m.GetPredicateName() != Predicate_replyInlineMarkup {
		return
	}

	for _, row := range m.GetRows() {
		for _, button := range row.GetButtons() {
			if button.GetPredicateName() != Predicate_keyboardButtonUrl {
				return
			}
		}
	}

	replyMarkup = m
	return
}
