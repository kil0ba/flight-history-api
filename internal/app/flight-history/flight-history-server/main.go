package flighthistoryserver

import "github.com/sirupsen/logrus"

type FlightHistoryServer struct {
	Log    *logrus.Logger
	Config *Config
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

	return &FlightHistoryServer{
		Config: config,
		Log:    log,
	}
}
