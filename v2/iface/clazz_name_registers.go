// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package iface

var (
	clazzNameRegisters2   = make(map[string]map[int]int32)
	clazzIdNameRegisters2 = make(map[int32]string)
)

func RegisterClazzName(clazzName string, layer int, clazzId int32) {
	if _, ok := clazzNameRegisters2[clazzName]; !ok {
		clazzNameRegisters2[clazzName] = make(map[int]int32)
	}
	clazzNameRegisters2[clazzName][layer] = clazzId
	clazzIdNameRegisters2[clazzId] = clazzName
}

func GetClazzIDByName(clazzName string, layer int) int32 {
	if m, ok := clazzNameRegisters2[clazzName]; ok {
		m2, ok2 := m[layer]
		if ok2 {
			return m2
		}
		m2, ok2 = m[0]
		if ok2 {
			return m2
		}
	}
	return 0
}

func RegisterClazzIDName(clazzName string, clazzId int32) {
	clazzIdNameRegisters2[clazzId] = clazzName
}

func GetClazzNameByID(clazzId int32) string {
	if clazzName, ok := clazzIdNameRegisters2[clazzId]; ok {
		return clazzName
	}
	return ""
}
