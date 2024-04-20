// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

import (
	"fmt"
)

func FromBirthdayString(birthday string) (v *Birthday) {
	if len(birthday) == 0 {
		return
	}

	// 1990-01-01
	var (
		year, month, day int32
	)

	_, err := fmt.Sscanf(birthday, "%d-%d-%d", &year, &month, &day)
	if err != nil {
		year = 0
		month = 0
		day = 0
	}

	v = MakeTLBirthday(&Birthday{
		Year:  MakeFlagsInt32(year),
		Month: month,
		Day:   day,
	}).To_Birthday()

	return
}

func (m *Birthday) ToBirthdayString() (v string) {
	if m == nil {
		return
	}

	v = fmt.Sprintf("%04d-%02d-%02d", m.GetYear().GetValue(), m.Month, m.Day)

	return
}
