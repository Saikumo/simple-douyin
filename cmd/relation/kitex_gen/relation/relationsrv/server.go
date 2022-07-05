// Code generated by Kitex v0.3.2. DO NOT EDIT.
package relationsrv

import (
	"github.com/cloudwego/kitex/server"
	"saikumo.org/simple-douyin/cmd/relation/kitex_gen/relation"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler relation.RelationSrv, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
