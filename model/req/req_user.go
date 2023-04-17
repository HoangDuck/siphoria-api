package req

type RequestSignUp struct {
	Email     string `json:"email,omitempty" validate:"required"`
	FirstName string `json:"first_name,omitempty" validate:"required"`
	LastName  string `json:"last_name,omitempty" validate:"required"`
	Password  string `json:"password,omitempty" validate:"required"`
}

type RequestSignIn struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type RequestUpdateProfile struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty"`
}

type RequestUpdatePassword struct {
	OldPassword string `json:"password,omitempty" validate:"required"`
	NewPassword string `json:"new_password,omitempty" validate:"required"`
}

type RequestResetPassword struct {
	Email string `json:"email" validate:"required"`
}

type RequestNewPasswordReset struct {
	Password string `json:"password" validate:"required"`
}

type RequestChangeSettings struct {
	AllowNotification bool `json:"allow_notification" validate:"required"`
}

type RequestCheckEmail struct {
	Email string `json:"email,omitempty" validate:"required"`
}

type RequestChangeAvatar struct {
}

type RequestGetAccountInfo struct {
	AccountID string `json:"account_id" validate:"required"`
}

type RequestGetProfileByAccountId struct {
	CustomerID string `json:"customer_id" validate:"required"`
}

type RequestSignInGoogle struct {
	IDToken string `query:"id_token"`
}

type RequestGetOauthInfo struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	Iat           int    `json:"iat"`
	Exp           int    `json:"exp"`
}

type RequestUpdateRank struct {
	RankTo string `json:"rank_to" validate:"required"`
}

type RequestRefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}
