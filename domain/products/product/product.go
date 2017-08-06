package DmProduct

import (
	"encoding/json"

	"github.com/valyala/fasthttp"

	IFSession "github.com/snowheat/order-transaction-go/infrastructure/session"
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func GetProduct(id int, ctx *fasthttp.RequestCtx) Product {
	sess := IFSession.InitSession(ctx)

	var product Product

	products := make([]Product, 0)

	json.Unmarshal([]byte(sess.Get("Products").(string)), &products)

	for _, v := range products {
		if v.Id == id {
			product = v
		}
	}

	return product
}
