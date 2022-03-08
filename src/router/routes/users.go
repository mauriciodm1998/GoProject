package routes

import (
	controllers "API/src/controllers"
	"net/http"
)

var UserRoutes = []Route{
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Func:         controllers.Create,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodGet,
		Func:         controllers.Get,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Func:         controllers.GetById,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodPut,
		Func:         controllers.Update,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodDelete,
		Func:         controllers.Delete,
		AuthRequired: true,
	},
	{
		URI:          "/users/{userId}/changepassword",
		Method:       http.MethodPost,
		Func:         controllers.ChangePassword,
		AuthRequired: true,
	},
}
