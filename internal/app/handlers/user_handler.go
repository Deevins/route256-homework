package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
)

var (
	errUserDuplicate = errors.New("ERROR: duplicate key value violates unique constraint \"users_email_key\" (SQLSTATE 23505)")
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UserResponseDTO struct {
	ID       *models.UserID `json:"id,omitempty"`
	Username *string        `json:"username,omitempty"`
	Email    *string        `json:"email,omitempty"`
}

func (rh *RootHandler) createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var unmarshalledUser UserResponseDTO

	_ = json.NewDecoder(r.Body).Decode(&unmarshalledUser)

	w.Header().Set("Content-Type", "application/json")

	if unmarshalledUser.Username == nil || unmarshalledUser.Email == nil {
		log.Printf("error while unmarshalling request body, err: [%s]", "invalid input body")
		w.WriteHeader(http.StatusBadRequest)
		res := Response{
			Message: "invalid input body",
			Status:  http.StatusBadRequest,
		}
		jsonStr, err := json.Marshal(&res)
		if err != nil {
			return
		}

		_, _ = w.Write(jsonStr)
		return
	}

	id, err := rh.services.User.CreateUser(rh.ctx, models.Username(*unmarshalledUser.Username), models.UserEmail(*unmarshalledUser.Email))
	if err != nil {
		if errors.Is(err, errUserDuplicate) {
			w.WriteHeader(http.StatusConflict)

			res := Response{
				Message: "user with this email already exists",
				Status:  http.StatusConflict,
			}
			jsonStr, err := json.Marshal(&res)
			if err != nil {
				return
			}

			_, _ = w.Write(jsonStr)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		res := Response{
			Message: "service failure",
			Status:  http.StatusInternalServerError,
		}
		jsonStr, err := json.Marshal(&res)
		if err != nil {
			return
		}

		_, _ = w.Write(jsonStr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	res := Response{
		Message: fmt.Sprintf("user with id:%s created", strconv.FormatUint(uint64(id), 10)),
		Status:  http.StatusCreated,
	}
	jsonStr, err := json.Marshal(&res)
	if err != nil {
		return
	}

	_, _ = w.Write(jsonStr)
}

func (rh *RootHandler) getAllUsers(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	users, err := rh.services.User.GetAll(rh.ctx)
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

		_, _ = w.Write(jsonStr)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	jsonStr, err := json.Marshal(users)
	if err != nil {
		return
	}

	_, _ = w.Write(jsonStr)
}

func (rh *RootHandler) getUserByID(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	// ask how to fix calling GetAllUsers handler in case param empty
	userID := params.ByName("id")

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedID, _ := models.ParseValueToUserID(userID)

	user, err := rh.services.User.GetByID(rh.ctx, parsedID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(err.Error())

		res := Response{
			Message: "service failure",
			Status:  http.StatusInternalServerError,
		}
		jsonStr, err := json.Marshal(&res)
		if err != nil {
			return
		}

		_, _ = w.Write(jsonStr)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// create mapper for DTO later
	m := &models.User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	jsonStr, err := json.Marshal(m)
	if err != nil {
		return
	}

	_, _ = w.Write(jsonStr)

}

func (rh *RootHandler) deleteUser(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
	userID := params.ByName("id")

	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedID, _ := models.ParseValueToUserID(userID)
	fmt.Println(parsedID)

	isDeleted, _ := rh.services.User.DeleteUser(rh.ctx, parsedID)
	if !isDeleted {
		//res := Response{
		//	Message: "ID is empty",
		//	Status:  http.StatusInternalServerError,
		//}
		//jsonStr, _ := json.Marshal(&res)
		//w.Write(jsonStr)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Message: fmt.Sprintf("user with ID %s successfully deleted", userID),
		Status:  http.StatusOK,
	}
	jsonStr, err := json.Marshal(&res)
	if err != nil {
		return
	}

	_, _ = w.Write(jsonStr)
}
