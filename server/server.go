package server

import (
	"context"
	"encoding/json"
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	pb "github.com/klever-io/gcp-logging/helloworld"
	"github.com/klever-io/kloud-sdk/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	// Get the logger from the context
	logger := ctxzap.Extract(ctx)

	logger.Info("SayHello INFO message", zap.String("name", in.Name))

	if in.Name == "error" {
		logger.Error("SayHello ERROR message")
	}

	if in.Name == "warn" {
		logger.Warn("SayHello WARN message")
	}

	reply := &pb.HelloReply{Message: "Hello " + in.GetName()}

	// For DEBUG level print both the request and the response objects
	logger.Debug("SayHello DEBUG message",
		zap.Any("request", json.RawMessage(logging.PrototoJson(in))),
		zap.Any("reply", json.RawMessage(logging.PrototoJson(reply))),
	)
	return reply, nil
}

func Serve() {
	logger := logging.NewLogger()

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		logger.Fatal("Failed to listen: %v", zap.String("error", err.Error()))
	}

	serverOpts := make([]grpc.ServerOption, 0)
	serverOpts = append(serverOpts,
		grpc.ChainUnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
		grpc.ChainStreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(logger),
		)),
	)

	s := grpc.NewServer(serverOpts...)
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)

	logger.Info("Server listening on port 50051")
	if err := s.Serve(lis); err != nil {
		logger.Fatal("Failed to serve: %v", zap.String("error", err.Error()))
	}
}
