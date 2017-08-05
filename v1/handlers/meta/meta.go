package MetaHandler

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Author(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html; charset=utf-8")

	fmt.Fprint(ctx, "Muhammad Insan Al-Amin\n")
}
