package registrator

import (
	"context"
	"strconv"

	telegram_utils "github.com/nikita5637/quiz-registrator-api/utils/telegram"
	users_utils "github.com/nikita5637/quiz-registrator-api/utils/users"

	"google.golang.org/grpc/metadata"
)

// AuthFuncOverride ...
func (r *Registrator) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	telegramClientID := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		telegramClientIDs := md.Get(telegram_utils.TelegramClientID)
		if len(telegramClientIDs) > 0 {
			telegramClientID = telegramClientIDs[0]
		}
	}

	id, err := strconv.ParseInt(telegramClientID, 10, 64)
	if err == nil {
		ctx = telegram_utils.NewContextWithClientID(ctx, id)
		user, err := r.usersFacade.GetUser(ctx)
		if err == nil {
			ctx = users_utils.NewContextWithUser(ctx, user)
		}
	}

	return ctx, nil
}
