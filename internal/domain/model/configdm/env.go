package configdm

import "fmt"

type Env struct {
	value string
}

const Dev = "dev"
const Staging = "staging"
const Production = "production"

func NewEnv(value string) (*Env, error) {
	err := validateEnv(value)
	if err != nil {
		return nil, err
	}
	return &Env{value: value}, nil
}

func (v Env) String() string {
	return v.value
}

func validateEnv(value string) error {
	if value != Dev && value != Staging && value != Production {
		return fmt.Errorf("the env %s value is invalid", value)
	}
	return nil
}
