package xaddress

import (
	"context"

	"github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"
	"github.com/yumo001/simple-learn-rpc/types/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteXAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteXAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteXAddressLogic {
	return &DeleteXAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteXAddressLogic) DeleteXAddress(req *types.IDsReq) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.DeleteXAddress(l.ctx, &example.IDsReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, err
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
