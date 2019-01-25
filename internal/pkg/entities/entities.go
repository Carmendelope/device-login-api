package entities

import (
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-authx-go"
)

const (
	emptyOrganizationId = "organization_id cannot be empty"
	emptyDeviceApiKey = "device_api_key cannot be empty"
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