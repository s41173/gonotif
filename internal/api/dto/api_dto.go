package dto

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"admin"`
	Password string `json:"password" binding:"required" example:"123456"`
}

type OauthRequest struct {
	TokenID string `json:"id_token" binding:"required" example:""`
}

type LoginResponse struct {
	Token string `json:"token"`
	Type  string `json:"type"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

type DecodeResponse struct {
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"invalid request"`
}

type ReqOtpRequest struct {
	Username string `json:"username" example:"0812345"`
}

type VerifyRequest struct {
	Phone string `json:"username" example:"0812345678"`
	Otp   int    `json:"otp" example:"1234"`
}

type ForgotRequest struct {
	Username string `json:"username" example:"admin"`
	Password string `json:"password" example:"123456"`
	Otp      int    `json:"otp" example:"1234"`
}

// Customer Session untuk redis
type CustomerSession struct {
	UserID      int    `json:"user_id"`
	Code        string `json:"code"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Chapter     int    `json:"chapter"`
	ChapterCode string `json:"chapter_code"`
	Token       string `json:"token"`
	Device      string `json:"device"`
	Image       string `json:"image"`
	LoginAt     string `json:"login_at"`
}
