
/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

 package register

import (
	"context"
	"github.com/nalej/grpc-device-manager-go"
)

// Manager structure with the required clients for roles operations.
type Manager struct {
	devManager grpc_device_manager_go.DevicesClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(accessClient grpc_device_manager_go.DevicesClient) Manager {
	return Manager{devManager:accessClient}
}

func (m * Manager) RegisterDevice (registerRequest *grpc_device_manager_go.RegisterDeviceRequest) (*grpc_device_manager_go.RegisterResponse, error) {
	return m.devManager.RegisterDevice(context.Background(), registerRequest)
}

