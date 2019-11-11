/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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