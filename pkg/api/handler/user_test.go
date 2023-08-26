package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/nadeem1815/premium-watch/pkg/usecase/mockusecase"
	"github.com/nadeem1815/premium-watch/pkg/utils/model"
	"github.com/nadeem1815/premium-watch/pkg/utils/response"
	"github.com/stretchr/testify/assert"
)

func TestUserRegister(t *testing.T) {
	//NewController creates a mock controller for testing
	ctrl := gomock.NewController(t)
	//NewMockUserUseCase creates a new mock instance of the user use case
	userUseCase := mockusecase.NewMockUserUseCase(ctrl)
	//NewUserHandler create a new user user handler
	userHandler := NewUserHandler(userUseCase)

	//testData is a slice of anonymous structs which are initialized
	testData := []struct {
		name             string // name of the test case
		userData         model.UsarDataInput
		buildStub        func(userUsecase mockusecase.MockUserUseCase) // function to generate the mock user use case
		expectedCode     int                                           // expected status code
		expectedResponse response.Response                             // expected response for the request
		expectedData     model.UserDataOutput                          //expected data in output
		expectedError    error                                         // expected error in the response
	}{
		{ // test case for checking user sign up for a non-registered users
			name: "non existing user",
			userData: model.UsarDataInput{
				Name:     "Nadeem",
				Surname:  "Fahad",
				EmailId:  "nadeemf408@gmail.com",
				Phone:    "8129487958",
				Password: "Nadeem@123",
			},
			buildStub: func(userUsecase mockusecase.MockUserUseCase) {
				userUsecase.EXPECT(). //setting the expected behaviour of the usecase method
							UserRegister(gomock.Any(), model.UsarDataInput{ //CreateUser usecase receives two arguments, one is context.so we can use gomock.Any(). next one is user signup information
						Name:     "Nadeem",
						Surname:  "Fahad",
						EmailId:  "nadeemf408@gmail.com",
						Phone:    "8129487958",
						Password: "Nadeem@123",
					}).
					Times(1).                    //how many times the CreateUser usecase should be called
					Return(model.UserDataOutput{ //what should CreateUser usecase return. Here it should return user info and nil (error)
						ID:      1,
						Name:    "Nadeem",
						Surname: "Fahad",
						EmailId: "nadeemf408@gmail.com",
						Phone:   "8129487958",
					}, nil)
			},
			expectedCode: 201,
			expectedResponse: response.Response{
				StatusCode: 201,
				Message:    "User created successfully",
				Data: model.UserDataOutput{
					ID:      1,
					Name:    "Nadeem",
					Surname: "Fahad",
					EmailId: "nadeemf408@gmail.com",
					Phone:   "8129487958",
				},
				Errors: nil,
			},
			expectedData: model.UserDataOutput{
				ID:      1,
				Name:    "Nadeem",
				Surname: "Fahad",
				EmailId: "nadeemf408@gmail.com",
				Phone:   "8129487958",
			},
		},
		// {
		// 	//	test case for checking signup of duplicate user
		// 	name: "duplicate user",
		// 	userData: model.UserDataInput{
		// 		FName:    "Amal",
		// 		LName:    "Madhu",
		// 		Email:    "amalmadhu@gmail.com",
		// 		Phone:    "7902631234",
		// 		Password: "password@123",
		// 	},
		// 	buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
		// 		userUsecase.EXPECT().
		// 			CreateUser(gomock.Any(), model.UserDataInput{
		// 				FName:    "Amal",
		// 				LName:    "Madhu",
		// 				Email:    "amalmadhu@gmail.com",
		// 				Phone:    "7902631234",
		// 				Password: "password@123",
		// 			}).
		// 			Times(1).
		// 			Return(
		// 				model.UserDataOutput{},
		// 				errors.New("user already exists"),
		// 			)
		// 	},
		// 	expectedCode: 400,
		// 	expectedResponse: response.Response{
		// 		StatusCode: 400,
		// 		Message:    "failed to create user",
		// 		Data:       model.UserDataOutput{},
		// 		Errors:     "failed to create user",
		// 	},
		// 	expectedError: errors.New("failed to create user"),
		// 	expectedData:  model.UserDataOutput{},
		// },
	}

	// looping through the test cases and running them individually.
	// for _, tt := range testData {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		//passing mock use case to buildStub function which is
	// 		tt.buildStub(*userUseCase)
	// 		// gin.Default will create a new engine instance with logger middleware by default
	// 		engine := gin.Default()
	// 		//NewRecorder from httptest package will create a recorder which records the response
	// 		recorder := httptest.NewRecorder()
	// 		//create new route for testing
	// 		engine.POST("/signup", userHandler.UserRegister)
	// 		//body will hold a slice of bytes. It is used for Marshaling json data and passing to the request body
	// 		var body []byte
	// 		//marshaling userdata in testcase to body
	// 		body, err := json.Marshal(tt.userData)
	// 		//validating no error occurred while marshaling userdata to body
	// 		assert.NoError(t, err)
	// 		//url for the test
	// 		url := "/signup"
	// 		// req is a pointer to http.Request . With httptest.NewRequest we are mentioning the http method, endpoint and body
	// 		req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	// 		//recorder will record the response, and req is the mock req that we created with test data
	// 		engine.ServeHTTP(recorder, req)
	// 		//actual will hold the actual response
	// 		var actual response.Response
	// 		//unmarshalling json data to response.Response format
	// 		err = json.Unmarshal(recorder.Body.Bytes(), &actual)
	// 		assert.NoError(t, err)
	// 		assert.Equal(t, tt.expectedCode, recorder.Code)
	// 		assert.Equal(t, tt.expectedResponse.Message, actual.Message)

	// 		//check if data is of type map[string]interface{}
	// 		fmt.Printf("type of actual data %t\n", actual.Data)
	// 		data, ok := actual.Data.(map[string]interface{})
	// 		if ok {
	// 			userData := model.UserDataOutput{
	// 				ID:    uint(int(data["user_id"].(float64))),
	// 				FName: data["f_name"].(string),
	// 				LName: data["l_name"].(string),
	// 				Email: data["email"].(string),
	// 				Phone: data["phone"].(string),
	// 			}
	// 			if !reflect.DeepEqual(tt.expectedData, userData) {
	// 				t.Errorf("got %q, but want %q", userData, tt.expectedData)
	// 			}
	// 		} else {
	// 			t.Errorf("actual.Data is not of type map[string]interface{}")
	// 		}

	// 	})
	// }
	for _, tc := range testData {

		t.Run(tc.name, func(t *testing.T) {

			tc.buildStub(*userUseCase)

			server := gin.New()
			server.POST("/signup", userHandler.UserRegister)

			jsonData, err := json.Marshal(&tc.userData)
			assert.NoError(t, err)
			body := bytes.NewBuffer(jsonData)

			mockReq, err := http.NewRequest(http.MethodPost, "/signup", body)
			assert.NoError(t, err)

			responseRec := httptest.NewRecorder()

			server.ServeHTTP(responseRec, mockReq)

			//validate
			assert.Equal(t, tc.expectedCode, responseRec.Code)
		})
	}

}
