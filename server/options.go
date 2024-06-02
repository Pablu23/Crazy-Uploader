package server

import (
	"net/http"

	"github.com/rs/zerolog"
)

type Options struct {
	TLS      *TlsOptions
	Port     int
	FilePath string
	Secret   *string
	Mux      *http.ServeMux
	Logger   *LoggerOptions
}

type TlsOptions struct {
	KeyPath  string
	CertPath string
}

type LoggerOptions struct {
	Level       zerolog.Level
	PrettyPrint bool
	FileLogs    *FileLogOptions
}

type FileLogOptions struct {
	Path        string
	MaxAge      int
	MaxBackupts int
}

func NewDefaultOptions() *Options {
	return &Options{
		TLS:      nil,
		Port:     8080,
		FilePath: "~/files",
		Secret:   nil,
		Mux:      http.NewServeMux(),
		Logger: &LoggerOptions{
			Level:       zerolog.TraceLevel,
			PrettyPrint: true,
			FileLogs:    nil,
		},
	}
}
