package echo

import "github.com/springmove/sptty"

type Config struct {
	sptty.BaseConfig

	Port string `yaml:"port"`
}

func (s *Config) ConfigName() string {
	return ServiceEcho
}
