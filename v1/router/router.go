package OTRouter

import (
	"github.com/buaazp/fasthttprouter"
	OTMetaHandler "github.com/snowheat/order-transaction-go/v1/handlers/meta"
)

//Set ...
func Set(router *fasthttprouter.Router) {
	router.GET("/meta/author", OTMetaHandler.Author)
}
