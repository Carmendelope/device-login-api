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

package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-device-manager-go"
)

const (
	emptyOrganizationId = "organization_id cannot be empty"
	emptyDeviceId = "device_id cannot be empty"
	emptyDeviceGroupId = "device_id cannot be empty"
	emptyDeviceApiKey = "device_api_key cannot be empty"
	emptyDeviceGroupApiKey = "device_group_api_key cannot be empty"
)

func ValidLoginRequest(loginRequest *grpc_authx_go.DeviceLoginRequest) derrors.Error {

	// OrganizationId
	if loginRequest.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	// DeviceApiKey
	if loginRequest.DeviceApiKey == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceApiKey)
	}

	return nil

}

func ValidRegisterRequest(registerRequest *grpc_device_manager_go.RegisterDeviceRequest) derrors.Error  {
	// OrganizationId
	if registerRequest.OrganizationId == "" {
		return derrors.NewInvalidArgumentError(emptyOrganizationId)
	}
	if registerRequest.DeviceGroupId == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceGroupId)
	}
	if registerRequest.DeviceId == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceId)
	}
	if registerRequest.DeviceGroupApiKey == "" {
		return derrors.NewInvalidArgumentError(emptyDeviceGroupApiKey)
	}
	return nil

}