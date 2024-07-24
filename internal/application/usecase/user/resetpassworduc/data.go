package resetpassworduc

type Input struct {
	Email string
}

type resetPasswordTemplateVars struct {
	Name              string
	ResetPasswordLink string
}
