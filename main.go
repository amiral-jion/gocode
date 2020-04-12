package main

import (
	"github.com/amiraljion/gocode/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	_ = http.ListenAndServe(":3000", nil)
}
