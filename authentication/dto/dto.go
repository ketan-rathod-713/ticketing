package dto

// NOTE: It is used for mongodb data insertion on signup directly.
type SignupReq struct {
	EmailId   string `json:"emailId" bson:"emailId" validate:"required,email"`
	Password  string `json:"password" bson:"password" validate:"required,min=8"`
	FirstName string `json:"firstName" bson:"firstName" validate:"required"`
	LastName  string `json:"lastName" bson:"lastName" validate:"required"`
}

// signup request data to store in mongodb directly
type SignupData struct {
	EmailId         string `json:"emailId" bson:"emailId"`
	Password        string `json:"password" bson:"password"`
	FirstName       string `json:"firstName" bson:"firstName"`
	LastName        string `json:"lastName" bson:"lastName"`
	Role            string `json:"role" bson:"role"`
	CreatedAt       string `json:"createdAt" bson:"createdAt"`
	IsEmailVerified bool   `json:"isEmailVerified" bson:"isEmailVerified"`
}

// signin request body
type SigninReq struct {
	EmailId  string `json:"emailId" bson:"emailId" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}

// signin response
type SigninRes struct {
	Success bool `json:"success"`
}

// signup response
type SignupRes struct {
	Success bool   `json:"success"`
	Id      string `json:"id"`
}

// what to return to user // for eg. /currentuser
type UserInfo struct {
	EmailId   string `json:"emailId"`
	Role      string `json:"role"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
