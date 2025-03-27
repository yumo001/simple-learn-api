package xaddress

import (
    "context"

    "github.com/yumo001/simple-learn-api/internal/svc"
    "github.com/yumo001/simple-learn-api/internal/types"
    "github.com/yumo001/simple-learn-rpc/types/example"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type GetXAddressByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetXAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetXAddressByIdLogic {
	return &GetXAddressByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetXAddressByIdLogic) GetXAddressById(req *types.IDReq) (resp *types.XAddressInfoResp, err error) {
	data, err := l.svcCtx.ExampleRpc.GetXAddressById(l.ctx, &example.IDReq{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &types.XAddressInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  l.svcCtx.Trans.Trans(l.ctx, i18n.Success),
		},
		Data: types.XAddressInfo{
            BaseIDInfo: types.BaseIDInfo{
                Id:        data.Id,
                CreatedAt: data.CreatedAt,
                UpdatedAt: data.UpdatedAt,
            },
        	UserId: data.UserId,
        	Default: data.Default,
        	FirstName: data.FirstName,
        	LastName: data.LastName,
        	CountryId: data.CountryId,
        	Street: data.Street,
        	Province: data.Province,
        	City: data.City,
        	PostalCode: data.PostalCode,
        	Phone: data.Phone,
		},
	}, nil
}

