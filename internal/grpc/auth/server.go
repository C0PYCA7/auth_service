package auth

import (
	"context"
	"github.com/C0PYCA7/protosAuth/gen/go/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	auth.UnimplementedAuthServer
	auth Auth
}

type Auth interface {
	Register(ctx context.Context, name, surname, login, password, mail string) (int64, error)
	Login(ctx context.Context, login, password string) (int64, bool, error)
	Generate(ctx context.Context, username string) (string, error)
	VerifyQr(ctx context.Context, username, code string) (int64, string, error)
}

func Register(gRPC *grpc.Server, authh Auth) {
	auth.RegisterAuthServer(gRPC, &Server{auth: authh})
}

func (s *Server) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if req.GetMail() == "" {
		return nil, status.Error(codes.InvalidArgument, "mail is empty")
	}
	if req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "password is empty")
	}

	uid, err := s.auth.Register(ctx, req.GetName(), req.GetSurname(), req.GetLogin(), req.GetPassword(), req.GetMail())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &auth.RegisterResponse{Uid: uid}, nil
}

func (s *Server) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	if req.GetLogin() == "" || req.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "login or password is empty")
	}

	uid, enable, err := s.auth.Login(ctx, req.GetLogin(), req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &auth.LoginResponse{Uid: uid, Enable: enable}, nil
}

func (s *Server) GenerateQR(ctx context.Context, req *auth.GenerateRequest) (*auth.GenerateResponse, error) {
	if req.GetLogin() == "" {
		return nil, status.Error(codes.InvalidArgument, "login or password is empty")
	}
	code, err := s.auth.Generate(ctx, req.GetLogin())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &auth.GenerateResponse{QrCode: code}, nil
}

func (s *Server) VerifyQR(ctx context.Context, req *auth.VerifyRequest) (*auth.VerifyResponse, error) {
	if req.GetCode() == "" {
		return nil, status.Error(codes.InvalidArgument, "code is empty")
	}

	uid, jwt, err := s.auth.VerifyQr(ctx, req.GetCode(), req.GetCode())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &auth.VerifyResponse{Uid: uid, Jwt: jwt}, nil
}
