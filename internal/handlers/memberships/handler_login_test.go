package memberships

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/kautsarhasby/katalog-musik/internal/models/memberships"
	"github.com/stretchr/testify/assert"
)

func TestHandler_Login(t *testing.T) {
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()

	mockService := NewMockservice(ctrlMock)

	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		api     *gin.Engine
		service service
		// Named input parameters for target function.
		c *gin.Context

		wantErr            bool
		expectedStatusCode int
		expectedBody       memberships.LoginResponse
		mockFn             func()
	}{
		// TODO: Add test cases.

		{
			name:    "Success",
			wantErr: false,
			mockFn: func() {
				mockService.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "test",
				}).Return("accessToken", nil)
			},
			expectedStatusCode: 200,
			expectedBody: memberships.LoginResponse{
				AccessToken: "accessToken",
			},
		},
		{
			name:    "false",
			wantErr: true,
			mockFn: func() {
				mockService.EXPECT().Login(memberships.LoginRequest{
					Email:    "test@gmail.com",
					Password: "test",
				}).Return("", assert.AnError)
			},
			expectedStatusCode: 400,
			expectedBody:       memberships.LoginResponse{},
		},
	}
	for _, tt := range tests {
		tt.mockFn()
		t.Run(tt.name, func(t *testing.T) {
			api := gin.New()
			h := NewHandler(api, mockService)
			h.RegisterRoute()
			w := httptest.NewRecorder()
			endpoint := `/memberships/login`

			request := &memberships.LoginRequest{
				Email:    "test@gmail.com",
				Password: "test",
			}

			val, err := json.Marshal(request)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)
			h.ServeHTTP(w, req)

			if !tt.wantErr {

				response := memberships.LoginResponse{}
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}

		})
	}
}
