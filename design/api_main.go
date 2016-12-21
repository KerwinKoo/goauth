package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// This is the goauth application API design used by goa to generate
// the application code, client, tests, documentation etc.
var _ = API("goauth", func() {
	Title("The virtual wine cellar")
	Description("A basic example of a WiFiDog Auth-server API implemented with goa")
	Contact(func() {
		Name("goauth")
		Email("gukaiqiang@gmail.com")
	})
	Host("localhost:8081")
	Scheme("http")
	BasePath("/wifidog")
})
