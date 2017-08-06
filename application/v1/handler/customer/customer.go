package OTCustomerHdl

import (
	"fmt"

	OTUtilities "github.com/snowheat/order-transaction-go/application/v1/utilities"
	DmProduct "github.com/snowheat/order-transaction-go/domain/products/product"
	//DmOrder "github.com/snowheat/order-transaction-go/domain/sales/order"
	"github.com/valyala/fasthttp"
)

type Product struct {
	Id       int
	Quantity int
}

func GetProductById(ctx *fasthttp.RequestCtx) {

	OTUtilities.GetJsonHttpHeader(ctx)

	id := OTUtilities.GetUserValueAsInt("id", ctx)
	jsonString := OTUtilities.GetJsonStringFromStruct(DmProduct.GetProduct(id, ctx))

	fmt.Fprint(ctx, jsonString)
}

func GetCouponById(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)

	id := OTUtilities.GetUserValueAsInt("id", ctx)
	jsonString := OTUtilities.GetJsonStringFromStruct(DmProduct.GetProduct(id, ctx))

	fmt.Fprint(ctx, jsonString)
}

func CreateOrder(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "CreateOrder")
}

func AddProductToOrder(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "AddProductToOrder")
}

func AddCouponToOrder(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "AddCouponToOrder")
}

func SubmitOrder(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "SubmitOrder")
}

func AddPaymentProofToOrder(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "AddPaymentProofToOrder")
}

func GetOrderDetail(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "AddPaymentProofToOrder")
}

func GetShippingDetail(ctx *fasthttp.RequestCtx) {
	OTUtilities.GetJsonHttpHeader(ctx)
	fmt.Fprint(ctx, "AddPaymentProofToOrder")
}
