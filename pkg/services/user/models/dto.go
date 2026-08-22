package models

// ==============================
// Login
// ==============================

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

// ==============================
// Logout
// ==============================

type LogoutRequest struct {
	LogoutFlag string `json:"logoutFlag" binding:"required,oneof=Y N"`
}

// ==============================
// Register
// ==============================

type RegisterRequest struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" binding:"required,email"`
	Mobile    string `json:"mobile" binding:"required,len=10,numeric"`
	Password  string `json:"password" binding:"required,min=8"`

	Address  string `json:"address"`
	State    string `json:"state"`
	District string `json:"district"`
	ZipCode  string `json:"zipcode"`

	Language string `json:"language"`
	SoilType string `json:"soilType"`
	Role     string `json:"role"`
}

// ==============================
// Refresh Token
// ==============================

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// ==============================
// Update User
// ==============================

type UpdateUserRequest struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email" binding:"required,email"`
	Address  string `json:"address"`
	State    string `json:"state"`
	District string `json:"district"`
	ZipCode  string `json:"zipcode"`
	Language string `json:"language"`
	SoilType string `json:"soilType"`
	Avatar   string `json:"avatar"`
}

// ==============================
// Update Location
// ==============================

type UpdateLocationRequest struct {
	Lat float64 `json:"lat" binding:"required"`
	Lng float64 `json:"lng" binding:"required"`
}

// ==============================
// Change Password
// ==============================

type ChangePasswordRequest struct {
	Email	   string `json:"email" binding:"required,email"`
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=8"`
}

// ==============================
// Forgot Password
// ==============================

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// ==============================
// Reset Password
// ==============================

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=8"`
}

// ==============================
// Update Avatar
// ==============================

type UpdateAvatarRequest struct {
	Avatar string `json:"avatar" binding:"required"`
}

// ==============================
// Common Response
// ==============================

type CommonResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==============================
// Login
// ==============================

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
	User          User   `json:"user"`
}

// ==============================
// Logout
// ==============================

type LogoutResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==============================
// Register
// ==============================

type RegisterResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`

	UserID int64 `json:"userId"`
}

// ==============================
// User Details
// ==============================

type UserDetailsResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`

	User User `json:"user"`
}

// ==============================
// Update User
// ==============================

type UpdateUserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==============================
// Refresh Token
// ==============================

type RefreshTokenResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`

	Token string `json:"token"`
}

// ==============================
// Change Password
// ==============================

type ChangePasswordResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==============================
// Forgot Password
// ==============================

type ForgotPasswordResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==============================
// Reset Password
// ==============================

type ResetPasswordResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==============================
// Update Avatar
// ==============================

type UpdateAvatarResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Avatar  string `json:"avatar"`
}

// ==============================
// Update Location
// ==============================

type UpdateLocationResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

// ==============================
// Delete User
// ==============================

type DeleteUserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
