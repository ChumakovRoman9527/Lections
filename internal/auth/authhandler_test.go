package auth_test

import (
	"14-TestingAPI/configs"
	"14-TestingAPI/internal/auth"
	"14-TestingAPI/internal/user"
	"14-TestingAPI/pkg/db"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func bootstrap() (*auth.AuthHandler, sqlmock.Sqlmock, error) {
	database, mock, err := sqlmock.New()
	if err != nil {

		return nil, nil, err
	}
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: database}))
	if err != nil {
		return nil, nil, err
	}
	userRepo := user.NewUserRepository(&db.Db{DB: gormDb})
	handler := auth.AuthHandler{
		Config: &configs.Config{
			Auth: configs.AuthConfig{Secret: "test_secret"}},
		AuthService: &auth.AuthService{
			UserRepository: userRepo,
		},
	}
	return &handler, mock, nil
}
func TestLoginSuccess(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatalf("Не удалось получить handler и mock: %s", err.Error())
		return
	}

	rows := sqlmock.NewRows([]string{"email", "password"}).
		AddRow("test_test@test.test", "$2a$10$PJi66i/azH2OL5rt18RXlOd.LMYYNN0k49yFMtVNPuDWTG2L3eAjK")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test_test@test.test",
		Password: "123",
	})
	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/login", reader)
	handler.Login()(wr, req)
	if wr.Result().StatusCode != http.StatusOK {
		t.Fatalf("Статус код не %d а: %d", http.StatusOK, wr.Result().StatusCode)
		return
	}
}
func TestRegisterHandler(t *testing.T) {
	handler, mock, err := bootstrap()
	if err != nil {
		t.Fatalf("Не удалось получить handler и mock: %s", err.Error())
		return
	}

	rows := sqlmock.NewRows([]string{"email", "password", "name"})
	// .AddRow("test_test@test.test", "$2a$10$PJi66i/azH2OL5rt18RXlOd.LMYYNN0k49yFMtVNPuDWTG2L3eAjK")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// data, _ := json.Marshal(auth.RegisterRequest{
	// 	Email:    "test_test@test.test",
	// 	Password: "123",
	// 	Name:     "testing",
	// })
	data, _ := json.Marshal(&auth.RegisterRequest{
		LoginRequest: auth.LoginRequest{
			Email:    "test_test@test.test",
			Password: "123",
		},
		Name: "test",
	})
	reader := bytes.NewReader(data)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/auth/register", reader)
	handler.Register()(wr, req)
	if wr.Result().StatusCode != http.StatusOK {
		t.Fatalf("Статус код не %d а: %d", http.StatusOK, wr.Result().StatusCode)
		return
	}
}
