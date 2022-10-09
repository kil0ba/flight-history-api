package flighthistoryserver

import "github.com/sirupsen/logrus"

type FlightHistoryServer struct {
	Log    *logrus.Logger
	Config *Config
}
