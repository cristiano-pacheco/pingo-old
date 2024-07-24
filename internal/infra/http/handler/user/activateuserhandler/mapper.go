package activateuserhandler

import (
	"encoding/base64"
	"fmt"

	"github.com/cristiano-pacheco/pingo/internal/application/usecase/user/activateuseruc"
)

func mapInputToUseCaseInput(in *input) (*activateuseruc.Input, error) {
	decodedToken, err := base64.StdEncoding.DecodeString(in.Token)
	if err != nil {
		err := fmt.Errorf("not possible to decode the token")
		return nil, err
	}

	useCaseInput := &activateuseruc.Input{
		ID:               in.ID,
		AccountConfToken: string(decodedToken),
	}

	return useCaseInput, nil
}
