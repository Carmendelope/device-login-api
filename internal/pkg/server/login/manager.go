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
	"github.com/nalej/grpc-authx-go"
)

// Manager structure with the required clients for roles operations.
type Manager struct {
	accessClient grpc_authx_go.AuthxClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(accessClient grpc_authx_go.AuthxClient) Manager {
	return Manager{accessClient: accessClient}
}

func (m *Manager) DeviceLogin(loginRequest *grpc_authx_go.DeviceLoginRequest) (*grpc_authx_go.LoginResponse, error) {
	return m.accessClient.DeviceLogin(context.Background(), loginRequest)
}
