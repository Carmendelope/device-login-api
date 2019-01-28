package register

import (
	"context"
	"github.com/google/uuid"
	"github.com/nalej/device-login-api/internal/pkg/utils"
	"github.com/nalej/grpc-device-login-api-go"
	"github.com/nalej/grpc-device-manager-go"
	"github.com/nalej/grpc-organization-go"
	"github.com/nalej/grpc-utils/pkg/test"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"math/rand"
	"os"
)


var _ = ginkgo.Describe("Applications", func() {

	var runIntegration = os.Getenv("RUN_INTEGRATION_TEST")

	if runIntegration != "true" {
		log.Warn().Msg("Integration tests are skipped")
		return
	}

	var (
		devManagerAddress = os.Getenv("IT_DEV_MNG_ADDRESS")
		systemModelAddress= os.Getenv("IT_SM_ADDRESS")
	)

	if systemModelAddress == "" || devManagerAddress == "" {
		ginkgo.Fail("missing environment variables")
	}

	var devManagerClient grpc_device_manager_go.DevicesClient
	var orgClient grpc_organization_go.OrganizationsClient


	var server *grpc.Server
	var listener *bufconn.Listener
	var client grpc_device_login_api_go.RegisterClient
	var smConn *grpc.ClientConn

	ginkgo.BeforeSuite(func() {

		listener = test.GetDefaultListener()
		server = grpc.NewServer()

		smConn = utils.GetConnection(systemModelAddress)
		devConn, err := grpc.Dial(devManagerAddress, grpc.WithInsecure())
		gomega.Expect(err).To(gomega.Succeed())

		devManagerClient = grpc_device_manager_go.NewDevicesClient(devConn)
		orgClient = grpc_organization_go.NewOrganizationsClient(smConn)

		manager := NewManager(devManagerClient)
		handler := NewHandler(manager)

		grpc_device_login_api_go.RegisterRegisterServer(server, handler)
		test.LaunchServer(server, listener)

		conn, err := test.GetConn(*listener)
		gomega.Expect(err).Should(gomega.Succeed())
		client = grpc_device_login_api_go.NewRegisterClient(conn)
		rand.Seed(ginkgo.GinkgoRandomSeed())



	})

	ginkgo.AfterSuite(func() {
		server.Stop()
		listener.Close()
	})

	ginkgo.Context("register", func() {
		var targetOrganization * grpc_organization_go.Organization
		testHelper := utils.NewDeviceTestHepler()

		ginkgo.BeforeEach(func() {
			targetOrganization = testHelper.AddOrganization("register test", orgClient)
		})

		ginkgo.It("Should be able to register a device", func() {

			// add device group
			group, err := devManagerClient.AddDeviceGroup(context.Background(), testHelper.CreateDeviceGroup(targetOrganization.OrganizationId))
			gomega.Expect(group).NotTo(gomega.BeNil())
			gomega.Expect(err).To(gomega.Succeed())

			toRegister := testHelper.CreateRegisterDeviceRequest(group)
			response, err := client.RegisterDevice(context.Background(), toRegister)
			gomega.Expect(response).NotTo(gomega.BeNil())
			gomega.Expect(err).To(gomega.Succeed())

		})
		ginkgo.It("Should not be able to register a device on a disable group", func() {
			// add device group
			request := testHelper.CreateDeviceGroup(targetOrganization.OrganizationId)
			request.Enabled = false

			group, err := devManagerClient.AddDeviceGroup(context.Background(), request)
			gomega.Expect(group).NotTo(gomega.BeNil())
			gomega.Expect(err).To(gomega.Succeed())

			toRegister := testHelper.CreateRegisterDeviceRequest(group)
			_, err = client.RegisterDevice(context.Background(), toRegister)
			gomega.Expect(err).NotTo(gomega.Succeed())

		})
		ginkgo.It("Should not be able to register a device on a no existing group", func() {

			request := testHelper.CreateDeviceGroup(targetOrganization.OrganizationId)

			group, err := devManagerClient.AddDeviceGroup(context.Background(), request)
			gomega.Expect(group).NotTo(gomega.BeNil())
			gomega.Expect(err).To(gomega.Succeed())

			toRegister := testHelper.CreateRegisterDeviceRequest(group)
			toRegister.DeviceGroupId = uuid.New().String()
			_, err = client.RegisterDevice(context.Background(), toRegister)
			gomega.Expect(err).NotTo(gomega.Succeed())

		})

	})

})