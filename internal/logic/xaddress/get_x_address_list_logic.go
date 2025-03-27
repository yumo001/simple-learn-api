package xaddress

import (
	"context"

    "github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"
	"github.com/yumo001/simple-learn-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetXAddressListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetXAddressListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetXAddressListLogic {
	return &GetXAddressListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetXAddressListLogic) GetXAddressList(req *types.XAddressListReq) (resp *types.XAddressListResp, err error) {
	data, err := l.svcCtx.ExampleRpc.GetXAddressList(l.ctx,
		&example.XAddressListReq{
			Page:        req.Page,
			PageSize:    req.PageSize,
			CreatedAt: req.CreatedAt,
			UpdatedAt: req.UpdatedAt,
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
	resp = &types.XAddressListResp{}
	resp.Msg = l.svcCtx.Trans.Trans(l.ctx, i18n.Success)
	resp.Data.Total = data.GetTotal()

	for _, v := range data.Data {
		resp.Data.Data = append(resp.Data.Data,
			types.XAddressInfo{
				BaseIDInfo: types.BaseIDInfo{
					Id:        v.Id,
					CreatedAt: v.CreatedAt,
					UpdatedAt: v.UpdatedAt,
				},
        	UserId: v.UserId,
        	Default: v.Default,
        	FirstName: v.FirstName,
        	LastName: v.LastName,
        	CountryId: v.CountryId,
        	Street: v.Street,
        	Province: v.Province,
        	City: v.City,
        	PostalCode: v.PostalCode,
        	Phone: v.Phone,
			})
	}
	return resp, nil
}
