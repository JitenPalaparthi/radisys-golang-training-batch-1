package main

import (
	"context"
	"customers/models"
	pb "customers/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCustomerServiceServer
}

func (s *Server) Create(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
	customer := new(models.Customer)
	customer.Email = in.Email
	customer.Address = in.Address
	customer.Name = in.Name

	out := new(pb.CustomerResponse)
	out.Code = 201
	out.Message = "Customer successfully created"
	return out, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
