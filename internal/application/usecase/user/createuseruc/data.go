package createuseruc

type Input struct {
	Name     string
	Email    string
	Password string
}

type Output struct {
	ID    string
	Name  string
	Email string
}
