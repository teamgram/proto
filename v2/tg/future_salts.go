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
	FutureSalt    = mt.FutureSalt
	TLFutureSalt  = mt.TLFutureSalt
	FutureSalts   = mt.FutureSalts
	TLFutureSalts = mt.TLFutureSalts
)

var (
	MakeFutureSalts = mt.MakeFutureSalts
	MakeFutureSalt  = mt.MakeFutureSalt
)
