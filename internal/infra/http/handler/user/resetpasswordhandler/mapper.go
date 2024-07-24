package resetpasswordhandler

import (
	"encoding/base64"
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/resetpassworduc"
)

func mapInputToUseCaseInput(in *input) (*resetpassworduc.Input, error) {
	decodedToken, err := base64.StdEncoding.DecodeString(in.Token)
	if err != nil {
		err := fmt.Errorf("not possible to decode the token")
		return nil, err
	}

	useCaseInput := &resetpassworduc.Input{
		ID:                 in.ID,
		ResetPasswordToken: string(decodedToken),
		Password:           in.Password,
	}

	return useCaseInput, nil
}
