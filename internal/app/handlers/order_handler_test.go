package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/magiconair/properties/assert"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	mock_service "gitlab.ozon.dev/daker255/homework-8/internal/app/services/mocks"
)

func TestRootHandler_createOrder(t *testing.T) {
	//arrange
	type mockBehavior func(s *mock_service.MockOrder, order models.Order)
	ctx := context.Background()

	testTable := []struct {
		name                 string
		inputBody            string
		inputOrder           models.Order
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"product_name":"qwe","quantity":1}`,
			inputOrder: models.Order{
				UserID:      models.UserID(1),
				ProductName: models.ProductName("qwe"),
				Quantity:    models.Quantity(1),
			},
			mockBehavior: func(s *mock_service.MockOrder, order models.Order) {
				s.EXPECT().CreateOrder(ctx, models.UserID(1), models.ProductName("qwe"), models.Quantity(1)).Return(models.OrderID(1), nil)

			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"message":"order with ID 1 created","status":201}`,
		},
		//{
		//	name:      "empty userID header",
		//	inputBody: `{"product_name":"qwe","quantity":1}`,
		//	inputOrder: models.Order{
		//		ProductName: "qwe",
		//		Quantity:    1,
		//	},
		//	mockBehavior: func(s *mock_service.MockOrder, order models.Order) {
		//		s.EXPECT().CreateOrder(ctx, models.UserID(0), models.ProductName("qwe"), models.Quantity(1)).Return(models.OrderID(0), errors.New("not found"))
		//	},
		//	expectedStatusCode:   http.StatusNotFound,
		//	expectedResponseBody: `{"message":"empty userID header","status":404}`,
		//},
		//{
		//	name:      "Service failure",
		//	inputBody: `{"username":"tests", "email":"tests@email.com"}`,
		//	inputUser: models.User{
		//		Username: "tests",
		//		Email:    "tests@email.com",
		//	},
		//	mockBehavior: func(s *mock_service.MockUser, user models.User) {
		//		s.EXPECT().CreateUser(ctx, models.Username(user.Username), models.UserEmail(user.Email)).Return(models.UserID(1), errors.New("service failure"))
		//
		//	},
		//	expectedStatusCode:   500,
		//	expectedResponseBody: `{"message":"service failure","status":500}`,
		//},
	}

	//act
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			c := gomock.NewController(t)
			defer c.Finish()

			orderService := mock_service.NewMockOrder(c)
			testCase.mockBehavior(orderService, testCase.inputOrder)

			services := &service.CoreService{Order: orderService}
			handlers := NewRootHandler(ctx, services)

			router := httprouter.New()
			router.POST("/api/orders.proto", handlers.createOrder)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST",
				"/api/orders.proto",
				bytes.NewBufferString(testCase.inputBody))
			req.Header.Set("userID", "1")
			router.ServeHTTP(w, req)

			//assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}

}

