package sendresetpasswordemailuc

type Input struct {
	Email string
}

type resetPasswordTemplateVars struct {
	Name              string
	ResetPasswordLink string
}
