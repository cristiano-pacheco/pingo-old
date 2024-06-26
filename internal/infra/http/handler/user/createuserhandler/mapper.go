package createuserhandler

import "github.com/cristiano-pacheco/pingo/internal/application/usecase/user/createuseruc"

func mapInputToUseCaseInput(in input) *createuseruc.Input {
	return &createuseruc.Input{
		Name:     in.Name,
		Email:    in.Email,
		Password: in.Password,
	}
}

func mapUseCaseOutputToOutput(out *createuseruc.Output) *output {
	return &output{
		ID:    out.ID,
		Name:  out.Name,
		Email: out.Email,
	}
}
