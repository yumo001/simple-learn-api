package xaddress

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/yumo001/simple-learn-api/internal/logic/xaddress"
	"github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"
)

// swagger:route post /x_address xaddress GetXAddressById
//
// Get x address by ID | 通过ID获取XAddress信息
//
// Get x address by ID | 通过ID获取XAddress信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: IDReq
//
// Responses:
//  200: XAddressInfoResp

func GetXAddressByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := xaddress.NewGetXAddressByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetXAddressById(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
