package main

import (
	"context"
	"log"
	"net"

	shape "github.com/ryanbeau/grpc-go-server/api/shape/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct {
	shape.UnimplementedShapeServer
}

func (s *server) TrianglePerimeter(ctx context.Context, req *shape.TrianglePerimeterRequest) (*shape.TrianglePerimeterResponse, error) {
	response := &shape.TrianglePerimeterResponse{}

	log.Printf("a:%d b:%d c:%d", req.SideA, req.SideB, req.SideC)

	// NOTE: basic triangle test for grpc server example - does not account for int overflow
	if req.SideA+req.SideB <= req.SideC {
		return response, status.Error(codes.InvalidArgument, "Invalid input: make sure a+b>c")
	}
	if req.SideA+req.SideC <= req.SideB {
		return response, status.Error(codes.InvalidArgument, "Invalid input: make sure a+c>b")
	}
	if req.SideB+req.SideC <= req.SideA {
		return response, status.Error(codes.InvalidArgument, "Invalid input: make sure b+c>a")
	}

	response.Value = req.SideA + req.SideB + req.SideC

	return response, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	shape.RegisterShapeServer(s, &server{})

	log.Printf("listening on port %s", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
