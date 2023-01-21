package registrator

import (
	"context"
	"time"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/logger"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-registrator-api/utils/time"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	moduleNameHeader = "x-module-name"
)

func logInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time_utils.TimeNow()

	h, err := handler(ctx, req)

	if err != nil {
		st := status.Convert(err)
		logger.Errorf(ctx, "Request - Method:%s Duration:%s Error:%v Details: %v",
			info.FullMethod,
			time.Since(start),
			err,
			st.Details(),
		)
	} else {
		logger.Debugf(ctx, "Request - Method:%s Duration:%s",
			info.FullMethod,
			time.Since(start),
		)
	}

	return h, err
}

func userStateInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	user := users_utils.UserFromContext(ctx)
	if user.State == model.UserStateBanned {
		st := status.New(codes.PermissionDenied, "permission denied")
		errorInfo := errdetails.ErrorInfo{
			Reason: "banned",
		}

		st, err := st.WithDetails(&errorInfo)
		if err != nil {
			panic(err)
		}

		return nil, st.Err()
	}

	if info.FullMethod == "/users.RegistratorService/CreateUser" {
		return handler(ctx, req)
	}

	switch info.FullMethod {
	case "/users.RegistratorService/AddGames",
		"/users.RegistratorService/GetPlaceByNameAndAddress":
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			headers := md.Get(moduleNameHeader)
			if len(headers) > 0 && headers[0] == "fetcher" {
				return handler(ctx, req)
			}
		}
	}

	if info.FullMethod == "/users.RegistratorService/GetPlaceByNameAndAddress" {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		headers := md.Get("x-inline-query")
		if len(headers) > 0 && headers[0] == "true" {
			return handler(ctx, req)
		}
	}

	if user.ID == 0 {
		st := status.New(codes.Unauthenticated, "user not found")
		errorInfo := errdetails.ErrorInfo{
			Reason: "user not found",
		}

		st, err := st.WithDetails(&errorInfo)
		if err != nil {
			panic(err)
		}

		return nil, st.Err()
	}

	h, err := handler(ctx, req)

	return h, err
}
