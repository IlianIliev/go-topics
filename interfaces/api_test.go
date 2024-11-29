package interfaces

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
	"to-topics/interfaces/services/users"
)

func TestCreateUserHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	usersServiceMock := users.NewMockUserService(ctrl)

	testCases := []struct {
		name           string
		payload        string
		setUpMocks     func()
		expectedStatus int
	}{
		{
			name:           "Success",
			expectedStatus: http.StatusCreated,
			payload: `{
				"name": "John Doe",
				"email": "john_doe@example.org",
				"password": "password"
			}`,
			setUpMocks: func() {
				usersServiceMock.EXPECT().CreateUser(users.CreateUserInput{
					Name:     "John Doe",
					Email:    "john_doe@example.org",
					Password: "password",
				}).Return(nil)
			},
		},
		{
			name:           "UnableToCreateUser",
			expectedStatus: http.StatusInternalServerError,
			payload: `{
				"name": "John Doe",
				"email": "john_doe@example.org",
				"password": "password"
			}`,
			setUpMocks: func() {
				usersServiceMock.EXPECT().CreateUser(users.CreateUserInput{
					Name:     "John Doe",
					Email:    "john_doe@example.org",
					Password: "password",
				}).Return(errors.New("something went wrong"))
			},
		},
		{
			name:           "BadRequest",
			expectedStatus: http.StatusBadRequest,
			payload: `{
				"name": "John Doe"
			}`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if testCase.setUpMocks != nil {
				testCase.setUpMocks()
			}
			request, err := http.NewRequest(
				http.MethodGet,
				"https://example.com/api/users",
				bytes.NewReader([]byte(testCase.payload)))
			if err != nil {
				t.Fatal("Error creating request:", err)
				return
			}

			writer := httptest.NewRecorder()

			handler := createUserHandler(usersServiceMock)
			handler(writer, request)

			require.Equal(t, testCase.expectedStatus, writer.Code)
		})
	}
}
