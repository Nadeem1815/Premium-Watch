package repository

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestUserRegister(t *testing.T) {
	tests := []struct {
		name           string
		input          model.UsarDataInput
		expectedOutput model.UserDataOutput
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{
			//test case for creating a new user
			name: "successful create user",
			input: model.UsarDataInput{
				Name:     "Nadeem",
				Surname:  "F",
				EmailId:  "nadeem@gmail.com",
				Password: "nadeem@123",
				Phone:    "8129487958",
			},
			expectedOutput: model.UserDataOutput{
				ID:      "1",
				Name:    "Nadeem",
				Surname: "F",
				EmailId: "nadeem@gmail.com",
				Phone:   "8129487958",
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "surname", "email_id", "password", "phone"}).
					AddRow("1", "Nadeem", "F", "nadeem@gmail.com", "nadeem@123", "8129487958")

				mock.ExpectQuery("^INSERT INTO users (.+)$").
					WithArgs("Nadeem", "F", "nadeem@gmail.com", "nadeem@123", "8129487958").
					WillReturnRows(rows)

				// mock.ExpectExec("^INSERT INTO user_infos (.+)$").
				// 	WithArgs(2).
				// 	WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: nil,
		},
		{
			// test case for trying to insert a user duplicate id's
			name: "duplicate user",
			input: model.UsarDataInput{
				Name:     "Nadeem",
				Surname:  "F",
				EmailId:  "nadeem@gmail.com",
				Password: "nadeem@123",
				Phone:    "8129487958",
			},
			expectedOutput: model.UserDataOutput{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO users(.+)$").
					WithArgs("Nadeem", "F", "nadeem@gmail.com", "nadeem@123", "8129487958").
					WillReturnError(errors.New("email or phone is already used"))

			},
			expectedErr: errors.New("email or phone is already used"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//New() method from sqlmock package create sqlmock database connection and a mock to manage expectations.
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			//close the mock db connection after testing.
			defer db.Close()

			//initialize a mock db session
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			//create NewUserRepository mock by passing a pointer to gorm.DB
			userRepository := NewUserRepository(gormDB)

			// before we actually execute our function, we need to expect required DB actions
			tt.buildStub(mock)

			//call the actual method
			actualOutput, actualErr := userRepository.UserRegister(context.TODO(), tt.input)
			// validate err is nil if we are not expecting to receive an error
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else { //validate whether expected and actual errors are same
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
			}

			// Check that all expectations were met
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
