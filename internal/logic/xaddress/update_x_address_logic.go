package xaddress

import (
	"context"

	"github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"
	"github.com/yumo001/simple-learn-rpc/types/example"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateXAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateXAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateXAddressLogic {
	return &UpdateXAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateXAddressLogic) UpdateXAddress(req *types.XAddressInfo) (resp *types.BaseMsgResp, err error) {
	data, err := l.svcCtx.ExampleRpc.UpdateXAddress(l.ctx,
		&example.XAddressInfo{
			Id:          req.Id,
        	UserId: req.UserId,
        	Default: req.Default,
        	FirstName: req.FirstName,
        	LastName: req.LastName,
        	CountryId: req.CountryId,
        	Street: req.Street,
        	Province: req.Province,
        	City: req.City,
        	PostalCode: req.PostalCode,
        	Phone: req.Phone,
		})
	if err != nil {
		return nil, err
	}
	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, data.Msg)}, nil
}
