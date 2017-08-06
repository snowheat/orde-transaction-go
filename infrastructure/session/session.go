package IFSession

import (
	"encoding/json"
	"fmt"

	"github.com/kataras/go-sessions"
	"github.com/valyala/fasthttp"
)

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type Coupon struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	DiscountType  int     `json:"discount_type"` // 1 Percentage , 2 Nominal
	DiscountValue float32 `json:"discount_value"`
	DateStart     string  `json:"date_start"`
	DateEnd       string  `json:"date_end"`
	Quantity      int     `json:"quantity"`
}

type Customer struct {
	Id          int
	Name        string
	PhoneNumber string
	Email       string
	Address     string
}

func InitSession(ctx *fasthttp.RequestCtx) sessions.Session {
	sess := sessions.StartFasthttp(ctx)

	if sess.Get("Products") == nil {

		var productsInit []Product
		productsInit = append(productsInit, Product{Id: 1, Name: "Jilbab Merah Muda All Size", Price: 70000, Quantity: 31})
		productsInit = append(productsInit, Product{Id: 2, Name: "Blouse Kuning Strip Biru All Size", Price: 125000, Quantity: 17})
		productsInit = append(productsInit, Product{Id: 3, Name: "Lovera Fruit Duffel Sling Bag", Price: 230000, Quantity: 6})
		productsInit = append(productsInit, Product{Id: 4, Name: "Griava Two Tone Ankle Socks", Price: 18000, Quantity: 22})

		productsInitJson, _ := json.Marshal(productsInit)

		sess.Set("Products", string(productsInitJson))

	}

	if sess.Get("Coupons") == nil {

		var couponsInit []Coupon
		couponsInit = append(couponsInit, Coupon{Id: 1, Name: "Kupon Diskon Kemerdekaan", DiscountType: 1, DiscountValue: 15.00, DateStart: "2017-08-01 00:00:01", DateEnd: "2017-08-20 11:59:59", Quantity: 10})
		couponsInit = append(couponsInit, Coupon{Id: 2, Name: "Kupon Hari Buruh", DiscountType: 2, DiscountValue: 20000.00, DateStart: "2017-05-01 00:00:01", DateEnd: "2017-05-10 11:59:59", Quantity: 5})

		couponsInitJson, _ := json.Marshal(couponsInit)

		sess.Set("Coupons", string(couponsInitJson))
	}

	if sess.Get("Customer") == nil {

		var customersInit []Customer
		customersInit = append(customersInit, Customer{Id: 1, Name: "Ruppina", PhoneNumber: "022-91236982", Email: "ruppi@gmail.com", Address: "Bandung karawaci garden tower 5A/13 Bandung"})
		customersInit = append(customersInit, Customer{Id: 2, Name: "Luffy", PhoneNumber: "021-3287332", Email: "luffy@live.com", Address: "Jalan Kesederahaan Kaum Buruh No. 43 Jakarta"})

		customersInitJson, _ := json.Marshal(customersInit)

		sess.Set("Customer", string(customersInitJson))

		fmt.Println("Init session based database (for example only) :", sess.GetAll())
	}

	return sess
}
