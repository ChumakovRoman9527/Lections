package auth

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	TOKEN string `json:"token"`
}

type RegisterRequest struct {
	LoginRequest
	Name string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	TOKEN string `json:"token"`
}
