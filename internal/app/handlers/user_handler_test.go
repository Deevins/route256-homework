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

	"github.com/golang/mock/gomock"
	"github.com/julienschmidt/httprouter"
	"github.com/magiconair/properties/assert"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	models_dto "gitlab.ozon.dev/daker255/homework-8/internal/app/repository/models"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	mock_service "gitlab.ozon.dev/daker255/homework-8/internal/app/services/mocks"
)

func TestRootHandler_createUser(t *testing.T) {
	//arrange
	type mockBehavior func(s *mock_service.MockUser, user models_dto.UserDTO)
	ctx := context.Background()

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models_dto.UserDTO
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: `{"username":"tests", "email":"tests@email.com"}`,
			inputUser: models_dto.UserDTO{
				Username: "tests",
				Email:    "tests@email.com",
			},
			mockBehavior: func(s *mock_service.MockUser, user models_dto.UserDTO) {
				s.EXPECT().CreateUser(ctx, user.Username, user.Email).Return(models.UserID(1), nil)

			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"message":"user with id:1 created","status":201}`,
		},
		{
			name:      "Service failure",
			inputBody: `{"username":"tests", "email":"tests@email.com"}`,
			inputUser: models_dto.UserDTO{
				Username: "tests",
				Email:    "tests@email.com",
			},
			mockBehavior: func(s *mock_service.MockUser, user models_dto.UserDTO) {
				s.EXPECT().CreateUser(ctx, user.Username, user.Email).Return(models.UserID(0), errors.New("service failure"))

			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"service failure","status":500}`,
		},
	}

	//act
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_service.NewMockUser(c)
			testCase.mockBehavior(userService, testCase.inputUser)

			services := &service.CoreService{User: userService}
			handlers := NewRootHandler(ctx, services)

			router := httprouter.New()
			router.POST("/api/users", handlers.createUser)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("POST",
				"/api/users",
				bytes.NewBufferString(testCase.inputBody))

			router.ServeHTTP(w, req)

			//assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}

}

func TestRootHandler_getAllUsers(t *testing.T) {
	//arrange
	type mockBehavior func(s *mock_service.MockUser, user models_dto.UserDTO)
	ctx := context.Background()

	expectedAllUsers, _ := json.Marshal([]*models_dto.UserDTO{
		{
			ID:       1,
			Username: "qwe",
			Email:    "qwe",
		},
		{
			ID:       2,
			Username: "qwee",
			Email:    "qwee",
		},
	})

	testTable := []struct {
		name                 string
		inputBody            string
		inputUser            models_dto.UserDTO
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: "",
			inputUser: models_dto.UserDTO{},
			mockBehavior: func(s *mock_service.MockUser, user models_dto.UserDTO) {
				s.EXPECT().GetAll(ctx).Return([]*models_dto.UserDTO{
					{
						ID:       1,
						Username: "qwe",
						Email:    "qwe",
					},
					{
						ID:       2,
						Username: "qwee",
						Email:    "qwee",
					},
				}, nil)

			},
			expectedStatusCode:   200,
			expectedResponseBody: string(expectedAllUsers),
		},
		{
			name:      "service failure",
			inputBody: "",
			inputUser: models_dto.UserDTO{},
			mockBehavior: func(s *mock_service.MockUser, user models_dto.UserDTO) {
				s.EXPECT().GetAll(ctx).Return([]*models_dto.UserDTO{}, errors.New("service error"))

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

			userService := mock_service.NewMockUser(c)
			testCase.mockBehavior(userService, testCase.inputUser)

			services := &service.CoreService{User: userService}
			handlers := NewRootHandler(ctx, services)

			router := httprouter.New()

			router.GET("/api/users", handlers.getAllUsers)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET",
				"/api/users", nil)

			router.ServeHTTP(w, req)

			//assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}

}

func TestRootHandler_getUserByID(t *testing.T) {
	//arrange
	type mockBehavior func(s *mock_service.MockUser, userID models.UserID)
	ctx := context.Background()

	testTable := []struct {
		name                 string
		inputBody            string
		userID               models.UserID
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "ok",
			inputBody: "",
			userID:    models.UserID(1),
			mockBehavior: func(s *mock_service.MockUser, userID models.UserID) {
				s.EXPECT().GetByID(ctx, userID).Return(&models_dto.UserDTO{
					ID:       models.UserID(1),
					Username: "qwe",
					Email:    "qwe",
				}, nil)
			},
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"id":1,"username":"qwe","email":"qwe"}`,
		},
	}

	//act
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctx := context.Background()
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_service.NewMockUser(c)

			services := &service.CoreService{User: userService}
			handlers := NewRootHandler(ctx, services)

			router := httprouter.New()
			router.GET("/api/users/:id", handlers.getUserByID)

			testCase.mockBehavior(userService, testCase.userID)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET",
				fmt.Sprintf("/api/users/%d", testCase.userID), nil)

			router.ServeHTTP(w, req)

			//assert
			assert.Equal(t, w.Code, testCase.expectedStatusCode)
			assert.Equal(t, w.Body.String(), testCase.expectedResponseBody)
		})
	}

}

//func TestRootHandler_deleteUser(t *testing.T) {
//	//arrange
//	type mockBehavior func(s *mock_service.MockUser, userID models.UserID)
//	ctx := context.Background()
//
//	testTable := []struct {
//		name                 string
//		inputBody            string
//		userID               models.UserID
//		mockBehavior         mockBehavior
//		expectedStatusCode   int
//		expectedResponseBody string
//	}{
//		{
//			name:      "ok",
//			inputBody: `{"id":"1"}`,
//			userID:    models.UserID(1),
//			mockBehavior: func(s *mock_service.MockUser, userID models.UserID) {
//				s.EXPECT().Delete(ctx, userID).Return(nil)
//
//			},
//			expectedStatusCode:   200,
//			expectedResponseBody: `{"message":"user with ID 1 successfully deleted","status":200}`,
//		},
//		{
//			name:                 "Empty fields",
//			inputBody:            ``,
//			mockBehavior:         func(s *mock_service.MockUser, user models.UserID) {},
//			expectedStatusCode:   400,
//			expectedResponseBody: `{"message":"ID is empty","status":400}`,
//		},
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
//			ctx := context.Background()
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			userService := mock_service.NewMockUser(c)
//			testCase.mockBehavior(userService, testCase.userID)
//
//			services := &service.CoreService{User: userService}
//			handlers := NewRootHandler(ctx, services)
//
//			router := httprouter.New()
//			router.DELETE("/api/users/:id", handlers.deleteUser)
//
//			w := httptest.NewRecorder()
//
//			req := httptest.NewRequest("DELETE",
//				fmt.Sprintf("/api/users/%d", testCase.userID), nil)
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
