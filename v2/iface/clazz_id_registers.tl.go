// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package iface

var (
	clazzIdRegisters2 = make(map[int32]func() TLObject)
)

func RegisterClazzID(clazzId int32, f func() TLObject) {
	clazzIdRegisters2[clazzId] = f
}

func NewTLObjectByClazzID(clazzId int32) TLObject {
	f, ok := clazzIdRegisters2[clazzId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClazzID(clazzId int32) (ok bool) {
	_, ok = clazzIdRegisters2[clazzId]
	return
}
