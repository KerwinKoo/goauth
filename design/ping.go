package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("ping", func() {
	BasePath("/ping")

	Action("AnswerPing", func() {
		Routing(GET("/:args"))
		Description("answer wifidog ping GET-Request")
		Response(OK, "text/plain")
	})

})
