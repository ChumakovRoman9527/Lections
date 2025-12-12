package auth

import (
	"14-TestingAPI/configs"
	"14-TestingAPI/pkg/jwt"
	"14-TestingAPI/pkg/req"
	"14-TestingAPI/pkg/res"
	"net/http"
)

type AuthHandler struct {
	*configs.Config
	*AuthService
}

type AuthHandlerDeps struct {
	*configs.Config
	*AuthService
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
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

		token, err := j.Create(jwt.JWTData{Email: body.Email})
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

func (handler *AuthHandler) Register() http.HandlerFunc {
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

		token, err := j.Create(jwt.JWTData{Email: email})
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
