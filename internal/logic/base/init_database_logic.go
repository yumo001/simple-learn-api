package base

import (
	"context"
	"github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitDatabaseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInitDatabaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitDatabaseLogic {
	return &InitDatabaseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx}
}

func (l *InitDatabaseLogic) InitDatabase() (resp *types.BaseMsgResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.BaseMsgResp{
		Msg:  "success",
		Code: 0,
	}

	baseResp, err := l.svcCtx.ExampleRpc.InitDatabase(l.ctx, nil)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
	}

	resp.Msg = baseResp.Msg
	return resp, err
}
