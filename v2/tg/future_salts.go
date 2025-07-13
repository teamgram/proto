// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package tg

import (
	"github.com/teamgram/proto/v2/mt"
)

type (
	FutureSalt       = mt.FutureSalt
	FutureSaltClazz  = mt.FutureSaltClazz
	TLFutureSalt     = mt.TLFutureSalt
	FutureSalts      = mt.FutureSalts
	FutureSaltsClazz = mt.FutureSaltsClazz
	TLFutureSalts    = mt.TLFutureSalts
)

var (
	MakeFutureSalts       = mt.MakeFutureSalts
	MakeFutureSalt        = mt.MakeFutureSalt
	DecodeFutureSaltClazz = mt.DecodeFutureSaltClazz
)