func TestRootHandler_getAllOrders(t *testing.T) {
	//arrange
	type mockBehavior func(s *mock_service.MockOrder, order models.Order)
	ctx := context.Background()

	timeNow := time.Now()

	expectedAllOrders, _ := json.Marshal([]*models.Order{
		{
			ID:          1,
			UserID:      1,
			ProductName: "qwe",
			Status:      "approved",
			Quantity:    1,
			OrderDate:   timeNow,
		},
		{
			ID:          2,
			UserID:      2,
			ProductName: "qwee",
			Status:      "approved",
			Quantity:    1,
			OrderDate:   timeNow,
		},
	})

	testTable := []struct {
		name                 string
		inputBody            string
		inputOrder           models.Order
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:       "ok",
			inputBody:  "",
			inputOrder: models.Order{},
			mockBehavior: func(s *mock_service.MockOrder, order models.Order) {
				s.EXPECT().GetAll(ctx).Return([]*models.Order{
					{
						ID:          1,
						UserID:      1,
						ProductName: "qwe",
						Status:      "approved",
						Quantity:    1,
						OrderDate:   timeNow,
					},
					{
						ID:          2,
						UserID:      2,
						ProductName: "qwee",
						Status:      "approved",
						Quantity:    1,
						OrderDate:   timeNow,
					},
				}, nil)

			},
			expectedStatusCode:   200,
			expectedResponseBody: string(expectedAllOrders),
		},
		{
			name:       "service failure",
			inputBody:  "",
			inputOrder: models.Order{},
			mockBehavior: func(s *mock_service.MockOrder, order models.Order) {
				s.EXPECT().GetAll(ctx).Return([]*models.Order{}, errors.New("service error"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"internal server error","status":500}`,
		},
	}

	//act
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			c := gomock.NewController(t)
			defer c.Finish()

			orderService := mock_service.NewMockOrder(c)
			testCase.mockBehavior(orderService, testCase.inputOrder)

			services := &service.CoreService{Order: orderService}
			handlers := NewRootHandler(ctx, services)

			router := httprouter.New()

			router.GET("/api/orders.proto", handlers.getAllOrders)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET",
				"/api/orders.proto", nil)

			router.ServeHTTP(w, req)

			//assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}

}

func TestRootHandler_getOrderByID(t *testing.T) {
	//arrange
	type mockBehavior func(s *mock_service.MockOrder, orderID models.OrderID)
	ctx := context.Background()

	timeNow := time.Now()

	expectedOrder, _ := json.Marshal(models.Order{
		ID:          1,
		UserID:      1,
		ProductName: "qwe",
		Status:      "approved",
		Quantity:    1,
		OrderDate:   timeNow,
	})

	testTable := []struct {
		name                 string
		inputBody            string
		orderID              models.OrderID
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: "",
			orderID:   models.OrderID(1),
			mockBehavior: func(s *mock_service.MockOrder, orderID models.OrderID) {
				s.EXPECT().GetByID(ctx, orderID).Return(&models.Order{
					ID:          1,
					UserID:      1,
					ProductName: "qwe",
					Status:      "approved",
					Quantity:    1,
					OrderDate:   timeNow,
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: string(expectedOrder),
		},
		//{
		//	name:                 "incorrect order ID",
		//	inputBody:            "",
		//	orderID:              models.OrderID(emptyOrderID),
		//	mockBehavior:         func(s *mock_service.MockOrder, orderID models.OrderID) {},
		//	expectedStatusCode:   http.StatusBadRequest,
		//	expectedResponseBody: `{"message":"invalid input body"}`,
		//},
	}

	//act
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			c := gomock.NewController(t)
			defer c.Finish()

			orderService := mock_service.NewMockOrder(c)

			services := &service.CoreService{Order: orderService}
			handlers := NewRootHandler(ctx, services)

			router := httprouter.New()
			router.GET("/api/orders.proto/:id", handlers.getOrderByID)

			testCase.mockBehavior(orderService, testCase.orderID)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET",
				fmt.Sprintf("/api/orders.proto/%d", testCase.orderID), nil)
			req.Header.Set("userID", "1")
			router.ServeHTTP(w, req)

			//assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}

}

//func TestRootHandler_deleteOrder(t *testing.T) {
//	//arrange
//	type mockBehavior func(s *mock_service.MockOrder, orderID models.OrderID)
//	ctx := context.Background()
//
//	testTable := []struct {
//		name                 string
//		inputBody            string
//		orderID              models.OrderID
//		mockBehavior         mockBehavior
//		expectedStatusCode   int
//		expectedResponseBody string
//	}{
//		{
//			name:      "ok",
//			inputBody: ``,
//			orderID:   models.OrderID(1),
//			mockBehavior: func(s *mock_service.MockOrder, orderID models.OrderID) {
//				s.EXPECT().DeleteOrder(ctx, models.OrderID(1)).Return(nil)
//
//			},
//			expectedStatusCode:   200,
//			expectedResponseBody: `{"message":"order with ID 1 successfully deleted","status":200}`,
//		},
//		//{
//		//	name:                 "Empty fields",
//		//	inputBody:            ``,
//		//	mockBehavior:         func(s *mock_service.MockUser, user models.UserID) {},
//		//	expectedStatusCode:   400,
//		//	expectedResponseBody: `{"message":"ID is empty","status":400}`,
//		//},
//		//{
//		//	name:      "Service failure",
//		//	inputBody: `{"username":"tests", "email":"tests@email.com"}`,
//		//	userID:    models.UserID(1),
//		//	mockBehavior: func(s *mock_service.MockUser, userID models.UserID) {
//		//		s.EXPECT().Delete(ctx, models.UserID(1)).Return(nil)
//		//
//		//	},
//		//	expectedStatusCode:   500,
//		//	expectedResponseBody: `{"message":"service failure","status":500}`,
//		//},
//	}
//
//	//act
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			orderService := mock_service.NewMockOrder(c)
//			testCase.mockBehavior(orderService, testCase.orderID)
//
//			//services := &service.CoreService{Order: orderService}
//			//handlers := NewRootHandler(ctx, services)
//
//			router := httprouter.New()
//			//router.DELETE("/api/orders.proto/:id", handlers.deleteUser)
//
//			w := httptest.NewRecorder()
//
//			req := httptest.NewRequest("DELETE",
//				fmt.Sprintf("/api/orders.proto/%d", testCase.orderID), nil)
//
//			router.ServeHTTP(w, req)
//
//			//assert
//			assert.Equal(t, w.Code, testCase.expectedStatusCode)
//			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
//		})
//	}
//
//}
