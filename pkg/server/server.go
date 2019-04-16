package server

import (
	"BangAuth/config"
	pb "BangAuth/pkg/protobuf"
	"context"
)

type Server struct{}

func (s *Server) CheckCookie(ctx context.Context, in *pb.CookieRequest) (*pb.UserInfoResponse, error) {
	config.Logger.Info("CheckCookie started")
	defer config.Logger.Info("CheckCookie finished")

	user, err := InfoFromCookie(in.JwtToken)
	if err != nil {
		return nil, err
	}

	return &pb.UserInfoResponse{
			Valid: true,
			User:  &user,
		},
		nil
}
