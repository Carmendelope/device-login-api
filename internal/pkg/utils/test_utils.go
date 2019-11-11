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

package utils

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/nalej/grpc-device-manager-go"
	"github.com/nalej/grpc-organization-go"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"google.golang.org/grpc"
	"math/rand"
	"context"
)

func GetConnection(address string) (* grpc.ClientConn) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	gomega.Expect(err).To(gomega.Succeed())
	return conn
}

type DeviceTestHelper struct {

}

func NewDeviceTestHepler() *DeviceTestHelper {
	return &DeviceTestHelper{}
}

func (d * DeviceTestHelper) AddOrganization (name string, provider grpc_organization_go.OrganizationsClient) * grpc_organization_go.Organization {
	targetOrganization, err := provider.AddOrganization(context.Background(), d.CreateOrganization(name))
	gomega.Expect(err).To(gomega.Succeed())
	gomega.Expect(targetOrganization).NotTo(gomega.BeNil())
	return targetOrganization
}

func (d * DeviceTestHelper) CreateOrganization(name string) *grpc_organization_go.AddOrganizationRequest {
	toAdd := &grpc_organization_go.AddOrganizationRequest{
		Name:                 fmt.Sprintf("%s-%d-%d", name, ginkgo.GinkgoRandomSeed(), rand.Int()),
	}

	return toAdd
}

func (d * DeviceTestHelper) CreateDeviceGroup(organizationID string) *grpc_device_manager_go.AddDeviceGroupRequest  {
	return &grpc_device_manager_go.AddDeviceGroupRequest{
		OrganizationId: organizationID,
		Name: "test group",
		Enabled: true,
		DefaultDeviceConnectivity: true,
	}
}

func (d * DeviceTestHelper) CreateRegisterDeviceRequest (group *grpc_device_manager_go.DeviceGroup) *grpc_device_manager_go.RegisterDeviceRequest{

	labels := make (map[string]string, 0)
	for i:= 1; i<= rand.Intn(4)+1; i++ {
		labels[fmt.Sprintf("label_%d", i)] = fmt.Sprintf("value_%d", i)
	}

	registerDevice := &grpc_device_manager_go.RegisterDeviceRequest {
		OrganizationId: group.OrganizationId,
		DeviceGroupId: group.DeviceGroupId,
		DeviceGroupApiKey:group.DeviceGroupApiKey,
		DeviceId: uuid.New().String(),
		Labels:labels,
	}
	return registerDevice

}