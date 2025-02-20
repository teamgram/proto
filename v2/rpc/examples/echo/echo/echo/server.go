// Code generated by Kitex v0.7.1. DO NOT EDIT.
package echo

import (
	api "github.com/teamgram/proto/v2/rpc/examples/echo/echo"

	server "github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler api.RPCEcho, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
