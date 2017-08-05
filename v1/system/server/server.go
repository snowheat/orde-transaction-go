package OTServer

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	SDConfig "github.com/snowheat/order-transaction-go/v1/config"
)

//Set ...
func Set(s *fasthttp.Server, router *fasthttprouter.Router) {
	s.Handler = router.Handler
	s.Name = SDConfig.SERVER_NAME
	//s.Concurrency = 500
	s.MaxConnsPerIP = SDConfig.MAX_CONNECTION_PER_IP
	s.DisableHeaderNamesNormalizing = true
	s.MaxRequestsPerConn = SDConfig.MAX_REQUEST_PER_CONNECTION
}

//Run ...
func Run(s *fasthttp.Server) {
	fmt.Println(`
   _____       __          __             __  
  / ___/____ _/ /__  _____/ /_____  _____/ /__
  \__ \/ __ '/ / _ \/ ___/ __/ __ \/ ___/ //_/
 ___/ / /_/ / /  __(__  ) /_/ /_/ / /__/ '<   
/____/\__,_/_/\___/____/\__/\____/\___/_/|_|


 *** Service(s) for handling Order Transaction ***

 Powered by Go & Valyala/Fasthttp

 Developed by Muhammad Insan Al-Amin

 Now listening on: http://localhost:` + SDConfig.PORT + `
 API documentation: http://localhost:` + SDConfig.PORT + `/v1/documentation

 Application started. Press CTRL+C to shut down.
	`)
	log.Fatal(s.ListenAndServe(":" + SDConfig.PORT))
}
