// Code generated by goctl. DO NOT EDIT!
// Source: demo2.proto

package demo2

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Demo2 interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultDemo2 struct {
		cli zrpc.Client
	}
)

func NewDemo2(cli zrpc.Client) Demo2 {
	return &defaultDemo2{
		cli: cli,
	}
}

func (m *defaultDemo2) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := NewDemo2Client(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
