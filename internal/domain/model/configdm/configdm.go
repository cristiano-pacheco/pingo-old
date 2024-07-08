// Package configdm contains the configuration domain model.
package configdm

type Config struct {
	Env     Env
	BaseURL BaseURL
}

func New(env, baseURL string) (*Config, error) {
	envVo, err := NewEnv(env)
	if err != nil {
		return nil, err
	}

	baseURLVo, err := NewBaseURL(baseURL)
	if err != nil {
		return nil, err
	}

	return &Config{
		Env:     *envVo,
		BaseURL: *baseURLVo,
	}, nil
}
