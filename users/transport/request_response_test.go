package transport

import (
	"bootcampProject/users/domain"
	"errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/validator.v2"
	"testing"
)

func TestCreateUserRequest_Validate(t *testing.T) {

	t.Run("Ok: Request is correct", func(t *testing.T) {
		reqOk := CreateUserRequest{User: domain.Users{
			PwdHash: "password",
			Name:    "Name test",
			Age:     25,
			Email:   "test@test.com",
		}}

		errsGot := reqOk.Validate()
		assert.NoError(t, errsGot)
	})

	t.Run("Error: empty fields", func(t *testing.T) {
		reqError := CreateUserRequest{User: domain.Users{
			PwdHash: "password",
			Name:    "",
			Age:     0,
			Email:   "test@test.com",
		}}
		errsGot := reqError.Validate()
		assert.Error(t, errsGot)
	})

	t.Run("Error: email format", func(t *testing.T) {
		reqEmailError := CreateUserRequest{User: domain.Users{
			PwdHash: "password",
			Name:    "Name test",
			Age:     25,
			Email:   "email",
		}}

		errExpected := errors.New("Email: " + validator.ErrRegexp.Error())

		errsGot := reqEmailError.Validate()
		assert.EqualError(t, errsGot, errExpected.Error())
	})
}
