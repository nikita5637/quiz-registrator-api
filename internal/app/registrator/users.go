package registrator

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/i18n"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	userAlreadyExistsLexeme = i18n.Lexeme{
		Key:      "user_already_exists",
		FallBack: "User already exists",
	}
	errEmailValidateLexeme = i18n.Lexeme{
		Key:      "err_email_validation",
		FallBack: "Invalid email format",
	}
	errNameAlphabetValidateLexeme = i18n.Lexeme{
		Key:      "err_name_alphabet_validation",
		FallBack: "Only Russian character set are allowed",
	}
	errNameLengthValidateLexeme = i18n.Lexeme{
		Key:      "err_name_length_validation",
		FallBack: "Name length must be between 1 and 100 characters",
	}
	errNameRequiredhValidateLexeme = i18n.Lexeme{
		Key:      "err_name_required_validation",
		FallBack: "User name is required",
	}
	errPhoneValidateLexeme = i18n.Lexeme{
		Key:      "err_phone_validation",
		FallBack: "Invalid phone format",
	}
	errStateValidateLexeme = i18n.Lexeme{
		Key:      "err_state_validation",
		FallBack: "Invalid user state",
	}
)

// CreateUser ...
func (r *Registrator) CreateUser(ctx context.Context, req *registrator.CreateUserRequest) (*registrator.CreateUserResponse, error) {
	newUser := model.User{
		Name:       req.GetName(),
		TelegramID: req.GetTelegramId(),
		State:      model.UserState(req.GetState()),
	}

	id, err := r.usersFacade.CreateUser(ctx, newUser)
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrUserAlreadyExists) {
			reason := fmt.Sprintf("user with telegram id %d already exists", req.GetTelegramId())
			st = getStatus(ctx, codes.AlreadyExists, err, reason, userAlreadyExistsLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.CreateUserResponse{
		Id: id,
	}, nil
}

// GetUser ...
func (r *Registrator) GetUser(ctx context.Context, req *registrator.GetUserRequest) (*registrator.GetUserResponse, error) {
	user, err := r.usersFacade.GetUser(ctx)
	if err != nil {
		st := status.New(codes.Internal, err.Error())
		return nil, st.Err()
	}

	return &registrator.GetUserResponse{
		User: &registrator.User{
			Id:         user.ID,
			Name:       user.Name,
			TelegramId: user.TelegramID,
			Email:      user.Email,
			Phone:      user.Phone,
			State:      registrator.UserState(user.State),
		},
	}, nil
}

// GetUserByID ...
func (r *Registrator) GetUserByID(ctx context.Context, req *registrator.GetUserByIDRequest) (*registrator.GetUserByIDResponse, error) {
	user, err := r.usersFacade.GetUserByID(ctx, req.GetId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrUserNotFound) {
			reason := fmt.Sprintf("user with id %d not found", req.GetId())
			st = getStatus(ctx, codes.NotFound, err, reason, i18n.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.GetUserByIDResponse{
		User: &registrator.User{
			Id:         user.ID,
			Name:       user.Name,
			TelegramId: user.TelegramID,
			Email:      user.Email,
			Phone:      user.Phone,
			State:      registrator.UserState(user.State),
		},
	}, nil
}

// GetUserByTelegramID ...
func (r *Registrator) GetUserByTelegramID(ctx context.Context, req *registrator.GetUserByTelegramIDRequest) (*registrator.GetUserByTelegramIDResponse, error) {
	user, err := r.usersFacade.GetUserByTelegramID(ctx, req.GetTelegramId())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrUserNotFound) {
			reason := fmt.Sprintf("user with talegram id %d not found", req.GetTelegramId())
			st = getStatus(ctx, codes.NotFound, err, reason, i18n.UserNotFoundLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.GetUserByTelegramIDResponse{
		User: &registrator.User{
			Id:         user.ID,
			Name:       user.Name,
			TelegramId: user.TelegramID,
			Email:      user.Email,
			Phone:      user.Phone,
			State:      registrator.UserState(user.State),
		},
	}, nil
}

// UpdateUserEmail ...
func (r *Registrator) UpdateUserEmail(ctx context.Context, req *registrator.UpdateUserEmailRequest) (*registrator.UpdateUserEmailResponse, error) {
	err := r.usersFacade.UpdateUserEmail(ctx, req.GetUserId(), req.GetEmail())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrUserNotFound) {
			st = getUserNotFoundStatus(ctx, err, req.GetUserId())
		} else if errors.Is(err, model.ErrUserEmailValidate) {
			reason := fmt.Sprintf("invalid user email: %s", req.GetEmail())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, errEmailValidateLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.UpdateUserEmailResponse{}, nil
}

// UpdateUserName ...
func (r *Registrator) UpdateUserName(ctx context.Context, req *registrator.UpdateUserNameRequest) (*registrator.UpdateUserNameResponse, error) {
	err := r.usersFacade.UpdateUserName(ctx, req.GetUserId(), req.GetName())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrUserNotFound) {
			st = getUserNotFoundStatus(ctx, err, req.GetUserId())
		} else if errors.Is(err, model.ErrUserNameValidateAlphabet) ||
			errors.Is(err, model.ErrUserNameValidateLength) ||
			errors.Is(err, model.ErrUserNameValidateRequired) {
			reason := fmt.Sprintf("invalid user name: %s", req.GetName())

			lexeme := i18n.Lexeme{}
			if errors.Is(err, model.ErrUserNameValidateAlphabet) {
				lexeme = errNameAlphabetValidateLexeme
			} else if errors.Is(err, model.ErrUserNameValidateLength) {
				lexeme = errNameLengthValidateLexeme
			} else if errors.Is(err, model.ErrUserNameValidateRequired) {
				lexeme = errNameRequiredhValidateLexeme
			}

			st = getStatus(ctx, codes.InvalidArgument, err, reason, lexeme)
		}

		return nil, st.Err()
	}

	return &registrator.UpdateUserNameResponse{}, nil
}

// UpdateUserPhone ...
func (r *Registrator) UpdateUserPhone(ctx context.Context, req *registrator.UpdateUserPhoneRequest) (*registrator.UpdateUserPhoneResponse, error) {
	err := r.usersFacade.UpdateUserPhone(ctx, req.GetUserId(), req.GetPhone())
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrUserNotFound) {
			st = getUserNotFoundStatus(ctx, err, req.GetUserId())
		} else if errors.Is(err, model.ErrUserPhoneValidate) {
			reason := fmt.Sprintf("invalid user phone: %s", req.GetPhone())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, errPhoneValidateLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.UpdateUserPhoneResponse{}, nil
}

// UpdateUserState ...
func (r *Registrator) UpdateUserState(ctx context.Context, req *registrator.UpdateUserStateRequest) (*registrator.UpdateUserStateResponse, error) {
	err := r.usersFacade.UpdateUserState(ctx, req.GetUserId(), model.UserState(req.GetState()))
	if err != nil {
		st := status.New(codes.Internal, err.Error())

		if errors.Is(err, model.ErrUserNotFound) {
			st = getUserNotFoundStatus(ctx, err, req.GetUserId())
		} else if errors.Is(err, model.ErrUserStateValidate) {
			reason := fmt.Sprintf("invalid user state: %d", req.GetState())
			st = getStatus(ctx, codes.InvalidArgument, err, reason, errStateValidateLexeme)
		}

		return nil, st.Err()
	}

	return &registrator.UpdateUserStateResponse{}, nil
}

func getUserNotFoundStatus(ctx context.Context, err error, userID int32) *status.Status {
	reason := fmt.Sprintf("user with id %d not found", userID)
	return getStatus(ctx, codes.NotFound, err, reason, i18n.UserNotFoundLexeme)
}
