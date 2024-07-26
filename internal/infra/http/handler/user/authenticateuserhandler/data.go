package authenticateuserhandler

type input struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type output struct {
	Token string `json:"token"`
}
