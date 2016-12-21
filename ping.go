package main

import (
	"log"

	"github.com/KerwinKoo/goauth/app"
	"github.com/goadesign/goa"
)

// PingController implements the ping resource.
type PingController struct {
	*goa.Controller
}

// NewPingController creates a ping controller.
func NewPingController(service *goa.Service) *PingController {
	return &PingController{Controller: service.NewController("PingController")}
}

// AnswerPing runs the AnswerPing action.
func (c *PingController) AnswerPing(ctx *app.AnswerPingPingContext) error {
	resp := "hello wifidog"
	log.Println("Receive wifidog's ping method and response:", resp)
	return ctx.OK([]byte(resp))
	return nil
}
