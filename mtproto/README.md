# mtproto 手动修改

## move to sub messages

## user

```
    // FIXME(@benqi):
    if m.GetRestricted() == true {
        // if m.GetRestrictionReason_FLAGVECTORRESTRICTIONREASON() != nil {
            x.Int(int32(CRC32_vector))
            x.Int(int32(len(m.GetRestrictionReason_FLAGVECTORRESTRICTIONREASON())))
            for _, v := range m.GetRestrictionReason_FLAGVECTORRESTRICTIONREASON() {
                x.Bytes((*v).Encode(layer))
            }
        x.Int(int32(CRC32_vector))
        x.Int(int32(len(m.GetRestrictionReason_FLAGVECTORRESTRICTIONREASON())))
        for _, v := range m.GetRestrictionReason_FLAGVECTORRESTRICTIONREASON() {
            x.Bytes((*v).Encode(layer))
        }
        // }
    }

```

## pollResults

```
				x.Int(m.GetTotalVoters().Value)
			}

			x.VectorInt(m.GetRecentVoters())
			if m.GetRecentVoters() != nil {
				x.VectorInt(m.GetRecentVoters())
			}

			return x.GetBuf()
		},
		0x5755785a: func() []byte {
				m.SetTotalVoters(&types.Int32Value{Value: dBuf.Int()})
			}

			m.SetRecentVoters(dBuf.VectorInt())
			if (flags & (1 << 3)) != 0 {
				m.SetRecentVoters(dBuf.VectorInt())
			}

			return dBuf.GetError()
		},
		0x5755785a: func() error {
			x.UInt(flags)

			x.Bytes(m.GetPoll().Encode(layer))
			x.VectorBytes(m.GetCorrectAnswers())
			
			if m.GetCorrectAnswers() != nil {
				x.VectorBytes(m.GetCorrectAnswers())
			}

			return x.GetBuf()
		},
		0x6b3765b: func() []byte {
			m28.Decode(dBuf)
			m.SetPoll(m28)

			m.SetCorrectAnswers(dBuf.VectorBytes())
			if (flags & (1 << 0)) != 0 {
				m.SetCorrectAnswers(dBuf.VectorBytes())
			}

			return dBuf.GetError()
		},
		0x6b3765b: func() error {
```

## poll

```
		0xd5529d06: func() []byte {
			// poll#d5529d06 id:long flags:# closed:flags.0?true public_voters:flags.1?true multiple_choice:flags.2?true quiz:flags.3?true question:string answers:Vector<PollAnswer> = Poll;
			x.UInt(0xd5529d06)
			
			x.Long(m.GetId())

			// set flags
			var flags uint32 = 0


			x.UInt(flags)

			x.Long(m.GetId())
			x.String(m.GetQuestion())

			x.Int(int32(CRC32_vector))
	var decodeF = map[uint32]func() error{
		0xd5529d06: func() error {
			// poll#d5529d06 id:long flags:# closed:flags.0?true public_voters:flags.1?true multiple_choice:flags.2?true quiz:flags.3?true question:string answers:Vector<PollAnswer> = Poll;
			m.SetId(dBuf.Long())
			flags := dBuf.UInt()
			_ = flags

			m.SetId(dBuf.Long())
			if (flags & (1 << 0)) != 0 {
				m.SetClosed(true)
			}
```

## future_salt
```
func (m *TLFutureSalt) Encode(layer int32) []byte {
	x := NewEncodeBuf(512)

	var encodeF = map[uint32]func() []byte{
		0x949d9dc: func() []byte {
			// future_salt#0949d9dc valid_since:int valid_until:int salt:long = FutureSalt;
			// FIXME(@benqi)
			// x.UInt(0x949d9dc)
			// no flags

			x.Int(m.GetValidSince())
			x.Int(m.GetValidUntil())
			x.Long(m.GetSalt())
			return x.GetBuf()
		},
	}

	clazzId := GetClazzID(Predicate_future_salt, int(layer))
	if f, ok := encodeF[uint32(clazzId)]; ok {
		return f()
	} else {
		// TODO(@benqi): handle error
		log.Errorf("not found clazzId by (%s, %d)", Predicate_future_salt, layer)
		return x.GetBuf()
	}

	return x.GetBuf()
}
```

## fileLocation

```
func (m *FileLocation) Encode(layer int32) []byte {
	predicateName := m.PredicateName
	if predicateName == "" {
		if n, ok := clazzIdNameRegisters2[int32(m.Constructor)]; ok {
			predicateName = n
		}
	}

	if layer >= 98 {
		predicateName = Predicate_fileLocationToBeDeprecated
	}

	switch predicateName {
	case Predicate_fileLocationToBeDeprecated:
		t := m.To_FileLocationToBeDeprecated()
		return t.Encode(layer)
	case Predicate_fileLocationUnavailable:
		t := m.To_FileLocationUnavailable()
		return t.Encode(layer)
	case Predicate_fileLocation:
		t := m.To_FileLocation()
		return t.Encode(layer)

	default:
		err := fmt.Errorf("invalid predicate error: %s", m.PredicateName)
		log.Errorf(err.Error())
		return []byte{}
	}
}
```