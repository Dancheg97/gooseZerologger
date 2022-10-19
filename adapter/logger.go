package adapter

import (
	"github.com/rs/zerolog"
)

func Get(logger *zerolog.Logger) *adapter {
	return &adapter{logger: logger}
}

type adapter struct {
	logger *zerolog.Logger
}

func (s *adapter) Fatal(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (s *adapter) Fatalf(format string, v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (s *adapter) Print(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (s *adapter) Println(v ...interface{}) {
	panic("not implemented") // TODO: Implement
}

func (s *adapter) Printf(format string, v ...interface{}) {
	panic("not implemented") // TODO: Implement
}
