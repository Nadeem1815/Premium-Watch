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
		{ //test case for creating a new user
			name: "successful create user",
			input: model.UsarDataInput{
				Name:     "Muhammed",
				Surname:  "S",
				EmailId:  "Muhammed@gmail.com",
				Phone:    "7902638845",
				Password: "muhammed@123",
			},
			expectedOutput: model.UserDataOutput{
				ID:      6,
				Name:    "Muhammed",
				Surname: "S",
				EmailId: "Muhammed@gmail.com",
				Phone:   "7902638845",
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "surname", "email_id", "password", "phone"}).
					AddRow(6, "Muhammed", "S", "Muhammed@gmail.com", "muhammed@123", "7902638845")

				// mock.ExpectQuery("^INSERT INTO users (.+)$").
				// 	WithArgs("Nadeem", "F", "nadeem@gmail.com", "nadeem@123", "8129487958").
				// 	WillReturnRows(rows)
				mock.ExpectQuery(`^INSERT INTO users\(name,surname,email_id,password,phone,created_at\) VALUES\(.+\) RETURNING id,name,surname,email_id,phone$`).
					WithArgs("Muhammed", "S", "Muhammed@gmail.com", "muhammed@123", "7902638845").
					WillReturnRows(rows)

				mock.ExpectExec("^INSERT INTO user_infos (.+)$").
					WithArgs(6).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			expectedErr: nil,
		},
		{ //test case for trying to insert a user with duplicate phone id
			name: "duplicate phone",
			input: model.UsarDataInput{
				Name:     "Muhammed",
				Surname:  "S",
				EmailId:  "Muhammed@gmail.com",
				Phone:    "7902638845",
				Password: "muhammed@123",
			},
			expectedOutput: model.UserDataOutput{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`^INSERT INTO users\(name,surname,email_id,password,phone,created_at\) VALUES\(.+\) RETURNING id,name,surname,email_id,phone$`).
					WithArgs("Muhammed", "S", "Muhammed@gmail.com", "muhammed@123", "7902638845").
					WillReturnError(errors.New("duplicate key value violates unique constraint 'phone'"))
			},
			expectedErr: errors.New("duplicate key value violates unique constraint 'phone'"),
		},
		{ //test case for trying to insert a user with duplicate email_id
			name: "duplicate email",
			input: model.UsarDataInput{
				Name:     "Muhammed",
				Surname:  "S",
				EmailId:  "Muhammed@gmail.com",
				Phone:    "7902638845",
				Password: "muhammed@123",
			},
			expectedOutput: model.UserDataOutput{},
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(`^INSERT INTO users\(name,surname,email_id,password,phone,created_at\) VALUES\(.+\) RETURNING id,name,surname,email_id,phone$`).
					WithArgs("Muhammed", "S", "Muhammed@gmail.com", "muhammed@123", "7902638845").
					WillReturnError(errors.New("duplicate key value violates unique constraint 'email'"))
			},
			expectedErr: errors.New("duplicate key value violates unique constraint 'email'"),
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

// func TestUserRegister(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		input          model.UsarDataInput
// 		expectedOutput model.UserDataOutput
// 		buildStub      func(mock sqlmock.Sqlmock)
// 		expectedErr    error
// 	}{
// 		// ... existing test case ...

// 		{ // test case for database error during user insertion
// 			name: "database error during user insertion",
// 			input: model.UsarDataInput{
// 				Name:     "Muhammed",
// 				Surname:  "S",
// 				EmailId:  "muhammed@example.com",
// 				Phone:    "1234567890",
// 				Password: "muhammed123",
// 			},
// 			expectedOutput: model.UserDataOutput{},
// 			buildStub: func(mock sqlmock.Sqlmock) {
// 				rows := sqlmock.NewRows([]string{"id", "name", "surname", "email_id", "phone", "password"}).
// 					AddRow(3, "Muhammed", "S", "muhammed@gmail.com", "7902638845", "muhammed@123")
// 				mock.ExpectQuery("^INSERT INTO users (.+)$").
// 					WillReturnRows(rows)

// 				// No expectations for the second query (user_infos)
// 			},
// 			expectedErr: fmt.Errorf("database error"),
// 		},

