package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/services/server/srvcalendars"
	"github.com/athomecomar/athome/backend/services/server/srvregister"
	"github.com/athomecomar/athome/backend/services/server/srvviewer"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/athomecomar/athome/pb/pbservices"
	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Services
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)

	pbservices.RegisterViewerServer(s, &srvviewer.Server{})
	pbservices.RegisterRegisterServer(s, &srvregister.Server{})
	pbservices.RegisterCalendarsServer(s, &srvcalendars.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
