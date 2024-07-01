package dto

// NOTE: It is used for mongodb data insertion on signup directly.
type SignupReq struct {
	EmailId   string `json:"emailId" bson:"emailId"`
	Password  string `json:"password" bson:"password"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}

type SigninReq struct {
	EmailId  string `json:"emailId"`
	Password string `json:"password"`
}

type SigninRes struct {
	Success bool `json:"success"`
}
