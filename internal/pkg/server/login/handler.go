/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package login

import (
	"context"
	"github.com/nalej/grpc-authx-go"
)

type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler{
	return &Handler{manager}
}


func (h * Handler) DeviceLogin(ctx context.Context, loginRequest *grpc_authx_go.DeviceLoginRequest) (*grpc_authx_go.LoginResponse, error) {
	return nil, nil
}