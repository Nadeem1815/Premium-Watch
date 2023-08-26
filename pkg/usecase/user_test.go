package usecase

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/nadeem1815/premium-watch/pkg/repository/mockrepo"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      model.UsarDataInput
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(model.UsarDataInput)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}
func EqCreateUserParams(arg model.UsarDataInput, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

// func TestUserRegister(t *testing.T) {

// 	// constTime := time.Now()

// 	testData := []struct {
// 		name           string
// 		input          model.UsarDataInput
// 		buildStub      func(userRepo *mockrepo.MockUserRepository, user model.UsarDataInput)
// 		expectedOutput model.UserDataOutput
// 		expectedError  error
// 	}{
// 		{
// 			name: "FailedToSaveUserOnDatabase",
// 			input: model.UsarDataInput{
// 				Name:     "Nadeem",
// 				Surname:  "Fahad",
// 				EmailId:  "nadeemf408@gmail.com",
// 				Phone:    "8129487958",
// 				Password: "Nadeem@123",
// 			},
// 			buildStub: func(userRepo *mockrepo.MockUserRepository, user model.UsarDataInput) {

// 				userRepo.EXPECT().UserRegister(gomock.Any(), EqCreateUserParams(user, user.Password)).Times(1).
// 					Return(model.UserDataOutput{}, errors.New("error on database"))
// 			},
// 			expectedOutput: model.UserDataOutput{},
// 			expectedError:  errors.New("error on database"),
// 		},
// 		{
// 			name: "SuccessSignup",
// 			input: model.UsarDataInput{
// 				Name:     "Nadeem",
// 				Surname:  "Fahad",
// 				EmailId:  "nadeemf408@gmail.com",
// 				Phone:    "8129487958",
// 				Password: "Nadeem@123",
// 			},
// 			buildStub: func(userRepo *mockrepo.MockUserRepository, user model.UsarDataInput) {
// 				userRepo.EXPECT().UserRegister(gomock.Any(), EqCreateUserParams(user, user.Password)).Times(1).
// 					Return(model.UserDataOutput{
// 						ID:      1,
// 						Name:    "Nadeem",
// 						Surname: "Fahad",
// 						EmailId: "nadeemf408@gmail.com",
// 						Phone:   "8129487958",
// 					}, nil)
// 			},
// 			expectedOutput: model.UserDataOutput{
// 				ID:      1,
// 				Name:    "Nadeem",
// 				Surname: "Fahad",
// 				EmailId: "nadeemf408@gmail.com",
// 				Phone:   "8129487958",
// 			},
// 			expectedError: nil,
// 		},
// 	}
// 	for _, tt := range testData {
// 		t.Run(tt.name, func(t *testing.T) {

// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()
// 			userRepo := mockrepo.NewMockUserRepository(ctrl)
// 			userUseCase := NewUserUseCase(userRepo)
// 			tt.buildStub(userRepo, tt.input)

// 			user, err := userUseCase.UserRegister(context.TODO(), tt.input)

// 			assert.Equal(t, tt.expectedError, err)
// 			assert.Equal(t, user, tt.expectedOutput)
// 		})
// 	}

// }

func TestUserRegister(t *testing.T) {
	// NewController from gomock package returns a new controller for testing
	ctrl := gomock.NewController(t)
	//NewMockUserRepository creates a new mock instance of the user repo
	userRepo := mockrepo.NewMockUserRepository(ctrl)
	//NewMockOrderRepository creates a new mock instance of the order repo
	// orderRepo := mockRepo.NewMockOrderRepository(ctrl)

	userUseCase := NewUserUseCase(userRepo)
	mockPassword := "password@123"

	testData := []struct {
		name           string
		input          model.UsarDataInput
		buildStub      func(userRepo mockrepo.MockUserRepository)
		expectedOutput model.UserDataOutput
		expectedError  error
	}{
		{
			name: "new user",
			input: model.UsarDataInput{
				Name:     "Nadeem",
				Surname:  "Fahad",
				EmailId:  "nadeemf408@gmail.com",
				Phone:    "8129487958",
				Password: mockPassword,
			},
			buildStub: func(userRepo mockrepo.MockUserRepository) {

				//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(mockPassword), 10)
				//if err != nil {
				//	t.Errorf("failed to hash password for testing : %q", err)
				//}

				userRepo.EXPECT().
					UserRegister(
						gomock.Any(),
						EqCreateUserParams(
							model.UsarDataInput{
								Name:     "Nadeem",
								Surname:  "Fahad",
								EmailId:  "nadeemf408@gmail.com",
								Phone:    "8129487958",
								Password: mockPassword},
							mockPassword),
					).
					Times(1).
					Return(model.UserDataOutput{
						ID:      1,
						Name:    "Nadeem",
						Surname: "Fahad",
						EmailId: "nadeemf408@gmail.com",
						Phone:   "8129487958",
					}, nil)
			},
			expectedOutput: model.UserDataOutput{
				ID:      1,
				Name:    "Nadeem",
				Surname: "Fahad",
				EmailId: "nadeemf408@gmail.com",
				Phone:   "8129487958",
			},
			expectedError: nil,
		},
		{
			name: "duplicate user",
			input: model.UsarDataInput{
				Name:     "Nadeem",
				Surname:  "Fahad",
				EmailId:  "nadeemf408@gmail.com",
				Phone:    "8129487958",
				Password: mockPassword,
			},
			buildStub: func(userRepo mockrepo.MockUserRepository) {
				userRepo.EXPECT().
					UserRegister(
						gomock.Any(),
						EqCreateUserParams(
							model.UsarDataInput{
								Name:     "Nadeem",
								Surname:  "Fahad",
								EmailId:  "nadeemf408@gmail.com",
								Phone:    "8129487958",
								Password: mockPassword},
							mockPassword),
					).
					Times(1).
					Return(model.UserDataOutput{}, errors.New("user already exists"))
			},
			expectedOutput: model.UserDataOutput{},
			expectedError:  errors.New("user already exists"),
		},
	}

	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userRepo)
			actualUser, err := userUseCase.UserRegister(context.TODO(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, actualUser, tt.expectedOutput)
		})
	}

}
