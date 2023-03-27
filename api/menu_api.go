package api

import "gokyrie/service"

type MenuApi struct {
	BaseApi,
	Service *service.MenuService
}
