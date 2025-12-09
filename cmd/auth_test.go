package main

import (
	"14-TestingAPI/internal/auth"
	"14-TestingAPI/internal/user"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "test_test@test.test",
		Password: "$2a$10$PJi66i/azH2OL5rt18RXlOd.LMYYNN0k49yFMtVNPuDWTG2L3eAjK",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().Where("Email = ?", "test_test@test.test").Delete(&user.User{})
}
func TestLoginSuccess(t *testing.T) {
	//Prepare
	db := initDB()
	initData(db)
	//Test
	ts := httptest.NewServer(App())
	defer removeData(db)
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test_test@test.test",
		Password: "123",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d", 200, res.StatusCode)
	}

	//1 проверка на токен
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.RegisterResponse

	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}

	if resData.TOKEN == "" {
		t.Fatal("Token nil !")
	}

}
func TestLoginFailed(t *testing.T) {
	db := initDB()
	initData(db)
	ts := httptest.NewServer(App())
	defer removeData(db)
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "test_test@test.test",
		Password: "123_",
	})

	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d", 401, res.StatusCode)
	}

}
