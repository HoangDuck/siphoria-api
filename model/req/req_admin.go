package req

type RequestSignInStaff struct {
	Email    string `json:"email,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}

type RequestUpdatePasswordStaff struct {
	OldPassword string `json:"old_password,omitempty" validate:"required"`
	NewPassword string `json:"new_password,omitempty" validate:"required"`
}

type RequestGetAccountInfoStaff struct {
	AccountID string `json:"account_id" validate:"required"`
}

type RequestGetStaffProfileByAccountId struct {
	StaffId string `json:"staff_id" validate:"required"`
}

type RequestCreateStaffAccount struct {
	Email         string `json:"email,omitempty" validate:"required"`
	Password      string `json:"password,omitempty" validate:"required"`
	Phone         string `json:"phone,omitempty" validate:"required"`
	Role          string `json:"role"`
	StaffID       string `json:"staff_id"`
	StatusAccount int    `json:"status_account"`
}

type RequestUpdateAccount struct {
	ID        string `json:"id"`
	Status    int    `json:"status"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RequestActivateAccountStaff struct {
	AccountID string `json:"ID" validate:"required"`
}

type RequestDeactivateAccountStaff struct {
	AccountID string `json:"ID" validate:"required"`
}

type RequestAddStaffProfile struct {
	Position         string `json:"position"`
	StatusWork       int    `json:"status_work"`
	HomeTown         string `json:"home_town"`
	Ethnic           string `json:"ethnic"`
	IdentifierNumber string `json:"identifier_number"`
	Email            string `json:"email"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Phone            string `json:"phone"`
	Gender           bool   `json:"gender"`
	DateOfBirth      string `json:"dob"`
	Address          string `json:"address"`
}

type RequestUpdateStaffProfile struct {
	StaffID          string `json:"staff_id"`
	Position         string `json:"position"`
	StatusWork       int    `json:"status_work"`
	HomeTown         string `json:"home_town"`
	Ethnic           string `json:"ethnic"`
	IdentifierNumber string `json:"identifier_number"`
	Email            string `json:"email"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Phone            string `json:"phone"`
	Gender           bool   `json:"gender"`
	DateOfBirth      string `json:"dob"`
	Address          string `json:"address"`
}

type RequestChangeRoleName struct {
	AccountID string `json:"ID" validate:"required"`
	Role      string `json:"role" validate:"required"`
}

type RequestChangeAvatarStaff struct {
	AccountID string `json:"ID" validate:"required"`
	AvatarUrl string `json:"avatar_url" validate:"required"`
}

type RequestAddStatusBooking struct {
	StatusCode  string `json:"status_code" validate:"required"`
	StatusName  string `json:"status_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type RequestAddStatusPayment struct {
	StatusCode  string `json:"status_code" validate:"required"`
	StatusName  string `json:"status_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type RequestAddStatusAccount struct {
	StatusCode  string `json:"status_code" validate:"required"`
	StatusName  string `json:"status_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type RequestAddStatusWork struct {
	StatusCode  string `json:"status_code" validate:"required"`
	StatusName  string `json:"status_name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type RequestCreateAccountByAdmin struct {
	Email     string `json:"email,omitempty" validate:"required"`
	Password  string `json:"password,omitempty" validate:"required"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role" validate:"required"`
}

type RequestUpdateCommissionRating struct {
	CommissionRate float32 `json:"commission_rate,omitempty" validate:"required"`
}

type RequestUpdateRating struct {
	Rating float32 `json:"rating,omitempty" validate:"required"`
}

type RequestPushNotificationAdmin struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	FCMKey      string `json:"fcm_key" validate:"required"`
}

type RequestApprovePayout struct {
	Resolve bool `json:"resolve,omitempty"`
}

type RequestSaveHotelWorkByEmployee struct {
	UserId  string `json:"user_id" validate:"required"`
	HotelId string `json:"hotel_id" validate:"required"`
}

type RequestDeleteHotelWorkByEmployee struct {
	UserId  string `json:"user_id" validate:"required"`
	HotelId string `json:"hotel_id" validate:"required"`
}
