// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package main

import (
	"github.com/teamgram/proto/example/internal/server"
)

func main() {
	svr := server.New()

	svr.Initialize()
	svr.RunLoop()
	svr.Destroy()
}
