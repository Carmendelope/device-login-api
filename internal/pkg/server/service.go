package server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nalej/derrors"
	"github.com/nalej/device-login-api/internal/pkg/server/login"
	"github.com/nalej/device-login-api/internal/pkg/server/register"
	"github.com/nalej/grpc-authx-go"
	"github.com/nalej/grpc-device-manager-go"
	"github.com/nalej/grpc-device-login-api-go"
	"github.com/nalej/grpc-utils/pkg/tools"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strings"
)

type Service struct {
	Configuration Config
	Server * tools.GenericGRPCServer
}

// NewService creates a new system model service.
func NewService(conf Config) *Service {
	return &Service{
		conf,
		tools.NewGenericGRPCServer(uint32(conf.Port)),
	}
}

type Clients struct {
	devManager grpc_device_manager_go.DevicesClient
	authxManager grpc_authx_go.AuthxClient
}

func (s * Service) GetClients() (* Clients, derrors.Error) {

	devConn, err := grpc.Dial(s.Configuration.DeviceManagerAddress, grpc.WithInsecure())
	if err != nil{
		return nil, derrors.AsError(err, "cannot create connection with the device manager")
	}
	authxConn, err := grpc.Dial(s.Configuration.AuthxAddress, grpc.WithInsecure())
	if err != nil{
		return nil, derrors.AsError(err, "cannot create connection with the authx manager")
	}

	dClient := grpc_device_manager_go.NewDevicesClient(devConn)
	aClient := grpc_authx_go.NewAuthxClient(authxConn)

	return &Clients{dClient, aClient}, nil
}

func (s * Service) Run () error {
	vErr := s.Configuration.Validate()
	if vErr != nil {
		log.Fatal().Str("err", vErr.DebugReport()).Msg("invalid configuration")
	}

	s.Configuration.Print()

	go s.LaunchGRPC()
	return s.LaunchHTTP()

}

func (s * Service) LaunchGRPC() error {
	clients, cErr := s.GetClients()
	if cErr != nil {
		log.Fatal().Str("err", cErr.DebugReport()).Msg("cannot generate clients")
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Configuration.Port))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to listen")
	}

	// Create handlers
	loginManager := login.NewManager(clients.authxManager)
	loginHandler := login.NewHandler(loginManager)

	registerManager := register.NewManager(clients.devManager)
	registerHandler := register.NewHandler(registerManager)

	grpcServer := grpc.NewServer()
	grpc_device_login_api_go.RegisterLoginServer(grpcServer, loginHandler)
	grpc_device_login_api_go.RegisterRegisterServer(grpcServer, registerHandler)

	reflection.Register(grpcServer)
	log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}

	return nil
}

func (s *Service) allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
}

func (s * Service) LaunchHTTP() error {

	addr := fmt.Sprintf(":%d", s.Configuration.HTTPPort)
	clientAddr := fmt.Sprintf(":%d", s.Configuration.Port)
	opts := []grpc.DialOption{grpc.WithInsecure()}
	mux := runtime.NewServeMux()

	if err := grpc_device_login_api_go.RegisterLoginHandlerFromEndpoint(context.Background(), mux, clientAddr, opts); err != nil {
		log.Fatal().Err(err).Msg("failed to start device manager handler")
	}
	if err := grpc_device_login_api_go.RegisterRegisterHandlerFromEndpoint(context.Background(), mux, clientAddr, opts) ; err != nil{
		log.Fatal().Err(err).Msg("failed to start authx manager handler")
	}

	server := &http.Server{
		Addr:    addr,
		Handler: s.allowCORS(mux),
	}

	log.Info().Str("address", addr).Msg("HTTP Listening")
	return server.ListenAndServe()

}