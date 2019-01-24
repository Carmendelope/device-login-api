/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package login

import "github.com/nalej/grpc-authx-go"

// Manager structure with the required clients for roles operations.
type Manager struct {
	accessClient grpc_authx_go.AuthxClient
}

// NewManager creates a Manager using a set of clients.
func NewManager(accessClient grpc_authx_go.AuthxClient) Manager {
	return Manager{accessClient:accessClient}
}
