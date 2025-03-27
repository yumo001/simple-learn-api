package xaddress

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/yumo001/simple-learn-api/internal/logic/xaddress"
	"github.com/yumo001/simple-learn-api/internal/svc"
	"github.com/yumo001/simple-learn-api/internal/types"
)

// swagger:route post /x_address/create xaddress CreateXAddress
//
// Create x address information | 创建XAddress信息
//
// Create x address information | 创建XAddress信息
//
// Parameters:
//  + name: body
//    require: true
//    in: body
//    type: XAddressInfo
//
// Responses:
//  200: BaseMsgResp

func CreateXAddressHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.XAddressInfo
		if err := httpx.Parse(r, &req, true); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := xaddress.NewCreateXAddressLogic(r.Context(), svcCtx)
		resp, err := l.CreateXAddress(&req)
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
