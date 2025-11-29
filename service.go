package echo

import (
	"fmt"

	"github.com/springmove/sptty"
)

const (
	ServiceEcho = "echo"
)

type Service struct {
	sptty.BaseService

	cfg Config
	e   *Echo
}

func (s *Service) ServiceName() string {
	return ServiceEcho
}

func (s *Service) Init(app sptty.ISptty) error {

	if err := sptty.GetApp().GetConfig(s.ServiceName(), &s.cfg); err != nil {
		return err
	}

	go func() {
		if err := s.Srv().Start(s.cfg.Port); err != nil {
			sptty.Log(sptty.ErrorLevel, fmt.Sprintf("Echo Server Err: %s", err.Error()), s.ServiceName())
			return
		}
	}()

	return nil
}

func (s *Service) Srv() *Echo {
	if s.e == nil {
		s.e = New()
	}

	return s.e
}
