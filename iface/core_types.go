// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package iface

import (
	"fmt"

	"github.com/teamgram/proto/bin"
)

const (
	MTPROTO_VERSION = 2
)

type TLObject interface {
	Encode(x *bin.Encoder, layer int32)
	Decode(d *bin.Decoder) error
	// ClazzID() uint32
	// ClazzName() string
	// String() string
}

func DecodeObject(d *bin.Decoder) (TLObject, error) {
	clazzID, err := d.ClazzID()
	if err != nil {
		return nil, err
	}

	r := NewTLObjectByClazzID(clazzID)
	if r == nil {
		err = fmt.Errorf("can't find registered classId: %x", clazzID)
		return nil, err
	}

	// log.Infof("newTLObjectByClassID, classID: %x", uint32(classID))
	err = r.Decode(d)
	if err != nil {
		err = fmt.Errorf("object(%x) decode error: %v", uint32(clazzID), err)
		return nil, err
	}

	return r, nil
}

func WriteObjectList[T TLObject](e *bin.Encoder, layer int32, vList []T) {
	e.PutInt(len(vList))
	for _, obj := range vList {
		obj.Encode(e, layer)
	}
}
