package authenticateuseruc

type Input struct {
	Email    string
	Password string
}

type Output struct {
	Token string
}
