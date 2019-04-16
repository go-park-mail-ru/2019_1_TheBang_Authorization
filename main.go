package main

import (
	"BangAuth/config"
	pb "BangAuth/pkg/protobuf"
	"BangAuth/pkg/server"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {
	defer config.Logger.Sync()
	config.Logger.Info(fmt.Sprintf("PORT: %v", config.PORT))

	lis, err := net.Listen("tcp", ":"+config.PORT)
	if err != nil {
		config.Logger.Fatalw("Listen port",
			"error:", err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterCookieCheckerServer(s, &server.Server{})

	if err = s.Serve(lis); err != nil {
		config.Logger.Fatalw("serve port",
			"error:", err.Error())
	}
}
