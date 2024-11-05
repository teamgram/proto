// Copyright (c) 2021-present,  Teamgram Studio (https://teamgram.io).
//  All rights reserved.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package server

import (
	"flag"
	"log"

	"github.com/cloudwego/kitex/server"
	"github.com/teamgram/proto/example/example/codec"
	"github.com/teamgram/proto/example/example/echo"
	"github.com/teamgram/proto/example/internal/config"
	"github.com/teamgram/proto/example/internal/server/mtproto"
	"github.com/teamgram/proto/example/internal/svc"
)

var configFile = flag.String("f", "etc/example.yaml", "the config file")

type Server struct {
	server.Server
}

func New() *Server {
	return new(Server)
}

func (s *Server) Initialize() error {
	var c config.Config
	ctx := svc.NewServiceContext(c)
	_ = ctx

	cCodec := codec.NewJsonCodec(true)
	s.Server = echo.NewServer(new(mtproto.EchoImpl), server.WithCodec(cCodec))
	return nil
}

func (s *Server) RunLoop() {
	if err := s.Server.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}

func (s *Server) Destroy() {
}
