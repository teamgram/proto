// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

type Options struct {
	// store method name mapping of java -> go.
	// use the annotation method name + parameter types as the unique identifier.
	MethodNames map[string]string
}
