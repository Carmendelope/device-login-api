
/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package register

import (
	"context"
	"github.com/nalej/grpc-device-manager-go"
)

type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler{
	return &Handler{manager}
}

func (h * Handler) RegisterDevice(context.Context, *grpc_device_manager_go.RegisterDeviceRequest) (*grpc_device_manager_go.RegisterResponse, error) {
	return nil, nil
}