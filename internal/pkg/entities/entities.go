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