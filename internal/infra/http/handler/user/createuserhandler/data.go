package createuserhandler

type input struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type output struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
