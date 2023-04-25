package handlers

import (
	"context"
	"github.com/julienschmidt/httprouter"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
)

const (
	ordersURL = "/api/orders"
	orderURL  = "/api/orders/:id"
	usersURL  = "/api/users"
	userURL   = "/api/users/:id"
)

type RootHandler struct {
	ctx      context.Context
	services *service.CoreService
}

func NewRootHandler(ctx context.Context, services *service.CoreService) *RootHandler {
	return &RootHandler{
		ctx:      ctx,
		services: services,
	}
}

func (rh *RootHandler) InitRoutes() *httprouter.Router {
	router := httprouter.New()

	// order end-points
	router.POST(ordersURL, rh.createOrder)
	router.GET(ordersURL, rh.getAllOrders)
	router.GET(orderURL, rh.getOrderByID)
	router.DELETE(orderURL, rh.deleteOrder)
	//

	// user end-points
	router.POST(usersURL, rh.createUser)
	router.GET(usersURL, rh.getAllUsers)
	router.GET(userURL, rh.getUserByID)
	router.DELETE(userURL, rh.deleteUser)
	//

	return router
}