// 		// Add more test cases here for different scenarios...

// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// ... existing test setup ...
// 			db, mock, err := sqlmock.New()
// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 			}
// 			//close the mock db connection after testing.
// 			defer db.Close()

// 			//initialize a mock db session
// 			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
// 			if err != nil {
// 				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
// 			}

// 			//create NewUserRepository mock by passing a pointer to gorm.DB
// 			userRepository := NewUserRepository(gormDB)

// 			// before we actually execute our function, we need to expect required DB actions
// 			tt.buildStub(mock)

// 			actualOutput, actualErr := userRepository.UserRegister(context.TODO(), tt.input)

// 			if tt.expectedErr == nil {
// 				assert.NoError(t, actualErr)
// 			} else {
// 				assert.Equal(t, tt.expectedErr, actualErr)
// 			}

// 			if !reflect.DeepEqual(tt.expectedOutput, actualOutput) {
// 				t.Errorf("got %v, but want %v", actualOutput, tt.expectedOutput)
// 			}

// 			// ... existing expectation check ...
// 			err = mock.ExpectationsWereMet()
// 			if err != nil {
// 				t.Errorf("Unfulfilled expectations: %s", err)
// 			}
// 		})
// 	}
// }

func TestFindbyEmail(t *testing.T) {

	testData := []struct {
		name           string
		email          string
		expectedOutput model.UserLoginVarifier
		buildStub      func(mock sqlmock.Sqlmock)
		expectedErr    error
	}{
		{ // test case finding valid user
			name:  "valid email",
			email: "muhammed@gmail.com",
			expectedOutput: model.UserLoginVarifier{
				ID:        1,
				Name:      "Nadeem",
				Surname:   "Fahad",
				EmailId:   "muhammed@gmail.com",
				Phone:     "8129487958",
				Password:  "nadeem@123",
				IsBlocked: false,
			},
			buildStub: func(mock sqlmock.Sqlmock) {
				// columns := []string{"id", "name", "surname", "email_id", "phone", "password", "is_blocked"}
				// expectedQuery := `SELECT u.id,u.name,u.surname,u.email_id,u.password,u.phone,info.is_blocked FROM users as u FULL OUTER JOIN user_infos as info ON u.id=info.users_id WHERE u.email_id=$1`
				// mock.ExpectQuery(expectedQuery).
				// 	WithArgs("muhammed@gmail.com").
				// 	WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Nadeem", "Fahad", "muhammed@gmail.com", "8129487958", "nadeem@123", false))
				mock.ExpectQuery("SELECT u.id, u.name, u.surname, u.email_id, u.phone, u.password, info.is_blocked FROM users FULL OUTER JOIN user_infos as info ON u.id=info.users_id WHERE u.email_id=$1;").
					WithArgs("muhammed@gmail.com").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "surname", "email_id", "phone", "password", "is_blocked"}).
						AddRow(1, "Nadeem", "Fahad", "muhammed@gmail.com", "8129487958", "nadeem@123", false))
				// WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Nadeem", "Fahad", "muhammed@gmail.com", "8129487958", "nadeem@123", true))
				// mock.ExpectQuery("SELECT u.id, u.name, u.surname, u.email_id, u.phone, u.password, infos.is_blocked FROM users FULL OUTER JOIN user_infos as info ON u.id = info.users_id WHERE u.email_id=$1").
				// 	WithArgs("muhammed@gmail.com").
				// 	WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Nadeem", "Fahad", "muhammed@gmail.com", "8129487958", "nadeem@123", false))
				// mock.ExpectQuery(`SELECT u.id,u.name,u.surname,u.email_id,u.password,u.phone,info.is_blocked FROM users as u FULL OUTER JOIN user_infos as info ON u.id=info.users_id WHERE u.email_id=$1`).
				// 	WithArgs("muhammed@gmail.com").
				// 	WillReturnRows(sqlmock.NewRows(columns).AddRow(1, "Nadeem", "Fahad", "muhammed@gmail.com", "8129487958", "nadeem@123", false))

			},
			expectedErr: nil,
		},
	}
	for _, tt := range testData {
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
			actualOutput, actualErr := userRepository.FindByEmail(context.TODO(), tt.email)
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
