//go:generate goagen bootstrap -d github.com/KerwinKoo/goauth/design

package main

import (
	"github.com/KerwinKoo/goauth/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("goauth")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "ping" controller
	c := NewPingController(service)
	app.MountPingController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}
}
