package main

import (
	"log"
	"net"

	"github.com/athomecomar/athome/backend/auth/authconf"
	"github.com/athomecomar/athome/backend/auth/server"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbconf"
	"github.com/go-redis/redis/v8"

	"google.golang.org/grpc"
)

func main() {
	svc := pbconf.Auth
	port := svc.GetPort()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	log.Println("listening on port " + port)
	pbauth.RegisterAuthServer(s, &server.Server{Redis: redisCli()})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func redisCli() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     authconf.GetDATABASE_HOST() + authconf.GetDATABASE_PORT(),
		Password: authconf.GetDATABASE_PASSWORD(),
		DB:       authconf.GetDATABASE_NUMBER(),
	})
}
