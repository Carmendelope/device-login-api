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

package commands

import (
	"github.com/nalej/device-login-api/internal/pkg/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var config = server.Config{}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Launch the server API",
	Long:  `Launch the server API`,
	Run: func(cmd *cobra.Command, args []string) {
		SetupLogging()
		log.Info().Msg("Launching API!")
		server := server.NewService(config)
		server.Run()
	},
}

func init() {
	runCmd.Flags().IntVar(&config.Port, "port", 6030, "Port to launch the Public gRPC API")
	runCmd.Flags().IntVar(&config.HTTPPort, "httpPort", 6031, "Port to launch the Public HTTP API")
	runCmd.PersistentFlags().StringVar(&config.AuthxAddress, "authxAddress", "localhost:8810",
		"Authx address (host:port)")
	runCmd.PersistentFlags().StringVar(&config.DeviceManagerAddress, "deviceManagerAddress", "localhost:6010",
		"Device Manager Address (host:port)")
	rootCmd.AddCommand(runCmd)
}
