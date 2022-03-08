package routes

import (
	"API/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	Func:         controllers.Login,
	AuthRequired: false,
}
