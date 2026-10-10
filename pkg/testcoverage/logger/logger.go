package logger

import (
	"bytes"
	"time"

	"github.com/rs/zerolog"
)

//nolint:gochecknoglobals // relax
var (
	buffer bytes.Buffer
	L      zerolog.Logger
)

func Init() { // coverage-ignore
	zerolog.DurationFieldFormat = zerolog.DurationFormatString
	w := zerolog.ConsoleWriter{Out: &buffer, TimeFormat: time.DateTime}
	L = zerolog.New(w).With().Timestamp().Logger()
}

func Destruct() {
	L = zerolog.Logger{}
	buffer = bytes.Buffer{}
}

func Bytes() []byte {
	return buffer.Bytes()
}
