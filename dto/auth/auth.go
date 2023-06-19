package authdto

type RegisterRequest struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender string `json:"gender" validate:"required"`
	Phone string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}

type LoginRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Gender string `json:"gender"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	Role string `json:"role"`
	Avatar string `json:"avatar"`
	Token string `json:"token"`
}