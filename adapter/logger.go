package adapter

import (
	"fmt"

	"github.com/rs/zerolog"
)

func Get(logger *zerolog.Logger) *adapter {
	return &adapter{logger: logger}
}

type adapter struct {
	logger *zerolog.Logger
}

func (s *adapter) Fatal(v ...interface{}) {
	s.logger.Fatal().Msg(fmt.Sprint(v...))
}

func (s *adapter) Fatalf(format string, v ...interface{}) {
	s.logger.Fatal().Msgf(format, v...)
}

func (s *adapter) Print(v ...interface{}) {
	s.logger.Info().Msg(fmt.Sprint(v...))
}

func (s *adapter) Println(v ...interface{}) {
	s.logger.Info().Msg(fmt.Sprintln(v...))
}

func (s *adapter) Printf(format string, v ...interface{}) {
	s.logger.Info().Msgf(format, v...)
}
