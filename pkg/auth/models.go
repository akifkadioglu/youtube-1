package auth

type BodyRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BodyLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
