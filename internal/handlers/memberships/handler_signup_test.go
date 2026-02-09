package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
	"github.com/stretchr/testify/assert"
)

func TestHandler_SignUp(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	mockService := NewMockservice(ctrlMock)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		api     *gin.Engine
		service service
		// Named input parameters for target function.
		c                  *gin.Context
		mockFn             func()
		expectedStatusCode int
	}{
		// TODO: Add test cases.
		{name: "Success",
			mockFn: func() {
				mockService.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "test",
				}).Return(nil)
			}, expectedStatusCode: 200},
		{name: "Failed",
			mockFn: func() {
				mockService.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "test",
				}).Return(errors.New("Email or username exists"))
			}, expectedStatusCode: 400},
	}
	for _, tt := range tests {
		tt.mockFn()
		t.Run(tt.name, func(t *testing.T) {
			api := gin.New()
			h := NewHandler(api, mockService)
			h.RegisterRoute()
			w := httptest.NewRecorder()
			endpoint := `/memberships/sign-up`
			request := memberships.SignUpRequest{
				Email:    "test@gmail.com",
				Username: "testusername",
				Password: "test",
			}

			val, err := json.Marshal(request)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
		})
	}
}
