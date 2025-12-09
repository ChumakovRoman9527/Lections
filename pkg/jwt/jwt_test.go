package jwt_test

import (
	"14-TestingAPI/pkg/jwt"
	"testing"
)

func TestJWTSign(t *testing.T) {
	const email = "test_test@test.test"
	jwtService := jwt.NewJWT("jc)!G<[p*cqRo8,9euU1Y(^Jgr!O<%EDJ6pwQE|h{Fe")
	token, err := jwtService.Create(jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(token)
	if !isValid {
		t.Fatal("Token is Invalid !")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s !", data.Email, email)
	}
}
