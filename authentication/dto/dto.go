package dto

// NOTE: It is used for mongodb data insertion on signup directly.
type SignupReq struct {
	EmailId  string `json:"emailId" bson:"emailId" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}

type SignupData struct {
	EmailId  string `json:"emailId" bson:"emailId"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type SigninReq struct {
	EmailId  string `json:"emailId" bson:"emailId" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}

type SigninRes struct {
	Success bool `json:"success"`
}

type SignupRes struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}

type UserInfo struct {
	EmailId string `json:"emailId"`
	Role    string `json:"role"`
}
