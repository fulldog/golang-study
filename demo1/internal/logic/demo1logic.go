package logic

import (
	"context"

	"demo1/internal/svc"
	"demo1/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Demo1Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemo1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Demo1Logic {
	return &Demo1Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Demo1Logic) Demo1(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
