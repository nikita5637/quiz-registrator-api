package gamephotos

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/nikita5637/quiz-registrator-api/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mocks"
	database "github.com/nikita5637/quiz-registrator-api/internal/pkg/storage/mysql"
	"github.com/nikita5637/quiz-registrator-api/internal/pkg/tx"
	timeutils "github.com/nikita5637/quiz-registrator-api/utils/time"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

type fixture struct {
	ctx    context.Context
	db     *tx.Manager
	dbMock sqlmock.Sqlmock
	facade *Facade

	gameStorage      *mocks.GameStorage
	gamePhotoStorage *mocks.GamePhotoStorage
}

func tearUp(t *testing.T) *fixture {
	db, dbMock, err := sqlmock.New()
	assert.NoError(t, err)

	fx := &fixture{
		ctx:    context.Background(),
		db:     tx.NewManager(db),
		dbMock: dbMock,

		gameStorage:      mocks.NewGameStorage(t),
		gamePhotoStorage: mocks.NewGamePhotoStorage(t),
	}

	fx.facade = &Facade{
		db: fx.db,

		gameStorage:      fx.gameStorage,
		gamePhotoStorage: fx.gamePhotoStorage,
	}

	t.Cleanup(func() {
		db.Close()
	})

	return fx
}

func Test_convertDBGamePhotoToModelGamePhoto(t *testing.T) {
	type args struct {
		game database.GamePhoto
	}
	tests := []struct {
		name string
		args args
		want model.GamePhoto
	}{
		{
			name: "test case 1",
			args: args{
				game: database.GamePhoto{
					ID:       1,
					FkGameID: 1,
					URL:      "url",
					CreatedAt: sql.NullTime{
						Time:  timeutils.TimeNow(),
						Valid: true,
					},
				},
			},
			want: model.GamePhoto{
				ID:       1,
				FkGameID: 1,
				URL:      "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDBGamePhotoToModelGamePhoto(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertDBGamePhotoToModelGamePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertModelGamePhotoToDBGamePhoto(t *testing.T) {
	type args struct {
		game model.GamePhoto
	}
	tests := []struct {
		name string
		args args
		want database.GamePhoto
	}{
		{
			name: "test case 1",
			args: args{
				game: model.GamePhoto{
					ID:       1,
					FkGameID: 1,
					URL:      "url",
				},
			},
			want: database.GamePhoto{
				ID:       1,
				FkGameID: 1,
				URL:      "url",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertModelGamePhotoToDBGamePhoto(tt.args.game); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertModelGamePhotoToDBGamePhoto() = %v, want %v", got, tt.want)
			}
		})
	}
}
