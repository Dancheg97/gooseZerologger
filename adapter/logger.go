package adapter

import (
	"fmt"
	"strings"

	"github.com/rs/zerolog"
)

// Transform zerolog logger to goose compatible logger with by using adapter.
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
	s.logger.Fatal().Msgf(fmt.Sprintf(format, v...))
}

func (s *adapter) Print(v ...interface{}) {
	s.logger.Info().Msg(fmt.Sprint(v...))
}

func (s *adapter) Println(v ...interface{}) {
	s.logger.Info().Msg(fmt.Sprint(v...))
}

func (s *adapter) Printf(format string, v ...interface{}) {
	if strings.HasSuffix(format, `\n`) {
		format = format[:len(format)-2]
	}
	s.logger.Info().Msg(fmt.Sprintf(format, v...))
}
