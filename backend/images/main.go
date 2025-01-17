package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/images/imageconf"
	"github.com/athomecomar/athome/backend/images/server"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbimages"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Images
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("listening on port " + port)

	s := grpc.NewServer()
	store := imageconf.GetSTORE()
	pbimages.RegisterImagesServer(s, &server.Server{Store: store})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
