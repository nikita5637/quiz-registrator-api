package certificates

import (
	"context"
	"fmt"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	quizlogger "github.com/nikita5637/quiz-registrator-api/internal/pkg/quiz_logger"
	usersutils "github.com/nikita5637/quiz-registrator-api/utils/users"
)

// ListCertificates ...
func (f *Facade) ListCertificates(ctx context.Context) ([]model.Certificate, error) {
	var modelCertificates []model.Certificate
	err := f.db.RunTX(ctx, "ListCertificates", func(ctx context.Context) error {
		dbCertificates, err := f.certificateStorage.GetCertificates(ctx)
		if err != nil {
			return fmt.Errorf("get certificates error: %w", err)
		}

		modelCertificates = make([]model.Certificate, 0, len(dbCertificates))
		for _, dbCertificate := range dbCertificates {
			modelCertificates = append(modelCertificates, convertDBCertificateToModelCertificate(dbCertificate))
		}

		userID := maybe.Nothing[int32]()
		userFromContext := usersutils.UserFromContext(ctx)
		if userFromContext != nil {
			userID = maybe.Just(userFromContext.ID)
		}

		if err := f.quizLogger.Write(ctx, quizlogger.Params{
			UserID:     userID,
			ActionID:   quizlogger.ReadingActionID,
			MessageID:  quizlogger.GotCertificatesList,
			ObjectType: maybe.Nothing[string](),
			ObjectID:   maybe.Nothing[int32](),
			Metadata:   nil,
		}); err != nil {
			return fmt.Errorf("write log error: %w", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("ListCertificates error: %w", err)
	}

	return modelCertificates, nil
}
