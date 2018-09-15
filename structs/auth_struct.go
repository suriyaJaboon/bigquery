package structs

type Users struct {
	UUID     string `json:"uuid" form:"-"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type TokenAuthentication struct {
	Status int `json:"status" form:"status"`
	Token string `json:"token" form:"token"`
	Message string `json:"message" form:"message"`
}
