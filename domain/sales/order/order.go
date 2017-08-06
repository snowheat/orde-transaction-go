package DmOrder

type Order struct {
	Id                      int     `json:"id"`
	CreatedDateTime         string  `json:"created_date_time"`
	CustomerID              int     `json:"customer_id"`
	DateTime                string  `json:"order_date_time"`
	HasCoupon               int     `json:"has_coupon"`
	CouponID                int     `json:"coupon_id"`
	TotalPrice              float32 `json:"total_price"`
	TotalPriceAfterDiscount float32 `json:"total_price_after_discount"`
	SubmittedDateTime       string  `json:"submitted_date_time"`

	PaymentProof         string `json:"payment_proof"`
	PaymentProofDateTime string `json:"payment_proof_date_time"`

	IsValid              string `json:"is_valid"`
	VerificationDateTime string `json:"verification_date_time"`

	IsShipped                  int    `json:"is_shipped"`
	ShippingId                 int    `json:"shipping_id"`
	ShippingDateTime           string `json:"shipping_date_time"`
	ShippingNotificationIsSent int    `json:"shipping_notification_is_sent"`

	Products []Product
}

func Create() {

}

func AddProduct() {

}

func AddCoupon() {

}

func Submit() {

}

func AddPaymentProof() {

}

func Get() {

}

func Validate() {

}

func AddShipping() {

}
