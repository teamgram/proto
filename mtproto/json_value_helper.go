// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

func FromJSONValue(v *JSONValue) any {
	switch v.GetPredicateName() {
	case Predicate_jsonNull:
		return nil
	case Predicate_jsonBool:
		return FromBool(v.Value_BOOL)
	case Predicate_jsonNumber:
		return v.Value_FLOAT64
	case Predicate_jsonString:
		return v.Value_STRING
	case Predicate_jsonArray:
		var (
			jArr = make([]any, 0)
		)
		for _, v2 := range v.Value_VECTORJSONVALUE {
			jArr = append(jArr, FromJSONValue(v2))
		}

		return jArr
	case Predicate_jsonObject:
		var (
			jMap = make(map[string]any)
		)
		for _, v2 := range v.Value_VECTORJSONOBJECTVALUE {
			jMap[v2.Key] = FromJSONValue(v2.Value)
		}

		return jMap
	default:
		panic("unknown json value")
		return nil
	}
}
