package server

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Server struct {
	mux      *http.ServeMux
	tls      *TlsOptions
	filePath string
	secret   *string
	port     int
	// log      *zerolog.Logger
}

func New(opts ...func(*Options)) *Server {
	options := NewDefaultOptions()

	for _, opt := range opts {
		opt(options)
	}

	var console io.Writer = os.Stderr
	if options.Logger.PrettyPrint {
		console = zerolog.ConsoleWriter{Out: os.Stderr}
	}

	zerolog.SetGlobalLevel(options.Logger.Level)
	if options.Logger.FileLogs != nil {
		log.Logger = log.Output(zerolog.MultiLevelWriter(console, &lumberjack.Logger{
			Filename:   options.Logger.FileLogs.Path,
			MaxAge:     options.Logger.FileLogs.MaxAge,
			MaxBackups: options.Logger.FileLogs.MaxBackupts,
		}))
	} else {
		log.Logger = log.Output(console)
	}

	return &Server{
		mux:      options.Mux,
		tls:      options.TLS,
		filePath: options.FilePath,
		secret:   options.Secret,
		port:     options.Port,
		// log:      &logger,
	}
}

func (s *Server) RegisterRoutes() {
	s.mux.HandleFunc("GET /", s.GetList)
	s.mux.HandleFunc("GET /test", s.GetTest)
}

func (s *Server) Start() error {
	server := http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: s.mux,
	}
	if s.tls != nil {
		log.Info().Str("key", s.tls.KeyPath).Str("cert", s.tls.CertPath).Int("port", s.port).Msg("Starting Server with TLS")
		return server.ListenAndServeTLS(s.tls.KeyPath, s.tls.CertPath)
	} else {
		log.Info().Int("port", s.port).Msg("Starting Server")
		return server.ListenAndServe()
	}
}
