package flighthistoryserver

import (
	"github.com/kil0ba/flight-history-api/internal/app/store"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type FlightHistoryServer struct {
	Log    *logrus.Logger
	Config *Config
	Store  *store.Store
}

func New(config *Config) *FlightHistoryServer {
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

	log.Info("Logger Initialized")

	flightStore := store.New(config.Db, log)

	return &FlightHistoryServer{
		Config: config,
		Log:    log,
		Store:  flightStore,
	}
}
