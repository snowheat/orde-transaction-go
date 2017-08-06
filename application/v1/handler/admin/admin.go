package OTAdminHdl

import (
	"fmt"

	//DmOrder "github.com/snowheat/order-transaction-go/v1/utilities"
	//DmShipping "github.com/snowheat/order-transaction-go/v1/utilities"
	OTUtilities "github.com/snowheat/order-transaction-go/v1/utilities"

	"github.com/valyala/fasthttp"
)

func GetOrderDetail(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "GetOrderDetail")
}

func ValidateOrder(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "ValidateOrder")
}

func CreateShipping(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "CreateShipping")
}

func AddShippingToOrder(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "AddShippingToOrder")
}
