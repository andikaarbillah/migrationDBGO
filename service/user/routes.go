package user

import (
	"ecom/service/auth"
	"ecom/types"
	"ecom/utils"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct{
	store types.UserStore
}

func NewHandler(store types.UserStore)* Handler{
	return &Handler{store: store}
}

func (h *Handler)RegisterRoutes(router *mux.Router){
	router.HandleFunc("login",h.handleLogin).Methods("POST")
	router.HandleFunc("register",h.handleRegister).Methods("POST")
}

func (h *Handler)handleLogin(w http.ResponseWriter, r *http.Request){

}

func (h *Handler)handleRegister(w http.ResponseWriter, r *http.Request){
	var payload types.RegisterUserPayload

	if err:= utils.ParseJSON(r, &payload); err != nil{
		utils.WriterError(w, http.StatusBadRequest, err)
		return
	}

	if err:= utils.Validate.Struct(payload);err != nil{
		errors:= err.(validator.ValidationErrors)
		utils.WriterError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	_, err:= h.store.GetUserByEmail(payload.Email)

	if err == nil{
		utils.WriterError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}

	hashedPassword, err:= auth.HashPassword(payload.Password)
	if err != nil{
		utils.WriterError(w, http.StatusInternalServerError, err)
	}
	err = h.store.CreateUser(types.User{
		FisrtName: payload.FisrtName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})

	if err != nil{
		utils.WriterError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}