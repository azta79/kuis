package repository

import (
    "context"
    "errors"
    "log"
    "regexp"
    "testing"

    "github.com/Calmantara/go-kominfo-2024/go-middleware/internal/infrastructure/mocks"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func newMockGorm() (*gorm.DB, sqlmock.Sqlmock) {
    db, mock, err := sqlmock.New()
    if err != nil {
        log.Fatalf("Tidak ada kesalahan yang diharapkan saat membuka koneksi basis data tiruan: %s", err)
    }
    gormDB, err := gorm.Open(postgres.New(postgres.Config{
        Conn: db,
    }), &gorm.Config{})

    if err != nil {
        log.Fatalf("Tidak ada kesalahan yang diharapkan saat membuka basis data gorm: %s", err)
    }
    return gormDB, mock
}

func TestCreateUser(t *testing.T) {
    t.Run("kesalahan membuat pengguna", func(t *testing.T) {
        db, mock := newMockGorm()
        // mock infra
        postgresMock := mocks.NewGormPostgres(t)
        postgresMock.On("GetConnection").Return(db)
        // mock query
        mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO "users" ("username") VALUES ($1)
        `)).WillReturnError(errors.New("beberapa kesalahan"))

        userRepo := userQueryImpl{db: postgresMock}
        user := model.User{Username: "testuser"}
        _, err := userRepo.CreateUser(context.Background(), user)
        assert.NotNil(t, err)
    })

    t.Run("berhasil membuat pengguna", func(t *testing.T) {
        db, mock := newMockGorm()
        // mock infra
        postgresMock := mocks.NewGormPostgres(t)
        postgresMock.On("GetConnection").Return(db)
        // mock query
        mock.ExpectExec(regexp.QuoteMeta(`
            INSERT INTO "users" ("username") VALUES ($1)
        `)).WillReturnResult(sqlmock.NewResult(1, 1))

        userRepo := userQueryImpl{db: postgresMock}
        user := model.User{Username: "testuser"}
        _, err := userRepo.CreateUser(context.Background(), user)
        assert.Nil(t, err)
    })
}
