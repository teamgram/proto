// Copyright 2025 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

import (
	"encoding/json"
)

func UnmarshalDialogFilter(data string) (*DialogFilter, error) {
	dialogFilter := &DialogFilter{}
	if err := json.Unmarshal([]byte(data), dialogFilter); err != nil {
		// c.Logger.Errorf("jsonx.UnmarshalFromString(%v) - error: %v", v, err)
		// continue
		return nil, err
	}

	return dialogFilter.FixData(), nil
}

func (m *DialogFilter) FixData() *DialogFilter {
	if m.Title != "" {
		if m.Title_STRING == "" {
			m.Title_STRING = m.Title
		}
		if m.Title_TEXTWITHENTITIES == nil {
			m.Title_TEXTWITHENTITIES = MakeTLTextWithEntities(&TextWithEntities{
				Text:     m.Title,
				Entities: []*MessageEntity{},
			}).To_TextWithEntities()
		}
	}
	if m.Title_STRING != "" {
		if m.Title == "" {
			m.Title = m.Title_STRING
		}
		if m.Title_TEXTWITHENTITIES == nil {
			m.Title_TEXTWITHENTITIES = MakeTLTextWithEntities(&TextWithEntities{
				Text:     m.Title_STRING,
				Entities: []*MessageEntity{},
			}).To_TextWithEntities()
		}
	}
	if m.Title_TEXTWITHENTITIES != nil {
		if m.Title == "" {
			m.Title = m.Title_TEXTWITHENTITIES.GetText()
		}
		if m.Title_STRING == "" {
			m.Title_STRING = m.Title_TEXTWITHENTITIES.GetText()
		}
	}

	return m
}
