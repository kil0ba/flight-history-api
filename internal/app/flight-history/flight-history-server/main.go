package flighthistoryserver

import (
	"context"
	"time"

	server_config "github.com/kil0ba/flight-history-api/internal/app/flight-history/flight-history-server/server-config"
	"github.com/kil0ba/flight-history-api/internal/app/store"
	jwt_utils "github.com/kil0ba/flight-history-api/internal/app/utils/jwt-utils"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type Duration int64

const (
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute
)

func New(ctx context.Context, config *server_config.Config) *server_config.FlightHistoryServer {
	log := logrus.New()

	switch config.DebugLevel {
	case "Trace":
		log.SetLevel(logrus.TraceLevel)
	case "Debug":
		log.SetLevel(logrus.DebugLevel)
	case "Info":
		log.SetLevel(logrus.InfoLevel)
	case "Warning":
		log.SetLevel(logrus.WarnLevel)
	case "Error":
		log.SetLevel(logrus.ErrorLevel)
	case "Fatal":
		log.SetLevel(logrus.FatalLevel)
	case "Panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		logrus.Warn("No debug level provided, setting 'Error'")
		log.SetLevel(logrus.ErrorLevel)
	}

	if config.Secret == "" {
		log.Fatal("Secret not provided!")
	}

	log.Info("Logger Initialized")

	flightStore := store.New(ctx, config.Db, log)

	if config.JwtSecret == "" {
		log.Panic("No JWT secret provided")
	}

	jwtManager := &jwt_utils.JWTManager{
		SecretKey:     config.JwtSecret,
		TokenDuration: time.Duration(24 * Hour),
	}

	return &server_config.FlightHistoryServer{
		Config:     config,
		Log:        log,
		Store:      flightStore,
		JwtManager: jwtManager,
	}
}
