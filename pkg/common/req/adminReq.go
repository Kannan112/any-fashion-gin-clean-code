package req

type SuperAdmin struct {
	Email    string ` json:"email" validate:"required" binding:"email"`
	Password string ` json:"password" validate:"required"`
}

type CreateAdmin struct {
	Name     string ` json:"user_name" validate:"required"`
	Email    string ` json:"email" validate:"required" binding:"email"`
	Password string ` json:"password" validate:"required"`
}

// {
//     "user_name":"testadmin",
//     "email":"testdmin@gmail.com",
//     "password":"1234567"
// }

type BlockData struct {
	UserId uint   ` json:"userid" validate:"required"`
	Reason string ` json:"reason" validate:"required"`
}
type UserEmail struct {
	Email string `json:"email"`
}
