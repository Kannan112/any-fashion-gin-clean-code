package req

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password"`
}
type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type GoogleAuth struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type OTPData struct {
	PhoneNumber string
}
type VerifyOtp struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}
type Address struct {
	Id           int    `json:"id"`
	House_number string `json:"house_number" binding:"required"`
	Street       string `json:"street" binding:"required"`
	City         string `json:"city" binding:"required"`
	District     string `json:"district" binding:"required"`
	Landmark     string `json:"landmark" binding:"required"`
	Pincode      int    `json:"pincode" binding:"required"`
	IsDefault    bool   `json:"isdefault" `
}
type Otpverifier struct {
	Pin   string `json:"pin,omitempty" validate:"required"`
	Phone string `json:"phoneNumber,omitempty" validate:"required"`
}

type AddAddress struct {
	House_number string `json:"house_number" binding:"required"`
	Street       string `json:"street" binding:"required"`
	City         string `json:"city" binding:"required"`
	District     string `json:"district" binding:"required"`
	Landmark     string `json:"landmark" binding:"required"`
	Pincode      int    `json:"pincode" binding:"required"`
	IsDefault    bool   `json:"isdefault" `
}
