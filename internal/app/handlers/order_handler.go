package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	"log"
	"net/http"
	"strconv"
)

type orderRequest struct {
	ProductName models.ProductName `json:"product_name"`
	Status      models.OrderStatus `json:"status,omitempty"`
	Quantity    models.Quantity    `json:"quantity"`
}

func (rh *RootHandler) createOrder(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.Header.Get("userID")

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("userID header not provided")
		return
	}

	parsedUserID, err := models.ParseValueToUserID(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		print(userID)
		json.NewEncoder(w).Encode("incorrect userID provided")
		return
	}

	var unmarshalledOrder orderRequest

	if err := json.NewDecoder(r.Body).Decode(&unmarshalledOrder); err != nil {
		log.Printf("error while unmarshalling request body, err: [%s]", err.Error())
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	id, _ := rh.services.Order.CreateOrder(
		rh.ctx,
		parsedUserID,
		unmarshalledOrder.ProductName,
		unmarshalledOrder.Quantity)

	w.WriteHeader(http.StatusCreated)

	res := Response{
		Message: fmt.Sprintf("order with ID %s created", strconv.FormatUint(uint64(id), 10)),
		Status:  http.StatusCreated,
	}
	jsonStr, err := json.Marshal(&res)
	if err != nil {
		return
	}

	w.Write(jsonStr)
}

func (rh *RootHandler) getAllOrders(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	orders, err := rh.services.Order.GetAll(rh.ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := Response{
			Message: "internal server error",
			Status:  http.StatusInternalServerError,
		}
		jsonStr, err := json.Marshal(&res)
		if err != nil {
			return
		}

		w.Write(jsonStr)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	jsonStr, err := json.Marshal(orders)
	if err != nil {
		return
	}

	w.Write(jsonStr)
}

func (rh *RootHandler) getOrderByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	// ask how to fix calling GetAllUsers handler in case param empty
	orderID := params.ByName("id")

	parsedID, _ := models.ParseValueToOrderID(orderID)

	order, err := rh.services.Order.GetByID(rh.ctx, parsedID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	userID := r.Header.Get("userID")

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("userID header not provided")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonStr, err := json.Marshal(order)
	if err != nil {
		return
	}

	w.Write(jsonStr)

}

func (rh *RootHandler) deleteOrder(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	orderID := params.ByName("id")

	if orderID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedID, _ := models.ParseValueToOrderID(orderID)

	// check userID later
	_, _ = rh.services.Order.DeleteOrder(rh.ctx, parsedID)

	w.Header().Set("Content-Type", "application/json")

	m := make(map[string]string)
	m["status"] = "ok"

	jsonStr, err := json.Marshal(&m)
	if err != nil {
		return
	}
	w.Write(jsonStr)
}
