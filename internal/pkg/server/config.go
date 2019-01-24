package server

import (
	"github.com/nalej/derrors"
	"github.com/nalej/device-login-api/version"
	"github.com/rs/zerolog/log"
)

type Config struct {
	// Port where the gRPC API service will listen requests.
	Port int
	// HTTPPort where the HTTP gRPC gateway will be listening.
	HTTPPort int
	// AuthxAddress with the host:port to connect to the Authx manager.
	AuthxAddress string
	// DeviceManagerAddress with the host:port to connect to Device Manager
	DeviceManagerAddress string
}

func (conf * Config) Validate() derrors.Error {

	if conf.Port <= 0 || conf.HTTPPort <= 0 {
		return derrors.NewInvalidArgumentError("ports must be valid")
	}
	if conf.AuthxAddress == "" {
		return derrors.NewInvalidArgumentError("authx must be set")
	}
	if conf.DeviceManagerAddress == "" {
		return derrors.NewInvalidArgumentError("deviceManager must be set")
	}

	return nil
}

func (conf * Config) Print()  {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.Port).Msg("gRPC port")
	log.Info().Int("port", conf.HTTPPort).Msg("HTTP port")
	log.Info().Str("URL", conf.AuthxAddress).Msg("Authx")
	log.Info().Str("URL", conf.DeviceManagerAddress).Msg("Device Manager")
}