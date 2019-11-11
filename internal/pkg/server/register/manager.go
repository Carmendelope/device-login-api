
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

