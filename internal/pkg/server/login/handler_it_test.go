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
	"github.com/nalej/grpc-device-login-api-go"
	"github.com/nalej/grpc-utils/pkg/test"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"math/rand"
	"github.com/google/uuid"
	"os"
)

var _ = ginkgo.Describe("Applications", func() {

	var runIntegration= os.Getenv("RUN_INTEGRATION_TEST")

	if runIntegration != "true" {
		log.Warn().Msg("Integration tests are skipped")
		return
	}

	var (
		authxAddress= os.Getenv("IT_AUTHX_ADDRESS")
	)

	if authxAddress == "" {
		ginkgo.Fail("missing environment variables")
	}

	var authxClient grpc_authx_go.AuthxClient
	var server *grpc.Server
	var listener *bufconn.Listener
	var client grpc_device_login_api_go.LoginClient

	ginkgo.BeforeSuite(func() {

		listener = test.GetDefaultListener()
		server = grpc.NewServer()

		authxConn, err := grpc.Dial(authxAddress, grpc.WithInsecure())
		gomega.Expect(err).To(gomega.Succeed())

		authxClient = grpc_authx_go.NewAuthxClient(authxConn)

		manager := NewManager(authxClient)
		handler := NewHandler(manager)

		grpc_device_login_api_go.RegisterLoginServer(server, handler)
		test.LaunchServer(server, listener)

		conn, err := test.GetConn(*listener)
		gomega.Expect(err).Should(gomega.Succeed())
		client = grpc_device_login_api_go.NewLoginClient(conn)
		rand.Seed(ginkgo.GinkgoRandomSeed())

	})

	ginkgo.AfterSuite(func() {
		server.Stop()
		listener.Close()
	})

	ginkgo.Context("login", func() {

		ginkgo.It("should be able to login", func() {

			// add a group
			groupToAdd :=grpc_authx_go.AddDeviceGroupCredentialsRequest{
				OrganizationId: uuid.New().String(),
				DeviceGroupId: uuid.New().String(),
				Enabled: true,
			}
			groupAdded, err := authxClient.AddDeviceGroupCredentials(context.Background(), &groupToAdd)
			gomega.Expect(err).To(gomega.Succeed())
			gomega.Expect(groupAdded).NotTo(gomega.BeNil())

			// add a device
			toAdd := grpc_authx_go.AddDeviceCredentialsRequest{
				OrganizationId: groupToAdd.OrganizationId,
				DeviceGroupId: groupToAdd.DeviceGroupId,
				DeviceId:  uuid.New().String(),
			}
			added, err := authxClient.AddDeviceCredentials(context.Background(), &toAdd)
			gomega.Expect(err).To(gomega.Succeed())
			gomega.Expect(added).NotTo(gomega.BeNil())

			// send a login
			loginRequest := grpc_authx_go.DeviceLoginRequest {
				OrganizationId: toAdd.OrganizationId,
				DeviceApiKey: added.DeviceApiKey,
			}
			loginResponse, err := client.DeviceLogin(context.Background(), &loginRequest)
			gomega.Expect(err).To(gomega.Succeed())
			gomega.Expect(loginResponse).NotTo(gomega.BeNil())

		})
		ginkgo.It("should not be able to login, wrong api_key", func() {

			// add a group
			groupToAdd :=grpc_authx_go.AddDeviceGroupCredentialsRequest{
				OrganizationId: uuid.New().String(),
				DeviceGroupId: uuid.New().String(),
				Enabled: true,
			}
			groupAdded, err := authxClient.AddDeviceGroupCredentials(context.Background(), &groupToAdd)
			gomega.Expect(err).To(gomega.Succeed())
			gomega.Expect(groupAdded).NotTo(gomega.BeNil())

			// add a device
			toAdd := grpc_authx_go.AddDeviceCredentialsRequest{
				OrganizationId: groupToAdd.OrganizationId,
				DeviceGroupId: groupToAdd.DeviceGroupId,
				DeviceId:  uuid.New().String(),
			}
			added, err := authxClient.AddDeviceCredentials(context.Background(), &toAdd)
			gomega.Expect(err).To(gomega.Succeed())
			gomega.Expect(added).NotTo(gomega.BeNil())

			// send a login
			loginRequest := grpc_authx_go.DeviceLoginRequest {
				OrganizationId: toAdd.OrganizationId,
				DeviceApiKey: uuid.New().String(),
			}
			_, err = client.DeviceLogin(context.Background(), &loginRequest)
			gomega.Expect(err).NotTo(gomega.Succeed())

		})
	})
})
