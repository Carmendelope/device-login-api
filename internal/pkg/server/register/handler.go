
/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package register

import (
	"context"
	"github.com/nalej/device-login-api/internal/pkg/entities"
	"github.com/nalej/grpc-device-manager-go"
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

func (h * Handler) RegisterDevice(context context.Context, registerRequest *grpc_device_manager_go.RegisterDeviceRequest) (*grpc_device_manager_go.RegisterResponse, error) {
	vErr := entities.ValidRegisterRequest(registerRequest)

	if vErr != nil {
		return nil, conversions.ToDerror(vErr)
	}

	response, err := h.Manager.RegisterDevice(registerRequest)
	if err != nil {
		log.Error().Str("trace", conversions.ToDerror(err).DebugReport()).Msg("device register error")
		return nil, err
	}
	return response, nil
}