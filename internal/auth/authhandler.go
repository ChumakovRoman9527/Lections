package auth

import (
	"11-JWTAUTH/configs"
	"11-JWTAUTH/pkg/jwt"
	"11-JWTAUTH/pkg/req"
	"11-JWTAUTH/pkg/res"
	"net/http"
)

type authHandler struct {
	*configs.Config
	*AuthService
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &authHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *authHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.AuthService.Login(body.Email, body.Password)
		if err != nil {
			res.Json(w, err.Error(), http.StatusUnauthorized)
			return
		}

		j := jwt.NewJWT(
			handler.Auth.Secret,
		)

		token, err := j.Create(body.Email)
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := LoginResponse{
			TOKEN: token, //handler.Config.Auth.Secret,
		}
		res.Json(w, data, http.StatusOK)
	}
}

func (handler *authHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}
		email, err := handler.AuthService.Register(body.Email, body.Password, body.Name)
		if err != nil {
			res.Json(w, err.Error(), http.StatusBadRequest)
			return
		}

		j := jwt.NewJWT(
			handler.Auth.Secret,
		)

		token, err := j.Create(email)
		if err != nil {
			res.Json(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := RegisterResponse{
			TOKEN: token, //handler.Config.Auth.Secret,
		}

		res.Json(w, data, http.StatusOK)
	}
}
