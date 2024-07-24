package resetpasswordhandler

type input struct {
	ID       string `json:"id"`
	Token    string `json:"token"`
	Password string `json:"password"`
}
