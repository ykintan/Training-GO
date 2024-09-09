package middleware

import (
	"context"
	"encoding/base64"
	"strings"
	"training-go/go-session10-crud-grpc-gateway/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {
		//check for public methods that don't need auth
		publicmethods := []string{
			"/proto.UserService/GetUsers",
			"/proto.UserService/GetUserByID",
		}
		for _, method := range publicmethods {
			if info.FullMethod == method {
				return handler(ctx, req)
			}
		}

		//extract metadata from context
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
		}

		authHeader, ok := md["authorization"]
		if !ok || len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "authorization header is missing")
		}

		//check for basic auth scheme
		if !strings.HasPrefix(authHeader[0], "Basic ") {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization scheme ")
		}
		//Decode base64 credentials
		decoded, err := base64.StdEncoding.DecodeString((authHeader[0])[6:])
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token ")
		}

		//split username and password
		creds := strings.SplitN(string(decoded), ":", 2)

		if len(creds) != 2 {
			return nil, status.Errorf(codes.Unauthenticated, "invalid authorization token ")
		}

		username, password := creds[0], creds[1]

		//validate the credential
		if username != config.AuthBasicUsername || password != config.AuthBasicPassword {
			return nil, status.Errorf(codes.Unauthenticated, "invalid username or password")
		}

		return handler(ctx, req)

	}
}
