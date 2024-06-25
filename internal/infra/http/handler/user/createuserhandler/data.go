package createuserhandler

type input struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
