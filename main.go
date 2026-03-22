package main

import (
	"fmt"
	"pet2/company"
	"pet2/server"
)

func main() {
	company := company.NewCompany()
	httphandlers := server.NewHTTPhandlers(company)
	httpserver := server.NewHTTPServer(httphandlers)

	if err := httpserver.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}
