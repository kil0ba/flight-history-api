package serverConfig

import (
	"github.com/kil0ba/flight-history-api/internal/app/store"
	jwt_utils "github.com/kil0ba/flight-history-api/internal/app/utils/jwt-utils"
	"github.com/sirupsen/logrus"
)

type FlightHistoryServer struct {
	Log        *logrus.Logger
	Config     *Config
	Store      *store.Store
	JwtManager *jwt_utils.JWTManager
}
