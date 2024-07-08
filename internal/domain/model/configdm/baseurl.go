package configdm

import "net/url"

type BaseURL struct {
	value string
}

func NewBaseURL(value string) (*BaseURL, error) {
	err := validateBaseURL(value)
	if err != nil {
		return nil, err
	}
	return &BaseURL{value: value}, nil
}

func (v BaseURL) String() string {
	return v.value
}

func validateBaseURL(value string) error {
	_, err := url.ParseRequestURI(value)
	if err != nil {
		return err
	}
	return nil
}
