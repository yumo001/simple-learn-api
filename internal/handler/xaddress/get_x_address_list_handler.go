package xaddress

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/yumo001/simple-learn-api/internal/logic/xaddress"
	"github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"
)

// swagger:route post /x_address/list xaddress GetXAddressList
//
// Get x address list | 获取XAddress信息列表
//
// Get x address list | 获取XAddress信息列表
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: XAddressListReq
//
// Responses:
//  200: XAddressListResp

func GetXAddressListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.XAddressListReq
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := xaddress.NewGetXAddressListLogic(r.Context(), svcCtx)
		resp, err := l.GetXAddressList(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
