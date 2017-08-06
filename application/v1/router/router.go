package OTRouter

import (
	"github.com/buaazp/fasthttprouter"
	OTAdminHdl "github.com/snowheat/order-transaction-go/application/v1/handler/admin"
	OTCustomerHdl "github.com/snowheat/order-transaction-go/application/v1/handler/customer"
)

func Set(router *fasthttprouter.Router) {
	router.GET("/v1/customer/product/:id", OTCustomerHdl.GetProductById)
	router.GET("/v1/customer/coupon/:id", OTCustomerHdl.GetCouponById)

	router.POST("/v1/customer/order", OTCustomerHdl.CreateOrder)
	router.POST("/v1/customer/order/:id/product", OTCustomerHdl.AddProductToOrder)
	router.POST("/v1/customer/order/:id/coupon", OTCustomerHdl.AddCouponToOrder)
	router.POST("/v1/customer/order/:id", OTCustomerHdl.SubmitOrder)
	router.POST("/v1/customer/order/:id/paymentproof", OTCustomerHdl.AddPaymentProofToOrder)
	router.GET("/v1/customer/order/:id", OTCustomerHdl.GetOrderDetail)

	router.GET("/v1/customer/shipping/:id", OTCustomerHdl.GetShippingDetail)

	router.GET("/v1/admin/order/:id", OTAdminHdl.GetOrderDetail)
	router.POST("/v1/admin/order/:id/validate", OTAdminHdl.ValidateOrder)

	router.POST("/v1/admin/shipping", OTAdminHdl.CreateShipping)
	router.POST("/v1/admin/order/:id/shipping", OTAdminHdl.AddShippingToOrder)

}
