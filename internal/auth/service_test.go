package auth_test

import (
	"14-TestingAPI/internal/auth"
	"14-TestingAPI/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error) {
	return &user.User{Email: "fake@fake.test"}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*user.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "user@test.test"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "123", "Тест userRepository")
	if err != nil {
		t.Fatal(err.Error())
	}
	if email != initialEmail {
		t.Fatalf("Email don't match %s != %s", initialEmail, email)
	}
}
