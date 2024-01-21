package quizlogger

import "github.com/nikita5637/quiz-registrator-api/internal/pkg/model"

const (
	// GameRegistered ...
	GameRegistered = iota + 1
	// GameUnregistered ...
	GameUnregistered
	// GamePaymentChanged ...
	GamePaymentChanged
	// GotCompleteListOfGames ...
	GotCompleteListOfGames
	// GotListOfUserGameIDs ...
	GotListOfUserGameIDs
	// GotListOfPassedAndRegisteredGames ...
	GotListOfPassedAndRegisteredGames
	// GotGamePhotos ...
	GotGamePhotos
)

// GamePaymentChangedMetadata ...
type GamePaymentChangedMetadata struct {
	OldPayment model.Payment
	NewPayment model.Payment
}

// GotListOfUserGameIDsMetadata ...
type GotListOfUserGameIDsMetadata struct {
	UserID int32
}

// GotListOfPassedAndRegisteredGamesMetadata ...
type GotListOfPassedAndRegisteredGamesMetadata struct {
	Offset uint64
	Limit  uint64
}
