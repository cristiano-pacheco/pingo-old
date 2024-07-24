// Package configdm contains the configuration domain model.
package configdm

type Config struct {
	Env             Env
	APIBaseURL      BaseURL
	FrontEndBaseURL BaseURL
}

func New(env, apiBaseURL, frontendBaseURL string) (*Config, error) {
	envVo, err := NewEnv(env)
	if err != nil {
		return nil, err
	}

	apiBaseURLVo, err := NewBaseURL(apiBaseURL)
	if err != nil {
		return nil, err
	}

	frontEndBaseURLVo, err := NewBaseURL(frontendBaseURL)
	if err != nil {
		return nil, err
	}

	return &Config{
		Env:             *envVo,
		APIBaseURL:      *apiBaseURLVo,
		FrontEndBaseURL: *frontEndBaseURLVo,
	}, nil
}
