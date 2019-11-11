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

package login

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/device-login-api/internal/pkg/entities"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler{
	return &Handler{manager}
}


func (h * Handler) DeviceLogin(ctx context.Context, loginRequest *grpc_authx_go.DeviceLoginRequest) (*grpc_authx_go.LoginResponse, error) {
	vErr := entities.ValidLoginRequest(loginRequest)

	if vErr != nil {
		return nil, conversions.ToDerror(vErr)
	}

	response, err := h.Manager.DeviceLogin(loginRequest)
	if err != nil {
		log.Error().Str("trace", conversions.ToDerror(err).DebugReport()).Str("OrganizationId", loginRequest.OrganizationId).Str("DeviceApiKey", loginRequest.DeviceApiKey).Msg("device login error")
		return nil, conversions.ToGRPCError(derrors.NewGenericError("Invalid device credentials"))
	}
	return response, nil
}